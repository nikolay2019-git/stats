package stats

import "github.com/nikolay2019-git/bank/v2/pkg/types"

//TotalInCategory находит сумму покупок в определённой категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	//ToDo
	sum := types.Money(0)
	for _, payment := range payments {

		if payment.Category == category {
			if payment.Status != "Fail" {
				sum += payment.Amount
			}

		}

	}
	return sum
}

// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	//ToDo
	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Status != "Fail" {
			sum += payment.Amount
		}

	}
	result := types.Money(float64(sum) / float64(len(payments)))
	return result

}
