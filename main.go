package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Operation struct {
	Operand1 int
	Operand2 int
	Operator string
	isArabic bool
}

var romanNumerals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// Перевод в арабскую систему счисления
func convertToArabic(roman string) int {
	total := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value, correct := romanNumerals[rune(roman[i])]

		if !correct {
			log.Panic("Ошибка конвертации")
		}

		if value < prevValue {
			total -= value
		} else {
			total += value
		}

		prevValue = value
	}
	return total
}

// Перевод в римскую систему счисления
func convertToRoman(num int) string {
	if num <= 0 || num > 100 {
		panic("Число должно быть от 1 до 100")
	}

	var result string

	for _, numeral := range "MDCLXVI" {
		value := romanNumerals[numeral]
		for num >= value {
			result += string(numeral)
			num -= value
		}
	}

	return result
}

func checkCorrectInputValue(str []string) Operation {
	// Проверяем длинну примера, тк может быть всего 2 переменные и операнд то только - 3
	if len(str) != 3 {
		panic("Ошибка в примере")
	}

	operator := str[1]

	// Переводим обе переменные в числа
	operand1, err1 := strconv.Atoi(str[0])
	operand2, err2 := strconv.Atoi(str[2])

	if operand1 > 10 || operand2 > 10 {
		panic("Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более")
	}

	// Если оба числа не римские то вернем пример
	if err1 == nil && err2 == nil {
		return Operation{operand1, operand2, operator, false}
	}

	// Если одно из чисел не прошло конвертацию
	if err1 == nil || err2 == nil {
		panic("Используются одновременно разные системы счисления.")
	}

	// В остальных случая у нас две строки, переведем их в арабскую и если не получится - словим панику
	operand1 = convertToArabic(str[0])
	operand2 = convertToArabic(str[2])

	return Operation{operand1, operand2, operator, true}
}

func calculating(op Operation) int {
	switch op.Operator {
	case "+":
		return op.Operand1 + op.Operand2
	case "-":
		return op.Operand1 - op.Operand2
	case "/":
		return op.Operand1 / op.Operand2
	case "*":
		return op.Operand1 * op.Operand2
	default:
		panic("Ошибка в примере: операнд не допустим!")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите Ваш пример:")

	userInput, _ := reader.ReadString('\n')

	regex := regexp.MustCompile(`(\d+|[IVXLCDM]+|[+\-*/])`)
	tokens := regex.FindAllString(userInput, -1)

	// Создаем массив где будем хранить отфильтрованные значения
	filteredSymbols := make([]string, 0)
	for _, symbol := range tokens {
		if symbol != "" {
			filteredSymbols = append(filteredSymbols, symbol)
		}
	}

	// Проверяем корректность примера
	formattedValue := checkCorrectInputValue(filteredSymbols)

	// Решаем пример
	result := calculating(formattedValue)

	if !formattedValue.isArabic {
		fmt.Printf("Ответ примера: %d", result)
		return
	}

	if result < 1 {
		panic("В римской системе нет отрицательных чисел.")
	}

	convertedResult := convertToRoman(result)
	fmt.Printf("Ответ примера: %s", convertedResult)
}
