package main

//fmt Format package, allows input and output functionalities.
import (
	"fmt" 
	//"math"
)
//import multiple packages

func main(){
	fmt.Println("Hello")
	//variables
	var i, a int
	i, a = 8, 12
	fmt.Println(i, a)
	//short notation
	v := 8
	fmt.Println(v)
	//constants, these cannot be declared in short form
	const pi = 3.14
	input()
}
func input(){
	//takes user input
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
func conditions(){
	if x:=42; x>18 {
		fmt.Println("Allowed")
	}else {
		fmt.Println("Not Allowed")
	}

	var fever float32
    fmt.Scanln(&fever)
    if fever >99.5{
        fmt.Println("Fever")
    }else {
        fmt.Println("Allowed")
    }
}

func switches(){
	var f int
	fmt.Scanln(&f)
	//your code goes here
	switch {
	  case f < 0:
		fmt.Println("Wrong Input")
	  case f<20 :
		fmt.Println("Infrasound")
	  
	  case f>20000:
		fmt.Println("Ultrasound")
	 default :
		fmt.Println("Audible")
	}
}

func loops(){
	//for
	sum :=1
	res :=0
	for sum <=1000{
		res+= sum
		sum ++
	}
	fmt.Println(res)

	var years int
    fmt.Scanln(&years)

    s:= 7
    for i:=0; i<years; i++{
        s *=2
    }
    fmt.Println(s)
}


func loopSwitch() {
    //your code goes here
	var numbers [3]int

	for i := 0; i < 3; i++ {
		fmt.Scanln(&numbers[i])
	}

	// Convert numbers to corresponding texts and output them
	for i := 0; i < 3; i++ {
		text := numberToText(numbers[i])
		fmt.Println(text)
	}
}

// Function to convert a number to its corresponding text in English
func numberToText(num int) string {
	switch num {
	case 0:
		return "Zero"
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	case 10:
		return "Ten"
	default:
		return "Unknown"
	}
}

//functions
//calling functions from another function
func calling() {
    talk()
}

func talk(){
    fmt.Println("GO")
}
func start() {
    var w string
    fmt.Scanln(&w)
    var x int
    fmt.Scanln(&x)

    repeat(w, x)
}

func repeat(w string, x int) {
    for i:= 0; i < x; i++ {
        fmt.Println(w)
    } 
}