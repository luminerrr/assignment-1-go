package main

import (
	"fmt"
	"strconv"
	"os"
)

type Student struct {
	id uint16
	name string
	address string
	occupancy string
	reason string
}

func main() {
	students := []Student{
		{
			id:1,
			name:"Jason",
			address:"Street 1",
			occupancy:"Politician",
			reason:"Because he didn't have anything to do",
		},
		{
			id:2,
			name:"Jason",
			address:"Street 2",
			occupancy:"Politician",
			reason:"Because he didn't have anything to do",
		},
		{
			id:3,
			name:"Jason",
			address:"Street 3",
			occupancy:"Politician",
			reason:"Because he didn't have anything to do",
		},
		{
			id:4,
			name:"Jason",
			address:"Street 4",
			occupancy:"Politician",
			reason:"Because he didn't have anything to do",
		},
		{
			id:5,
			name:"Jason",
			address:"Street 5",
			occupancy:"Politician",
			reason:"Because he didn't have anything to do",
		},
	}

	input := os.Args
	val, err := convertInput(input[1])
	selectedStudents := students[val]

	if err == " "{
		fmt.Printf(`error %s`, err)
	} else{
		fmt.Printf(`Name : %s Address : %s Job : %s Reasons for joining Golang : %s`,
		selectedStudents.name, selectedStudents.address, selectedStudents.occupancy, selectedStudents.reason,)

	}
	
	

	
	_ = err
}

func convertInput(input string)( i int, errMsg string) {
	i, err := strconv.Atoi(input)
	if err != nil {
		errMsg = "Error occurred, please try again"
	} else if i >= 6 || i <= 0{
		errMsg = "There's only 5 students. Try again"
	}
	i = i-1
	return
}

