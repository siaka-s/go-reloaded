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

// convertir le mot en majuscule
func convertirUp(upper string) string {

	upper = strings.ToUpper(upper)

	return upper
}

// convertir le mot en minuscule
func convertirLow(lower string) string {

	lower = strings.ToLower(lower)

	return lower
}

func convertirCap(cap string) string {

	cap = strings.Title(cap)

	return cap
}

// Convertion du texte en tableau
func wordTabExe(sample string) []byte {

	// récuperation du contenu du filchier
	mots, err := ioutil.ReadFile(sample)

	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
	}
	return mots
}

func nombreDeMot(s []string) string {
	str := ""

	for i, tag := range s {
		if tag == "(cap," || tag == "(low," || tag == "(up," {
			s[i] = ""
			s[i+1] = ""
		} else if tag != "(up)" && tag != "(hex)" && tag != "(bin)" && tag != "(cap)" && tag != "(low)" && tag != "" {
			if i == 0 {
				str = str + tag
			} else {
				str = str + " " + tag
			}
		}
	}
	return str
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

			bin := convertirbin(mots[a])

			binf := strconv.FormatInt(bin, 10)

			mots[a] = binf

		case "(up)":
			mots[i] = ""

			a := i - 1

			up := convertirUp(mots[a])

			mots[a] = up
		case "(low)":
			mots[i] = ""

			a := i - 1

			low := convertirLow(mots[a])

			mots[a] = low
		case "(cap)":
			mots[i] = ""

			a := i - 1

			cap := convertirCap(mots[a])

			mots[a] = cap

			break
		}

	}
	// convertir le tableau en chaine de caractère avec un espace de séparation
	result := strings.Join(mots, " ")
	// enlever les espaces après traitement
	result = strings.Replace(result, "  ", " ", -1)

	return result
}

// Migration du contenu modifier dans le fichier affichage des modification
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
