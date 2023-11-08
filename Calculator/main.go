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
	var valid1, valid2, rome1, rome2 bool
	for romanNum, arabNum := range validNumbers { // Проверка первой цифры
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
	for romanNum, arabNum := range validNumbers { // Проверка второй цифры
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
	return
}

// ОБРАТНОЕ ПРЕОБРАЗОВАНИЕ РЕЗУЛЬТАТА ВЫЧИСЛЕНИЯ
// Функция преобразования арабской цифры/числа (int) в римскую цифру (string)
func arabNumToRomanNum(num int) (romanString string) {
	for i := 0; i < num; i++ {
		romanString = romanString + "I"
	}
	romanString = strings.Replace(romanString, "IIIII", "V", -1)  // 1. Преобразование "IIIII" в "V"
	romanString = strings.Replace(romanString, "VV", "X", -1)     // 2. Преобразование "VV" в "X"
	romanString = strings.Replace(romanString, "IIII", "IV", -1)  // 3. Преобразование "IIII" в "IV"
	romanString = strings.Replace(romanString, "VIV", "IX", -1)   // 4. Преобразование "VIV" в "IX"
	romanString = strings.Replace(romanString, "XXXX", "XL", -1)  // 5. Преобразование "XXXX" в "XL"
	romanString = strings.Replace(romanString, "XLX", "L", -1)    // 6. Преобразование "XLX" в "L"
	romanString = strings.Replace(romanString, "LL", "LXXX", -1)  // 7. Преобразование "LL" в "LXXX"
	romanString = strings.Replace(romanString, "LXXXX", "XC", -1) // 8. Преобразование "LXXXX" в "XC"
	romanString = strings.Replace(romanString, "XCX", "C", -1)    // 9. Преобразование "XCX" в "C"
	return
}

func main() {
	var result int
	operand := []string{"+", "-", "/", "*"}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите значение")
		text, _ := reader.ReadString('\n') // Ввод данных в формате строки
		valid := false
		rome := false
		for i := range operand {
			num1, num2, containOperand := strings.Cut(text, operand[i])
			num1 = strings.TrimSpace(num1) // Очищает все пробелы, табуляцию
			num2 = strings.TrimSpace(num2)
			if containOperand {
				a, _ := strconv.Atoi(num1) // Преобразование string в int
				b, _ := strconv.Atoi(num2)
				valid, rome, a, b = validNum(a, b, num1, num2) // Проверка валидности введенных данных и определения системы исчисления
				switch i {
				case 0:
					result = a + b
				case 1:
					result = a - b
				case 2:
					if b != 0 {
						result = a / b
					}
				case 3:
					result = a * b
				}
			}
		}

		if valid {
			switch {
			case rome && result > 0:
				fmt.Println(arabNumToRomanNum(result)) // Вывод результата в римской системе исчисления
			case rome && result < 0:
				fmt.Println("Исключение, в римской системе исчесления отсутствуют отрицательные числа")
			case rome && result == 0:
				fmt.Println("Исключение, в римской системе исчесления отсутствуeт ноль")
			default:
				fmt.Println(strconv.Itoa(result)) // Вывод результата в арабской системе исчисления
			}
		} else {
			fmt.Println("Ошибка введенных данных") // Вывод ошибки введенных данных
			break
		}
	}
}
