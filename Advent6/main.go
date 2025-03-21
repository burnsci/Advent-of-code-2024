package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func check(e error){
	if e != nil{
		panic(e)
	}
}


func main(){
    file, err := os.Open("./maze.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

    maze := [][]string{}

    for scanner.Scan(){
		row := scanner.Text()
        mazeRow := strings.Split(row, "")
		maze = append(maze, mazeRow)
	}

	mz := Maze{maze: maze}
	guardPos := guardPosition{}
	for i, mazeRow:=range mz.maze{
		for j, mazeCell := range mazeRow{
			if mazeCell == "^"{
				guardPos = *newGuardPosition(i,j,mazeCell)
			}
		}
	}

	for ;!guardPos.patrolComplete;{
		guardPos.move(mz)
		mz.updateVisit(guardPos)
	}

	fmt.Println("Unique visits: ", mz.countUniqueVisits())
}