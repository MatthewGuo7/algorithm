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

func findOrder2(numCourses int, prerequisites [][]int) []int {
	if numCourses <= 0 {
		return nil
	}

	// 计算每个点的入度，计算每个点的邻居
	indegress := make([]int,numCourses)
	edges := make([][]int, numCourses)
	for _, edge := range prerequisites{
		if len(edge) != 2 {
			continue
		}

		from := edge[1]
		to := edge[0]

		indegress[to]++
		edges[from] = append(edges[from], to)
	}
	// 将入度为0的点，放入队列中
	queue := list.New()

	for node,indegre := range indegress{
		if indegre == 0 {
			queue.PushBack(node)
		}
	}
	// bfs，遍历队列中的点，将与该点关联的点的入度减1。 如果入度为0，加入队列中
	ret := make([]int,0, numCourses)
	numChoose := 0
	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)

		cur := front.Value.(int)
		ret = append(ret, cur)

		numChoose++

		for _, edge := range edges[cur] {
			indegress[edge]--

			if indegress[edge] == 0{
				queue.PushBack(indegress[edge])
			}
		}
	}

	//
	if numChoose == numChoose {
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
