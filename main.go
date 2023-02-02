package main

import (
	"fmt"
	"strconv"
)

// Fonction hecadecimal en entier
func hexconvert() {

	// initialisation d'une variable  pour effectuer le test
	hexadecimal_num := "23a"

	// la fonction parseInt() nous permet de réaliser la conversion
	decimal_num, err := strconv.ParseInt(hexadecimal_num, 16, 64)

	// Gestion de l'erreur
	if err != nil {
		panic(err)
	}

	// Affichage du résultat
	fmt.Println(decimal_num)

}
