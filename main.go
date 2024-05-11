package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	id 		int
	name	string
	address	string
	job		string
	reason 	string
	
}

func main () {

	var students = []Student {
		{
			id : 1,
			name: "Dwi Pasca",
			address: "Palu, Sulawesi Tengah",
			job: "Frontend Dev",
			reason: "want to become full stack and the demand for Golang is very high",
		},
		{
			id : 2,
			name: "Jane Shepard",
			address: "Earth",
			job: "Guardian",
			reason: "I'm commander shepard not programmer",
		},
		{
			id : 3,
			name: "Benezia",
			address: "Coral Island",
			job: "Gardener",
			reason: "never enroll to golang class",
		},
	}

	var names = []string{}
	var ids = []int{}
	var foundStudent []Student 

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Search data by Name or ID : ")
	search, _ := reader.ReadString('\n')
	search = strings.TrimSpace(search)

	if search == "" {
		fmt.Println("Please insert name or id student")
		fmt.Println("Example: Type Dwi Pasca or 1 in your keyboard")
		return 
	}
	// convert string to int
	searchId, _ := strconv.Atoi(search)

	for _, std := range students {
		names = append(names, std.name)
		ids = append(ids, std.id)
	}

	for index, name := range names {
		if (name == search) {
			foundStudent = append(foundStudent, students[index])
		}
	}

	for index, id := range ids {
		if (id == searchId) {
			foundStudent = append(foundStudent, students[index])
		}
	}
	
	printResultDataSearch(foundStudent)
}

func printResultDataSearch (foundStudent []Student) {
	if len(foundStudent) > 0 {
		fmt.Println("\nResult : ")
		fmt.Println(strings.Repeat("-", 20))
		for _, student := range foundStudent {
			fmt.Println("ID : ", student.id)
			fmt.Println("Name : ", student.name)
			fmt.Println("Job : ", student.job)
			fmt.Println("Address : ", student.address)
			fmt.Println("Reason : ", student.reason)
		}
		fmt.Println( strings.Repeat("-", 20))
	} else {
		fmt.Println(strings.Repeat("-", 20))
		fmt.Println("No student found with the Name or ID you type")
		fmt.Println(strings.Repeat("-", 20))
	}
}