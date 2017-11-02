package chief

// Pool a pool of primary keys
type Pool struct {
	Name string
}

// NewPool create a new pool
func NewPool(name string) *Pool {
	return &Pool{Name: name}
}
