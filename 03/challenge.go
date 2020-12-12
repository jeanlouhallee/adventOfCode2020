package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func loadGridIntoArray(file string) [][]string {
	var grid = make([][]string, 0)

	rawFileBytes, _ := ioutil.ReadFile(file)
	r := bufio.NewReader(strings.NewReader(string(rawFileBytes)))

	var j = 0
	grid = append(grid, make([]string, 0))
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			stringChar := string(c)

			if(stringChar == "\n") {
				grid = append(grid, make([]string, 0))
				j++
			} else {
				grid[j] = append(grid[j], stringChar)
			}
		}
	}
	return grid
}

func calculateNumberOfTrees(grid [][]string, down int, side int) int {
	var numberOfTrees = 0
	var lengthX = len(grid[0])
	var lengthY = len(grid)

	var i = 0
	var j = 0
	for  {
		if i >= lengthY {
			break
		}
		if j >= lengthX {
			j -= lengthX
		}

		position := grid[i][j]

		if(position == "#") {
			numberOfTrees++
		}
		i += down
		j += side
	}
	return numberOfTrees
}

func challenge1(grid [][]string) int {
	return calculateNumberOfTrees(grid, 1, 3)
}

func challenge2(grid [][]string) int {
	return calculateNumberOfTrees(grid, 1, 1) *
		calculateNumberOfTrees(grid, 1, 3) *
		calculateNumberOfTrees(grid, 1, 5) *
		calculateNumberOfTrees(grid, 1, 7) *
		calculateNumberOfTrees(grid, 2, 1)

}

func main(){
	grid := loadGridIntoArray("./03/input.txt")
	fmt.Println(challenge1(grid))
	fmt.Println(challenge2(grid))

}
