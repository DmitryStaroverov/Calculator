package main

import "fmt"

func main() {

	//Привет, нубарь!
	//Привет, нубарь!
	//Привет, нубарь!
	//Привет, нубарь!//Привет, нубарь!
	//Привет, нубарь!
	//Привет, нубарь!
	//Привет, нубарь!//Привет, нубарь!
	//Привет, нубарь!
	//Привет, нубарь!
	//Привет, нубарь!
	var (
		num1, num2 float64
		sigh       string
	)
	fmt.Print("Введите первое значение: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Неверно первое числовое значение.")
		return
	}
	fmt.Print("Введите знак для вычисления: ")
	_, err = fmt.Scan(&sigh)
	if err != nil {
		fmt.Println("Неверный знак для вычисления.")
		return
	}
	fmt.Print("Введите второе числовое значение: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Неверно второе числовое значение.")
		return
	}

	switch sigh {

	case "-":
		fmt.Println("Ответ:", num1-num2)
	case "+":
		fmt.Println("Ответ:", num1+num2)
	case "/":
		if num2 != 0 {
			fmt.Println("Ответ:", num1/num2)
		} else {
			fmt.Println("Делить на ноль нельзя.")
		}
	case "*":
		fmt.Println("Ответ:", num1*num2)
	default:
		fmt.Println("Неверное выражение")
	}

}
