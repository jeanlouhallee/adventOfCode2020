package main

import (
	"bufio"
	"fmt"
	"os"
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

func passportRespectsCriterias(passport map[string]string) bool{
	criterias := []string{"ecl", "pid", "eyr", "hcl", "iyr", "byr", "hgt"}

	for _, criteria := range criterias {
		_, exists := passport[criteria]
		if exists == false {
			return false
		}
	}
	return true
}

func formatPassports(rawPassports []string) []map[string]string {
	passports:= make([]map[string]string, 0)

	var m = make(map[string]string)
	passports = append(passports, m)
	for _, line := range rawPassports {
		if line == "" {
			m = make(map[string]string)
			passports = append(passports, m)
		} else {
			a := strings.Fields(line)
			for _, element := range a {
				pair := strings.Split(element, ":")
				m[pair[0]] = pair[1]
			}
		}
	}
	return passports
}

func countNbOfValidPassports(passports []map[string]string) int {
	var nbOfValidPassports = 0
	for _, passport := range passports {
		fmt.Println(passport)
		if passportRespectsCriterias(passport) {
			nbOfValidPassports++
		}
	}
	return nbOfValidPassports
}

func main() {
	rawPassports, _:= readLinesFromFile("./04/input.txt")
	nbOfValidPassports := countNbOfValidPassports(formatPassports(rawPassports))
	fmt.Println(nbOfValidPassports)
}
