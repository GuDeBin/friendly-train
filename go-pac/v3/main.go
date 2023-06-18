package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/danicat/simpleansi"
)

var maze []string

func loadMaze(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	return nil
}

func printScreen() {
	simpleansi.ClearScreen()
	for _, line := range maze {
		fmt.Println(line)
	}
}

func readInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	}

	return "", nil

}

func initialise() {
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalln("Unable to activate cbreak mode:", err)
	}

}

func cleanup() {
	cbTerm := exec.Command("stty", "-cbreak", "echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalln("Unable to restore cooked mode:", err)
	}

}

func main() {

	initialise()
	defer cleanup()

	err := loadMaze("maze.txt")
	if err != nil {
		log.Fatalln("Unable to load maze:", err)
		return
	}

	for {
		printScreen()
		input, err := readInput()
		if err != nil {
			log.Fatalln("Unable to read input:", err)
			break
		}

		if input == "ESC" {
			break
		}
	}

}
