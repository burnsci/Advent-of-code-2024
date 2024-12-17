package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var X string = "X"
var M string = "M"
var A string = "A"
var S string = "S"

func check(e error){
    if e != nil{
        panic(e)
    }
}

func validateTop(i int) bool{
	return i>=3
}
func validateBottom(i int, grid [][]string) bool{
	return i <= len(grid)-4
}
func validateLeft(j int) bool{
	return j>=3
}
func validateRight(j int, row []string) bool{
	return j <= len(row)-4
}

func checkForward(i int, j int, grid [][]string) bool{
	if !validateRight(j, grid[i]){
		return false
	}
	
	if grid[i][j+1] == M{
		if grid[i][j+2] == A{
			if grid[i][j+3] == S{
				return true
			}
		}
	}
	return false
}
func checkReverse(i int, j int, grid [][]string) bool{
	if !validateLeft(j){
		return false
	}

	if grid[i][j-1] == M{
		if grid[i][j-2] == A{
			if grid[i][j-3] == S{
				return true
			}
		}
	}
	return false
}
func checkUp(i int, j int, grid [][]string) bool{
	if !validateTop(i){
		return false
	}

	if grid[i-1][j] == M{
		if grid[i-2][j] == A{
			if grid[i-3][j] == S{
				return true
			}
		}
	}
	return false
}
func checkDown(i int, j int, grid [][]string) bool{
	if !validateBottom(i, grid){
		return false
	}

	if grid[i+1][j] == M{
		if grid[i+2][j] == A{
			if grid[i+3][j] == S{
				return true
			}
		}
	}
	return false
}
func checkDiagonalUpLeft(i int, j int, grid [][]string) bool{
	if !validateLeft(j) || !validateTop(i){
		return false
	}
	if grid[i-1][j-1] == M{
		if grid[i-2][j-2] == A{
			if grid[i-3][j-3] == S{
				return true
			}
		}
	}
	return false
}
func checkDiagonalUpRight(i int, j int, grid [][]string) bool{
	if !validateRight(j, grid[i]) || !validateTop(i){
		return false
	}
	if grid[i-1][j+1] == M{
		if grid[i-2][j+2] == A{
			if grid[i-3][j+3] == S{
				return true
			}
		}
	}
	return false
}
func checkDiagonalDownLeft(i int, j int, grid [][]string) bool{
	if !validateLeft(j) || !validateBottom(i, grid){
		return false
	}
	if grid[i+1][j-1] == M{
		if grid[i+2][j-2] == A{
			if grid[i+3][j-3] == S{
				return true
			}
		}
	}
	return false
}
func checkDiagonalDownRight(i int, j int, grid [][]string) bool{
	if !validateRight(j, grid[i]) || !validateBottom(i, grid){
		return false
	}
	if grid[i+1][j+1] == M{
		if grid[i+2][j+2] == A{
			if grid[i+3][j+3] == S{
				return true
			}
		}
	}
	return false
}

func validBoundary(i int, j int, grid [][]string) bool{
	return (i>0 && i<len(grid)-1 && j>0 && j<len(grid[i])-1)
}

func checkDiagonalTlBr(i int, j int, grid [][]string) bool{
	if !validBoundary(i,j,grid){
		return false
	}
	if grid[i-1][j-1]==M{
		if grid[i+1][j+1]==S{
			return true
		}
	}
	if grid[i-1][j-1]==S{
		if grid[i+1][j+1]==M{
			return true
		}
	}
	return false
}

func checkDiagonalBlTr(i int, j int, grid [][]string) bool{
	if !validBoundary(i,j,grid){
		return false
	}
	if grid[i+1][j-1]==M{
		if grid[i-1][j+1]==S{
			return true
		}
	}
	if grid[i+1][j-1]==S{
		if grid[i-1][j+1]==M{
			return true
		}
	}
	return false
}

func main(){

	file, err := os.Open("./wordSearch.txt")
    check(err)
    
    defer file.Close()

    scanner := bufio.NewScanner(file)

	grid := [][]string{}

	for scanner.Scan(){
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	finds := 0
	xmas := 0
	for i:=0; i<len(grid[0]); i++{
		for j:=0; j<len(grid); j++{
			if grid[i][j] == A{
				if checkDiagonalBlTr(i,j,grid) && checkDiagonalTlBr(i,j,grid){
					xmas++
				}
			}

			if grid[i][j]==X{
				if checkForward(i,j,grid){
					finds++
				}
				if checkReverse(i,j,grid){
					finds++
				}
				if checkUp(i,j,grid){
					finds++
				}
				if checkDown(i,j,grid){
					finds++
				}
				if checkDiagonalUpLeft(i,j,grid){
					finds++
				}
				if checkDiagonalUpRight(i,j,grid){
					finds++
				}
				if checkDiagonalDownLeft(i,j,grid){
					finds++
				}
				if checkDiagonalDownRight(i,j,grid){
					finds++
				}
			}
		}
	}

	fmt.Println(finds)
	fmt.Println(xmas)
}