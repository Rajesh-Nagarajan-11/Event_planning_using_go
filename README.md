# Event_planning_using_go
Backend RestApi using Gin framework

Download all dependencies If the project already has a go.mod file (which it should if it's a Go module project), just run:

go mod tidy
This will download all the needed packages.
It will also remove any unused ones automatically.

To Run the Project Command : go run . 

# My Go notes 
Go is a Open Source Programming Language Developed By Google 
Static Typing Language 
Sample Code 
package Main <-- Main package is entry point of go Program

import "fmt" <-- Importing Necessary Built in Packages 

func main(){ <-- Main Function
 fmt.println("Hello World") <--statement
}

To run Go , Command 'go run filename.go'

else need to made module , go mod init example.com/My-app
then 'go build' <-- this create a .exe file in windows and executable file in Linux and Mac os 
then ./Filename 
or 'go run .' also runs

# Variables :
   Two ways to Declare Variables in Go 
   1. Var keyword 
   2. :=
   
   Var name="Rajesh"
   name:="Rajesh"
   Both works 
   
  Four Types of Data Types :
  1.int (int8,int16,int32,int64)
  2.Float (float32,float64)
  3.string
  4.Boolean (bool)
  
  Sample Code :
		package main
		import "fmt"
		func main() {
			var IsTrue bool = true
			var num int = 89
			var name string = "Rajesh"
			var dec float64 = 89.76
			fmt.Println(IsTrue)
			fmt.Println(num)
			fmt.Println(name)
			fmt.Println(dec)
			
		}
 Constant using Const Keyword example : const pi =3.12 , it wont reassigned 
 # Getting Input from a User  <---console 
    fmt.scan(..any) is function 
    parameter should be set as address of the variable as C programming (&) ampersand symbol 
    Note : First Word of String only get Recoginized By the fmt.scan() function
    <--Example-->
    func main() {

	var x, y int64
	fmt.Println("Enter X value")
	fmt.Scan(&x)
	fmt.Println("Enter Y Value")
	fmt.Scan(&y)
	
	fmt.Println("add value", x+y)
}

Printf <--function , Same as in C program 

 Format specifier In go Refer ---> https://pkg.go.dev/fmt
 
 Sprintf,Sprintln , Sprint <--- built in functions under fmt package used to format a string which returns the string ,
 usecase ---> instead of Directly printing , format in right way and stores in a string variable and print that variable 
 
 Example :
	x := fmt.Sprintln("Expenses Before Tax :", ebt)
	y := fmt.Sprint("Expenses After Tax :", eat)
	z := fmt.Sprintf("Profit : %v \n", profit)
	fmt.Print(x, y, z)
	
Multi-Line Print Using (Backticks) ``
        Example :
        x := fmt.Sprintln(`Expenses 
	Before Tax :`, ebt)
Note : Escape Sequences Not Validate in Multi Line print Statements
# Functions 
	Func is keyword used to create a Function ,
	 Func add (x,y int64)
	 {
	 //statements
	 }
	 if you have return type then 
	 func add(x,y int64) (int64)
	 {
	   res=x+y
	   return res
	 }
	 
	 alternative Apporach 
	 func add(x,y int64) (res int64)
	 {
	   res=x+y
	   retrun // simple return keyword Returns the res 
	 }
	 
	 Addtionally In go lang You can able to return more than one value 
	 For Example :
	 func main() {
		var x int64
		var y int64
		fmt.Println("Enter X Value")
		fmt.Scan(&x)
		fmt.Println("Enter Y Value")
		fmt.Scan(&y)
		var add, sub, mul, div = arithmetric(x, y)
		fmt.Println(add, sub, mul, div) 
		}

	func arithmetric(x, y int64) (add, sub, mul, div int64) {
		add = x + y
		sub = x - y
		mul = x * y
		div = x / y
		return
	}
	 
# Control Sturctures 
if , if else , Else if , Nested If else etc...are same as other c, c++ ,java programing language only difference is it not mandatory to define conditions inside parenthesis ()  , simply 
if x ==5 {
fmt.println("You entered 5")
}

Example Code :
	var choice int64
	fmt.Scan(&choice)
	if choice <= 3 {
		if choice == 1 {
			fmt.Println("Check Balance")
		} else if choice == 2 {
			fmt.Println("Deposit Money")
		} else {
			fmt.Println("withdraw Money")
		}
	} else {
		fmt.Println("Invalid Input")
	}
	
** For Loop ** 
Same as C,C++ and java 
 example :
 for i:=0;i<10;i++{
  fmt.println(i)
  }

Keywords : break , Countinue 

# Switch Case  
	func main() {
	var choice int64
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("1")
		fallthrough <---  Fallthrough is keyword it will execute the next case also , in this code , if choice is 1 it will print 1 as well as 2 also next
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Greater")
	}
}

In go lang no need to add break statement for each case explicitly..


# Packages 
We can use all the functions and attributes across files within the same package.

For example:
package main

Accessing functions and variables across multiple packages:

