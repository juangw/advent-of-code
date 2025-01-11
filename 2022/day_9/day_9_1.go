package day_9

import (
	"github.com/juangw/advent-of-code/math"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	count     int
}

func RunPart1(logger *log.Logger) {
	motions, err := ioutil.ReadFile("./2022/day_9/day_9.txt")
	if err != nil {
		logger.Fatalf("unable to read file: %v", err)
	}

	cleanedMotions := strings.TrimSpace(string(motions))
	motionsSlice := strings.Split(cleanedMotions, "\n")

	head := [2]int{0, 0}
	tail := [2]int{0, 0}

	// mapping of directions
	diffs := map[string][2]int{
		"U": {0, 1},
		"D": {0, -1},
		"L": {-1, 0},
		"R": {1, 0},
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

		instruct := instruction{direction: direction, count: intCount}
		for instruct.count > 0 {
			// move head
			diff := diffs[instruct.direction]
			head[0] += diff[0]
			head[1] += diff[1]

			// update tail
			xDiff := head[0] - tail[0]
			yDiff := head[1] - tail[1]

			if math.AbsInt(xDiff) > 1 {
				// move tail x coord if it needs to catch up
				tail[0] += diff[0]
				// account for diagonal adjustment, same math... add y diff
				if yDiff != 0 {
					tail[1] += yDiff
				}
			} else if math.AbsInt(yDiff) > 1 {
				// move tail y coord if it needs to catch up
				tail[1] += diff[1]
				// account for diagonal adjustment, same math... add x diff
				if xDiff != 0 {
					tail[0] += xDiff
				}
			}

			// add unique locations where tail has been
			visited[tail] += 1
			instruct.count--
		}
	}

	// return spots TAIL visited at least once, map[[2]int]bool
	logger.Println(len(visited))
}
