package stats

import (
	"fmt"

	"github.com/nikolay2019-git/bank/v2/pkg/types"
)

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       0,
			Amount:   1_000_00,
			Category: "food",
			Status:   "OK",
		},
		{
			ID:       1,
			Amount:   2_000_00,
			Category: "food",
			Status:   "OK",
		},
		{
			ID:       2,
			Amount:   3_000_00,
			Category: "food",
			Status:   "OK",
		},
	}
	fmt.Println(TotalInCategory(payments, "food"))
	//Output:600000

}
