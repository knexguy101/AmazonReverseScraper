package main

import (
	"amazonReverse/tasks"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	printCommands()

	for scanner.Scan() {

		printCommands()

		items := strings.Split(scanner.Text(), " ")
		if len(items) < 2 {
			continue
		}

		switch items[0] {
		case "image":
			tasks.SearchByImage(items[1])
			break
		case "product":
			tasks.SearchByProduct(items[1])
			break
		case "store":
			maxItems, err := strconv.Atoi(items[1])
			if err != nil {
				fmt.Println("[INPUT ERROR]", items[1], "is not a number")
				continue
			}

			tasks.SearchByStore(items[2], maxItems)
			break
		}
	}
}

func printCommands() {
	fmt.Println("[COMMANDS]")
	fmt.Println("image [image link]\nproduct [amazon product link]\nstore [max # of items searched] [amazon store link]")
	fmt.Printf("->")
}
