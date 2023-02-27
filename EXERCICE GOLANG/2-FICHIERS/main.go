package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	AjouterTxt()
}

func lireTxt() {
	file, err := ioutil.ReadFile("EXERCICE GOLANG/2-FICHIERS/test.txt") // lire le fichier text.txt
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file)) // conversion de byte en string
}

func AjouterTxt() {
	var crit string

	fmt.Print("choisissez un truc à écrire :")
	fmt.Scanln(&crit)
	file, err := os.OpenFile("EXERCICE GOLANG/2-FICHIERS/test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600){
		panic(err)
	}
	_, err = file.WriteString(crit) // écrire dans le fichier
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile("EXERCICE GOLANG/2-FICHIERS/test.txt") 
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(data))
}
