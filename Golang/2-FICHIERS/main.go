package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	for {

		fmt.Println(" _________________________________________")
		fmt.Println("|                 MENU:                   |")
		fmt.Println("|         1. Lire le fichier              |")
		fmt.Println("|         2. Ajouter du texte             |")
		fmt.Println("|         3. Supprimer le texte           |")
		fmt.Println("|         4. Remplacer le texte           |")
		fmt.Println("|         5. Quitter                      |")
		fmt.Println("|_________________________________________|")

		var choice int
		fmt.Print("Choisissez une option :")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			lireTxt()
		case 2:
			AjouterTxt()
		case 3:
			SupprimerTxt()
		case 4:
			RemplacerTxt()
		case 5:
			fmt.Println("Au revoir")
			return
		default:
			fmt.Println("Mauvais choix, réesayez avec un nombre entre 1 et 5.")
		}
	}
}

func lireTxt() {
	file, err := ioutil.ReadFile("EXERCICE GOLANG/2-FICHIERS/test.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}

func AjouterTxt() {
	var ecrit string

	fmt.Print("Choisissez quelque chose à dire :")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	ecrit = scanner.Text()
	file, err := os.OpenFile("EXERCICE GOLANG/2-FICHIERS/test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString(ecrit)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile("EXERCICE GOLANG/2-FICHIERS/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(data))
}

func SupprimerTxt() {
	monfichier := "EXERCICE GOLANG/2-FICHIERS/test.txt"

	err := ioutil.WriteFile(monfichier, []byte(""), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Le fichier %s a été vidé\n", monfichier)
}

func RemplacerTxt() {
	monfichier := "EXERCICE GOLANG/2-FICHIERS/test.txt"

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entrez le nouveau contenu du fichier : ")
	nouveauTxt, _ := reader.ReadString('\n')

	err := ioutil.WriteFile(monfichier, []byte(nouveauTxt), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Le fichier %s a été mis à jour\n", monfichier)
}
