package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func sieveOfAtkin(ch chan int, limit int) {

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
		log.Fatalf("Provide an unsigned integer. Error: %v", err)
	}

	if limit <= 5 {
		log.Fatal("Select a limit greater than 5!")
	}

	fileName := num + ".txt"
	f, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("Error creating file: %v", err)
		return
	}

	defer f.Close()

	primes := make([]uint64, 0)
	primeSum := uint64(0)
	ch := make(chan int)

	start := time.Now()

	go sieveOfAtkin(ch, limit)
	for prime := range ch {
		primeSum += uint64(prime)
		primes = append(primes, uint64(prime))
	}

	info := "\nFind all primes up to %d\nNumber of primes: %d\nSum of primes: %d"

	s := fmt.Sprintf(info, limit, len(primes), primeSum)
	_, err = f.WriteString(fmt.Sprintf("%v %v", primes, s))

	if err != nil {
		log.Fatalf("Error writing string to file. Error: %v", err)
	}

	elapsed := time.Since(start)
	fmt.Printf("%v \nTime since start: %v\n\n", s, elapsed)
}
