package bfs

import (
	"container/list"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	//1.找到所有的点
	allNodes := findAllNodesByDFS(node)
	//2.复制所有的点
	//mapping := copyNodes(allNodes)
	//3.复制所有的边
	copyEdges(allNodes)

	return allNodes[node]
}

func findAllNodesByDFS(node *Node)  map[*Node]*Node {
	if node == nil {
		return nil
	}

	queue := list.New()
	visted := map[*Node]*Node{}

	queue.PushBack(node)
	visted[node] = &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}

	for queue.Len() > 0 {
		curNode := queue.Front()
		queue.Remove(curNode)


		cur := curNode.Value.(*Node)

		for _, neighbour := range cur.Neighbors {
			if _, ok := visted[neighbour]; ok {
				continue
			}

			queue.PushBack(neighbour)
			visted[neighbour]= &Node{
				Val:       neighbour.Val,
				Neighbors: make([]*Node, 0),
			}
		}
	}

	return visted
}

func copyNodes(nodes map[*Node]struct{}) map[*Node]*Node{
	ret := make(map[*Node]*Node)
	for node ,_ := range nodes {
		newNode := &Node{
			Val:       node.Val,
			Neighbors: make([]*Node,0),
		}

		ret[node] = newNode
	}

	return ret
}

func copyEdges(mapping map[*Node]*Node) {
	for node,_ := range mapping{
		if newNode, ok := mapping[node]; ok {
			for _, neighbour := range node.Neighbors {
				neighbourOfNewNode := mapping[neighbour]
				newNode.Neighbors = append(newNode.Neighbors, neighbourOfNewNode)
			}
		}
	}
}