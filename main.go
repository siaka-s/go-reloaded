package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Fonction de convertion hecadecimal en entier
func convertirhex(hexnum string) int64 {

	decimal_num, err := strconv.ParseInt(hexnum, 16, 64)

	if err != nil {
		panic(err)
	}
	return decimal_num
}

// convertion de binaire en entier
func convertirbin(binaire string) int64 {

	n, err := strconv.ParseInt(binaire, 2, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return n
}

// stockage de la phrase dans un tableau et execution de cas

func wordTabExe(sample string) []byte {

	// récuperation du contenu du filchier
	mots, err := ioutil.ReadFile(sample)

	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
	}
	return mots
}

// Fonction de traitement
func traitement(phrase string) string {

	mots := strings.Fields(phrase)

	for i := 0; i < len(mots); i++ {

		switch mots[i] {
		case "(hex)":
			mots[i] = ""

			a := i - 1

			dec := convertirhex(mots[a])
			// converti ne nombre en chaine de caratère
			decf := strconv.FormatInt(dec, 10)

			// assigné la valeur a la position
			mots[a] = decf
			// convertir le chiffre dec en chaine de caractère
		case "(bin)":
			mots[i] = ""

			a := i - 1

			dec := convertirbin(mots[a])

			decf := strconv.FormatInt(dec, 10)

			mots[a] = decf

		case "(up)":

		case "(low)":

		case "(cap)":

		default:
			fmt.Println("Aucune modification à éffectuer")
		}
	}
	// convertir le tableau en chaine de caractère avec un espace de séparation
	result := strings.Join(mots, " ")
	// enlever les espaces après traitement
	result = strings.Replace(result, "  ", " ", -1)

	return result
}

// Migration du contenu modifier dans un nouveau fichier
func remplacerContenuFichier(nvFichier string, newContenu string) error {
	return ioutil.WriteFile(nvFichier, []byte(newContenu), 0666)
}

func main() {

	sample := os.Args[1]

	result := os.Args[2]

	// strockage du contenu du fichier dans une variable
	phrase := string(wordTabExe(sample))

	// exécution de la fonction de traitement
	contenu := (traitement(phrase))

	// execution de la fontion remplacement avec affichage de message d'exécution réussie ou d'erreur
	err := remplacerContenuFichier(result, contenu)
	if err != nil {
		fmt.Println("Erreur lors de la réécriture du fichier :", err)
	} else {
		fmt.Println("Le contenu du fichier a été remplacé avec succès")
	}

}
