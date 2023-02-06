package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Fonction hecadecimal en entier
func convertirhex(hexnum string) int64 {

	// la fonction parseInt() nous permet de réaliser la conversion
	decimal_num, err := strconv.ParseInt(hexnum, 16, 64)

	// Gestion de l'erreur du nombre recu en paramètre
	if err != nil {
		panic(err)
	}
	return decimal_num
}

// stocage de la phrase dans un tableau et execution de cas

func wordTabExe(sample string) []byte {

	// récuperation du contenu du filchier
	mots, err := ioutil.ReadFile(sample)

	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
	}
	return mots
}

func traitement(phrase string) string {

	mots := strings.Fields(phrase)

	for i := 0; i < len(mots); i++ {

		if mots[i] == "(hex)" {

			mots[i] = ""

			a := i - 1

			dec := convertirhex(mots[a])

			decf := strconv.FormatInt(dec, 10)

			// assigné la valeur a la position
			mots[a] = decf
			// convertir le chiffre dec en chaine de caractère
		}
	}
	// convertir le tableau en chaine de caractère
	result := strings.Join(mots, " ")
	// enlever les espaces après traitement
	result = strings.Replace(result, "  ", " ", -1)

	return result
}

func remplacerContenuFichier(nomFichier string, contenu string) error {
	return ioutil.WriteFile(nomFichier, []byte(contenu), 0666)
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
