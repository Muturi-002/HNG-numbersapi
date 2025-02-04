package main
import (
	"fmt"
	/* "errors"
	"net/http"
	"encoding/json" */
	"strconv"
	"math"
)
var prime, perfect bool
var num, sum, even, odd int
var properties []string
var armstrong string

func Armstrong(number int) string{
	digits:= strconv.Itoa(number)
	digLength:= len(digits)
	for _, digit:= range digits{
		digitInt, _:= strconv.Atoi(string(digit))
		sum+= int(math.Pow(float64(digitInt), float64(digLength)))
	}
	if sum==number{
		armstrong= "armstrong"
	}
	return armstrong
}
func Prime(number int) bool{
	if number <=1{
		prime=false
	}
	if number==2{
		prime=true
	}
	if number%2==0{
		prime=false
	}
	for i:=3; i*1<=number;i+=2{
		if number%i==0{
			prime=false
		}
	}
	fmt.Print("Prime?: ")
	return prime
}
func Perfect(number int) bool{
	if number<=1{
		return false
	}
	sum:=0
	for i:=1; i<=number/2;i++{
		if number%i==0{
			sum+=i
		}
	}
	if sum==number{
		perfect=true
	}else{
		perfect=false
	}
	fmt.Print("Perfect?: ")
	return perfect
}
func SumDigits(number int) int{
	sum=0
	digits:= strconv.Itoa(number)
	for _, digit:= range digits{
		digitInt, _:= strconv.Atoi(string(digit))
		sum+=digitInt
	}
	return sum
}
func EvenOdd(number int) string {
	if number%2 == 0 {
		even = number
		return "Even"
	} else {
		odd = number
		return "Odd"
	}
}
func Properties(number int) []string{
	if Armstrong(number)=="armstrong"{
		properties= append(properties, armstrong)
	}
	if EvenOdd(number)=="Even"{
		properties= append(properties, "Even")
	}else{
		properties= append(properties, "Odd")
	}
	return properties
}
func Xtics (number int){//method calling all declared methods
	prime:=Prime(number)
	perfect:=Perfect(number)
	prop:=Properties(number)
	sum:=SumDigits(number)
	fmt.Print(prime,"\n",perfect,"\n",prop,"\n",sum)
}
func api(number int) {
	// Use API to get a fun facts of a number- http://numbersapi.com/#42

}

func main() {
	var num int
	fmt.Println("Enter a number: ")
	fmt.Scan(&num)
	fmt.Println("The number is:", num,".Check Properties---")
	Xtics(num)
	fmt.Println()
}
