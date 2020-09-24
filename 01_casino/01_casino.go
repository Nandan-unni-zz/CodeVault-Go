package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msg(m string) {
	fmt.Print("\n")
	for i := 0; i < len(m)+4; i++ {
		fmt.Print("_")
	}
	fmt.Printf("\n\n  %s  \n", m)
	for i := 0; i < len(m)+4; i++ {
		fmt.Print("_")
	}
}

func end(init int, now int) {
	if init == now {
		fmt.Print("\n\nYou have no gain and no loss.")
	} else if init > now {
		fmt.Printf("\n\nYou lost %d", init-now)
	} else {
		fmt.Printf("\n\nYou gained %d", now-init)
	}
}

func main() {
	var initPlr, amtPlr, betPlr, numPlr int
	var amtDlr, betDlr, numDlr, numLuck int
	fmt.Print("\nYour balance will be the double of what you bet. ")
	fmt.Print("\n\nEnter the bet amount : $ ")
	fmt.Scanf("%d", &betPlr)
	betDlr = 100
	if betPlr > 1000 {
		fmt.Print("\nMax amount is 1000")
		betPlr = 1000
	} else if betPlr < 500 {
		fmt.Print("\nMin amount is 500")
		betPlr = 500
	}

	amtPlr = 2 * betPlr
	initPlr = amtPlr
	amtDlr = 2 * betDlr
	fmt.Printf("\n\nPlayer Bet Amount : $ %d", betPlr)
	fmt.Printf("\nDealer Bet Amount : $ %d", betDlr)

	for true {
		fmt.Print("\n\nEnter your number (1-8) : ")
		fmt.Scanf("%d", &numPlr)
		if numPlr > 8 {
			fmt.Print("\nTop number is 8")
			numPlr = 8
		} else if numPlr < 1 {
			fmt.Print("\nLowest number is 1")
			numPlr = 1
		}

		fmt.Printf("\n\nPlayer's Number : %d", numPlr)

		for true {
			numDlr = rand.Intn(7) + 1
			if numDlr != numPlr {
				break
			}
		}

		fmt.Printf("\nDealer Number : %d", numDlr)

		fmt.Print("\n\nRolling . . .")
		fmt.Print(".")
		numLuck = rand.Intn(7) + 1
		time.Sleep(20)
		fmt.Printf("\n\nDrawn Number : %d", numLuck)
		if numPlr == numLuck {
			amtPlr += betDlr
			amtDlr -= betDlr
			msg("PLAYER WON !")
			if amtDlr <= 0 {
				fmt.Print("\nDealer lost all his money")
				end(initPlr, amtPlr)
				break
			}
		} else if numDlr == numLuck {
			amtPlr -= betPlr
			amtDlr += betDlr
			msg("DEALER WON !")
			if amtPlr <= 0 {
				fmt.Print("\nYour balance is 0.")
				end(initPlr, 0)
				break
			}
		} else {
			msg("DRAW !")
		}

		fmt.Printf("\n\n\nPlayer's Balance : $ %d", amtPlr)
		fmt.Printf("\nDealer's Balance : $ %d", amtDlr)
		fmt.Print("\n\n1. Continue \t 2. Quit \t : ")
		var cont int
		fmt.Scanf("%d", &cont)
		fmt.Print("___________________________________________________")
		if cont == 2 {
			end(initPlr, amtPlr)
			break
		}
	}
}
