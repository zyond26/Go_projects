package learn

import "fmt"

const str string = "String"

func man() {
	fmt.Println("str = ", str)
	const num = 10
	fmt.Println("num = ", num)

	const (
		name    = "Nguyen Van Huy"
		age     = 23
		country = "Viet Nam"
	)

	fmt.Println("name = ", name)
	fmt.Println("age = ", age)
	fmt.Println("country = ", country)
}