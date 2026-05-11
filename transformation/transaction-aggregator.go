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

func summarizeByUserAndCategory(transactions []Transaction) map[string]map[string]int {
	result := map[string]map[string]int{}

	for _, txn := range transactions {
		if txn.UserID == "" || txn.Category == "" {
			continue
		}

		if result[txn.UserID] == nil {
			result[txn.UserID] = make(map[string]int)
		}
		result[txn.UserID][txn.Category] += txn.Amount
	}

	return result
}
