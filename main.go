package main

import (
	"HOMEWORK-1-EZGIUSTUNEL/booksAction"
	"os"
)

func main() {
	//book list
	books := []string{"Simyaci",
		"Bab-i Esrar",
		"Nar Ağaci",
		"Fareler ve İnsanlar",
		"Kürk Mantolu Madonna",
		"Hayvan Çiftliği",
		"Şeker Portakali",
		"Uçurtma Avcisi",
		"Suç ve Ceza",
		"Serenad",
		"Yeraltindan Notlar",
		"Toprak Ana",
		"Fatih Harbiye",
		"Saatleri Ayarlama Enstitüsü",
		"Acimak",
		"Ateşten Gömlek",
		"Çocukluğum",
		"Aşk",
		"Kuyucakli Yusuf",
		"Arkadaş",
	}

	inputList := os.Args
	booksAction.PerformAction(inputList, books)
}
