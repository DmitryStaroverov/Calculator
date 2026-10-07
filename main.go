package main

import "fmt"

func main() {

	var (
		num1, num2 float64
		sigh       string
	)
	fmt.Print("Введите первое значение: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("Неверно первое числовое значение.")
		return
	}
	fmt.Print("Введите знак для вычисления: ")
	_, err = fmt.Scanln(&sigh)
	if err != nil {
		fmt.Println("Неверный знак для вычисления.")
		return
	}
	fmt.Print("Введите второе числовое значение: ")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Println("Неверно второе числовое значение.")
		return
	}

	switch sigh {

	case "-":
		answer, result := subtract(num1, num2)
		fmt.Println(answer, result)

	case "+":
		answer, result := add(num1, num2)
		fmt.Println(answer, result)

	case "/":
		answer, result, err := divide(num1, num2)

		if err != nil {

			fmt.Println(err)
			return
		}
		fmt.Println(answer, result)

	case "*":
		answer, result := multiply(num1, num2)
		fmt.Println(answer, result)

	default:

		fmt.Println("Неверное выражение")
	}

}

func add(a, b float64) (string, float64) {

	return "Ответ:", a + b
}

func subtract(a, b float64) (string, float64) {
	return "Ответ:", a - b

}

func multiply(a, b float64) (string, float64) {
	return "Ответ:", a * b

}
func divide(a, b float64) (string, float64, error) {
	if b == 0 {
		return "", 0, fmt.Errorf("ошибка: деление на ноль.")
	}
	return "Ответ:", a / b, nil
}
