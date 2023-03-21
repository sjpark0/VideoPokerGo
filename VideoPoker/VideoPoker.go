package VideoPoker

//import "fmt"

import "math/rand"

type VideoPoker struct {
	m_pChecker *CheckCARD
	m_pHand    []CARD
	m_pAvail   []CARD
	m_pRemain  []CARD
}

func NewVideoPoker() *VideoPoker {
	v := new(VideoPoker)
	v.m_pChecker = NewCheckCARD()
	v.m_pHand = make([]CARD, NUM_HAND)
	v.m_pAvail = make([]CARD, NUM_HAND)
	v.m_pRemain = make([]CARD, NUM_TOTAL-NUM_HAND)

	return v
}

func (v *VideoPoker) ConValToType(val int) TYPE {
	switch val {
	case 0:
		return CARD_SPADE
	case 1:
		return CARD_CLOVER
	case 2:
		return CARD_HEART
	case 3:
		return CARD_DIAMOND
	default:
		return UNKNOWN
	}

}
func (v *VideoPoker) GenerateNCard(cardIdx []int, numTotalCard int, numGenCard int) {
	for i := 0; i < numGenCard; i++ {
		cardIdx[i] = rand.Intn(numTotalCard)
		for j := 0; j < i; j++ {
			if cardIdx[i] == cardIdx[j] {
				i--
				break
			}
		}
	}
}

func (v *VideoPoker) PrintCard(card []CARD, numCard int) {
	for i := 0; i < numCard; i++ {
		card[i].Print()
	}
}
func (v *VideoPoker) GetHand() []CARD {
	return v.m_pHand
} // test
func (v *VideoPoker) PrintHand() {
	v.PrintCard(v.m_pHand, NUM_HAND)
} // test

func (v *VideoPoker) GenerateCARD() {
	var isHand []bool = make([]bool, NUM_TOTAL)
	var handIdx []int = make([]int, NUM_TOTAL)

	v.GenerateNCard(handIdx, NUM_TOTAL, NUM_HAND)

	for i := 0; i < NUM_HAND; i++ {
		v.m_pHand[i].m_number = handIdx[i]/4 + 1
		v.m_pHand[i].m_type = v.ConValToType(handIdx[i] % 4)
		isHand[handIdx[i]] = true
	}

	cntRemain := 0
	for i := 0; i < NUM_TOTAL; i++ {
		if !isHand[i] {
			v.m_pRemain[cntRemain].m_number = i/4 + 1
			v.m_pRemain[cntRemain].m_type = v.ConValToType(i % 4)
			cntRemain++
		}
	}
} // test
