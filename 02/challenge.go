package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	check(err)

	defer file.Close()

	var listOfStrings []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var line = scanner.Text()
		listOfStrings = append(listOfStrings, line)
	}
	return listOfStrings, err
}

func extractParts(line string) (int, int, string, string){

	seperatedLine := strings.Split(line, " ")
	minMax := strings.Split(seperatedLine[0], "-")

	firstNumber, err := strconv.Atoi(minMax[0])
	check(err)

	secondNumber, err := strconv.Atoi(minMax[1])
	check(err)

	letterToSearch := string(seperatedLine[1][0])

	subjectToAnalyse := seperatedLine[2]

	return firstNumber, secondNumber, letterToSearch, subjectToAnalyse
}

func challenge1(listOfStrings []string) int {
	var nbOfValidPwd = 0

	for _, line := range listOfStrings {
		var nbOfOccurence int
		minOccurence, maxOccurence, letter, subject := extractParts(line)
		nbOfOccurence = len(regexp.MustCompile(letter).FindAllStringIndex(subject, -1))

		if nbOfOccurence >= minOccurence && nbOfOccurence <= maxOccurence {
			nbOfValidPwd++
		}
	}
	return nbOfValidPwd
}

func challenge2(listOfStrings []string) int {
	var nbOfValidPwd = 0

	for _, line := range listOfStrings {
		firstNumber, secondNumber, letter, subject := extractParts(line)

		x := letter == string(subject[firstNumber - 1])
		y := letter == string(subject[secondNumber - 1])
		if  (x || y) && !(x && y)  {
			nbOfValidPwd++
		}
	}

	return nbOfValidPwd
}

func main() {

	listOfStrings, err := readLinesFromFile("./02/input.txt")
	check(err)
	fmt.Println(challenge1(listOfStrings))
	fmt.Println(challenge2(listOfStrings))
}
