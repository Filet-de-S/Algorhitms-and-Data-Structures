package main

import (
	pqHeap "../Data_Structures/priority_queue(bin_heap)"
	unDirGraph "../Data_Structures/graph/undirected"
	"errors"
	"testing"
)

type route struct {
	via    *unDirGraph.Node
	weight unDirGraph.Weight
}

type path struct {
	node *unDirGraph.Node
	via *unDirGraph.Node
}

func addToQue(visited map[unDirGraph.Node]*route, pq *pqHeap.Heap, edges []*unDirGraph.Edge, via *unDirGraph.Edge) {
	for e := range edges {
		if _, ok := visited[ *edges[e].Node ]; ok {
			continue
		}
		//TODO: at the moment we store duplicates in pq until visited
		// could be improvement in pq.DS: new func(changeIfPresent)
		pq.Insert(&pqHeap.Data{
			CmpData: via.Weight + edges[e].Weight,
			Data:    &path{edges[e].Node, via.Node},
		})
	}
}

type ArgIsNil error
type NoConnectivity error

func ShortestPathDijkstra(g *unDirGraph.Graph, from, to *unDirGraph.Node) ([]*finRoute, error) {
	if from == nil || to == nil || g == nil {
		return nil, ArgIsNil(errors.New(""))
	} else if *from == *to {
		return []*finRoute{{via: from}}, nil
	}
	edges, err := g.GetEdges(from); if err != nil {
		return nil, NoConnectivity(errors.New(""))
	}
	visited := map[unDirGraph.Node]*route{}
	visited[*from] = &route{}

	pq := pqHeap.New("min")
	addToQue(visited, pq, edges, &unDirGraph.Edge{Node: from})
	var found bool
	var least *route

	for pq.Size() > 0 {
		data, err := pq.Pop(); if err != nil {
			return nil, UndefinedError(errors.New("some issue with priority queue implementation"))
		}
		node := data.Data.(*path)
		if _, ok := visited[ *node.node ]; ok {
			continue
		}

		visited[ *node.node ] = &route{
			via: node.via,
			weight:  data.CmpData,
		}

		if *node.node == *to || found {
			found = true
			if *node.node == *to && (least == nil || (least != nil && least.weight > data.CmpData )) {
				least = &route{
					via:    node.node,
					weight: data.CmpData,
				}
			}
			if data.CmpData > least.weight {
				route, err := genRoute(visited, least, to, from); if err != nil {
					return nil, err
				}
				return route, nil
			}
		}

		edges, err := g.GetEdges(node.node); if err == nil {
			addToQue(visited, pq, edges, &unDirGraph.Edge{Node: node.node, Weight: data.CmpData})
		}
	}
	if found == true {
		route, err := genRoute(visited, least, to, from); if err != nil {
			return nil, err
		}
		return route, nil
	}

	return nil, NoConnectivity(errors.New(""))
}

type UndefinedError error

type finRoute struct {
	node *unDirGraph.Node
	via *unDirGraph.Node
	weight int
}

func genRoute(visited map[unDirGraph.Node]*route, least *route, to, from *unDirGraph.Node) (fin []*finRoute, err error) {
	//fin = append(fin, least)
	if node, ok := visited[ *least.via ]; ok {
		fin = append(fin, &finRoute{
			node:   to,
			via:    node.via,
			weight: node.weight,
		})
		least = node
	}
	for {
		if node, ok := visited[ *least.via ]; ok {
			fin = append(fin, &finRoute{
				node:   least.via,
				via:    node.via,
				weight: node.weight,
			})
			if node.via == nil {
				fin[len(fin)-1].node = from
				return fin, nil
			}
			least = node
		} else {
			return nil, UndefinedError(errors.New("expected node for route"))
		}
	}
}

func TestDijkstraNoRoute(t *testing.T) {
	g := genRouteForDijsktra()

	s := &unDirGraph.Node{"s"}
	e := &unDirGraph.Node{"e"}

	if _, err := ShortestPathDijkstra(g, s, e); err != nil {
		t.Fatal("expected true, got false")
	} else {
		//for p := range path {
		//	t.Log(path[p].node, "via", path[p].via, path[p].weight)
		//}
	}
}

func TestDijkstra(t *testing.T) {
	g := genRouteForDijsktra()

	s := &unDirGraph.Node{"s"}
	e := &unDirGraph.Node{"e"}

	if _, err := ShortestPathDijkstra(g, s, e); err != nil {
		t.Fatal("expected true, got false")
	} else {
		//for p := range path {
		//	t.Log(path[p].node, "via", path[p].via, path[p].weight)
		//}
	}
}


func genRouteForDijsktra() *unDirGraph.Graph {
	gr := unDirGraph.New()

	s := &unDirGraph.Node{"s"}
	a := &unDirGraph.Node{"a"}
	b := &unDirGraph.Node{"b"}
	d := &unDirGraph.Node{"d"}
	h := &unDirGraph.Node{"h"}
	f := &unDirGraph.Node{"f"}
	g := &unDirGraph.Node{"g"}
	e := &unDirGraph.Node{"e"}
	k := &unDirGraph.Node{"k"}
	i := &unDirGraph.Node{"i"}
	j := &unDirGraph.Node{"j"}
	l := &unDirGraph.Node{"l"}
	c := &unDirGraph.Node{"c"}
	gr.AddNode(s, a, b, d, h, f, g, e, k, i, j, l, c)

	gr.AddEdge(s, &unDirGraph.Edge{Node: a, Weight: 7},
		&unDirGraph.Edge{Node: b, Weight: 2})
	gr.AddEdge(a, &unDirGraph.Edge{Node: b, Weight: 3},
		&unDirGraph.Edge{Node: d, Weight: 4})
	gr.AddEdge(b, &unDirGraph.Edge{Node: h, Weight: 1},
		&unDirGraph.Edge{Node: d, Weight: 4})
	gr.AddEdge(d, &unDirGraph.Edge{Node: f, Weight: 5})
	gr.AddEdge(f, &unDirGraph.Edge{Node: h, Weight: 3})
	gr.AddEdge(h, &unDirGraph.Edge{Node: g, Weight: 2})
	gr.AddEdge(g, &unDirGraph.Edge{Node: e, Weight: 2})
	gr.AddEdge(e, &unDirGraph.Edge{Node: k, Weight: 5})
	gr.AddEdge(k, &unDirGraph.Edge{Node: i, Weight: 4},
		&unDirGraph.Edge{Node: j, Weight: 4})
	gr.AddEdge(i, &unDirGraph.Edge{Node: j, Weight: 6},
		&unDirGraph.Edge{Node: l, Weight: 4})
	gr.AddEdge(j, &unDirGraph.Edge{Node: l, Weight: 4})
	gr.AddEdge(l, &unDirGraph.Edge{Node: c, Weight: 2})
	gr.AddEdge(c, &unDirGraph.Edge{Node: s, Weight: 3})

	return gr
}

//	⟷
//	⭡
//	⭣
//	⭥
//	⭠
//	⭢
// ⬉ ⬈
// ⬊ ⬋
