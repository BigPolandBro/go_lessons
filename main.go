package main

import (
	"context"
	"errors"
	"fmt"
	"math"
	"mymath/average"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"unicode"
)

func ConvertCelsiusToFahrenheit(celsius float32) float32 {
	return celsius*1.8 + 32
}

func task1_1() {
	var celsius, fahrenheit float32
	fmt.Scan(&celsius)
	fahrenheit = ConvertCelsiusToFahrenheit(celsius)
	fmt.Printf("# Output: %.0f°C = %.2f°F\n", celsius, fahrenheit)
}

func task1_2() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("Fizz")
		}
		if i%5 == 0 {
			fmt.Printf("Buzz")
		}
		if i%3 != 0 && i%5 != 0 {
			fmt.Printf("%d", i)
		}
		fmt.Printf("\n")
	}
}

func SumDigits(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	if sum < 0 {
		sum *= -1
	}
	return sum
}

func task1_3() {
	var n int
	fmt.Scan(&n)
	fmt.Println(SumDigits(n))
}

type CurrencyConverter struct {
	Rate float64
}

func (ctu CurrencyConverter) ConvertToUSD(rubles float64) float64 {
	return rubles / ctu.Rate
}

func (ctu CurrencyConverter) ConvertToRUB(dollars float64) float64 {
	return dollars * ctu.Rate
}

func task2_1() {
	converter := CurrencyConverter{Rate: 80.5}
	fmt.Println(converter.ConvertToRUB(15))
	fmt.Println(converter.ConvertToUSD(1000))

	converter = CurrencyConverter{Rate: 75.5}
	fmt.Println(converter.ConvertToUSD(755)) // 10.0
}

func FilterEven(numbers ...int) []int {
	var result []int
	for _, number := range numbers {
		if number%2 == 0 {
			result = append(result, number)
		}
	}
	return result
}

func task2_2() {
	//numbers := []{1, 2, 3, 40, 50, -50, -51, 101, 0, 202, 404, 111, -1}
	fmt.Println(FilterEven(1, 2, 3, 40, 50, -50, -51, 101, 0, 202, 404, 111, -1))
}

type BankAccount struct {
	balance float64
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.balance += amount
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if ba.balance < amount {
		return errors.New("Insufficient funds in your bank account.")
	}

	ba.balance -= amount
	return nil
}

func (ba *BankAccount) Balance() float64 {
	return ba.balance
}

func task2_3() {
	ba := BankAccount{balance: 666.0}
	fmt.Println(ba.Balance()) // 666.0
	ba.Deposit(0.666)
	fmt.Println(ba.Balance())         // 666.666
	fmt.Println(ba.Withdraw(334.5))   // nil
	fmt.Println(ba.Balance())         // 332.166
	fmt.Println(ba.Withdraw(332.2))   // error message should be returned here
	fmt.Println(ba.Balance())         // no change in balance
	fmt.Println(ba.Withdraw(332.166)) // no error here
	fmt.Println(ba.Balance())         // bal == 0
}

func ValidatePassword(pass string) error {
	if len(pass) < 8 {
		return errors.New("Password has less than 8 characters!")
	}
	containsDigit := false
	containsCapital := false
	for _, symbol := range pass {
		if unicode.IsDigit(symbol) {
			containsDigit = true
		}
		if unicode.IsUpper(symbol) {
			containsCapital = true
		}
	}

	if !containsCapital {
		return errors.New("Password does not contain capital letter!")
	}

	if !containsDigit {
		return errors.New("Password does not contain digit!")
	}

	return nil
}

func task3_1() {
	password := "abac"
	fmt.Println(ValidatePassword(password))
	password = "kekakeka8"
	fmt.Println(ValidatePassword(password))
	password = "lOlkeklol"
	fmt.Println(ValidatePassword(password))
	password = "loLkek7_rofl"
	fmt.Println(ValidatePassword(password))
}

const MB = 1024 * 1024

func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", errors.New("File does not exist")
	}
	if len(data) > MB {
		return "", errors.New("File's size is bigger than 1MB")
	}
	//fmt.Printf("%T", data)
	return string(data), nil
}

func task3_2() {
	relativePath := "file.txt"
	data, err := ReadFile(relativePath)
	fmt.Println(data, err)

	relativePath = "file_not_existing.txt"
	data, err = ReadFile(relativePath)
	fmt.Println(data, err)

	absolutePath := "/Users/bigpolandbro/Desktop/New_employee_guide_ENG.pdf"
	data, err = ReadFile(absolutePath)
	if err == nil {
		fmt.Println("Successfully read", len(data), "bytes from", absolutePath)
		fmt.Printf("First 10 bytes: %v\n", data[:10])
	} else {
		fmt.Println(err)
	}
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func task4_1() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}
	fmt.Println(circle.Area())
	fmt.Println(rectangle.Area())
}

func task4_2() {
	fmt.Println(average.Average([]float64{1, 2, 3, 4}))
}

// goroutines

func go_greeting(n int, wg *sync.WaitGroup) {
	defer wg.Done() // right before function returns
	time.Sleep(time.Second)
	fmt.Println("Hello from goroutine ", n)
}

func conc_task1_1() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go go_greeting(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers have greeted!")
}

func go_greeting_(n int) {
	time.Sleep(time.Second)
	fmt.Println("Hello from goroutine ", n)
}

