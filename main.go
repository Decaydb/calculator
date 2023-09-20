package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// Мапы, чтоб не строчить кучу кода. Можно было конечно воспользоваться какой-нибудь библиотекой для преобразования...
var aNums = map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X", 11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV", 16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX"}
var rNums = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10, "XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20}

// var wichSign string
var wichSystem string
var numOne int
var numTwo int
var res int

// Преобразователь из арабской системы в римскую.
func arabicToRoman(arabicNum int) string {
	romanNum, exists := aNums[arabicNum]
	if !exists {
		return ""
	}
	return romanNum
}

// Преобразователь из римской в арабскую.
func romanToArabic(romanNum string) int {
	arabicNum, exists := rNums[romanNum]
	if !exists {
		return -1
	}
	return arabicNum
}

//Конвертация из string в целочисленный int.

func Convert(convFirst string, convSecond string) (int, int) {
	solOne, _ := strconv.Atoi(convFirst)
	solTwo, _ := strconv.Atoi(convSecond)
	return solOne, solTwo
}

// Функция проверки. Принимает первое число, знак, второе число. Возвращает уже систему и тип знака.
func Check(strFirst string, strSecond string) {

	check1, _ := regexp.MatchString("^[1-9]|10", strFirst)
	check2, _ := regexp.MatchString("^[1-9]|10", strSecond)
	if check1 == true && check2 == true {
		wichSystem = "arabic"
	} else if check1 == false && check2 == false {
		wichSystem = "roman"
	} else if check1 == true && check2 == false {
		panic("Калькулятор умеет работать только с арабскими или римскими цифрами одновременно!")
	} else if check1 == false && check2 == true {
		panic("Калькулятор умеет работать только с арабскими или римскими цифрами одновременно!")
	}
}

func Solution(solFirst string, solSign string, solSecond string) {

	if wichSystem == "arabic" {

		numOne, numTwo = Convert(solFirst, solSecond)
		if (numOne < 1) || (numOne > 10) == true {
			panic("Нельзя вводить числа меньше чем 1 и больше чем 10!")
		} else {

			switch solSign {
			case "+":
				res = (numOne + numTwo)
			case "-":
				res = (numOne - numTwo)
			case "*":
				res = (numOne * numTwo)
			case "/":
				res = (numOne / numTwo)
			}
		}
		println(res)
	}
	if wichSystem == "roman" {
		numOne = romanToArabic(solFirst)
		numTwo = romanToArabic(solSecond)
		print(numOne)

		switch solSign {
		case "+":
			res = numOne + numTwo
		case "-":
			res = numOne - numTwo
		case "*":
			res = numOne * numTwo
		case "/":
			res = numOne / numTwo
		}
		if res < 1 {
			panic("Результатом вычисления римских чисел не может быть ноль или отрицательное число")
		} else {

			println(arabicToRoman(res))
		}
	}

}

func main() {
	print("Введите выражение!\n")
	var first, second, sign string //Хоть и объявлено некрасиво, но нужны лишь пустые строки.
	fmt.Scan(&first, &sign, &second)
	Solution(first, sign, second)
	println(wichSystem)

}
