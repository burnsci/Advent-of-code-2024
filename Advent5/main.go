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

func getMiddleNumber(numbers []string) int{
	middleIndex := int((float64(len(numbers))/2)-0.5)
	middleNumberStr := numbers[middleIndex]
	middleNumberInt, _ := strconv.Atoi(middleNumberStr)
	return middleNumberInt
}

func splitRule(rule string) []string{
	return strings.Split(rule, "|")
}

func getIndex(numbers []string, value string) int{
	for index, current := range numbers{
		if current == value{
			return index
		}
	}
	return -1
}

func revertToOrder(numbers []string)string{
	order := ""
	for _, value := range numbers{
		if order == ""{
			order = value
		}else{
			order = order + "," + value
		}
	}
	return order
}

func fixNumbers(order string, rules []string) []string{
	modified := false
	fixedNumbers := strings.Split(order, ",")
	//fixNumberMarker:
	for _, rule := range rules{
		if !applyRule(rule, revertToOrder(fixedNumbers)){
			ruleComponents := splitRule(rule)
			posA := getIndex(fixedNumbers, ruleComponents[0])
			posB := getIndex(fixedNumbers, ruleComponents[1])
			temp := fixedNumbers[posA]
			fixedNumbers[posA] = fixedNumbers[posB]
			fixedNumbers[posB] = temp
			modified = true
			//break fixNumberMarker
		}
	}
	if modified{
		fixedNumbers = fixNumbers(revertToOrder(fixedNumbers), rules)
	}
	return fixedNumbers
}

func applyRule(rule string, order string) bool{
	numbers := strings.Split(order, ",")

	ruleComponents := splitRule(rule)
	if !strings.Contains(order, ruleComponents[0]) || !strings.Contains(order, ruleComponents[1]){
		return true //if it does not contain one part of the rule then it is a valid rule
	}
	posA, posB := 0,0
	for index, value := range numbers{
		if value == ruleComponents[0]{
			posA = index
		}
		if value == ruleComponents[1]{
			posB = index
		}
	}
	return posA<=posB
}

func main(){
	file, err := os.Open("./printing.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := []string{}
	orders := []string{}

	total := 0
	wrongTotals := 0

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
		valid := true
		//nextOrderMarker:
		for _, number := range numbers{
			for _, rule := range rules{
				if strings.Contains(rule, number){
					applicableRules = append(applicableRules, rule)
				}
			}
			for _, applicableRule := range applicableRules{
				if !applyRule(applicableRule, order) {
					valid = false
					//break nextOrderMarker
				}
			}	
		}

		if valid{
			total=total + getMiddleNumber(numbers)
		}else if !valid{
			wrongTotals = wrongTotals + getMiddleNumber(fixNumbers(order, applicableRules))
		}
		valid = true
	}

	fmt.Println("regular total", total)
	fmt.Println("corrected middle total: ", wrongTotals)
}