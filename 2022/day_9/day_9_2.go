package day_9

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type node struct {
	position [2]int
	next     *node
}

type rope struct {
	head *node
	tail *node
}

func (r rope) String() string {
	str := ""
	i := 0
	for itr := r.head; itr != nil; itr = itr.next {
		str += fmt.Sprintf("%d:[%d,%d]->", i, itr.position[0], itr.position[1])
		i++
	}
	return str
}

func RunPart2(logger *log.Logger) {
	motions, err := ioutil.ReadFile("./2022/day_9/day_9.txt")
	if err != nil {
		logger.Fatalf("unable to read file: %v", err)
	}

	cleanedMotions := strings.TrimSpace(string(motions))
	motionsSlice := strings.Split(cleanedMotions, "\n")

	// mapping of directions
	diffs := map[string][2]int{
		"U": {1, 0},
		"D": {-1, 0},
		"L": {0, -1},
		"R": {0, 1},
	}

	startNode := &node{}
	nextNode := startNode
	for i := 1; i <= 9; i++ {
		nextNode.next = &node{}
		nextNode = nextNode.next
	}
	rope := rope{
		head: startNode,
		tail: nextNode,
	}
	visited := map[[2]int]int{
		{0, 0}: 1,
	}
	for _, motion := range motionsSlice {
		directions := strings.Split(string(motion), " ")
		direction, count := directions[0], directions[1]
		intCount, err := strconv.Atoi(count)
		if err != nil {
			logger.Fatal("unable to convert direction count to int")
		}

		for intCount > 0 {
			// move head
			diff := diffs[direction]
			rope.head.position[0] += diff[0]
			rope.head.position[1] += diff[1]

			// recusively update tail nodes
			nextTailNode := rope.head
			for nextTailNode != nil {
				if nextTailNode.next == nil {
					break
				}
				xDiff := nextTailNode.position[0] - nextTailNode.next.position[0]
				yDiff := nextTailNode.position[1] - nextTailNode.next.position[1]

				if AbsInt(xDiff) > 1 && AbsInt(yDiff) > 1 {
					nextTailNode.next.position[0] += xDiff / 2
					nextTailNode.next.position[1] += yDiff / 2
				} else if AbsInt(xDiff) > 1 {
					nextTailNode.next.position[0] += xDiff / 2
					nextTailNode.next.position[1] += yDiff
				} else if AbsInt(yDiff) > 1 {
					nextTailNode.next.position[1] += yDiff / 2
					nextTailNode.next.position[0] += xDiff
				} else {
					// no need to continue updating children if movement is over
					break
				}
				nextTailNode = nextTailNode.next
			}
			visited[rope.tail.position] += 1
			intCount--
		}
	}

	// return spots TAIL visited at least once, map[[2]int]bool
	logger.Println(len(visited))
}
