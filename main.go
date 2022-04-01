package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func sieveOfAtkins(ch chan int, limit int) {

	sieve := make([]bool, limit)
	for i := 0; i < limit; i++ {
		sieve[i] = false
	}

	var n int
	for x := 1; x*x <= limit; x++ {
		for y := 1; y*y <= limit; y++ {
			n = (4 * x * x) + (y * y)
			if n <= limit && (n%12 == 1 || n%12 == 5) {
				sieve[n] = !sieve[n]
			}

			n = (3 * x * x) + (y * y)
			if n <= limit && n%12 == 7 {
				sieve[n] = !sieve[n]
			}

			n = (3 * x * x) - (y * y)
			if x > y && n <= limit && n%12 == 11 {
				sieve[n] = !sieve[n]
			}
		}
	}

	for r := 5; r*r < limit; r++ {
		if sieve[r] {
			for i := r * r; i < limit; i += r * r {
				sieve[i] = false
			}
		}
	}

	ch <- 2
	ch <- 3
	for a := 5; a < limit; a++ {
		if sieve[a] {
			ch <- a
		}
	}
	close(ch)
}

func main() {

	input := os.Args[1:]
	num := strings.Join(input[:], "")
	limit, err := strconv.Atoi(num)

	if err != nil {
		log.Fatal(err)
	}

	if limit <= 5 {
		fmt.Println("Select a limit greater than 5")
		os.Exit(1)
	}

	fileName := num + ".txt"
	f, err := os.Create(fileName)

	if err != nil {
		fmt.Printf("Error creating file: %v", err)
		return
	}

	defer f.Close()
	//primers := make([]uint64, 0)
	primeCount := 0
	primeSum := uint64(0)
	ch := make(chan int)

	start := time.Now()
	go sieveOfAtkins(ch, limit)
	for prime := range ch {
		primeCount += 1
		primeSum += uint64(prime)
		//primers = append(primers, uint64(prime))
		_, err = f.WriteString(fmt.Sprintf("%v ", uint64(prime)))
		if err != nil {
			fmt.Printf("Error writing string: %v", err)
		}
	}

	info := "\nFind all primes up to %s\nNumber of primes: %d\nSum of primes: %d"
	s := fmt.Sprintf(info, num, primeCount, primeSum)
	f.WriteString(s)
	elapsed := time.Since(start)
	fmt.Println(s, "Time since start: ", elapsed)
}
