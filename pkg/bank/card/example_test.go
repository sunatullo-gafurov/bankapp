package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {

	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)

	// Output:
	// 1000000
}

func ExampleDeposit_positive() {

	card := types.Card{Balance: -10_000_00, Active: true}
	Deposit(&card, 40_000_00)
	fmt.Println(card.Balance)

	// Output:
	// 3000000
}

func ExampleDeposit_inactive() {

	card := types.Card{Balance: 40_000_00, Active: false}
	Deposit(&card, 40_000_00)
	fmt.Println(card.Balance)

	// Output:
	// 4000000
}
func ExampleDeposit_limit() {

	card := types.Card{Balance: 30_000_00, Active: true}
	Deposit(&card, 60_000_00)
	fmt.Println(card.Balance)

	// Output:
	// 3000000
}

func ExampleAddBonus_positive() {

	card := types.Card{Balance: 20_000_00, MinBalance: 12_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)

	// Output:
	// 2002958
}

func ExampleAddBonus_inactive() {

	card := types.Card{Balance: 20_000_00, MinBalance: 12_000_00, Active: false}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)

	// Output:
	// 2000000
}

func ExampleAddBonus_noMoney() {

	card := types.Card{Balance: -20_000_00, MinBalance: 12_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)

	// Output:
	// -2000000
}

func ExampleAddBonus_limit() {

	card := types.Card{Balance: 20_000_00, MinBalance: 12_000_000_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)

	// Output:
	// 2000000
}

func ExampleTotal() {

	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 15_000_00,
			Active:  true,
		},
		{
			Balance: -15_000_00,
			Active:  true,
		},
		{
			Balance: 15_000_00,
			Active:  false,
		},
	}

	fmt.Println(Total(cards))

	//Output:
	//2500000
}

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Name:    "card",
			PAN:     "5058 xxxx xxxx 6198",
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Name:    "card",
			PAN:     "5058 xxxx xxxx 3254",
			Balance: 10_000_00,
			Active:  false,
		},
		{
			Name:    "card",
			PAN:     "5058 xxxx xxxx 4789",
			Balance: -10_000_00,
			Active:  true,
		},
		{
			Name:    "card",
			PAN:     "5058 xxxx xxxx 4789",
			Balance: 20_000_00,
			Active:  true,
		},
	}

	newCards := PaymentSources(cards)

	for _, newCard := range newCards {
		fmt.Println(newCard.Number)
	}

	//Output:
	//5058 xxxx xxxx 6198
	//5058 xxxx xxxx 4789
}
