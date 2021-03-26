package transfer

// Total returns transfer amount with bonus
func Total(amount int) (result int) {
	bonus := 5
	result = amount + amount*bonus/1000
	return result
}
