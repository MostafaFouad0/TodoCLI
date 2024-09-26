package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func PreProcess() {
	var wg sync.WaitGroup
	wg.Add(1)
	go CreateNecessaryFiles("TODOs.txt", &wg)
	wg.Add(1)
	go CreateNecessaryFiles("DONE.txt", &wg)
	wg.Wait()
}

func main() {
	PreProcess()
	fmt.Println()
	var in int
	for true {
		fmt.Println("Select Number :")
		fmt.Println("[1] - Show Todos")
		fmt.Println("[2] - Show What You Have Done So Far")
		fmt.Println("[3] - Add Task")
		fmt.Println("[4] - Mark Task As Done")
		fmt.Println("[0] - Exit")
		fmt.Scanln(&in)
		if in == 1 {
			OutputFormatedFile("TODOs.txt")
			fmt.Println("------------------------------------")
		} else if in == 2 {
			OutputFormatedFile("DONE.txt")
			fmt.Println("------------------------------------")
		} else if in == 3 {
			fmt.Println("Type Description for the task you want to add")
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			WriteToFile("TODOs.txt", input)
			fmt.Println("------------------------------------")
		} else if in == 4 {
			fmt.Println("Type The Number Of The Task")
			fmt.Scanln(&in)
			DataToDone, ok := DeleteLine("TODOs.txt", in)
			if ok {
				WriteToFile("DONE.txt", DataToDone)
			} else {
				fmt.Println("Error Happened.")
				continue
			}
		} else if in == 0 {
			break
		} else {
			fmt.Println("[!] Please Choose Between 0 and 4")
		}
	}
}
