package day8

import (
	"fmt"
	"strings"
)

func SolutionPart1(input string) int {
	lines := strings.Split(input, "\n")

	var rootNode *Node
	existingNodes := map[string]*Node{}
	for _, nodes := range lines[2:] {
		from := nodes[:3]
		toLeft := nodes[7:10]
		toRight := nodes[12:15]

		if rootNode == nil && from == "AAA" {
			rootNode = NewNode(from)
			existingNodes[from] = rootNode
		}

		parentNode, ok := existingNodes[from]
		if !ok {
			parentNode = NewNode(from)
			existingNodes[from] = parentNode
		}

		leftNode, ok := existingNodes[toLeft]
		if !ok {
			leftNode = parentNode.AddLeftChildren(toLeft)
		} else {
			parentNode.SetLeftChildren(leftNode)
		}
		existingNodes[toLeft] = leftNode

		rightNode, ok := existingNodes[toRight]
		if !ok {
			rightNode = parentNode.AddRightChildren(toRight)
		} else {
			parentNode.SetRightChildren(rightNode)
		}
		existingNodes[toRight] = rightNode
	}

	instructions := lines[0]
	numLoops := 0
	currNode := rootNode
	for {
		numLoops++
		for _, ins := range instructions {
			if currNode.left == nil || currNode.right == nil {
				fmt.Printf("running into nil children on %+v\n", currNode)
			}
			switch ins {
			case 'L':
				currNode = currNode.left
			case 'R':
				currNode = currNode.right
			default:
				panic(fmt.Sprintf("invalid instruction %c\n", ins))
			}
		}
		if currNode.value == "ZZZ" {
			break
		}
	}

	return numLoops * len(instructions)
}

func SolutionPart2(input string) int {
	lines := strings.Split(input, "\n")

	var rootNodes []*Node
	existingNodes := map[string]*Node{}
	for _, nodes := range lines[2:] {
		from := nodes[:3]
		toLeft := nodes[7:10]
		toRight := nodes[12:15]

		if from[2] == 'A' {
			rootNode := NewNode(from)
			rootNodes = append(rootNodes, rootNode)
			existingNodes[from] = rootNode
		}

		parentNode, ok := existingNodes[from]
		if !ok {
			parentNode = NewNode(from)
			existingNodes[from] = parentNode
		}

		leftNode, ok := existingNodes[toLeft]
		if !ok {
			leftNode = parentNode.AddLeftChildren(toLeft)
		} else {
			parentNode.SetLeftChildren(leftNode)
		}
		existingNodes[toLeft] = leftNode

		rightNode, ok := existingNodes[toRight]
		if !ok {
			rightNode = parentNode.AddRightChildren(toRight)
		} else {
			parentNode.SetRightChildren(rightNode)
		}
		existingNodes[toRight] = rightNode
	}

	instructions := lines[0]
	finishSteps := []int{}
	for _, rootNode := range rootNodes {

		numLoops := 0
		currNode := rootNode
		for {
			numLoops++
			for _, ins := range instructions {
				if currNode.left == nil || currNode.right == nil {
					fmt.Printf("running into nil children on %+v\n", currNode)
				}
				switch ins {
				case 'L':
					currNode = currNode.left
				case 'R':
					currNode = currNode.right
				default:
					panic(fmt.Sprintf("invalid instruction %c\n", ins))
				}
			}
			if currNode.value[2] == 'Z' {
				break
			}
		}

		finishSteps = append(finishSteps, numLoops*len(instructions))
	}

	return LCM(finishSteps[0], finishSteps[1], finishSteps[2:]...)
}
