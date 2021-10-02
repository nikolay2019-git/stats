package stats

import (
	"reflect"
	"testing"

	"github.com/nikolay2019-git/bank/v2/pkg/types"
)

func TestCategoriesAvg_empty(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1000},
		{ID: 2, Category: "food", Amount: 4000},
		{ID: 3, Category: "auto", Amount: 1000},
		{ID: 4, Category: "auto", Amount: 1000},
		{ID: 5, Category: "fun", Amount: 1000},
		{ID: 6, Category: "food", Amount: 2000},
	}

	expected := map[types.Category]types.Money{
		"auto": 1000,
		"fun":  1000,
		"food": 3000,
	}
	result := CategoriesAvg(payments)

	//fmt.Println(result)
	if len(result) == 0 {
		t.Error("result len ==0")
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodDynamic_empty(t *testing.T) {

	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"auto": 10,
		"food": 25,
		"mobile":5,
	}

	expected := map[types.Category]types.Money{
		"auto": 0,
		"food": 5,
		"mobile":5,
	}
	result := PeriodDynamic(first, second)

	if len(result) == 0 {
		t.Error("result len ==0")
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}
