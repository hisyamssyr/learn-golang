package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
	// "golang.org/x/tools/go/analysis/passes/sortslice"
)

type Expense struct {
	Id			string	`json:"id"`
	Description	string	`json:"desc"`
	Category	string	`json:"cat"`
	Amount		float64	`json:"amount"`
	Date		string	`json:"date"`
}

const filename = "expenses.json"
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for true {
		menuItems := []string{
			"View all expenses",
			"Add new expense",
			"Update an expense",
			"Delete an expense",
			"Show expenses summary",
			"Set monthly budget limit",
			"Export to .csv file",
		}
	
		prompt := promptui.Select{
			Label: "Select Menu",
			Items: menuItems,
		}
	
		index, _, _ := prompt.Run()
		processMenu(index)
	}
}

func processMenu(index int) {
	switch index {
		case 0: view()
		case 1: add()
		case 2: update()
		case 3: delete()
		case 4: summary()
		case 5: limit()
		case 6: export()
	}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func load()([]Expense) {
	data, err := os.ReadFile(filename)
	checkError(err)

	var exp []Expense
	err = json.Unmarshal(data, &exp)
	checkError(err)

	return exp
}

func save(exp []Expense) {
	data, err := json.MarshalIndent(exp, "", " ")
	checkError(err)

	err = os.WriteFile(filename, data, 0664)
	checkError(err)
}

func section() {
	fmt.Println(strings.Repeat("=", 40))
}

func generateId(exp []Expense)(string) {
	idx := len(exp)

	if idx == 0 {return "E0001"}
	lastId := exp[idx - 1].Id[1:]
	intId, _ := strconv.Atoi(lastId)

	return fmt.Sprintf("T%04d", intId + 1)
}

func printExp(exp Expense) {
	section()
	fmt.Println("ID: " + exp.Id)
	fmt.Println("Description: " + exp.Description)
	fmt.Println("Category: " + exp.Category)
	fmt.Println("Amount: " + strconv.FormatFloat(float64(exp.Amount), 'f', 2, 64))
	fmt.Println("Created at: " + exp.Date)
	section()
}

func view() {
	exp := load()

	if len(exp) == 0 {
		section()
		fmt.Println("Expense is empty...")
		section()
	} else {
		sort.Slice(exp, func(i, j int) bool { return exp[i].Date < exp[j].Date})
	}

	for _, val := range exp {
		printExp(val)
	}
}

func add() {
	exp := load()

	section()
	fmt.Println("Add new expense")
	section()

	var ex Expense
	ex.Id = generateId(exp)

	fmt.Println("Input expense description:")
	scanner.Scan()
	ex.Description = scanner.Text()

	fmt.Println("Input expense category:")
	scanner.Scan()
	ex.Category = scanner.Text()
	
	fmt.Println("Input expense amount:")
	scanner.Scan()
	ex.Amount, _ = strconv.ParseFloat(scanner.Text(), 64)
	
	ex.Date = time.Now().Format("31-10-2006")	

	exp = append(exp, ex)
	save(exp)
}

func update() {
	// TODO: search id, show if found, edit desc/amount/category, save 
}

func delete() {
	// TODO: search id, shw if found, del, save
}

func summary() {
	// TODO: show monthly statistic - open monthly detail
}

func limit() {
	// TODO: input limit, can show warning
}

func export() {
	// TODO: export csv
}
