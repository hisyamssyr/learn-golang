package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID     string	`json:"id"`
	Desc   string	`json:"desc"`
	Status string	`json:"status"`
	Create string	`json:"created_at"`
	Update string	`json:"updated_at"`
}

const filename = "tasks.json"
var scanner = bufio.NewScanner(os.Stdin)

func main() {

	for true {

		showMenu()

		var choice int
		fmt.Println("Enter your choice (1 - 9): ")
		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		checkError(err)

		process(choice)
	}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func showMenu() {
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("              TASK TRACKER")
	fmt.Println(strings.Repeat("=", 40))

	fmt.Println("Menu:")
	fmt.Println("1. Add new task")
	fmt.Println("2. Edit existing task")
	fmt.Println("3. Delete existing task")
	fmt.Println("4. Change task's status")
	fmt.Println("5. Show all tasks")
	fmt.Println("6. Show all done tasks")
	fmt.Println("7. Show all on progress tasks")
	fmt.Println("8. Show all not done tasks")
	fmt.Println("9. Close program")
}

func process(menu int) {
	switch menu {
	case 1: addTask()
	case 2: editTask()
	case 3: deleteTask()
	case 4: changeStatus()
	case 5: showAll()
	case 6: showDone()
	case 7: showOnProgress()
	case 8: showNotDone()
	case 9: close()
	}
}

func loadTasks()([]Task) {
	data, err := os.ReadFile(filename)
	checkError(err)

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	checkError(err)

	return tasks
}

func saveFile(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", " ")
	checkError(err)

	err = os.WriteFile(filename, data, 0644)
	checkError(err)
}

func getId(tasks []Task)(string) {
	length := len(tasks)

	if length == 0 {return "T0001"}
	prev := tasks[length - 1].ID
	num := prev[1:]
	n, _ := strconv.Atoi(num)

	return fmt.Sprintf("T%04d", n + 1)
}

func getIndex(tasks []Task, target string)(int) {
	for i, v := range tasks {
		if v.ID == target {
			return i
		}
	}
	return -1
}

func printTask(ts Task) {
	fmt.Println("UID Task: " + ts.ID)
	fmt.Println("Description: " + ts.Desc)
	fmt.Println("Status: " + ts.Status)
	fmt.Println("Created at: " + ts.Create)
	fmt.Println("Last Updated at: " + ts.Update + "\n")
}

func addTask() {
	tasks := loadTasks()
	var ts Task

	fmt.Printf("Input task description: ")
	scanner.Scan()
	ts.Desc = scanner.Text()

	ts.ID = getId(tasks)

	ts.Status = "Pending"
	ts.Create = time.Now().Format("02-01-06 15:04:05")
	ts.Update = time.Now().Format("02-01-06 15:04:05")

	tasks = append(tasks, ts)
	saveFile(tasks)
}

func editTask() {
	tasks := loadTasks()

	fmt.Println("Enter the task's UID to edit")
	scanner.Scan()
	uid := scanner.Text()

	index := getIndex(tasks, uid)
	if index >= 0 {
		 printTask(tasks[index])
	} else {
		fmt.Println("Look-Like UID Doesn't exist")
		return
	}
	fmt.Print("Option:\n 1. Edit task description\n 2. Edit task status\n")
	fmt.Println("Enter your choice: ")
	scanner.Scan()
	cho := scanner.Text()

	if cho == "1" {
		fmt.Println("Input new description")
		scanner.Scan()
		tasks[index].Desc = scanner.Text()
	} else {
		fmt.Println("Input new status")
		scanner.Scan()
		tasks[index].Status = scanner.Text()
	}

	saveFile(tasks)
	fmt.Println("Change saved successfully!!!")
}

func deleteTask() {
	tasks := loadTasks()

	fmt.Println("Enter the task's UID to delete")
	scanner.Scan()
	uid := scanner.Text()

	index := getIndex(tasks, uid)
	if index >= 0 {
		 printTask(tasks[index])
	} else {
		fmt.Println("Look-Like UID Doesn't exist")
		return
	}

	fmt.Println("Are you sure, you want to delete this task?")
	fmt.Print(" 1. Yes\n 2. No\nEnter your choice: ")
	scanner.Scan()
	cho := scanner.Text()

	if cho == "1" {
		tasks = append(tasks[:index], tasks[index+1:]...)
		saveFile(tasks)
		fmt.Println("Task deleted successfully!!!")
	} else {
		fmt.Println("Canceled")
	}
}

func changeStatus() {
	tasks := loadTasks()

	fmt.Println("Enter the task's UID to change the status")
	scanner.Scan()
	uid := scanner.Text()

	index := getIndex(tasks, uid)
	if index >= 0 {
		 printTask(tasks[index])
	} else {
		fmt.Println("Look-Like UID Doesn't exist")
		return
	}

	fmt.Println("Enter new status: ")
	scanner.Scan()
	tasks[index].Status = scanner.Text()

	saveFile(tasks)
	fmt.Println("Successfully made a change")
}

func showAll() {
	tasks := loadTasks()
	
	for _, v := range tasks {
		printTask(v)
	}
}

func showDone() {
	tasks := loadTasks()

	for _, v := range tasks {
		if v.Status == "Done" {printTask(v)}
	}
}

func showOnProgress() {
	tasks := loadTasks()

	for _, v := range tasks {
		if v.Status == "On Progress" {printTask(v)}
	}
}

func showNotDone() {
	tasks := loadTasks()

	for _, v := range tasks {
		if v.Status == "Pending" {printTask(v)}
	}
}

func close() {
	println("Program closed.")
	os.Exit(0)
}