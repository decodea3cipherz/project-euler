package main

import (
	"fmt"
	"project-euler/helpers"
	"strconv"
)

/*
The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.
There are thirteen such primes below 100: 2,3,5,7,11,13,17,31,37,71,73,79,97.
How many circular primes are there below one million?
*/
func compute(n int) int {
	count := 0
	sieve := helpers.SieveOfEratosthenes(n)

outer:
	for i := 2; i < n; i++ {
		num := strconv.Itoa(i)
		for j := 1; j <= len(num); j++ {
			subNum := num[j:len(num)] + num[:j]
			subInt, _ := strconv.Atoi(subNum)
			if !sieve[subInt] {
				continue outer
			}
		}
		count++
	}
	return count
}

func main() {
	fmt.Println(compute(1_000_000))
}
