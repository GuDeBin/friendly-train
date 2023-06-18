package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var maze []string

// 加载maze01.txt文件，读取到切片内
func loadMaze(filePath string) error {
	// 读取文件
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
		return err
	}
	defer file.Close()

	// read the file content into the maze slice
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		maze = append(maze, line[:len(line)-1])
	}

	return nil
}

// 打印maze切片
func printMaze() {
	for _, row := range maze {
		fmt.Println(row)
	}
}

func main() {
	loadMaze("maze01.txt")
	printMaze()
}
