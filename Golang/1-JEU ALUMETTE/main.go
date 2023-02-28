package main

import (
	"fmt"
)

func main() {
	var nombredallumettes int

	for {
		fmt.Print("Combien d'allumettes voulez-vous (4 min) :")
		fmt.Scanln(&nombredallumettes)
		if nombredallumettes >= 4 {
			break
		}
		fmt.Println("Le nombre d'allumettes doit être au moins 4.")
	}

	var joueur int = 1
	var nombrerestant int = nombredallumettes

	for nombrerestant > 0 {
		fmt.Println("Il reste", nombrerestant, "allumettes.")

		var choix int
		for {
			fmt.Print("Joueur ", joueur, ", combien d'allumettes voulez-vous prendre (1 à 3) : ")
			fmt.Scanln(&choix)
			if choix >= 1 && choix <= 3 && choix <= nombrerestant {
				break
			}
			fmt.Println("Le choix doit être entre 1 et 3 et ne doit pas dépasser le nombre d'allumettes restantes.")
		}

		nombrerestant -= choix

		if joueur == 1 {
			joueur = 2
		} else {
			joueur = 1
		}
	}

	fmt.Println("Le joueur", joueur, "a perdu !")
}