func conc_task1_1_() {

	for i := 1; i <= 5; i++ {
		go go_greeting_(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("All workers have greeted!")
}

func sendNumber(n int, ch chan int) {
	ch <- n
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func sendNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		sendNumber(i, ch)
	}
	close(ch)
}

func receiveNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for x := range ch {
		sum += x
		fmt.Println(x)
	}
	fmt.Println("Sum = ", sum)
}

func conc_task1_2_wait_concept() {
	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go receiveNumbers(ch, &wg)
	wg.Add(1)
	go sendNumbers(ch, &wg)

	wg.Wait()
}

func sendNumbers_(ch chan int) {
	for i := 1; i <= 100; i++ {
		sendNumber(i, ch)
	}
	close(ch)
}

func receiveNumbers_(ch chan int) {
	sum := 0
	for x := range ch {
		sum += x
		fmt.Println(x)
	}
	fmt.Println("Sum = ", sum)
}

func conc_task1_2_channels_concept() {
	ch := make(chan int)

	go sendNumbers_(ch)
	receiveNumbers_(ch)
}

func conc_task1_2_channels_concept_2() {
	ch := make(chan int)

	go receiveNumbers_(ch)
	sendNumbers_(ch)
}

func worker(ch chan int, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	for x := range ch {
		if x%2 == 0 {
			fmt.Println("Worker: ", i, " prints ", x)
		}
	}
}

func conc_task2_1() {
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ch, i, &wg)
	}

	wg.Add(1)
	go sendNumbers(ch, &wg)
	wg.Wait()
}

func sendNums(ch chan int, ctx context.Context, wg *sync.WaitGroup) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	defer wg.Done()

	num := 1

	//flag := false

Ticker:
	for range ticker.C {
		fmt.Println("Iteration ")
		// if flag {
		// 	fmt.Println("flag breaks the loop ")
		// 	break
		// }

		select {
		case info := <-ctx.Done():
			fmt.Println("Context cancelled, worker exiting")
			fmt.Println(info)
			close(ch)
			//flag = true
			break Ticker
		default:
			fmt.Println("Before to channel ")
			ch <- num
			fmt.Println("After to channel ")
			num++
		}
	}

	fmt.Println("Worker exits!")
}

func conc_task2_2() {
	ch := make(chan int)
	var wg sync.WaitGroup

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	wg.Add(1)
	go sendNums(ch, ctx, &wg)

	for num := range ch {
		fmt.Println("Received:", num)
	}

	<-ctx.Done()
	fmt.Println("Shutting down...")
	stop()
	fmt.Println("Context canceled")

	wg.Wait()
	fmt.Println("gracefully shutted down mthfckr!!!")
}

func process_err(err error) {
	if err != nil {
		panic(err)
	}
}

func req_avito() int64 {
	baseUrl := "https://www.avito.ru/moskva"
	query := "лабубу+coca-cola+оригинал"

	encodedQuery := url.QueryEscape(query)
	url := fmt.Sprintf("%s?q=%s", baseUrl, encodedQuery)
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	process_err(err)

	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 YaBrowser/25.4.0.0 Safari/537.36")

	start := time.Now()
	response, err := client.Do(request)
	process_err(err)

	defer response.Body.Close()

	duration := time.Since(start)

	if response.StatusCode != http.StatusOK {
		fmt.Println("Avito request error ", response.StatusCode)
		return duration.Milliseconds()
	}

	// body, err := ioutil.ReadAll(response.Body)
	// process_err(err)

	// fmt.Println(string(body))
	return duration.Milliseconds()
}

func req_common(url string) int64 {
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	process_err(err)

	start := time.Now()
	response, err := client.Do(request)
	process_err(err)

	defer response.Body.Close()

	duration := time.Since(start)

	if response.StatusCode != http.StatusOK {
		fmt.Println(url, " request error ", response.StatusCode)
		return duration.Milliseconds()
	}

	// body, err := ioutil.ReadAll(response.Body)
	// process_err(err)

	// fmt.Println(string(body))
	return duration.Milliseconds()
}

func conc_task3_real() {
	chans := make([]chan int64, 5)
	for i := range chans {
		chans[i] = make(chan int64)
	}

	go func() {
		chans[0] <- req_avito()
	}()

	go func() {
		url := "https://www.google.com/search?q=labubu"
		chans[1] <- req_common(url)
	}()

	go func() {
		url := "https://example.com"
		chans[2] <- req_common(url)
	}()

	go func() {
		url := "https://yandex.ru"
		chans[3] <- req_common(url)
	}()

	go func() {
		url := "https://www.bing.com"
		chans[4] <- req_common(url)
	}()

	for i := 0; i < 5; i++ {
		select {
		case dur := <-chans[0]:
			fmt.Println("Avito req duration, ms: ", dur)
		case dur := <-chans[1]:
			fmt.Println("Google req duration, ms: ", dur)
		case dur := <-chans[2]:
			fmt.Println("Example req duration, ms: ", dur)
		case dur := <-chans[3]:
			fmt.Println("Yandex req duration, ms: ", dur)
		case dur := <-chans[4]:
			fmt.Println("Bing req duration, ms: ", dur)
		}
	}
}

func main() {
	fmt.Println("This is to check bugfix branch")
	conc_task3_real()
}
