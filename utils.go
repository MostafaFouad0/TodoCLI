package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func FileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func WriteToFile(FileName, Data string) bool {
	file, err := os.OpenFile(FileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("[!] Error Happened %v", err)
		return false
	}
	defer file.Close()
	file.Write([]byte(Data))
	return true
}

func ReadFile(FileName string) (string, bool) {
	file, err := os.Open(FileName)
	if err != nil {
		return "", false
	}
	var Data string
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "\n" {
			continue
		}
		Data += scanner.Text() + "\n"
	}
	return Data, true
}

func WriteContentToFile(FileName, Data string) bool {
	file, err := os.OpenFile(FileName, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return false
	}
	defer file.Close()
	file.WriteString(Data)
	return true
}

func DeleteLine(FileName string, LineNumber int) (string, bool) {
	file, err := os.Open(FileName)
	if err != nil {
		fmt.Printf("[!] Error Happened %v", err)
		return "", false
	}
	var Data string
	var toDone string
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		if scanner.Text() == "\n" {
			continue
		}
		if i != LineNumber {
			Data += scanner.Text()
		} else {
			toDone += scanner.Text()
		}
	}
	toDone += "\n"
	return toDone, WriteContentToFile(FileName, Data)
}
func OutputFormatedFile(FileName string) {
	Data, err := ReadFile(FileName)
	if !err {
		fmt.Println("Error Happened reading the file.")
		return
	}
	DataArray := strings.Split(Data, "\n")
	if len(DataArray[0]) == 0 {
		return
	}
	DataArray = DataArray[:len(DataArray)-1]
	for i, value := range DataArray {
		fmt.Printf("%v - %v\n", i+1, value)
	}
}

func CreateFile(name string) bool {
	file, err := os.Create(name)
	defer file.Close()
	if err != nil {
		fmt.Printf("[!] Error happened when creating %v\n", name)
		return false
	}
	fmt.Printf("[!] %v Created successfully!\n", name)
	return true
}

func CreateNecessaryFiles(FileName string, wg *sync.WaitGroup) {
	defer wg.Done()
	if !FileExists(FileName) {
		fmt.Printf("[!] %v Doesn't exist\n", FileName)
		fmt.Printf("[!] Creating %v...\n", FileName)
		CreateFile(FileName)
	} else {
		fmt.Printf("[!] %v Already Exists\n", FileName)
	}
}
