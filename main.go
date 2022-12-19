package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	action := []string{"+", "-", "/", "*"}
	safeBox := []string{}

	conv := Converter{MapBox: make(map[string]int), arabianKeyMap: map[int]string{}}

	conv.arabianKeyMap[1] = "I"
	conv.arabianKeyMap[5] = "V"
	conv.arabianKeyMap[9] = "IX"
	conv.arabianKeyMap[10] = "X"
	conv.arabianKeyMap[40] = "XL"
	conv.arabianKeyMap[50] = "L"
	conv.arabianKeyMap[90] = "XC"
	conv.arabianKeyMap[100] = "C"
	conv.MapBox["I"] = 1
	conv.MapBox["V"] = 5
	conv.MapBox["X"] = 10

	actionIndex := -1

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	for i := 0; i < len(action); i++ {
		if strings.Contains(line, action[i]) {
			actionIndex = i
			break
		}
	}
	if actionIndex == -1 {
		fmt.Println("Некорректный оператор")
		return
	}
	conv.operator = action[actionIndex]
	safeBox = strings.Split(line, conv.operator)
	safeBox[0] = strings.Trim(safeBox[0], " \t")
	safeBox[1] = strings.Trim(safeBox[1], " \t")
	if conv.InRomanCheck(safeBox[0]) == conv.InRomanCheck(safeBox[1]) {

		var IsRoman bool
		IsRoman = conv.InRomanCheck(safeBox[0])
		if IsRoman {
			conv.x = conv.RomanToInt(safeBox[0])
			conv.y = conv.RomanToInt(safeBox[1])

			if conv.x > 10 || conv.y > 10 || conv.x <= 0 || conv.y <= 0{
				println("Вводные данные должны быть целыми числами в диапозоне от 1 до 10")
			} else {
				conv.res = conv.Calculation(conv.x, conv.y)
				if conv.res < 0 {
					conv.res = 0
					fmt.Printf("Результат с отрицательным значением, поэтому ответ - %d", conv.res)
				}
				romanRes := conv.inToRoman(conv.res)
				fmt.Println(romanRes)
			}

		} else {
			conv.x, _ = strconv.Atoi(safeBox[0])
			conv.y, _ = strconv.Atoi(safeBox[1])
			if conv.x > 10 || conv.y > 10 || conv.x <= 0 || conv.y <= 0{
				println("Вводные данные должны быть целыми числами в диапозоне от 1 до 10")
			} else {
				conv.res = conv.Calculation(conv.x, conv.y)
				println(conv.res)
			}
		}

	} else {
		fmt.Print("Числа должны быть в одном формате")
	}
}
