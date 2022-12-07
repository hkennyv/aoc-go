package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hkennyv/aoc-go/utils"
)

func main() {
	input := parseInput("input.txt")
	root := buildTree(input)
	dirs := getDirs(root)

	// part1
	sum := 0
	for _, d := range dirs {
		if d.Size < 100000 {
			sum += d.Size
		}
	}
	fmt.Println("part1:", sum)

	// part2
	diff := 30000000 - (70000000 - root.Size)
	closest := root.Size
	for _, d := range dirs {
		if d.Size > diff {
			closest = utils.Min(d.Size, closest)
		}
	}
	fmt.Println("part2:", closest)
}

func buildTree(input []string) *utils.Node {

	root := &utils.Node{
		Name:     "/",
		Size:     0,
		Parent:   nil,
		Children: []*utils.Node{},
	}

	cwd := root

	// skip first input, since it is root
	i := 1

	for i < len(input) {
		line := input[i]
		split := strings.Split(line, " ")

		if strings.HasPrefix(line, "$ cd") {
			name := split[2]

			if name == ".." {
				cwd = cwd.Parent
			} else {
				new, err := cwd.AddChild(name, 0)
				if err != nil {
					fmt.Println(err)
				}
				cwd = new
			}

		} else if strings.HasPrefix(line, "$ ls") {
			// do nothing
		} else if strings.HasPrefix(line, "dir ") {
			// do nothing
		} else {
			name := split[1]
			size, _ := strconv.Atoi(split[0])
			_, err := cwd.AddChild(name, size)
			if err != nil {
				fmt.Println(err)
			}
		}

		i++
	}

	root.CalculateSize()
	return root
}

func getDirs(n *utils.Node) []*utils.Node {
	ds := make([]*utils.Node, 0)

	if len(n.Children) > 0 {
		ds = append(ds, n)
	}

	for _, c := range n.Children {
		ds = append(ds, getDirs(c)...)
	}

	return ds
}

func parseInput(fp string) []string {
	b, _ := os.ReadFile(fp)
	return strings.Split(strings.TrimSpace(string(b)), "\n")
}
