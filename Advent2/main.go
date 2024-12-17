package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error){
    if e != nil{
        panic(e)
    }
}

func convert(items[] string) []int{
    numbers := []int{}
    for _, item := range items{
        value,err := strconv.Atoi(item)
        check(err)
        numbers = append(numbers, value)
    }
    return numbers
}

func acceptableDifference(a int, b int) bool{
    diff := a-b
    if diff>=-3 && diff<=3 && diff!=0{
        return true
    }
    return false
}

func isSafe(items[] int) bool{
    if(items[0] < items[1]){
        for index,item :=range items{
            if index<len(items)-1{
                if item>items[index+1] || !acceptableDifference(item, items[index+1]){
                    return false
                }
            }
        }
    }else{
        for index,item :=range items{
            if index<len(items)-1{
                if item<items[index+1] || !acceptableDifference(item, items[index+1]){
                    return false
                }
            }
        }
    }
    return true
}

func isRecoverable(items[] int) bool{
    totalNewSlices := len(items)
    for i:=0; i<totalNewSlices; i++{
        emptySlice := []int{}
        for j:=0; j<totalNewSlices;j++{
            if i != j{
                emptySlice = append(emptySlice, items[j])
            }
        }
        if isSafe(emptySlice){
            return true
        }
    }
    return false
}

func main() {
    file, err := os.Open("./inputAdvent2.txt")
    check(err)
    
    defer file.Close()

    scanner := bufio.NewScanner(file)
    totalSafeRecords :=0

    for scanner.Scan(){
        items := convert(strings.Split(scanner.Text(), " "))
        if isSafe(items){
            totalSafeRecords++
        }else if isRecoverable(items){
            totalSafeRecords++
        }
    }

    fmt.Println(totalSafeRecords)
    fmt.Println("end")
}
