package transformation

type Transaction struct {
	ID       string
	UserID   string
	Category string
	Amount   int
}

func summarizeByUser(transactions []Transaction) map[string]int {
	result := make(map[string]int)

	for _, txn := range transactions {
		if txn.UserID == "" {
			continue
		}

		result[txn.UserID] += txn.Amount
	}

	return result
}
