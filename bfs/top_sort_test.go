package bfs

import (
	"container/list"
	"fmt"
	"testing"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses <= 0 {
		return nil
	}

	ret := make([]int, 0, numCourses)
	//1. 计算每个点的入度；计算每个点的邻居
	indegrees := make([]int, numCourses)
	edges := make([][]int, numCourses)
	for _, edge := range prerequisites {
		if len(edge) != 2 {
			continue
		}

		edges[edge[1]] = append(edges[edge[1]], edge[0])
		indegrees[edge[0]]++
	}

	//2. 将入度为0的点加入到队列中
	queue := list.New()
	for node, indegree := range indegrees {
		if indegree == 0 {
			queue.PushBack(node)
		}
	}
	//3. 进行BFS.
	// 取出队列中的点，将该点关联的点的入度减一, 若该点的入度为0，则加入队列中
	numChoose := 0
	for queue.Len() > 0 {
		curNode := queue.Front()
		queue.Remove(curNode)

		cur := curNode.Value.(int)

		numChoose++
		ret = append(ret, cur)

		for _, edge := range edges[cur] {
			indegrees[edge]--

			if indegrees[edge] == 0 {
				queue.PushBack(edge)
			}
		}
	}
	//4. 判断，返回结果

	if numChoose == numCourses {
		return ret
	}

	return nil
}

func TestFindOrder(t *testing.T) {
	pre := [][]int{
	}

	ret := findOrder(2, pre)
	fmt.Println(ret)
}
