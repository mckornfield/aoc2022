package day7

import (
	"fmt"
	"strconv"
	"strings"

	s "github.com/mckornfield/aoc2022/shared"
)

type Directory struct {
	path       string
	parentPath string
	size       int
}

func (d Directory) AdjustSize(fileSize int, directoryMap map[string]Directory) {
	d.size += fileSize
	directoryMap[d.path] = d
	parent := d.parentPath
	if parent != "" {
		if _, ok := directoryMap[parent]; ok {
			directoryMap[parent].AdjustSize(fileSize, directoryMap)
		}
	}
}

func processDirectoryInput(input string) map[string]Directory {

	directoryStack := s.Stack{}
	directoryMap := make(map[string]Directory)
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "$ cd") {
			cdSplit := strings.Split(line, " ")
			directoryName := cdSplit[2]
			if directoryName == ".." {
				directoryStack.Pop()
			} else {
				var previousDir string
				var parentPath string
				var path string
				if !directoryStack.IsEmpty() {
					previousDir = directoryStack.Peek().(string)
					parentPath = directoryMap[previousDir].path
					path = parentPath + directoryName + "/"
				} else {
					parentPath = ""
					previousDir = ""
					path = "/"
				}
				currentDirectory := Directory{
					parentPath: parentPath,
					size:       0,
					path:       path,
				}
				directoryMap[path] = currentDirectory
				directoryStack.Push(path)
			}
			fmt.Println(directoryStack)
		} else if strings.HasPrefix(line, "$ ls") {
			// Ignore
		} else { // Process files, ignore files
			if !strings.HasPrefix(line, "dir ") {
				fileLine := strings.Split(line, " ")
				print(fileLine[1])
				size, err := strconv.Atoi(fileLine[0])
				if err != nil {
					panic(err)
				}
				currentDir := directoryMap[directoryStack.Peek().(string)]
				currentDir.AdjustSize(size, directoryMap)
			}
		}
	}
	return directoryMap
}

func getFolderDeletionCount(input string) int {
	result := 0
	directoryMap := processDirectoryInput(input)
	for _, val := range directoryMap {
		if val.size <= 100000 {
			fmt.Println("ahh found a directory " + val.path)
			result += val.size
		}
	}
	fmt.Println(directoryMap)
	return result
}

func getDirectoryToDelete(input string) int {
	directoryMap := processDirectoryInput(input)
	totalSpaceUsed := directoryMap["/"].size
	currentMin := totalSpaceUsed
	maxSpace := 70000000
	updateSize := 30000000
	spaceRequired := updateSize - (maxSpace - totalSpaceUsed)
	for _, val := range directoryMap {
		fmt.Println(val)
		if val.size >= spaceRequired && currentMin > val.size {
			currentMin = val.size
		}
	}
	fmt.Println(directoryMap)
	return currentMin
}
