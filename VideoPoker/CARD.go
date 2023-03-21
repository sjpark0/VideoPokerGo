package VideoPoker

import "fmt"

type TYPE int

const (
	UNKNOWN = iota
	CARD_SPADE
	CARD_CLOVER
	CARD_HEART
	CARD_DIAMOND
)

type CARD struct {
	m_number int
	m_type   TYPE
}

func NewCARD() *CARD {
	return &CARD{-1, UNKNOWN}
}

/*func NewCARD(number int, t TYPE) *CARD {
	return &CARD{number, t}
}*/

func (c *CARD) IsSameNumber(c2 CARD) bool {
	if c.m_number == c2.m_number {
		return true
	} else {
		return false
	}
}

func (c *CARD) IsSameType(c2 CARD) bool {
	if c.m_type == c2.m_type {
		return true
	} else {
		return false
	}
}

func (c *CARD) IsSameCard(c2 CARD) bool {
	if c.IsSameNumber(c2) && c.IsSameType(c2) {
		return true
	} else {
		return false
	}
}
func (c *CARD) AssignFrom(c2 CARD) {
	c.m_number = c2.m_number
	c.m_type = c2.m_type
}
func (c *CARD) Print() {
	switch c.m_number {
	case 1:
		fmt.Print("A,")
	case 11:
		fmt.Print("J,")
	case 12:
		fmt.Print("Q,")
	case 13:
		fmt.Print("K,")
	default:
		fmt.Printf("%d,", c.m_number)
	}

	switch c.m_type {
	case CARD_SPADE:
		fmt.Println("Spade")
	case CARD_CLOVER:
		fmt.Println("Clover")
	case CARD_DIAMOND:
		fmt.Println("Diamond")
	case CARD_HEART:
		fmt.Println("Heart")
	default:
		fmt.Println("Unknown")
	}
}
