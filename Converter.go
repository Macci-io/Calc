package main

type InRoman interface {
	InRomanCheck()
	Calculation()
	RomanToInt()
	IntToRoman()
	GetFloorKey()
}


type Converter struct {
	MapBox            map[string]int
	x                 int
	y                 int
	res               int
	saveStringElem    string
	saveStringElemTwo string
	operator          string
	arabianKeyMap     map[int]string
}

func (c *Converter) InRomanCheck(number string) bool {

	var arr []rune
	arr = []rune(number)
	if _ , ok := c.MapBox[string(arr[0])]; ok == true{
		return true
	} else {
		return false
	}

}

func (c *Converter) Calculation(x, y int) int {
	var res int
	switch c.operator {
	case "+":
		res = x + y
	case "-":
		res = x - y
	case "/":
		res = x / y
	case "*":
		res = x * y
	}
	return res
}

func (c *Converter) RomanToInt(roman string) int {
	end := len(roman) - 1

	var arabian int
	var arr []rune
	arr = []rune(roman)
	result, ok := c.MapBox[string(arr[end])]
	if ok == false {
		return -1
	}

	for i := end - 1; i >= 0; i-- {
		arabian, _ = c.MapBox[string(arr[i])]

		if arabian < c.MapBox[string(arr[i+1])] {
			result -= arabian
		} else {
			result += arabian
		}
	}
	return result
}

func (c *Converter) inToRoman(number int) string {
	var roman string

	for number != 0 {
		arabianKey := c.GetFloorKey(number)
		roman += c.arabianKeyMap[int(arabianKey)]
		number -= arabianKey
	}

	return roman
}

func (c *Converter) GetFloorKey(number int) int{
	var valueArab int
	for key := range c.arabianKeyMap {
		if number == key {
			valueArab = key
			break
		} else if number > key{
			valueArab = key
		}
	}
	return valueArab
}

