package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type histEntry struct {
	currX, currY, dirX, dirY int
	straightSteps            int
}

type queueEntry struct {
	currX, currY, dirX, dirY int
	straightSteps            int
	currCosts                int
}

const MAX_BUFFER = 1000000

func queueToHist(qu queueEntry) histEntry {
	var his histEntry
	his.currX = qu.currX
	his.currY = qu.currY
	his.dirX = qu.dirX
	his.dirY = qu.dirY
	his.straightSteps = qu.straightSteps
	return his
}

func checkBoundsHist(data [][]int, his histEntry) bool {
	if his.currX < 0 || his.currX >= len(data[0]) || his.currY < 0 || his.currY >= len(data) {
		return false
	}
	return true
}

func makeQueueEntry(currX, currY, dirX, dirY, straightSteps, currCosts int) queueEntry {
	var qu queueEntry
	qu.currX = currX
	qu.currY = currY
	qu.dirX = dirX
	qu.dirY = dirY
	qu.straightSteps = straightSteps
	qu.currCosts = currCosts
	return qu
}

func makeHistEntry(currX, currY, dirX, dirY, straightSteps int) histEntry {
	var his histEntry
	his.currX = currX
	his.currY = currY
	his.dirX = dirX
	his.dirY = dirY
	his.straightSteps = straightSteps
	return his
}

func histToQueue(his histEntry, cost int) queueEntry {
	var qu queueEntry
	qu.currX = his.currX
	qu.currY = his.currY
	qu.dirX = his.dirX
	qu.dirY = his.dirY
	qu.straightSteps = his.straightSteps
	qu.currCosts = cost
	return qu
}

func findSmallestHeatLoss(data [][]int) int {
	queue := make([]queueEntry, MAX_BUFFER)

	queue = append(queue, makeQueueEntry(1, 0, 1, 0, 1, data[0][1]))
	queue = append(queue, makeQueueEntry(0, 1, 0, 1, 1, data[1][0]))

	history := make(map[histEntry]bool)

	for len(queue) > 0 {
		currEntry := queue[0]
		queue = queue[1:]

		if currEntry.currX == len(data[0])-1 && currEntry.currY == len(data)-1 {
			return currEntry.currCosts
		}

		if history[queueToHist(currEntry)] {
			continue
		}

		history[queueToHist(currEntry)] = true

		if currEntry.straightSteps < 3 {
			en := makeHistEntry(currEntry.currX+currEntry.dirX, currEntry.currY+currEntry.dirY, currEntry.dirX, currEntry.dirY, currEntry.straightSteps+1)
			if checkBoundsHist(data, en) {
				queue = append(queue, histToQueue(en, currEntry.currCosts+data[en.currY][en.currX]))
				sort.Slice(queue, func(i, j int) bool {
					return queue[i].currCosts < queue[j].currCosts
				})
			}
		}

		en := makeHistEntry(currEntry.currX+currEntry.dirY, currEntry.currY+currEntry.dirX, currEntry.dirY, currEntry.dirX, 1)
		if checkBoundsHist(data, en) {
			queue = append(queue, histToQueue(en, currEntry.currCosts+data[en.currY][en.currX]))
			sort.Slice(queue, func(i, j int) bool {
				return queue[i].currCosts < queue[j].currCosts
			})
		}

		en = makeHistEntry(currEntry.currX-currEntry.dirY, currEntry.currY-currEntry.dirX, -currEntry.dirY, -currEntry.dirX, 1)
		if checkBoundsHist(data, en) {
			queue = append(queue, histToQueue(en, currEntry.currCosts+data[en.currY][en.currX]))
			sort.Slice(queue, func(i, j int) bool {
				return queue[i].currCosts < queue[j].currCosts
			})
		}
	}

	return -1
}

func handleData(data [][]int) int {
	return findSmallestHeatLoss(data)
}

func main() {
	// Open the file
	file, err := os.Open("17.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var data [][]int
	// Iterate through each lin
	for scanner.Scan() {
		line := scanner.Text()
		var lineArr []int
		for _, char := range line {
			lineArr = append(lineArr, int(char-'0'))
		}
		data = append(data, lineArr)
	}

	sum := handleData(data)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
