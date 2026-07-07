package main

import "fmt"

func debugPrintTiles(tiles []*MapTile, gameData GameData) {

	fmt.Println("===============================")

	for i, t := range tiles {
		if i%gameData.ScreenWidth == 0 {
			fmt.Println()
		}
		tileChar := "."
		if t.Blocked {
			tileChar = "#"
		}
		fmt.Printf("%s", tileChar)

	}

	fmt.Println("\n\n===============================")
}
