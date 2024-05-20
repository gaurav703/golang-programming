package main

import "fmt"

func main() {
	var number int = 10
	fmt.Println("Number is", number)

	// short variable declaration
	congrats := "Congratulations!"
	fmt.Println(congrats)

	// multiple variable declaration
	var (
		name string = "John"
		age  int    = 25
	)
	fmt.Println("Name is", name)
	fmt.Println("Age is", age)

	// same line variable declaration
	var a, b, c int = 1, 2, 3
	fmt.Println("a is", a)
	fmt.Println("b is", b)
	fmt.Println("c is", c)

	// short variable declaration
	x, y, z := 4, "hello", 6
	fmt.Println("x is", x)
	fmt.Println("y is", y)
	fmt.Println("z is", z)

	// type conversion example
	var i int = 10
	var f float64 = float64(i)
	fmt.Println("type conversion is", f)

	// data type in go
	var (
		// integer
		integer int = 10
		// float
		float float64 = 10.5
		// string
		str string = "Hello"
		// boolean
		boolean bool = true
		// complex number is not used frequently
		complex complex128 = 1 + 2i

		// unsigned integer
		// uint8  : 0 to 255
		// uint16 : 0 to 65535
		// uint32 : 0 to 4294967295

		// signed integer
		// int8  : -128 to 127
		// int16 : -32768 to 32767
		// int32 : -2147483648 to 2147483647
		// int64 : -9223372036854775808 to 9223372036854775807

	)

	fmt.Println("integer is", integer)
	fmt.Println("float is", float)
	fmt.Println("string is", str)
	fmt.Println("boolean is", boolean)
	fmt.Println("complex is", complex)

	// string length
	var str1 string = "Hello, World!"
	fmt.Println("Length of string is", len(str1))

	// string concatenation
	var str2 string = "Hello, "
	var str3 string = "World!"
	fmt.Println("Concatenated string is", str2+str3)

	// string comparison
	var str4 string = "Hello, World!"
	var str5 string = "Hello, World!"
	fmt.Println("Are strings equal?", str4 > str5)

}
