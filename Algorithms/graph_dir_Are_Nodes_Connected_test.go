package main

import (
	diGraph "../Data_Structures/graph/directed"
	Queue "../Data_Structures/queue"
	"testing"
)

func diGraphAreNodesConnected(g *diGraph.Graph, from, to *diGraph.Node) (found bool) {
	if from == nil || to == nil || g == nil {
		return false
	} else if *from == *to {
		return true
	}
	edges, err := g.GetEdges(from); if err != nil {
		return false
	}
	visited := map[diGraph.Node]struct{}{}
	visited[*from] = struct{}{}
	q := Queue.Queue{}
	q.Add(edges)

	for q.Size() > 0 {
		edges = q.Pop().([]*diGraph.Edge)
		for n := range edges {
			node := *edges[n].Node
			if _, ok := visited[node] ; ok {
				continue
			}
			if node == *to {
				return true
			}
			visited[node] = struct{}{}
			edges, err := g.GetEdges(&node); if err == nil {
				q.Add(edges)
			}
		}
	}
	return false
}

func TestDiGraphNodesConFalse(t *testing.T) {
	g := genGraph()
	// 1 ⟷ 2 ⭢ 5
	// ⭥   /	 /
	// ⭥ ⬋		/
	// 3 ⭢ 4 ⭠

	n5:= &diGraph.Node{"5"}
	n1:= &diGraph.Node{"1"}

	if ok := diGraphAreNodesConnected(g, n5, n1); ok {
		t.Fatal("expected false, got true")
	}
}

func TestDiGraphNodesConTrue(t *testing.T) {
	g := genGraph()
	// 1 ⟷ 2 ⭢ 5
	// ⭥   /	 /
	// ⭥ ⬋		/
	// 3 ⭢ 4 ⭠

	n5:= &diGraph.Node{"5"}
	n1:= &diGraph.Node{"1"}

	if ok := diGraphAreNodesConnected(g, n1, n5); !ok {
		t.Fatal("expected true, got false")
	}
}


func genGraph() *diGraph.Graph {
	// 1 ⟷ 2 ⭢ 5
	// ⭥   /	 /
	// ⭥ ⬋		/
	// 3 ⭢ 4 ⭠

	g := diGraph.New()
	n1 := &diGraph.Node{"1"}
	n2 := &diGraph.Node{"2"}
	n3 := &diGraph.Node{"3"}
	n4:= &diGraph.Node{"4"}
	n5:= &diGraph.Node{"5"}
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddEdge(n1, &diGraph.Edge{Node: n2})
	g.AddEdge(n1, &diGraph.Edge{Node: n3})
	g.AddEdge(n2, &diGraph.Edge{Node: n1})
	g.AddEdge(n2, &diGraph.Edge{Node: n3})
	g.AddEdge(n2, &diGraph.Edge{Node: n5})
	g.AddEdge(n3, &diGraph.Edge{Node: n1})
	g.AddEdge(n3, &diGraph.Edge{Node: n4})
	g.AddEdge(n5, &diGraph.Edge{Node: n4})
	return g
}

//	⟷
//	⭡
//	⭣
//	⭥
//	⭠
//	⭢
// ⬉ ⬈
// ⬊ ⬋
