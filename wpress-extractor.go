package main

import (
	"fmt"
	"os"

	"github.com/mabakach/wpress"
)

func main() {
	fmt.Printf("Wpress Extracter.\n")

	if len(os.Args) >= 2 {
		pathTofile := os.Args[1]
		outputPath := "."
		if len(os.Args) >= 3 {
			outputPath = os.Args[2]
			if fileExists(outputPath) {
				fmt.Println("Output path is a file! Please provide a path to a directory.")
				return
			}
			if !directoryExists(outputPath) {
				err := os.MkdirAll(outputPath, 0777)
				if err != nil {
					fmt.Println("Could not create output directory ", outputPath)
					fmt.Println("Error ", err)
					return
				}
			}
		}
		fmt.Println(pathTofile)
		archiver, _ := wpress.NewReader(pathTofile)
		_, err := archiver.ExtractToPath(outputPath)
		if err != nil {
			fmt.Println("Error = ")
			fmt.Println(err)
		} else {
			fmt.Println("All done!")
		}

	} else {
		printUsage()
		return
	}
	return
}

func fileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func directoryExists(dirName string) bool {
	info, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func printUsage() {
	fmt.Println("Usage: Wordpress-Extractor <path/to/wpress/file> [output/path]")
	fmt.Println("Default output path is the current directory")
}
