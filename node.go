package chief

// Node node is a manager of multiple pools
type Node struct {
	Pools map[string]*Pool
}

// NewNode create a new Node
func NewNode() *Node {
	return &Node{
		Pools: make(map[string]*Pool),
	}
}
