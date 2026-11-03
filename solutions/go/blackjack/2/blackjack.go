package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	value := 0
	switch card {
    case "two": value = 2
    case "three": value = 3
    case "four": value = 4
    case "five": value = 5
    case "six": value = 6
    case "seven": value = 7
    case "eight": value = 8
    case "nine": value = 9
    case "ten": value = 10
    case "jack": value = 10
    case "queen": value = 10
    case "king": value = 10
    case "ace": value = 11
	}
    return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var option string
    sumCard := ParseCard(card1) + ParseCard(card2)
    switch {
    case 21 < sumCard: option = "P"
    case 21 == sumCard: 
        if ParseCard(dealerCard) >= 10 {
            option = "S"
        } else {
            option = "W"
        }
    case 17 <= sumCard && 20 >= sumCard: option = "S"
    case 12 <= sumCard && 16 >= sumCard:
        if ParseCard(dealerCard) >= 7 { 
            option = "H" 
        } else { 
            option = "S" 
        }
    case 11 >= sumCard: option = "H"
    }
    return option
}
