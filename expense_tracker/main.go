package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

type Expense struct {
	Id			string	`json:"id"`
	Description	string	`json:"desc"`
	Category	string	`json:"cat"`
	Amount		float64	`json:"amount"`
	Date		string	`json:"date"`
}

// TODO: Class for menu list

const filename = "expenses.json"
var scanner = bufio.NewScanner(os.Stdin)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func pause() {
	fmt.Println("\nTekan ENTER untuk lanjut...")
	fmt.Scanln()
}

func exitProgram() {
	fmt.Println("Program is closed...")
	os.Exit(0)
}

func load() []Expense {
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

func generateId(exp []Expense) string {
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
	fmt.Println("Last update: " + exp.Date)
	section()
}

func parseMoney(val string) (bool, float64) {
	num, err := strconv.ParseFloat(val, 64)

	if err != nil || num < 0.0 {
		return false, 0
	} else {
		return true, num
	}
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
	pause()
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

	_, ex.Category = showMenu("Select Expense Category", categoryMenu())

	fmt.Println("Input expense amount:")
	for true {
		scanner.Scan()
		isValid, val := parseMoney(scanner.Text())

		if isValid {
			ex.Amount = val
			break
		} else {
			fmt.Println("Value doesn't valid, must an positive integer...")
		}

	}

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

func checkWarning() {
	// TODO: warning alert
}

func mainMenu() []string {
	menuItems := []string{
		"View all expenses",
		"Add new expense",
		"Update an expense",
		"Delete an expense",
		"Show expenses summary",
		"Set monthly budget limit",
		"Export to .csv file",
		"Exit",
	}

	return menuItems
}

func categoryMenu() []string {
	menuItems := []string{
		"Food & Drinks",   // makan, minum, nongkrong
		"Transportation",  // bensin, tiket, parkir
		"Utilities",       // listrik, air, internet
		"Housing",         // sewa, cicilan rumah
		"Entertainment",   // film, game, hobi
		"Shopping",        // belanja kebutuhan, pakaian
		"Health",          // obat, dokter, gym
		"Education",       // kursus, buku
		"Travel",          // liburan, hotel
		"Others",          // kategori tambahan
	}

	return menuItems
}

func showMenu(label string, menu []string) (int, string) {
	prompt := promptui.Select{
		Label: label,
		Items: menu,
	}
	idx, choice, _ := prompt.Run()

	return idx, choice
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
		case 7: exitProgram()
	}
}

func main() {
	for true {
		clearScreen()

		menu, _ := showMenu("SELECT MENU", mainMenu())
		processMenu(menu)
	}

	add()
}