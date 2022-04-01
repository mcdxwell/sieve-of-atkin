# Sieve of Atkin - Go Implementation

Last month, I revisited problem 10: `Summation of primes`, on projecteuler.net. I remembered completing the problem (inefficiently) in Python, and I wanted to try my hand at it with Go. 

And before doing so, I read through the threads section where brilliant programs and ideas flooded my screen. A commenter caught my eye with: 
>"Holy moly the difference. TIL Atkins Algorithm is a thing, very nifty took runtime from 10 minutes to 23 milliseconds."
>
>   --- RampageFillTheRedBar

(Note: They did not post their program, but a special thanks to them for leading me down the path of sieves/sieve theory.)

With the knowledge I gained from reading about sieves, wikis, forums, and the implementations of others, I was able to solve problem 10 efficiently with the `Sieve of Atkin`.  

I didn't want to stop there; I wanted to generate and visualize every prime up to a given limit, rather than just taking the sum of primes up to a given limit (problem 10).

----

## 💻 The Results of Writing Prime Numbers to a File

| Limit (n) | Time (s) | File Size (MB) | # of Primes | Sum of Primes     |
| --------- | -------- | -------------- | ----------- | ----------------- |
| 1e1       | 0.00     | 7.10e-5        | 4           | 17                |  
| 1e2       | 0.00     | 1.38e-4        | 25          | 1060              |  
| 1e3       | 5.06e-4  | 7.13e-4        | 168         | 76127             |
| 1e4       | 1.05e-3  | 6.02e-3        | 1229        | 5736396           |  
| 1e5       | 4.71e-3  | 5.62e-2        | 9592        | 454396537         |  
| 1e6       | 3.69e-2  | 5.38e-1        | 78498       | 37550402023       |  
| 1e7       | 3.23e-1  | 5.22           | 664579      | 3203324994356     |  
| 1e8       | 3.33     | 51.10          | 5761455     | 279209790387276   |  
| 1e9       | 31.26    | 501.96         | 50847534    | 24739512092254535 |


## 🧪 Testing?

I haven't made the time to create tests for this project, but check out
https://primes.utm.edu/howmany.html to compare my numbers and to learn more about primes.


## 👥 Usage

            $ git clone https://github.com/mcdxwell/sieve-of-atkin.git
            $ cd sieve-of-atkin
            $ go run main.go [provide an unsigned integer]
            $ [output]

## 🔢 Example
- Look at ___

## 📝 Notes

- This program may be the first implementation of the Sieve of Atkin in Go.
- There are better ways to write to a file; I am learning. 😉

## 📰 References

https://projecteuler.net/thread=10;page=7

https://en.wikipedia.org/wiki/Sieve_of_Atkin

https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes

https://en.wikipedia.org/wiki/Sieve_theory

https://www.geeksforgeeks.org/sieve-of-atkin/

https://github.com/jmoiron/euler.go/blob/master/010.go