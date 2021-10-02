package stats

import (
	"github.com/nikolay2019-git/bank/v2/pkg/types"
)

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

// CategoriesAvg считает среднюю сумму платежа по каждой категории
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	//TODO
	// Получаем массив payment`ов
	// проходим по массиву, получаем Categoty И Amount
	// Нужно пройти по всем категориям и получить сумму и количество в категории?
	/*
		за один проход? Можно ли получить все суммы за один проход?
		Можно.
		Но как посчитать среднюю сумму платежа?
		Нужно для каждой категории узнать количество вхождений
		чтобы потом разделить на это количество
		если у нас есть категория, у нас будет одно вхождение и одна сумма
		а нам нужно где-то хранить количество вхождений отдельно от основной map
		map - это ответ.

		Нам нужно пройтись по списку и подсчитать среднее, а в конце записать в мар

		а  если хранить прошлую сумму и текущую? а если их не 2 и не 3?
		как мы узнаем, что получили все суммы? нужно создавать список сумм для каждой категории
		или делать столько проходов, сколько категорий?





	*/
	countMap := map[types.Category]int{}

	categories := map[types.Category]types.Money{} // результирующая мар

	for _, payment := range payments { // один проход по списку payments
		/*
			Проверяем, есть ли в мар эта категория ставим не сумму а количество вхождений +1
			в итоге получаем мар с количеством вхождений этой категории
			тогда проходим второй раз, при первом обращении можем слить количество вхождений в переменную
			а где я тогда возьму амоунты?
			Для каждой категории можно подсчитывать сумму, а количество вхождений накапливать
			но надо накапливать количество*количество категорий
			неизвестное количество переменных, список что ли

			список количеств, зависящий от категории, ещё одна мар?
			параллельно.
			можно ли считать среднее на ходу?
			не знаю даже как

			пофиг на память и количество проходов
			нужно решить задачу хотя бы в лоб
			хоть как то
			за первый проход
			создаём список категорий, какие категории присутствуют и сколько их
			столько мы и вернём в ответе
			в маре мы каждый раз при встрече категории накапливаем количество
			откуда будем брать количество каждый раз, когда будем считать среднее
			нет, мы возьмём количество только раз, в конце,
			а накапливать будем сумму, прибавляя каждый раз к амоунт во второй мар

			когда пройдёмся по списку в последний раз, разделим на количество


			Первый проход:
			получаем список категорий и подсчитываем количество вхождений
			можно даже сумму считать за первый проход?
			и в конце просто посчитать среднее для каждой категории

			у нас три суммы
			почему три?
			от 1 и выше, неизвестное количество
			N категорий, N количеств
			N средних сумм
		*/
		_, ok := countMap[payment.Category]
		if ok {
			countMap[payment.Category] += 1

			categories[payment.Category] += payment.Amount
		} else {
			countMap[payment.Category] = 1
			categories[payment.Category] = payment.Amount
		}

	}

	for key, value := range categories {
		categories[key] = types.Money(float64(value) / float64(countMap[key]))
	}
	return categories
}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {

	resultMap := map[types.Category]types.Money{}

	/*
		Нужно ли создать объединение ключей из двух карт, если у них не совпадают ключи?
		Как получить из второй карты то, чего нет в первом?
		Лишний проход по всем ключам?
		Или один проход по каждой карте?
		Одновременный проход?




	*/


	for key, value := range first {
		if _, ok := second[key]; ok {
			resultMap[key] = second[key] - value
		} else {
			resultMap[key] = 0 - value
		}

	}
	for key, value := range second {
		if _, ok := resultMap[key]; !ok{
			if _, ok:= first[key]; ok{
				resultMap[key] = value - first[key]
			} else{
				resultMap[key] = value
			}
			
		}
	}

	return resultMap
}
