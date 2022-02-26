package booksAction

import (
	"fmt"
	"strings"
)

const search = "search"
const list = "list"

//Method for perform action according to user input on book list.
func PerformAction(inputList, itemList []string) {
	if len(inputList) == 1 {
		PrintMessagesToConsole()
		return
	}

	firstInput := inputList[1]

	if firstInput == list {
		ListBooks(itemList)
		return
	} else if firstInput == search {
		if len(inputList) == 2 {
			fmt.Printf("\nPlease enter the book name you want to search!\n\n")
			return
		}
		SearchBooks(itemList, inputList[2:])
	} else {
		PrintMessagesToConsole()
	}
}

//Method for print the books in the book list.
func ListBooks(books []string) {
	fmt.Println()
	for _, value := range books {
		fmt.Println(value)
	}
	fmt.Println()
}

//Method for search books according to given input and prints the matched books
func SearchBooks(books, keywordList []string) []string {
	var searchedBooks []string

	for _, keyword := range keywordList {
		for _, book := range books {
			if Contains(searchedBooks, strings.ToLower(book)) {
				continue
			}
			if strings.Contains(strings.ToLower(book), strings.ToLower(keyword)) {
				searchedBooks = append(searchedBooks, book)
			}
		}
	}

	PrintList(searchedBooks)
	return searchedBooks
}

//Method for print a list.
func PrintList(list []string) {
	fmt.Println()
	for _, value := range list {
		fmt.Println(value)
	}
	fmt.Println()
}

//Method for print messages to user in the console.
func PrintMessagesToConsole() {
	fmt.Printf("\n--Invalid Input--\n\n")
	fmt.Println("You can use the methods below to make some actions on book list")
	fmt.Println("list: Lists the books")
	fmt.Printf("search \"bookname\": searches the bookname given in the book list\n\n")
}

//Method for check if array contains the item or not.
func Contains(strList []string, str string) bool {
	for _, value := range strList {
		if value == str {
			return true
		}
	}
	return false
}
