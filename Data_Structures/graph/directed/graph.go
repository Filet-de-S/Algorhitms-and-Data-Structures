package directed

import (
	"errors"
	"sync"
)

type Node struct {
	Data string
}

type Edge struct {
	Node *Node
	Weight Weight
}

type Weight = int

func New() *Graph {
	return &Graph{
		edges: make(map[Node][]*Edge),
		lock:  sync.RWMutex{},
	}
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Edge
	lock  sync.RWMutex
}

func (g *Graph) GetEdges(n *Node) ([]*Edge, error) {
	g.lock.RLock()
	if edges, ok := g.edges[*n]; ok {
		g.lock.RUnlock()
		return edges, nil
	}
	g.lock.RUnlock()
	return nil, NoEdges(errors.New(""))
}

func (g *Graph) GetEdgesNodes(n *Node) ([]*Node, error) {
	g.lock.RLock()
	if edges, ok := g.edges[*n]; ok {
		g.lock.RUnlock()
		nodes := make([]*Node, 0, len(edges))
		for n := range edges {
			nodes = append(nodes, edges[n].Node)
		}
		return nodes, nil
	}
	g.lock.RUnlock()
	return nil, NoEdges(errors.New(""))
}

func (g *Graph) GetNodesLive(ch chan Node) error {
	g.lock.RLock()
	if g.nodes != nil {
		LOOP:
		for i := range g.nodes {
			select {
			case <-ch:
				break LOOP
			case ch <- *g.nodes[i]:
			}
		}
		g.lock.RUnlock()
		return nil
	}
	g.lock.RUnlock()
	return NoNodes(errors.New(""))
}

func (g *Graph) GetNodes() ([]*Node, error) {
	g.lock.RLock()
	defer g.lock.RUnlock()
	if g.nodes != nil {
		return g.nodes, nil
	}
	return nil, NoNodes(errors.New(""))
}


func (g *Graph) DelNode(n *Node) {
	g.lock.Lock()
	//delete(g.nodes, *n)
	for i := range g.nodes {
		if *g.nodes[i] == *n {
			copy(g.nodes[i:], g.nodes[i+1:])
			g.nodes[len(g.nodes)-1] = nil
			g.nodes = g.nodes[:len(g.nodes)-1]
			break
		}
	}
	delete(g.edges, *n)
	g.lock.Unlock()
}

func (g *Graph) AddNode(nodes... *Node) {
	g.lock.Lock()
	for n := range nodes {
		g.nodes = append(g.nodes, nodes[n])
		//g.nodes[ *nodes[n] ] = struct{}{}
	}
	g.lock.Unlock()
}

func (g *Graph) AddEdge(n *Node, edge... *Edge) {
	g.lock.Lock()
	for e := range edge {
		g.edges[*n] = append(g.edges[*n], edge[e])
	}
	g.lock.Unlock()
}

type NoEdges error
type NoNodes error
