package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func gsutilCommand(arguments ...string) error {
	cmd := exec.Command("gsutil", arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Need 3 arguments : file_pattern, num_files, destination")
		return
	}
	args := os.Args[1:]
	filePattern := args[0]
	fileDestination := args[2]

	nbFiles, err := strconv.Atoi(args[1])
	if err != nil || nbFiles <= 0 {
		fmt.Println("Argument 2, num_files needs to be a integer greater than 0")
		return
	}

	wildcarPosition := strings.Index(filePattern, "*")
	if wildcarPosition == -1 {
		fmt.Printf("file_pattern %v has no wildcard \n", filePattern)
		return
	}

	beginingPattern := filePattern[:wildcarPosition]
	endingPattern := filePattern[wildcarPosition+1:]

	gsutilCommand("cp", beginingPattern+"000000000000"+endingPattern, fileDestination)

	maxFiles := 31
	trailingLength := 12
	nbTour := nbFiles / maxFiles

	numFile := 0

	for tour := 0; tour <= nbTour; tour++ {
		countFiles := 0
		composeCommand := "compose " + fileDestination + " "
		for numFile < nbFiles && countFiles < maxFiles {
			pattern := strconv.Itoa(numFile)
			zeros := ""
			for i := 0; i < trailingLength-len(pattern); i++ {
				zeros += "0"
			}
			composeCommand += beginingPattern + zeros + pattern + endingPattern + " "
			numFile++
			countFiles++
		}
		if countFiles > 0 {
			composeCommand += fileDestination
			fmt.Println(composeCommand)
			arguments := strings.Split(composeCommand, " ")
			err = gsutilCommand(arguments...)
			if err != nil {
				return
			}
		}
	}
}
