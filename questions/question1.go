package algorithm

import (
	"fmt"
	"strconv"
	"strings"
)

func Addition() {
	var addition int

	n := 4
	m := 9

	addition = m + n
	fmt.Println(addition)

}

func EvenOrOdd() {
	var number int
	fmt.Println("give me an integer number")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("your number is even")
	} else {
		fmt.Println("your number is odd")
	}

}

func Addition1to100() {

	addition := 0
	for i := 0; i <= 100; i++ {

		addition = addition + i

	}
	fmt.Println("addition of all numbers between 1 to 100:", addition)
}

func Addition1to100odd() {
	addition := 0
	for i := 1; i <= 100; i += 2 {
		addition = addition + i

	}
	fmt.Println("addition of all odd numbers between 1 to 100:", addition)

}

func Addition1to100even() {
	addition := 0
	for i := 0; i <= 100; i += 2 {
		addition = addition + i

	}
	fmt.Println("addition of all even numbers between 1 to 100:", addition)

}

func SmallNumber() {

	numbers := [8]int{8, -1, 2, 3, 5, 6, 7, 9}
	SmallNumber := numbers[0]
	for i := 0; i < 8; i++ {
		a := i
		for j := (i + 1); j < len(numbers); j++ {
			if numbers[j] < numbers[a] {
				SmallNumber = numbers[j]
				a = j
			}
		}

	}
	fmt.Println("smallest number :", SmallNumber)
}

func FindAverages() {
	numbers := [9]int{8, 1, 2, 3, 4, 5, 6, 7, 9}

	Addition := 0

	for i := 0; i < len(numbers); i++ {

		Addition = Addition + numbers[i]

	}

	fmt.Println("averages of numbers:", Addition/len(numbers))

}

func Notes() {

	var midterm float64
	var final float64

	fmt.Println("enter your midterm score")
	fmt.Scanln(&midterm)

	fmt.Println("enter your final score")
	fmt.Scanln(&final)

	fmt.Println("your total grade:", midterm*3/10+final*7/10)

}

func Bodymassindex() {

	var mass float64
	var lenght float64
	fmt.Println("enter your mass as kg")
	fmt.Scanln(&mass)
	fmt.Println("enter your lenght m")
	fmt.Scanln(&lenght)
	Bodymassindex := mass / (lenght * lenght)

	fmt.Println("your BodyMass İndex:", Bodymassindex)

	if Bodymassindex < 18.5 {
		fmt.Println("BMİI: weak")
	} else if Bodymassindex > 18.5 && Bodymassindex < 24.9 {
		fmt.Println("BMI :Normal")
	} else if Bodymassindex > 25 && Bodymassindex < 29.9 {
		fmt.Println("BMİI: overweight")
	} else if Bodymassindex > 30 && Bodymassindex < 34.9 {
		fmt.Println("BMI : I degree obese")
	} else if Bodymassindex > 35 && Bodymassindex < 39.9 {
		fmt.Println("BMI : II degree obese")
	} else {
		fmt.Println("BMI : III degree obese")
	}
}

func Question9() {

	fmt.Println("enter your integer elements. Seperate them using ','")

	var elements string

	fmt.Scanln(&elements)

	Str := strings.Split(elements, ",")
	list := make([]int, len(Str))

	for i, Str := range Str {
		INT, _ := strconv.Atoi(strings.TrimSpace(Str))
		list[i] = INT

	}

	for i := 0; i < len(list); i++ {
		a := i
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[a] {
				a = j
			}
		}
		list[i], list[a] = list[a], list[i]
	}
	fmt.Println(list)
}

func Question10() {
	fmt.Println("enter your integer elements. Seperate them using ','")

	var elements string

	fmt.Scanln(&elements)

	Str := strings.Split(elements, ",")
	list := make([]int, len(Str))
	for i, Str := range Str {
		INT, _ := strconv.Atoi(strings.TrimSpace(Str))
		list[i] = INT

	}

	for i := 0; i < len(list); i++ {

		amountOfnumbers := 1
		response := fmt.Sprintf("amount of number %d:", list[i])

		for j := i + 1; j < len(list); j++ {

			if list[j] == list[i] {
				amountOfnumbers = (amountOfnumbers + 1)
				list = append(list[:j], list[j+1:]...)
				j = j - 1
			}

		}
		fmt.Println(response, amountOfnumbers)
	}
}
