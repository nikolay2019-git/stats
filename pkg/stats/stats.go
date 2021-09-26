package stats

import "github.com/nikolay2019-git/bank.git/pkg/bank/types"

//TotalInCategory находит сумму покупок в определённой категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	//ToDo
	sum := types.Money(0)
	for _, payment := range payments {

		if payment.Category == category {
			sum += payment.Amount
		}

	}
	return sum
}