You need to create a folder.
Note: The folder name will be treated as the package name.

To export a function in Go, it's a bit different. The function name must start with an uppercase letter. This indicates that the function is exported and can be accessed from other packages.

Using Third party packages :
  first we need to install a Package in our project 
   example : go get github.com/Pallinder/go-randomdata
then we need to import in our .go file
 Example Program: 
 package main

import (
	"fmt"
	"example.com/my_first_app/mathops" <--- this is a package that i defined (User defined )
	"github.com/Pallinder/go-randomdata" <--- this is a package downloaded and imported (Third party)
)

func main() {
	fmt.Println(mathops.Add(6, 9))
	fmt.Println(mathops.Sub(6, 9))
	fmt.Println(mathops.Mul(6, 9))
	fmt.Println(mathops.Div(6, 9))
	fmt.Println(randomdata.City())
}

# Pointers 

  Same as  C programing language (&) Single Ampersand , And (*) Asterisk
  Example Program (Passing Argument As Pointers)
		Package main

		import ("fmt")

		func main() {
			var x, y int64
			fmt.Scan(&x)
			fmt.Scan(&y)
			fmt.Println("X and Y Before Swap :", x, y)
			swap(&x, &y)
			fmt.Println("X and Y After Swap :", x, y)
		}

		func swap(x, y *int64) {
			temp := *x
			*x = *y
			*y = temp
		}
		
# Struct & Custom types 

Struct example code :
package main

	import ("fmt")

		type User struct {
			firstname string
			lastname  string
			age       int64
		}

		func main() {
			var fuser User
			fuser.firstname = "Rajesh"
			fuser.lastname = "Nagarajan"
			fuser.age = 89
			fuser.displaystruct()
		}

		func (c User) displaystruct() {
			fmt.Println(c.firstname, c.lastname, c.age)
		}


# Interfaces In Go 

type interface_name interface { <--- Syntax
}

Example : package main

import "fmt"

// Define the interface
type Shape interface {
    Area() float64
}

// Define a struct
type Rectangle struct {
    Width  float64
    Height float64
}

// Rectangle has the Area() method, so it fits the Shape interface
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Another struct
type Circle struct {
    Radius float64
}

// Circle also implements Area() -> fits the Shape interface
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// A function that accepts anything with the Shape interface
func PrintArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    circle := Circle{Radius: 3}

    // Both structs can be used as Shape
    PrintArea(rect)   // Output: Area: 50
    PrintArea(circle) // Output: Area: 28.26
}

# Arrays 
 Var myarr = [4] int64;
 other ->> mrayy:= [4] int64 {1,2,3,4}
 fmt.println(myarr) <-- prints whole array 
 fmt.println(myarr[0]) <--prints specific value using index 
 can able to use slice same as python list -> myarr[1:3]
 *** Dynamic Array ***
 using append built in function 
 prices:=[] int {15,13}
 fmt.println(prices)
 prices= append(prices,25}
 fmt.println(prices)
 
 Example :
 package main

import "fmt"

func main() {
	names := []string{"Rajesh", "Ramesh"}
	fmt.Println(names)
	fmt.Println(len(names))
	names = append(names, "mohan")
	fmt.Print(names)
	fmt.Print((cap(names)))
}
add two array in append ( use ... special syntax)
names =append(names,newnames...)
 
*** Map ***

package main

import "fmt"

func main() {
	x := map[string]string{"AWS": "www.aws.com"} <-- map decalaration
	fmt.Println(x)
	x["Azure"] = "www.Azure.com"
	fmt.Println(x)
	fmt.Println(x["AWS"])
}
 you can also use built in function -> delete()
 delete(x,"Azure") <-- it will delete azure key value pair from x map
 
*** Concurrency ***
Concurrency in Go (often called Golang) refers to the ability of a program to handle multiple tasks at once, making efficient use of system resources like CPU and memory.

In Go, when you want to run a function in the background, you use a goroutine.
 Syntax:
go myFunction()

Needs create a Wait group for wait main function until all the backgrounds completes , otherwise main function terminated i.e all background also terminated

Import sync 

    Step     | Code                  | Meaning
1. Create    | var wg sync.WaitGroup | Initialize the WaitGroup
2. Add count | wg.Add(n)             | "I'm waiting for n tasks"
3. Done      | wg.Done()             | Each task says "I'm done"
4. Wait      | wg.Wait()             | Pause until all Done() called

Sample Program :
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
	wg.Done() <-- its return that background service done 
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	wg.Done()
}

func main() {
	wg.Add(4) <-- we are actually waiitng for 4 background services
	go greet("Nice to meet you!")
	go greet("How are you?")
	go slowGreet("How ... are ... you ...?")
	go greet("I hope you're liking the course!")
	wg.Wait() <-- here its wait to complete all the background services , it holds the main function in waiting stage

}

*** Go REST API (Gin Framework ) ***
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/helloworld", helloworld)
	server.Run(":8080") // localhost:8080
}
func helloworld(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"Message": "helloworld",
	})
}

