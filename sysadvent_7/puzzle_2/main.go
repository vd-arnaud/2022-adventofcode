package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	name     string
	nodeType string
	size     int
	nodes    []*Node
}

func main() {

	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var currentPath string = ""
	var currentDir string = ""
	var currentNode *Node
	var mapNodes map[string]*Node = make(map[string]*Node)
	for fileScanner.Scan() {
		line := fileScanner.Text()

		////////////
		// cd line
		////////////
		if strings.HasPrefix(line, "$ cd") {
			// Set current path
			if string(line[5]) == "." && string(line[6]) == "." {
				pathSplit := strings.Split(currentPath, "/")
				previousPathSplit := pathSplit[:len(pathSplit)-1]
				currentPath = strings.Join(previousPathSplit, "/")
				currentDir = previousPathSplit[len(previousPathSplit)-1]
			} else {
				currentPath = currentPath + "/" + line[5:len(line)]
				currentDir = line[5:len(line)]
			}
			//fmt.Println(currentPath)
			// Set current Node
			if node, ok := mapNodes[currentPath]; ok {
				currentNode = node
			} else {
				n := Node{
					name:     currentDir,
					nodeType: "dir",
					size:     0,
					nodes:    []*Node{},
				}

				mapNodes[currentPath] = &n
				currentNode = &n
			}

			//fmt.Println(currentPath)
			//fmt.Println(currentNode)
		}

		////////////
		// dir line
		////////////
		if strings.HasPrefix(line, "dir ") {
			//fmt.Printf("Line = %s\n", line)
			//fmt.Printf("Current Node before = %v\n", currentNode)
			regexp_dir := `dir\s(?P<dirname>\S+)`
			r := regexp.MustCompile(regexp_dir)
			matches := r.FindStringSubmatch(line)
			dirname := matches[1]

			if node, ok := mapNodes[currentPath+"/"+dirname]; ok {
				currentNode.nodes = append(currentNode.nodes, node)
			} else {
				n := Node{
					name:     dirname,
					nodeType: "dir",
					size:     0,
					nodes:    []*Node{},
				}
				mapNodes[currentPath+"/"+dirname] = &n
				currentNode.nodes = append(currentNode.nodes, &n)
			}
			//fmt.Printf("Current Node after = %v\n", currentNode)
		}

		////////////
		// file line
		////////////
		regexp_file := `(?P<size>\d+)\s(?P<filename>\S+)`
		matched, _ := regexp.MatchString(regexp_file, line)
		if matched {
			r := regexp.MustCompile(regexp_file)
			matches := r.FindStringSubmatch(line)
			size, _ := strconv.Atoi(matches[1])
			filename := matches[2]
			//create Node
			fileNode := Node{
				name:     filename,
				nodeType: "file",
				size:     size,
				nodes:    nil,
			}
			//add to nodes of currentNode
			currentNode.nodes = append(currentNode.nodes, &fileNode)

		}
	}

	var candidatSize = 9999999
	var totalSize int = 0
	for path, node := range mapNodes {
		fmt.Printf("Path = %s || Node = ", path)
		fmt.Println(node)
		size := computeNodeSize(node)
		if size >= 2558312 {
			if size < candidatSize {
				candidatSize = size
			}
		}
		totalSize += size

	}
	fmt.Println(totalSize)
	var rootSize int = computeNodeSize(mapNodes["/root"])
	fmt.Println(rootSize)
	var freeSpace = 70000000 - rootSize
	fmt.Println(freeSpace)
	var spaceToFree = 30000000 - freeSpace
	fmt.Println(spaceToFree)
	fmt.Println(candidatSize)
	readFile.Close()
}

func computeNodeSize(node *Node) int {
	var size int = 0
	for _, childNode := range node.nodes {
		if childNode.nodeType == "file" {
			size += childNode.size
		} else {
			size += computeNodeSize(childNode)
		}
	}
	return size
}
