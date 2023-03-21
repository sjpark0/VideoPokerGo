package main

import (
	"./VideoPoker"
)

func main() {
	//t := UNKNOWN
	game := VideoPoker.NewVideoPoker()
	checker := VideoPoker.NewCheckCARD()
	//var hand []VideoPoker.CARD

	game.GenerateCARD()
	game.PrintHand()
	hand := game.GetHand()

	checker.Sorting(hand)
	game.PrintCard(hand, 5)

	checker.PrintHandCheck(hand)

}
