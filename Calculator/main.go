package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция проверки валидности введенных данных и определения двух цифр
func validNum(a, b int, s1, s2 string) (valid, rome bool, arabNum1, arabNum2 int) {
	validNumbers := map[string]int{ // Мап валидных значений вводимых данных
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10}
	valid = false
	valid1 := false
	valid2 := false
	rome = false
	rome1 := false
	rome2 := false
	for romanNum, arabNum := range validNumbers {
		switch {
		case validNumbers[romanNum] == a:
			arabNum1 = arabNum
			valid1 = true
		case romanNum == s1:
			arabNum1 = arabNum
			valid1 = true
			rome1 = true
		}
	}
	for romanNum, arabNum := range validNumbers {
		switch {
		case validNumbers[romanNum] == b:
			arabNum2 = arabNum
			valid2 = true
		case romanNum == s2:
			arabNum2 = arabNum
			valid2 = true
			rome2 = true
		}
	}
	valid = valid1 && valid2
	rome = rome1 && rome2
	if valid && rome1 == rome2 {
		valid = true
	} else {
		valid = false
	}
	return valid, rome, arabNum1, arabNum2
}

// ОБРАТНОЕ ПРЕОБРАЗОВАНИЕ РЕЗУЛЬТАТА ВЫЧИСЛЕНИЯ
// Функция преобразования арабской цифры/числа (int) в римскую цифру (string)
func arabNumToRomanNum(num int) string {
	var stickString string
	for i := 0; i < num; i++ {
		stickString = stickString + "I"
	}
	replacedString := strings.Replace(stickString, "IIIII", "V", -1)   // 1. Преобразование "IIIII" в "V"
	replacedString = strings.Replace(replacedString, "VV", "X", -1)    // 2. Преобразование "VV" в "X"
	replacedString = strings.Replace(replacedString, "IIII", "IV", -1) // 3. Преобразование "IIII" в "IV"
	replacedString = strings.Replace(replacedString, "VIV", "IX", -1)  // 4. Преобразование "VIV" в "IX"
	return replacedString
}

func main() {

	var result int
	operand := []string{"+", "-", "/", "*"}

	reader := bufio.NewReader(os.Stdin)

	for {

		//ВВОД ДАННЫХ
		fmt.Println("Введите значение")
		text, _ := reader.ReadString('\n') // Ввод данных в формате строки

		valid := false
		rome := false

		for i := range operand {
			num1, num2, containOperand := strings.Cut(text, operand[i])
			num1 = strings.TrimSpace(num1) // Очищает все пробелы, табуляцию
			num2 = strings.TrimSpace(num2)
			if containOperand {

				a, _ := strconv.Atoi(num1)
				b, _ := strconv.Atoi(num2)

				valid, rome, a, b = validNum(a, b, num1, num2)

				switch i {
				case 0:
					result = a + b
				case 1:
					result = a - b
				case 2:
					result = a / b
				case 3:
					result = a * b
				}
			}
		}

		if valid {
			switch {
			case rome && result > 0:
				fmt.Println(arabNumToRomanNum(result))
			case rome && result < 0:
				fmt.Println("Ошибка, в римской системе исчесления отсутствуют отрицательные числа")
			case rome && result == 0:
				fmt.Println("Ошибка, в римской системе исчесления отсутствуeт ноль")
			default:
				fmt.Println(strconv.Itoa(result)) // Вывод результата
			}
		} else {
			fmt.Println("Ошибка введенных данных") // Вывод ошибки введенных данных
			break
		}
	}
}
