package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error){
    if e != nil{
        panic(e)
    }
}

func main(){

	dat, err := os.ReadFile("corruptedMemory.txt")
	check(err)

	corruptedMemory := string(dat)

	regX, err := regexp.Compile(`mul\([0-9]{1,3}\,[0-9]{1,3}\)|do\(\)|don't\(\)`)
	//regX, err := regexp.Compile("mul\\([0-9]{1,3}\\,[0-9]{1,3}\\)")
	check(err)

	matches := regX.FindAllString(corruptedMemory, -1)

	total := 0
	active := true

	for _,item := range matches {
		if strings.HasPrefix(item, "mul"){
			if active{
				//strip 'mul(' & ')'
				snip, _ := strings.CutPrefix(item, "mul(")
				snip, _ = strings.CutSuffix(snip, ")")
				//split on ', '
				numbers := strings.Split(snip, ",")
				//calculate
				a, _ := strconv.Atoi(numbers[0])
				b, _ := strconv.Atoi(numbers[1])
				total = total + (a*b)
			}
		}else if strings.HasPrefix(item, "do()"){
			active = true
		}else if strings.HasPrefix(item, "don't()"){
			active = false
		}
	}

	fmt.Println("Total: ", total)


}