package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func readIntsFromFile(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// O(n^3) ... not good
func challenge_1(listOfInts []int, desiredResult int) int {

	for i, numberX := range listOfInts {
		for j, numberY := range listOfInts {
			for k, numberZ := range listOfInts {
				if i != j && i != k {
					if numberX + numberY + numberZ == desiredResult {
						return numberX * numberY * numberZ
					}
				}
			}
		}
	}
	return -1
}

// O(n^2) ... not good
func challenge_2(listOfInts []int, desiredResult int) int {

	for i, numberX := range listOfInts {
		for j, numberY := range listOfInts {
			if i != j {
				if numberX + numberY== desiredResult {
					return numberX * numberY
				}
			}
		}
	}
	return -1
}

func main() {
	file, err := os.Open("./01/input.txt")
	check(err)

	ints, err := readIntsFromFile(bufio.NewReader(file))

	fmt.Println(challenge_1(ints, 2020))
	fmt.Println(challenge_2(ints, 2020))
}
