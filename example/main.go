package main

import (
	"fmt"
	"sort"

	closure_table "github.com/guidoxie/closure-table"
)

type Node struct {
	ID       uint
	ParentID uint
}

func (n *Node) GetID() uint {
	return n.ID
}

func (n *Node) GetParentID() uint {
	return n.ParentID
}

func main() {
	nodes := []closure_table.INode{
		&Node{ID: 11, ParentID: 8},
		&Node{ID: 1, ParentID: 3},
		&Node{ID: 3, ParentID: 0},
		&Node{ID: 2, ParentID: 3},
		&Node{ID: 5, ParentID: 4},
		&Node{ID: 4, ParentID: 0},
		&Node{ID: 6, ParentID: 0},
		&Node{ID: 8, ParentID: 7},
		&Node{ID: 9, ParentID: 7},
		&Node{ID: 7, ParentID: 0},
		&Node{ID: 10, ParentID: 0}}
	ct := closure_table.Generate(nodes)
	sort.Slice(ct, func(i, j int) bool {
		if ct[i].TreeID == ct[j].TreeID {
			return ct[i].Ancestor < ct[j].Ancestor
		}
		return ct[i].TreeID < ct[j].TreeID
	})
	for _, c := range ct {
		fmt.Printf("TreeID: %d, Ancestor: %d, Descendant: %d, Distance: %d\n", c.TreeID, c.Ancestor, c.Descendant, c.Distance)
	}
}
