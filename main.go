package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("[+] Калькулятор запущен [+]")

	for {
		fmt.Print("Режим чтения : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		result, err := calculate(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println(result)
	}
}

func calculate(input string) (interface{}, error) {
	parts := strings.Split(input, " ")

	if len(parts) == 1 {
		panic("Не является математической операцией.")
	}

	if len(parts) != 3 {
		panic("Формат математической операции не удовлетворяет заданию")
	}

	num1, err := parseNumber(parts[0])
	if err != nil {
		return nil, err
	}

	num2, err := parseNumber(parts[2])
	if err != nil {
		return nil, err
	}

	isNum1Rim := isRoman(parts[0])
	isNum2Rim := isRoman(parts[2])

	if isNum1Rim != isNum2Rim {
		panic("Используются одновременно разные системы счисления.")
	}

	operator := parts[1]

	switch operator {
	case "+":
		result := num1 + num2
		if isNum1Rim {
			return toRoman(result), nil
		}
		return result, nil
	case "-":
		result := num1 - num2
		if isNum1Rim {
			if result <= 0 {
				panic("В римской системе нет отрицательных чисел.")
			}
			return toRoman(result), nil
		}
		return result, nil
	case "*":
		result := num1 * num2
		if isNum1Rim {
			return toRoman(result), nil
		}
		return result, nil
	case "/":
		if num2 == 0 {
			return nil, fmt.Errorf("деление на ноль")
		}
		result := num1 / num2
		if isNum1Rim {
			if result <= 0 {
				panic("В римской системе нет отрицательных чисел.")
			}
			return toRoman(result), nil
		}
		return result, nil
	default:
		panic("оператор != ( +, -, /, *)")

	}
}

func parseNumber(input string) (int, error) {
	if num, err := strconv.Atoi(input); err == nil && num >= 0 && num <= 10 {
		return num, nil
	}
	
	romanNumerals := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	if num, ok := romanNumerals[input]; ok {
		return num, nil
	}

	return 0, fmt.Errorf(`Калькулятор работает с целыми числами от 0 до 10 включительно.   Внимание ---> %s`, input)
}

func isRoman(input string) bool {
	romanNumerals := map[rune]bool{'I': true, 'V': true, 'X': true}
	for _, char := range input {
		if !romanNumerals[char] {
			return false
		}
	}
	return true
}

func toRoman(num int) string {

	romanNumerals := []struct {
		Value  int
		Symbol string
	}{

		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, rn := range romanNumerals {
		for num >= rn.Value {
			roman += rn.Symbol
			num -= rn.Value
		}
	}
	return roman
}
