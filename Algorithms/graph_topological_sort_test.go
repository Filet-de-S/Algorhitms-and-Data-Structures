package main

import (
	diGraph "../Data_Structures/graph/directed"
	_stack "../Data_Structures/stack"
	"testing"
)

func GraphTopologicalSort(g *diGraph.Graph) *_stack.Stack {
	if g == nil {
		return nil
	}
	visited := map[diGraph.Node]struct{}{}
	next := &_stack.Stack{}
	ordered := &_stack.Stack{}
	waitPlease := &_stack.Stack{}

	nodes, err := g.GetNodes()
	if err != nil {
		return nil
	}

	for i := range nodes {
		if _, ok := visited[*nodes[i]]; ok {
			continue
		}
		node := nodes[i]
		for node != nil {
			visited[*node] = struct{}{}

			edges, err := g.GetEdges(node)
			if err == nil {
				for e := range edges {
					next.Insert(edges[e].Node)
				}
			} else {
				ordered.Insert(node)
				node = nil
			}
			for {
				nex, err := next.Pop()
				if err == nil {
					if node != nil {
						waitPlease.Insert(node)
					}
					node = nex.(*diGraph.Node)
					if _, ok := visited[*node]; ok {
						node = nil
						continue
					}
					break
				} else {
					node = nil
					break
				}
			}
		}
		for {
			q, err := waitPlease.Pop()
			if err == nil {
				ordered.Insert(q)
			} else {
				break
			}
		}
	}
	return ordered
}

func TestTopSort(t *testing.T) {
	g := genRouteForTopSort()
	o := GraphTopologicalSort(g)
	app := []diGraph.Node{}
	for {
		ord, err := o.Pop()
		if err == nil {
			node := ord.(*diGraph.Node)
			app = append(app, *node)
		} else {
			break
		}
	}
	t.Log(app)
}

func genRouteForTopSort() *diGraph.Graph {
	gr := diGraph.New()

	a := &diGraph.Node{"a"}
	b := &diGraph.Node{"b"}
	c := &diGraph.Node{"c"}
	d := &diGraph.Node{"d"}
	e := &diGraph.Node{"e"}
	f := &diGraph.Node{"f"}
	g := &diGraph.Node{"g"}
	h := &diGraph.Node{"h"}
	i := &diGraph.Node{"i"}
	j := &diGraph.Node{"j"}
	gr.AddNode(a,b,c,d,e,f,g,h,i,j)

	gr.AddEdge(a, &diGraph.Edge{Node: b},
		&diGraph.Edge{Node: f})
	gr.AddEdge(b, &diGraph.Edge{Node: h})
	gr.AddEdge(d, &diGraph.Edge{Node: c},
		&diGraph.Edge{Node: e},
		&diGraph.Edge{Node: i},
		&diGraph.Edge{Node: h})
	gr.AddEdge(e, &diGraph.Edge{Node: i})
	gr.AddEdge(g, &diGraph.Edge{Node: a},
		&diGraph.Edge{Node: b},
		&diGraph.Edge{Node: c})
	gr.AddEdge(i, &diGraph.Edge{Node: c})
	gr.AddEdge(j, &diGraph.Edge{Node: e})

	return gr
}
