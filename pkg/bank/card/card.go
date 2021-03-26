package card

import "bank/pkg/bank/types"

// IssueCard creates a card
func IssueCard(currency types.Currency, color string, name string) types.Card {

	return types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
}

// Withdraw withdraws money
func Withdraw(card types.Card, amount types.Money) types.Card {

	if !card.Active || amount > card.Balance || amount <= 0 || amount > 20_000_00 {
		return card
	}

	card.Balance -= amount

	return card
}

// Deposit tops up money to the card
func Deposit(card *types.Card, amount types.Money) {

	topUpLimit := 50_000_00

	if !card.Active {
		return
	}

	if amount <= 0 {
		return
	}

	if amount > types.Money(topUpLimit) {
		return
	}

	card.Balance += amount
}

// AddBonus each month adds bonus to the card according to minimal balance
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {

	maxBonus := 5_000_00

	if !card.Active {
		return
	}

	if card.Balance < 0 {
		return
	}

	if card.MinBalance < 10_000_00 {
		return
	}

	bonus := card.MinBalance * types.Money(percent) * types.Money(daysInMonth) / (types.Money(daysInYear) * 100)

	if bonus > types.Money(maxBonus) {
		return
	}

	card.Balance += bonus
}

// Total calculates sum of balances in card
func Total(cards []types.Card) types.Money {

	sum := types.Money(0)

	for _, card := range cards {
		if card.Active && card.Balance >= 0 {
			sum += card.Balance
		}
	}

	return sum
}

// PaymentSources makes payment from a card
func PaymentSources(cards []types.Card) []types.PaymentSources {

	cardsToPay := []types.PaymentSources{}

	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			card := types.PaymentSources{
				Type:    card.Name,
				Number:  card.PAN,
				Balance: card.Balance,
			}
			cardsToPay = append(cardsToPay, card)
		}
	}
	return cardsToPay
}
