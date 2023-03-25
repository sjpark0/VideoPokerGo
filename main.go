package main

import (
	"VideoPoker/VideoPoker"
)

func main() {
	//t := UNKNOWN
	/*game := VideoPoker.NewVideoPoker()
	checker := VideoPoker.NewCheckCARD()

	game.GenerateCARDForTest()
	game.PrintHand()
	game.ChangeHand()
	game.PrintHand()
	hand := game.GetHand()

	checker.Sorting(hand)

	checker.PrintHandCheck(hand)*/
	game := VideoPoker.NewVideoPoker()
	game.TotalProbability()

}
