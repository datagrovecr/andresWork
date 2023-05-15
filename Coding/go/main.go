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

