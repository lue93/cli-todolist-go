package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type TodoItem struct {
	name        string
	description string
	status      string
}

type TodoItems struct {
	id    int
	items map[int]TodoItem
}

func (tds *TodoItems) size() int {
	return len(tds.items)
}

func (tds *TodoItems) idAutoIncrement() int {
	tds.id = tds.id + 1
	return tds.id
}

func (tds *TodoItems) add(tdi *TodoItem) {
	var position = tds.idAutoIncrement()
	tds.items[position] = *tdi
}
func (tds *TodoItems) rm(id int) {
	delete(tds.items, id)
}

func (tds *TodoItems) get(id int) (tdi *TodoItem) {
	if item, ok := tds.items[id]; ok {
		return &item
	}
	return nil
}

func (tds *TodoItems) set(position int, tdi *TodoItem) {
	tds.items[position] = *tdi
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func initTodoItemsOnMemory() *TodoItems {
	itemsOnMem := TodoItems{
		id:    0,
		items: make(map[int]TodoItem),
	}
	return &itemsOnMem
}

func printMainMenuActions() {
	fmt.Println("🧭 - CLI - TODO LIST")
	fmt.Println("A - ➕ Add new todo item to list")
	fmt.Println("D - 🗑️ Delete a todo item from list")
	fmt.Println("E - 📝 Edit a todo item from list")
	fmt.Println("S - 🔄 Change status of a todo item from list (🔵, 🟢, 🟡, 🔴)")
	fmt.Println("L - 🔍 View the todo list")
}

func handleMainMenuActions() (string, error) {
	set := map[string]bool{
		"A":       true,
		"D":       true,
		"E":       true,
		"S":       true,
		"L":       true,
		"CLEAR":   true,
		"RUNTIME": true,
	}

	var action string
	fmt.Print("Choice a action above (A,D,E,S,L): ")
	fmt.Scanln(&action)
	if set[strings.ToUpper(action)] {
		return action, nil
	} else {
		return action, fmt.Errorf("Invalid action")
	}
}

func main() {

	var todoItems = initTodoItemsOnMemory()

	printMainMenuActions()
	for { // Loop infinito (como while(true)) em java
		action, error := handleMainMenuActions()
		if error != nil {
			clearScreen()
			fmt.Println("Unknown action")
			printMainMenuActions()
		} else {
			switch action {
			case "A":
				fmt.Println("Adding a new todo item to list")

				var name string
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Give a name for this task:")
				name, _ = reader.ReadString('\n')

				var description string
				fmt.Print("Give a description for this task:")
				description, _ = reader.ReadString('\n')

				name = strings.ReplaceAll(name, "\n", "")
				description = strings.ReplaceAll(description, "\n", "")

				todoItems.add(&TodoItem{
					name:        name,
					description: description,
					status:      "🔵TODO",
				})
				fmt.Println("Added a new todo item to list")
			case "D":
				fmt.Println("Deleting a todo item from list ")
				var id int
				fmt.Print("Type id for delete the task:")
				id, error = fmt.Scanf("%d", &id)
				if error != nil {
					fmt.Println(fmt.Errorf("Invalid id"))
				} else {
					todoItems.rm(id)
					fmt.Println("Delete todo from list ")
				}
			case "E":
				fmt.Println("Editing a todo item from list")
				var id int
				fmt.Print("Type id for edit the task:")
				id, error = fmt.Scanf("%d", &id)
				if error != nil {
					fmt.Println(fmt.Errorf("Invalid id"))
				} else {
					var todoItem = todoItems.get(id)

					var updateName string
					reader := bufio.NewReader(os.Stdin)
					fmt.Print("Update the name for this task:")
					updateName, _ = reader.ReadString('\n')

					var updateDescription string
					fmt.Print("Update the description for this task:")
					updateDescription, _ = reader.ReadString('\n')

					todoItem.name = strings.ReplaceAll(updateName, "\n", "")
					todoItem.description = strings.ReplaceAll(updateDescription, "\n", "")

					todoItems.set(id, todoItem)
				}
			case "S":
				fmt.Println("Changing a todo item state")
				fmt.Println("You can change between TODO, DOING and DONE states")

				var id int
				fmt.Print("Type id for change state of the task:")
				id, error = fmt.Scanf("%d", &id)
				if error != nil {
					fmt.Println(fmt.Errorf("Invalid id"))
				} else {
					var todoItem = todoItems.get(id)

					var state string
					fmt.Print("Choice a state (TODO, DOING or DONE): ")
					fmt.Scanln(&action)
					switch state {
					case "TODO":
						var oldState = todoItem.status
						todoItem.status = "🔵TODO"
						fmt.Printf("change state from %s to TODO", oldState)
					case "DOING":
						var oldState = todoItem.status
						todoItem.status = "🟡DOING"
						fmt.Printf("change state from %s to DOING", oldState)
					case "DONE":
						var oldState = todoItem.status
						todoItem.status = "🟢DONE"
						fmt.Printf("change state from %s to DONE", oldState)
					case "CANCELLED":
						var oldState = todoItem.status
						todoItem.status = "🔴CANCELLED"
						fmt.Printf("change state from %s to CANCELLED", oldState)
					default:
						fmt.Println("Invalid state")

					}
					todoItems.set(id, todoItem)

				}
			case "L":
				fmt.Println("Listing the todo items from list ")
				for i := 1; i <= todoItems.size(); i++ {
					var todoItem = todoItems.get(i)
					fmt.Println("id: ", i)
					fmt.Println("name: ", todoItem.name)
					fmt.Println("description: ", todoItem.description)
					fmt.Println("status: ", todoItem.status)
					fmt.Println("")

				}
				fmt.Println("")
			case "clear":
				clearScreen()
				printMainMenuActions()
			case "runtime":
				var m runtime.MemStats
				runtime.ReadMemStats(&m)

				fmt.Printf("Memória alocada: %v KB\n", m.Alloc/1024)
				fmt.Printf("Total alocado: %v KB\n", m.TotalAlloc/1024)
				fmt.Printf("Número de goroutines: %d\n", runtime.NumGoroutine())
			default:
				fmt.Println("Unknown action")
			}

		}
	}

}
