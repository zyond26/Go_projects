package learn

import "fmt"

func main() {
	///Cach 1:
	var a = true
	fmt.Println("a = ", a)

	//Cach 2:
	b := "Hello World!"
	fmt.Println("b = ", b)

	//Cach 3:
	var c, d int = 10, 20
	fmt.Printf("c = %d\nd = %d\n", c, d)

	// Cach 4:
	var e float32
	e = 0.5
	fmt.Println("e = ", e)

	// Cach 5:
	str, num := "string", 10
	fmt.Printf("str = %s\nnum = %d\n", str, num)

	//Cach 6:
	var (
		name   = "Nguyen Van Huy"
		age    = 23
		height int
	)

	fmt.Println("my name is", name, ", age is", age, "and height is", height)
}

//OUT PUT: 
// a =  true
// b =  Hello World!
// c = 10
// d = 20
// e =  0.5
// str = string
// num = 10
// my name is Nguyen Van Huy , age is 23 and height is 0