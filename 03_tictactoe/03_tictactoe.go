package main

import (
	"fmt"
	"os"
	"os/exec"
)

func logo(pos []string) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Print("\n\t\tTic Tac Toe")
	fmt.Printf(
		`
	 _______________________
        |       |       |       |
        |   %s   |   %s   |   %s   |
        |_______|_______|_______|
        |       |       |       |
        |   %s   |   %s   |   %s   |
        |_______|_______|_______|
        |       |       |       |
        |   %s   |   %s   |   %s   |
        |_______|_______|_______|
		`, pos[0], pos[1], pos[2], pos[3], pos[4], pos[5], pos[6], pos[7], pos[8])
}

func main() {
	fmt.Print("NQ")
	var remote = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	logo(remote)
	var plr1, plr2 string
	fmt.Scan(&plr1)
	if plr1 == "X" || plr1 == "x" {
		plr1 = "X"
		plr2 = "O"
	} else {
		plr1 = "O"
		plr2 = "X"
	}
	logo(remote)

	processor := func(plr string, plrno int) bool {
		var chances = []string{"048", "246", "012", "345", "678", "036", "147", "258"}
		var pos string
		fmt.Printf("\n  Player %d can choose his position : ", plrno)
		fmt.Scanf(pos)
		for i, num := range remote {
			if pos == num {
				remote = copy(remote[:i], "X", remote[i:])
			}
		}
		return true
	}

	var win bool
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			win = processor(plr1, 1)
			if win {
				return
			}
		} else {
			win = processor(plr2, 2)
			if win {
				return
			}
		}
	}

}
