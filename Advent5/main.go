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

	rules := []string{}
	orders := []string{}

	for scanner.Scan(){
		row := scanner.Text()
		if strings.Contains(row, "|"){
			rules = append(rules, row)
		}else if strings.Contains(row, ","){
			orders = append(orders, row)
		}
	}

	for _, order := range orders{
		numbers := strings.Split(order, ",")
		applicableRules := []string{}
		
		for _, number := range numbers{
			for _, rule := range rules{
				if strings.Contains(rule, number){
					applicableRules = append(applicableRules, rule)
				}
			}
		}

		//for each applicable rule
		//split on |
		//find index of each number in the order
		//if index1>index2 invalid order
		//if all rules valid add middle number to total
	}
	
}