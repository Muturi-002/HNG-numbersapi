package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"encoding/json"
	_ "github.com/gin-gonic/gin"
)
type Number struct{
	Num int `json:"number"`
	Prime bool `json:"is_prime"`
	Perfect bool `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum int `json:"digit_sum"`
	API_Fact string `json:"fun_fact"`
}
var prime, perfect, err bool
var num, sum, even, odd int
var properties []string
var armstrong string

func api(number int) {
	// Use API to get a fun facts of a number- http://numbersapi.com/#42
	url := fmt.Sprintf("http://numbersapi.com/%d/math", number)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:",err)
		return
	}
	defer response.Body.Close()
	
}
func Armstrong(number int) string {
	digits := strconv.Itoa(number)
	digLength := len(digits)
	for _, digit := range digits {
		digitInt, _ := strconv.Atoi(string(digit))
		sum += int(math.Pow(float64(digitInt), float64(digLength)))
	}
	if sum == number {
		armstrong = "armstrong"
	}
	return armstrong
}
func Prime(number int) bool {
	if number <= 1 {
		prime = false
	}
	if number == 2 {
		prime = true
	}
	if number%2 == 0 {
		prime = false
	}
	for i := 3; i*1 <= number; i += 2 {
		if number%i == 0 {
			prime = false
		}
	}
	return prime
}
func Perfect(number int) bool {
	if number <= 1 {
		return false
	}
	sum := 0
	for i := 1; i <= number/2; i++ {
		if number%i == 0 {
			sum += i
		}
	}
	if sum == number {
		perfect = true
	} else {
		perfect = false
	}
	return perfect
}
func SumDigits(number int) int {
	sum = 0
	digits := strconv.Itoa(number)
	for _, digit := range digits {
		digitInt, _ := strconv.Atoi(string(digit))
		sum += digitInt
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
func Properties(number int) []string {
	if Armstrong(number) == "armstrong" {
		properties = append(properties, armstrong)
	}
	if EvenOdd(number) == "Even" {
		properties = append(properties, "Even")
	} else {
		properties = append(properties, "Odd")
	}
	return properties
}

func main() {
	var num Number
	fmt.Println("Enter a number: ")
	fmt.Scan(&num.Num)
	if num.Num <=0 {
		fmt.Println("Please enter a positive number")
		return
	}else{
		num.Prime = Prime(num.Num)
		num.Perfect= Perfect(num.Num)
		num.DigitSum = SumDigits(num.Num)
		num.Properties = Properties(num.Num)
		response := map[string]interface{}{
			"number": num.Num,
			"is_prime": num.Prime,
			"is_perfect": num.Perfect,
			"properties": num.Properties,
			"digit_sum":num.DigitSum,
		}
		jsonResponse, _ := json.MarshalIndent(response,"","\t")
		json.Unmarshal(jsonResponse, &response)
		fmt.Println(string(jsonResponse))
	}
	
}
