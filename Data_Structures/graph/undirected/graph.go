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
		nodes: make(map[Node]struct{}),
		edges: make(map[Node][]*Edge),
		lock:  sync.RWMutex{},
	}
}

type Graph struct {
	nodes map[Node]struct{}
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
		for node := range g.nodes {
			select {
			case <-ch:
				break LOOP
			case ch <- node:
			}
		}
		g.lock.RUnlock()
		return nil
	}
	g.lock.RUnlock()
	return NoNodes(errors.New(""))
}

func (g *Graph) DelNode(n *Node) {
	g.lock.Lock()
	delete(g.nodes, *n)
	delete(g.edges, *n)
	g.lock.Unlock()
}

func (g *Graph) AddNode(nodes... *Node) {
	g.lock.Lock()
	for n := range nodes {
		g.nodes[ *nodes[n] ] = struct{}{}
	}
	g.lock.Unlock()
}

func (g *Graph) AddEdge(n *Node, edge... *Edge) {
	g.lock.Lock()
	for e := range edge {
		g.edges[*n] = append(g.edges[*n], edge[e])
		g.edges[ *edge[e].Node ] = append(g.edges[ *edge[e].Node ], &Edge{Node: n,	Weight: edge[e].Weight})
	}
	g.lock.Unlock()
}

type NoEdges error
type NoNodes error
