package main

import (
	"fmt"
	"os"
	"os/exec"
)

type libBook struct {
	slno   int
	name   string
	author string
}

type issBook struct {
	slno   int
	name   string
	author string
	issuer string
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func pause() {
	pause := ""
	fmt.Printf("\nContinue...")
	fmt.Scanf("%s", &pause)
}

func main() {
	var libBooks [20]libBook
	var issBooks [20]issBook
	var choice int
	noOfLibBooks := 1
	noOfIssBooks := 1
	for true {
		clear()
		fmt.Printf("\n\t\tLIBRARY MANAGEMENT SYSTEM\n")
		fmt.Printf("\n\n\t  TASKS\n\t1. Add Book\n\t2. Edit Book\n\t3. Issue Book\n\t4. Return Book")
		fmt.Printf("\n\n\t  REFERENCE\n\t5. View Available Books\n\t6. View Issued Book\n\t7. Exit")
		fmt.Printf("\n\n\tChoice : ")
		fmt.Scanf("%d", &choice)
		if choice == 1 {
			clear()
			fmt.Printf("\n\t\tAdd Book\n\n")
			libBooks[noOfLibBooks].slno = noOfLibBooks
			fmt.Printf("Enter the name of book : ")
			fmt.Scanf("%s", &libBooks[noOfLibBooks].name)
			fmt.Printf("Enter the name of author : ")
			fmt.Scanf("%s", &libBooks[noOfLibBooks].author)
			noOfLibBooks++
		} else if choice == 2 {
			clear()
			fmt.Printf("\n\t\tEdit Book\n\n")
			findSlno := -1
			fmt.Printf("Enter the slno : ")
			fmt.Scanf("%d", &findSlno)
			if libBooks[findSlno].name != "" {
				fmt.Printf("\nBook %d\n", libBooks[findSlno].slno)
				fmt.Printf("%s by %s\n\n", libBooks[findSlno].name, libBooks[findSlno].author)
				var newName, newAuthor string
				fmt.Printf("Enter new name of book : ")
				fmt.Scanf("%s", &newName)
				fmt.Printf("Enter new name of author : ")
				fmt.Scanf("%s", &newAuthor)
				libBooks[findSlno].name = newName
				libBooks[findSlno].author = newAuthor
				fmt.Printf("\nUpdated Details\nBook %d\n", libBooks[findSlno].slno)
				fmt.Printf("%s by %s\n\n", libBooks[findSlno].name, libBooks[findSlno].author)
			} else {
				fmt.Printf("\nBook not found !")
			}
			pause()
		} else if choice == 3 {
			clear()
			fmt.Printf("\n\t\tIssue Book\n\n")
			findSlno := -1
			fmt.Printf("Enter the slno : ")
			fmt.Scanf("%d", &findSlno)
			if libBooks[findSlno].name != "" {
				fmt.Printf("\nBook %d\n", libBooks[findSlno].slno)
				fmt.Printf("%s by %s\n\n", libBooks[findSlno].name, libBooks[findSlno].author)
				fmt.Printf("\nEnter the name of the issuer : ")
				fmt.Scanf("%s", &issBooks[noOfIssBooks].issuer)
				issBooks[noOfIssBooks].slno = libBooks[findSlno].slno
				issBooks[noOfIssBooks].name = libBooks[findSlno].name
				issBooks[noOfIssBooks].author = libBooks[findSlno].author
				noOfIssBooks++
				libBooks[findSlno].slno = 0
				libBooks[findSlno].name = ""
				libBooks[findSlno].author = ""
			} else {
				fmt.Printf("\nBook not found !")
			}
		} else if choice == 4 {
			clear()
			fmt.Printf("\n\t\tReturn Book\n\n")
			findSlno := -1
			fmt.Printf("Enter the slno : ")
			fmt.Scanf("%d", &findSlno)
			for i := 0; i < noOfIssBooks; i++ {
				if issBooks[i].slno == findSlno {
					findSlno = i
				}
			}
			if issBooks[findSlno].name != "" {
				fmt.Printf("Book %d\n", issBooks[findSlno].slno)
				fmt.Printf("%s by %s\n", issBooks[findSlno].name, libBooks[findSlno].author)
				fmt.Printf("Issued to %s\n\n", issBooks[findSlno].issuer)
				pause()
				retSlNo := issBooks[findSlno].slno
				libBooks[retSlNo].slno = issBooks[findSlno].slno
				libBooks[retSlNo].name = issBooks[findSlno].name
				libBooks[retSlNo].author = issBooks[findSlno].author
				issBooks[findSlno].slno = 0
				issBooks[findSlno].name = ""
				issBooks[findSlno].author = ""
				issBooks[findSlno].issuer = ""
			}
		} else if choice == 5 {
			clear()
			fmt.Printf("\n\t\tAvailable Books\n\n")
			for i := 1; i < noOfLibBooks; i++ {
				if libBooks[i].name != "" {
					fmt.Printf("Book %d\n", libBooks[i].slno)
					fmt.Printf("%s by %s\n\n", libBooks[i].name, libBooks[i].author)
				}
			}
			pause()
		} else if choice == 6 {
			clear()
			fmt.Printf("\n\t\tIssued Books\n\n")
			for i := 1; i < noOfIssBooks; i++ {
				if issBooks[i].name != "" {
					fmt.Printf("Book %d\n", issBooks[i].slno)
					fmt.Printf("%s by %s\n", issBooks[i].name, libBooks[i].author)
					fmt.Printf("Issued to %s\n\n", issBooks[i].issuer)
				}
			}
			pause()
		} else {
			break
		}
		choice = 0
	}
}
