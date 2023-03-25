package VideoPoker

import (
	"fmt"
	"math/rand"
)

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

func (v *VideoPoker) ConTypeToVal(value TYPE) int {
	switch value {
	case CARD_SPADE:
		return 0
	case CARD_CLOVER:
		return 1
	case CARD_HEART:
		return 2
	case CARD_DIAMOND:
		return 3
	default:
		return -1
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

func (v *VideoPoker) ChangeOneCARD(card []CARD, handIdx int, remainIdx int) {
	card[handIdx].AssignFrom(v.m_pRemain[remainIdx])
}
func (v *VideoPoker) ComputeAvgCreditForCardChange(card []CARD, handIdx []int, numChangeCard int) float64 {
	credit, numComputeCredit := v.ComputeTotalCreditForCardChange(card, handIdx, 0, numChangeCard)

	return float64(credit) / float64(numComputeCredit)
}
func (v *VideoPoker) ComputeTotalCreditForCardChange(card []CARD, handIdx []int, startRemainIdx int, numChangeCard int) (int, int) {
	numComputeCredit := 0
	if numChangeCard == 0 {
		numComputeCredit = 1
		return v.ComputeCredit(card), numComputeCredit
	}

	credit := 0

	var tempCARD []CARD = make([]CARD, NUM_HAND)
	for i := 0; i < NUM_HAND; i++ {
		tempCARD[i].AssignFrom(card[i])
	}

	for i := startRemainIdx; i < NUM_TOTAL-NUM_HAND; i++ {
		v.ChangeOneCARD(tempCARD, handIdx[0], i)

		tempCredit, tempNumComputeCredit := v.ComputeTotalCreditForCardChange(tempCARD, handIdx[1:], i+1, numChangeCard-1)
		credit += tempCredit
		numComputeCredit += tempNumComputeCredit
	}

	return credit, numComputeCredit
}

func (v *VideoPoker) ComputeCredit(card []CARD) int {
	tempCARD := make([]CARD, NUM_HAND)
	for i := 0; i < NUM_HAND; i++ {
		tempCARD[i].AssignFrom(card[i])
	}
	v.m_pChecker.Sorting(tempCARD)

	return v.m_pChecker.ReturnCredit(tempCARD, 1)
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
func (v *VideoPoker) GenerateCARDWithInput(cards []CARD) {
	var isHand []bool = make([]bool, NUM_TOTAL)

	for i := 0; i < NUM_HAND; i++ {
		v.m_pHand[i].m_number = cards[i].m_number
		v.m_pHand[i].m_type = cards[i].m_type
		isHand[v.ConTypeToVal(cards[i].m_type)+(cards[i].m_number-1)*4] = true
	}

	cntRemain := 0
	for i := 0; i < NUM_TOTAL; i++ {
		if !isHand[i] {
			v.m_pRemain[cntRemain].m_number = i/4 + 1
			v.m_pRemain[cntRemain].m_type = v.ConValToType(i % 4)
			cntRemain++
		}
	}
}
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
func (v *VideoPoker) GenerateCARDForTest() {
	var isHand []bool = make([]bool, NUM_TOTAL)
	handIdx := [NUM_HAND]int{26, 44, 8, 19, 14}

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
}

func (v *VideoPoker) ComputeOptimumChange(handIdx []int) (int, float64) {
	tempHandIdx := make([]int, NUM_HAND)
	copy(tempHandIdx, handIdx)

	numChangeCard := 0
	credit := 0.0

	optimumCredit := v.ComputeAvgCreditForCardChange(v.m_pHand, handIdx, 0)
	//fmt.Println(numChangeCard, optimumCredit)

	for i := 0; i < NUM_HAND; i++ {
		tempHandIdx[0] = i
		credit = v.ComputeAvgCreditForCardChange(v.m_pHand, tempHandIdx, 1)
		if credit > optimumCredit {
			optimumCredit = credit
			numChangeCard = 1
			for idx := 0; idx < numChangeCard; idx++ {
				handIdx[idx] = tempHandIdx[idx]
			}
		}
	}
	//fmt.Println(numChangeCard, optimumCredit)

	for i := 0; i < NUM_HAND; i++ {
		for j := i + 1; j < NUM_HAND; j++ {
			tempHandIdx[0] = i
			tempHandIdx[1] = j
			credit = v.ComputeAvgCreditForCardChange(v.m_pHand, tempHandIdx, 2)

			if credit > optimumCredit {
				optimumCredit = credit
				numChangeCard = 2
				for idx := 0; idx < numChangeCard; idx++ {
					handIdx[idx] = tempHandIdx[idx]
				}
			}

		}
	}
	//fmt.Println(numChangeCard, optimumCredit)

	for i := 0; i < NUM_HAND; i++ {
		for j := i + 1; j < NUM_HAND; j++ {
			for k := j + 1; k < NUM_HAND; k++ {
				tempHandIdx[0] = i
				tempHandIdx[1] = j
				tempHandIdx[2] = k
				credit = v.ComputeAvgCreditForCardChange(v.m_pHand, tempHandIdx, 3)

				if credit > optimumCredit {
					optimumCredit = credit
					numChangeCard = 3
					for idx := 0; idx < numChangeCard; idx++ {
						handIdx[idx] = tempHandIdx[idx]
					}
				}
			}

		}
	}
	//fmt.Println(numChangeCard, optimumCredit)

	for i := 0; i < NUM_HAND; i++ {
		for j := i + 1; j < NUM_HAND; j++ {
			for k := j + 1; k < NUM_HAND; k++ {
				for l := k + 1; l < NUM_HAND; l++ {
					tempHandIdx[0] = i
					tempHandIdx[1] = j
					tempHandIdx[2] = k
					tempHandIdx[3] = l
					credit = v.ComputeAvgCreditForCardChange(v.m_pHand, tempHandIdx, 4)

					if credit > optimumCredit {
						optimumCredit = credit
						numChangeCard = 4
						for idx := 0; idx < numChangeCard; idx++ {
							handIdx[idx] = tempHandIdx[idx]
						}
					}
				}
			}

		}
	}
	//fmt.Println(numChangeCard, optimumCredit)

	for i := 0; i < NUM_HAND; i++ {
		for j := i + 1; j < NUM_HAND; j++ {
			for k := j + 1; k < NUM_HAND; k++ {
				for l := k + 1; l < NUM_HAND; l++ {
					for m := l + 1; m < NUM_HAND; m++ {
						tempHandIdx[0] = i
						tempHandIdx[1] = j
						tempHandIdx[2] = k
						tempHandIdx[3] = l
						tempHandIdx[4] = m
						credit = v.ComputeAvgCreditForCardChange(v.m_pHand, tempHandIdx, 5)

						if credit > optimumCredit {
							optimumCredit = credit
							numChangeCard = 5
							for idx := 0; idx < numChangeCard; idx++ {
								handIdx[idx] = tempHandIdx[idx]
							}
						}
					}
				}
			}

		}
	}
	//fmt.Println(numChangeCard, optimumCredit)

	return numChangeCard, optimumCredit
}
func (v *VideoPoker) ChangeHandIdx() ([]int, int) {
	handIdx := make([]int, NUM_HAND)

	numChangeCard, _ := v.ComputeOptimumChange(handIdx)

	return handIdx, numChangeCard
}
func (v *VideoPoker) ReplaceChangeHandIdx(handIdx []int, numChangeCard int) {
	changeIdx := make([]int, NUM_HAND)

	v.GenerateNCard(changeIdx, NUM_TOTAL-NUM_HAND, numChangeCard)

	for i := 0; i < numChangeCard-1; i++ {
		v.m_pHand[handIdx[i]].AssignFrom(v.m_pRemain[changeIdx[i]])
	}
}
func (v *VideoPoker) ChangeHand() {
	handIdx := make([]int, NUM_HAND)
	for i := 0; i < NUM_HAND; i++ {
		handIdx[i] = -1
	}
	numChangeCard, _ := v.ComputeOptimumChange(handIdx)
	//print(optimumCredit)
	fmt.Println(numChangeCard, handIdx)

	changeIdx := make([]int, NUM_HAND)
	for i := 0; i < NUM_HAND; i++ {
		changeIdx[i] = -1
	}
	v.GenerateNCard(changeIdx, NUM_TOTAL-NUM_HAND, numChangeCard)

	for i := 0; i < numChangeCard-1; i++ {
		v.m_pHand[handIdx[i]].AssignFrom(v.m_pRemain[changeIdx[i]])
	}
}

func (v *VideoPoker) TotalProbability() {
	isHand := make([]bool, NUM_TOTAL)
	handIdx := make([]int, NUM_HAND)

	cntRemain := 0
	numGame := 0
	totalCredit := 0.0

	for c1 := 0; c1 < NUM_TOTAL; c1++ {
		for c2 := c1 + 1; c2 < NUM_TOTAL; c2++ {
			for c3 := c2 + 1; c3 < NUM_TOTAL; c3++ {
				for c4 := c3 + 1; c4 < NUM_TOTAL; c4++ {
					for c5 := c4 + 1; c5 < NUM_TOTAL; c5++ {
						for i := 0; i < NUM_TOTAL; i++ {
							isHand[i] = false
						}

						v.m_pHand[0].m_number = c1/4 + 1
						v.m_pHand[0].m_type = v.ConValToType(c1 % 4)
						v.m_pHand[1].m_number = c2/4 + 1
						v.m_pHand[1].m_type = v.ConValToType(c2 % 4)
						v.m_pHand[2].m_number = c3/4 + 1
						v.m_pHand[2].m_type = v.ConValToType(c3 % 4)
						v.m_pHand[3].m_number = c4/4 + 1
						v.m_pHand[3].m_type = v.ConValToType(c4 % 4)
						v.m_pHand[4].m_number = c5/4 + 1
						v.m_pHand[4].m_type = v.ConValToType(c5 % 4)

						isHand[c1] = true
						isHand[c2] = true
						isHand[c3] = true
						isHand[c4] = true
						isHand[c5] = true

						cntRemain = 0
						for i := 0; i < NUM_TOTAL; i++ {
							if !isHand[i] {
								v.m_pRemain[cntRemain].m_number = i/4 + 1
								v.m_pRemain[cntRemain].m_type = v.ConValToType(i % 4)
								cntRemain += 1
							}
						}

						_, optimumCredit := v.ComputeOptimumChange(handIdx)
						totalCredit += optimumCredit
						numGame++
						//print("Probability", numGame, totalCredit, Float(totalCredit) / Float(numGame))
						fmt.Println("Probability ", numGame, " ", optimumCredit, " ", totalCredit)
						/*if(numGame % 10 == 0) {
							println("Probability" + numGame.toString() + "," + totalCredit.toString() + "," + ((totalCredit.toFloat()) / (numGame.toFloat())).toString())
						}*/
					}
				}
			}
		}
	}
	fmt.Println("Total Probability", numGame, ",", totalCredit, ",", float64(totalCredit)/float64(numGame))
}
