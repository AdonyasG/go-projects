package cli

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"strings"
	"github.com/AdonyasG/go-projects/grocery-list-organizer/usecases"
)

type CLIHandler struct {
	useCase usecases.ItemUseCase
}

func NewCLIHandler(u usecases.ItemUseCase) *CLIHandler {
	return &CLIHandler{useCase: u}
}

func (h *CLIHandler) Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)

	for {
		fmt.Fprintln(writer, "\nChoose an action: add, list, delete, or exit")
		fmt.Fprint(writer, "-> ")
		writer.Flush()

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		switch input {
		case "add":
			h.addItem(scanner, writer)
		case "list":
			h.listItems(writer)
		case "delete":
			h.deleteItem(scanner, writer)
		case "exit":
			fmt.Fprintln(writer, "Exiting...")
			writer.Flush()
			return
		default:
			fmt.Fprintln(writer, "Unknown command. Please choose add, delete, list, or exit.")
			writer.Flush()
		}
	}
}

func (h *CLIHandler) addItem(scanner *bufio.Scanner, writer io.Writer) {
	fmt.Fprint(writer, "Enter Item name: ")
	writer.(*bufio.Writer).Flush()

	if !scanner.Scan() {
		fmt.Fprintln(writer, "Error reading input.")
		return
	}

	name := strings.TrimSpace(scanner.Text())

	fmt.Fprint(writer, "Enter Item Category: ")
	writer.(*bufio.Writer).Flush()

	if !scanner.Scan() {
		fmt.Fprintln(writer, "Error reading input.")
		return
	}

	category := strings.TrimSpace(scanner.Text())

	item, err := h.useCase.AddItem(name, category)

	if err != nil {
		fmt.Fprintln(writer, "Error adding Item: ", err)
	} else {
		fmt.Fprintf(writer, "Item added with ID %d\n", item.ID)
	}

	writer.(*bufio.Writer).Flush()

}

func (h *CLIHandler) listItems(writer io.Writer) {

	items := h.useCase.ListItemsByCategory()

	if len(items) == 0 {
		fmt.Fprintln(writer, "No Items found.")
		writer.(*bufio.Writer).Flush()
		return
	}

	fmt.Fprintln(writer, "Items:")
	for _, item := range items {

		fmt.Fprintf(writer, "ID: %d, Name: %s, Category: %s\n", item.ID, item.Name, item.Category)
	}
	writer.(*bufio.Writer).Flush()
}

func (h *CLIHandler) deleteItem(scanner *bufio.Scanner, writer io.Writer) {

	fmt.Fprint(writer, "Enter ID of an Item to remove: ")
	writer.(*bufio.Writer).Flush()

	if !scanner.Scan() {
		fmt.Fprintln(writer, "Error reading input.")
		return
	}

	idStr := strings.TrimSpace(scanner.Text())

	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Fprintln(writer, "Invalid ID format")
		return
	}

	if err := h.useCase.DeleteItem(id); err != nil {
		fmt.Fprintln(writer, "Error deleting item:", err)
	} else {
		fmt.Fprintln(writer, "Item deleted")
	}

	writer.(*bufio.Writer).Flush()

}
