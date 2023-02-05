package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Fonction hecadecimal en entier
func convertirhex(hexnum string) {

	// la fonction parseInt() nous permet de réaliser la conversion
	decimal_num, err := strconv.ParseInt(hexnum, 16, 64)

	// Gestion de l'erreur du nombre recu en paramètre
	if err != nil {
		panic(err)
	}

	// Affichage du résultat
	fmt.Println(decimal_num)

}

// stocage de la phrase dans un tableau et execution de cas

func wordTabExe(sample string) []byte {

	// récuperation du contenu du filchier
	phrase, err := ioutil.ReadFile(sample)

	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
	}
	return phrase
}

func traitement(phrase string) {

	mots := strings.Fields(phrase)

	fmt.Println(mots[1])

	for _, v := range mots {
		if v == "(hex)" {
			convertirhex(mots[0])
		}
	}

}

func main() {

	sample := os.Args[1]

	// result := os.Args[2]

	// strockage du contenu du fichier dans une variable
	phrase := string(wordTabExe(sample))

	// exécution de la fonction de traitement
	traitement(phrase)

	// fmt.Printf("le texte de debut est %v \n le texte modifier est %v", sample, result)

}
