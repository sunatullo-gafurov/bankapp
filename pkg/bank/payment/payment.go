package payment

import "bank/pkg/bank/types"

// Max returns the maximum payment
func Max(payments []types.Payment) types.Payment {

	index := 0
	max := payments[0].Amount

	for i, payment := range payments {
		if max < payment.Amount {
			max = payment.Amount
			index = i
		}
	}
	return payments[index]
}
