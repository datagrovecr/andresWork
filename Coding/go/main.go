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
// defer
func welcome(){
	fmt.Println("welcome")
}
//outputs the code in the second function, then outputs welcome
func salutation(){
	defer welcome()
	fmt.Println("Hey")
}

//Scopes
// global
var myVar int

//local
func myFunc() {
	var h string = ""
}

//mars age challenge
func enter() {
    var age int
    fmt.Scanln(&age)

    mars := mars_age(age)
    fmt.Println(mars)
}
func mars_age(age int) int{
    marsYear := 1.88 // Mars year is approximately 1.88 Earth years
    marsAge := float64(age) / marsYear
    marsAgeRounded := int(marsAge) // Rounding to the nearest whole number
    return marsAgeRounded
}
//pointers
var p *int

func point() {
    var num float32
     var factor float32

    fmt.Scanln(&num)
    fmt.Scanln(&factor)

    scale(&num, factor)
    fmt.Println(num)
}

func scale(num *float32, factor float32){
    //mult num by factor
      *num = *num * factor
      }

//structs (objects)
func struc() {
    e1 := Employee{"James", 42, "Manager", 0}

    var x float32
    fmt.Scanln(&x)
    e1.salary = x
    
    fmt.Println("========")
    fmt.Println(e1.name + "(" + e1.position + ")")
    fmt.Println(e1.age)
    fmt.Println(e1.salary)
    fmt.Println("========")
}

type Employee struct{
    name string
    age int 
    position string
    salary float32
}

func ponitStruct() {
    var x int
    fmt.Scanln(&x)

    t := Timer{"timer1", 0}
    
    for i:=0;i<x;i++ {
        t.tick()
    }
}

type Timer struct{
    id string
    value int 
}
func (more *Timer) tick(){
     more.value ++
     fmt.Println(more.value)
}

//Array, Range, Map unit

//Array
var a [5] int
//b := [5]int{2, 4, 6, 8, 10}

//Maps
var mp map[string]int = map[string]int {
	"apple": 5,
	"pear": 6,
	"orange": 9,
} 


//maps of strings that point to an integer, does not keep track of the order only knows that a key points to a value


func maps(){
	var mp map[string]int = map[string]int {
		"apple": 5,
		"pear": 6,
		"orange": 9,
	} 
	//access values
	fmt.Println(mp["apple"])
	//change value
	mp["apple"] = 900
	//add value
	mp["kiwi"] = 20
	//delete value
	delete(mp, "apple")

	//check if a key exists
	val, ok:= mp["apple"]

	//get lenght of map
	fmt.Println(len(mp))
}