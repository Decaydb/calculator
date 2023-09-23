package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/imbue11235/roman"
)

var wichSystem string
var numOne int
var numTwo int
var res int
var first, sign, second string

// Преобразователь из арабской системы в римскую.
func arabicToRoman(arabicNum int) string {
	return roman.FromArabic(arabicNum)
}

// Преобразователь из римской в арабскую.
func romanToArabic(romanNum string) int {
	return roman.ToArabic(romanNum)
}

//Конвертация из string в целочисленный int.

func Convert(convFirst string, convSecond string) (int, int) {
	solOne, _ := strconv.Atoi(convFirst)
	convSecond = strings.ReplaceAll(convSecond, "\r", "")
	solTwo, err := strconv.Atoi(convSecond)
	if err != nil {
		panic(err)
	}
	return solOne, solTwo
}

// Функция проверки. Принимает первое число, знак, второе число. Возвращает уже систему и тип знака.
func Handling(target string) (string, string, string) {
	slice := strings.Split(target, " ")
	if len(slice) > 3 == true {
		panic("Можно вводить только два числа и знак между ними!")
	} else if len(slice) < 1 {
		panic("Так нельзя, введите больше!")
	} else if len(slice) == 3 {
		first = slice[0]
		second = slice[2]
		sign = slice[1]

	} else {
		panic("Строка не является математической операцией")
	}
	check1, _ := regexp.MatchString("^[1-9]|10", first)
	check2, _ := regexp.MatchString("^[1-9]|10", second)
	if check1 == true && check2 == true {
		wichSystem = "arabic"
	} else if (check1 == false) && (check2 == false) {
		wichSystem = "roman"
	} else if (check1 == true) && (check2 == false) {
		panic("Калькулятор умеет работать только с арабскими или римскими цифрами, если арабское то целое и положительное, не больше 10!")
	} else if (check1 == false) && (check2 == true) {
		panic("Калькулятор умеет работать только с арабскими или римскими цифрами если арабское то целое и положительное, не больше 10!")
	}
	//println("Какая используется система счисления: ", wichSystem, "\n")
	return first, sign, second
}

func Solution(solFirst string, solSign string, solSecond string) {

	if wichSystem == "arabic" {

		numOne, numTwo = Convert(solFirst, solSecond)
		if (numOne < 1) || (numOne > 10) == true {
			panic("Нельзя вводить числа меньше чем 1 и больше чем 10!")
		} else {
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
		}
		println(res)
	}
	if wichSystem == "roman" {
		numOne = romanToArabic(solFirst)
		numTwo = romanToArabic(solSecond)

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
	reader := bufio.NewReader(os.Stdin)
	for {
		print("Введите выражение!\n")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		first, sign, second = Handling(text)
		Solution(first, sign, second)
	}
}
