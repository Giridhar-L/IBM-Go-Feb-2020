//package declaration
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//import other packages

/* var x,y int8
var (
	id int = 100,
	name string = "Magesh",
	isEmployed bool = true
)

var (
	id = 100
	name = "Magesh"
	isEmployed = true
)

var (
	id, name, isEmployed = 100, "magesh", true
) */

var reader = bufio.NewReader(os.Stdin)

func readConsole(prompt string) string {
	fmt.Println(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func main() {
	//id := 100
	/*
		var a [2]string
		a[0] = "Hello"
		a[1] = "World"
	*/
	//Arrays
	/*
		a := [2]string{"Hello", "World"}
		fmt.Println(a[0], a[1])
		fmt.Println(a)
		fmt.Println(len(a))
		a[0], a[1] = a[1], a[0]
		fmt.Println(a)
	*/

	//Slice
	/* nos := []int{10, 20, 30, 40, 50}
	fmt.Println(nos)
	fmt.Println(nos[1:4])
	fmt.Println(nos[:2])
	fmt.Println(nos[2:]) */

	//appending to a slice
	/* nos = append(nos, 60)
	fmt.Println(nos) */

	//appending multiple values
	/* nos = append(nos, 70, 80)
	fmt.Println(nos) */

	//appending another slice
	/* integers := []int{1, 2, 3, 4, 5}
	nos = append(nos, integers...)
	fmt.Println(nos)

	*/

	/*
		no := 8
		if no%2 == 0 {
			fmt.Printf("%v is even\n", no)
		} else {
			fmt.Printf("%v is odd\n", no)
		}
	*/

	/*
		if no := 8; no%2 == 0 {
			fmt.Printf("%v is even\n", no)
		} else {
			fmt.Printf("%v is odd\n", no)
		}
	*/
	//for loop
	/* for i := 0; i < 10; i++ {
		fmt.Println(i)
	} */

	//nested loop
	/*
		for i := 0; i < 10; i++ {
			for j := 0; j < 5; j++ {
				fmt.Println(i, j)
			}
		}
	*/

	//for as a while loop
	/* k := 0
	for k < 10 {
		fmt.Println(k)
		k++
	} */

	//for with no condition
	/* k := 0
	for {
		fmt.Println(k)
		k++
		if k >= 10 {
			break
		}
	} */

	/* nos := []int{10, 20, 30, 40, 50}
	for _, no := range nos {
		fmt.Println(no)
	} */

	//switch case
	/* text := readConsole("Enter your name ")
	fmt.Println(text)

	day := "mon"
	switch day {
	case "mon":
		fmt.Println("Monday")

	case "tue":
		fmt.Println("Tuesday")

	case "wed":
		fmt.Println("Wednesday")

	default:
		fmt.Println("Not a weekday")
	} */

	/*
		x, y := dummy()
		fmt.Println(x, y)
	*/

	//Maps
	/* pincodes := map[string]int{
		"KA": 560000,
		"TN": 600000,
		"MH": 400000,
	}

	pincodes["DL"] = 100000

	for key, value := range pincodes {
		fmt.Println(key, value)
	}

	//delete(pincodes, "DL")
	fmt.Println(len(pincodes))

	pin, exists := pincodes["DL"]
	if !exists {
		fmt.Println("DL pincode does not exist")
	} else {
		fmt.Println("DL pincode is ", pin)
	}
	fmt.Println(pincodes["TN"]) */
	//

	//Pointers
	/* x := 100
	addressOfX := &x
	fmt.Println(*addressOfX) */

	//struct
	e := Employee{id: 100, name: "Magesh", salary: 10000}
	e.display()
}

func newEmployee(id int, name string, salary float32) *Employee {
	newEmp := Employee{id: id, name: name, salary: salary}
	return &newEmp
}

func (e *Employee) display() {
	fmt.Println(e.id, e.name, e.salary)
}

/* struct */

type Employee struct {
	id     int
	name   string
	salary float32
}

/* func dummy() (int, string) {
	return 10, "xyz"
} */

/* func dummy() (int, string) {
	x,y = 10, "xyz"
	return x, y
} */

func dummy() (x int, y string) {
	x, y = 10, "xyz"
	return
}
