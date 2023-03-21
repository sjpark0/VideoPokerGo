package VideoPoker

import "fmt"

const NUM_TOTAL int = 52
const NUM_HAND int = 5

type CheckCARD struct {
	m_creditJackBetter    int
	m_creditTwoPair       int
	m_creditTriple        int
	m_creditStraight      int
	m_creditFlush         int
	m_creditFullHouse     int
	m_creditFourCARD      int
	m_creditStraightFlush int
	m_creditRoyalFlush    int
}

func NewCheckCARD() *CheckCARD {
	return &CheckCARD{1, 2, 3, 4, 6, 9, 25, 50, 800}
}

func (checker *CheckCARD) IsOnePair(card []CARD) bool {
	if card[0].IsSameNumber(card[1]) ||
		card[1].IsSameNumber(card[2]) ||
		card[2].IsSameNumber(card[3]) ||
		card[3].IsSameNumber(card[4]) {
		return true
	} else {
		return false
	}
}
func (checker *CheckCARD) IsJackBetter(card []CARD) bool {
	if (card[0].IsSameNumber(card[1]) && (card[0].m_number >= 11 || card[0].m_number == 1)) ||
		(card[1].IsSameNumber(card[2]) && (card[1].m_number >= 11 || card[1].m_number == 1)) ||
		(card[2].IsSameNumber(card[3]) && (card[2].m_number >= 11 || card[2].m_number == 1)) ||
		(card[3].IsSameNumber(card[4]) && (card[3].m_number >= 11 || card[3].m_number == 1)) {
		return true
	} else {
		return false
	}
}
func (checker *CheckCARD) IsTwoPair(card []CARD) bool {
	if (card[0].IsSameNumber(card[1]) && card[2].IsSameNumber(card[3])) ||
		(card[0].IsSameNumber(card[1]) && card[3].IsSameNumber(card[4])) ||
		(card[1].IsSameNumber(card[2]) && card[3].IsSameNumber(card[4])) {
		return true
	} else {
		return false
	}
}
func (checker *CheckCARD) IsTriple(card []CARD) bool {
	if card[0].IsSameNumber(card[2]) ||
		card[1].IsSameNumber(card[3]) ||
		card[2].IsSameNumber(card[4]) {
		return true
	} else {
		return false
	}
}
func (checker *CheckCARD) IsFlush(card []CARD) bool {
	if card[0].IsSameType(card[1]) && card[0].IsSameType(card[2]) && card[0].IsSameType(card[3]) && card[0].IsSameType(card[4]) {
		return true
	} else {
		return false
	}
}
func (checker *CheckCARD) IsStraight(card []CARD) bool {
	if card[0].m_number == card[1].m_number-1 &&
		card[1].m_number == card[2].m_number-1 &&
		card[2].m_number == card[3].m_number-1 &&
		card[3].m_number == card[4].m_number-1 {
		return true
	} else if card[0].m_number == 1 &&
		card[1].m_number == 10 &&
		card[2].m_number == 11 &&
		card[3].m_number == 12 &&
		card[4].m_number == 13 {
		return true
	} else {
		return false
	}
}
func (checker *CheckCARD) IsFullHouse(card []CARD) bool {
	if (card[0].IsSameNumber(card[2]) && card[3].IsSameNumber(card[4])) ||
		(card[0].IsSameNumber(card[1]) && card[2].IsSameNumber(card[4])) {
		return true
	} else {
		return false
	}
}
func (checker *CheckCARD) IsFourCARD(card []CARD) bool {
	if card[0].IsSameNumber(card[3]) ||
		card[1].IsSameNumber(card[4]) {
		return true
	} else {
		return false
	}
}
func (checker *CheckCARD) IsStraightFlush(card []CARD) bool {
	if checker.IsFlush(card) && checker.IsStraight(card) {
		return true
	} else {
		return false
	}
}
func (checker *CheckCARD) IsRoyalFlush(card []CARD) bool {
	if checker.IsFlush(card) {
		if card[0].m_number == 1 &&
			card[1].m_number == 10 &&
			card[2].m_number == 11 &&
			card[3].m_number == 12 &&
			card[4].m_number == 13 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (checker *CheckCARD) ReturnCredit(card []CARD, bet int) int {
	if checker.IsRoyalFlush(card) {
		return bet * checker.m_creditRoyalFlush
	} else if checker.IsStraightFlush(card) {
		return bet * checker.m_creditStraightFlush
	} else if checker.IsFourCARD(card) {
		return bet * checker.m_creditFourCARD
	} else if checker.IsFullHouse(card) {
		return bet * checker.m_creditFullHouse
	} else if checker.IsFlush(card) {
		return bet * checker.m_creditFlush
	} else if checker.IsStraight(card) {
		return bet * checker.m_creditStraight
	} else if checker.IsTriple(card) {
		return bet * checker.m_creditTriple
	} else if checker.IsTwoPair(card) {
		return bet * checker.m_creditTwoPair
	} else if checker.IsJackBetter(card) {
		return bet * checker.m_creditJackBetter
	} else {
		return 0
	}
}

func (checker *CheckCARD) PrintHandCheck(card []CARD) {
	if checker.IsRoyalFlush(card) {
		fmt.Println("Royal Flush")
	} else if checker.IsStraightFlush(card) {
		fmt.Println("Straight Flush")
	} else if checker.IsFourCARD(card) {
		fmt.Println("Four CARD")
	} else if checker.IsFullHouse(card) {
		fmt.Println("Full House")
	} else if checker.IsFlush(card) {
		fmt.Println("Flush")
	} else if checker.IsStraight(card) {
		fmt.Println("Straight")
	} else if checker.IsTriple(card) {
		fmt.Println("Triple")
	} else if checker.IsTwoPair(card) {
		fmt.Println("Two Pair")
	} else if checker.IsJackBetter(card) {
		fmt.Println("Jack Better")
	} else if checker.IsOnePair(card) {
		fmt.Println("One Pair")
	} else {
		fmt.Println("Nothing")
	}
} // test

func (checker *CheckCARD) Sorting(card []CARD) {
	var tempCARD CARD

	for i := 0; i < NUM_HAND; i++ {
		for j := i + 1; j < NUM_HAND; j++ {
			if card[i].m_number > card[j].m_number {
				tempCARD.AssignFrom(card[i])
				card[i].AssignFrom(card[j])
				card[j].AssignFrom(tempCARD)
			}
		}
	}
} // test
