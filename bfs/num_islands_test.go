package bfs

import (
	"container/list"
	"fmt"
	"testing"
)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	if len(grid[0]) == 0 {
		return 0
	}

	visited := make(map[Coordinate]struct{}, len(grid)*len(grid[0]))

	islands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] == '1' {

				cur := Coordinate{x: i, y: j}
				if _, ok := visited[cur]; ok {
					continue
				}

				bfs(grid, cur, visited)
				islands++
			}
		}
	}

	return islands
}

type Coordinate struct {
	x, y int
}

func bfs(grid [][]byte, coordinate Coordinate, visited map[Coordinate]struct{}) {
	queue := list.New()
	queue.PushBack(coordinate)

	visited[coordinate] = struct{}{}

	for queue.Len() > 0 {
		curNode := queue.Front()
		queue.Remove(curNode)

		cur := curNode.Value.(Coordinate)

		//if isValidCoordinate(grid, cur) && grid[cur.x][cur.y] == '1'{


			//grid[cur.x][cur.y] = '0'

			left := Coordinate{x: cur.x - 1, y: cur.y}
			right := Coordinate{x: cur.x + 1, y: cur.y}
			up := Coordinate{x: cur.x, y: cur.y - 1}
			down := Coordinate{x: cur.x, y: cur.y + 1}

			//queue.PushBack(left)
			//queue.PushBack(right)
			//queue.PushBack(up)
			//queue.PushBack(down)
			canPushToQueue(grid, queue, left, visited)
			canPushToQueue(grid, queue, right, visited)
			canPushToQueue(grid, queue, up, visited)
			canPushToQueue(grid, queue, down, visited)

	}
}

func canPushToQueue(grid [][]byte, queue *list.List,
	coordinate Coordinate, visited map[Coordinate]struct{}) {

	if !(isValidCoordinate(grid, coordinate) && grid[coordinate.x][coordinate.y] == '1') {
		return
	}

	if _, ok := visited[coordinate]; !ok  {
		queue.PushBack(coordinate)
		visited[coordinate] = struct{}{}
	}
}

func isValidCoordinate(grid [][]byte, coordinate Coordinate) bool {
	n := len(grid)
	m := len(grid[0])

	return (coordinate.x >= 0 && coordinate.x < n) && (coordinate.y >= 0 && coordinate.y < m)
}

func TestIslands(t *testing.T) {
	grid := [][]byte{

		{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'},
	}

	fmt.Println(numIslands(grid))
}
