package main

import "fmt"

// Do not edit the class below except
// for the depthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	// Write your code here.
	array = append(array, n.Name)
	for _, child := range n.Children {
		array = child.DepthFirstSearch(array)
	}
	return array
}

func main() {
	root := Node{Name: "A"}
	b := Node{Name: "B"}
	c := Node{Name: "C"}
	d := Node{Name: "D"}
	e := Node{Name: "E"}
	f := Node{Name: "F"}
	g := Node{Name: "G"}
	h := Node{Name: "H"}
	i := Node{Name: "I"}
	j := Node{Name: "J"}
	k := Node{Name: "K"}

	root.Children = []*Node{&b, &c, &d}
	b.Children = []*Node{&e, &f}
	d.Children = []*Node{&g, &h}
	f.Children = []*Node{&i, &j}
	g.Children = []*Node{&k}
	result := []string{}
	fmt.Println(root.DepthFirstSearch(result))
}
