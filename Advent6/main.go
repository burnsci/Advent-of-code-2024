package main

import (
	"bufio"
	"os"
	"strings"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

func main(){
    file, err := os.Open("./printing.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

    maze := [][]string{}

    for scanner.Scan(){
		row := scanner.Text()
        mazeRow := strings.Split(row, "")
		maze = append(maze, mazeRow)
	}

    
}