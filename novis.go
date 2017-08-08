package novis

import (
	"path"
	"strings"
)

// Used for singleton implementations
var novis = New()

// Add adds a new branch to the root returning the new branch
func Add(name, path string, params ...string) *Branch {
	return novis.Add(name, path, params...)
}

// Rev reverses the dotted name lookup path and returns the
// url path for the name, if param values are provided they will
// be used in place of any params in the path
func Rev(name string, values ...string) string {
	return novis.Rev(name, values...)
}

// GetBranch returns a branch by look up name
// For example "foo.bar" would return the "bar" branch if
// the branch exists
func GetBranch(name string) *Branch {
	return novis.Get(name)
}

// GetBranch returns a branch by look up name
// For example "foo.bar" would return the "bar" branch if
// the branch exists
func Get(name string) *Branch {
	return novis.Get(name)
}

// Novis is the root all branches grow from
type Novis struct {
	Root *Branch
}

// GetBranch returns a branch by look up name
// For example "foo.bar" would return the "bar" branch if
// the branch exists
func (novis *Novis) GetBranch(name string) *Branch {
	return novis.Get(name)
}

// GetBranch returns a branch by look up name
// For example "foo.bar" would return the "bar" branch if
// the branch exists
func (novis *Novis) Get(name string) *Branch {
	var branch *Branch
	route := strings.Split(name, ".")
	last := route[len(route)-1]
	for branch = range novis.traverse(name) {
		if branch.name == last {
			break
		}
	}
	return branch
}

// Add adds a new branch to the root returning the new branch
func (novis *Novis) Add(name, path string, params ...string) *Branch {
	return novis.Root.Add(name, path, params...)
}

// Rev reverses the dotted name lookup path and returns the
// url path for the name, if param values are provided they will
// be used in place of any params in the path
func (novis *Novis) Rev(name string, values ...string) string {
	parts := []string{}
	params := []string{}
	for branch := range novis.traverse(name) {
		parts = append(parts, branch.path)
		params = append(params, branch.params...)
	}
	p := path.Join(parts...) // join the parts together
	for i := 0; i < len(values) && i < len(params); i++ {
		p = strings.Replace(p, params[i], values[i], -1)
	}
	return p
}

// traverse follows the lookup path to the end placing branches
// ontp a receive only channel callers can range over
func (novis *Novis) traverse(lookup string) <-chan *Branch {
	return novis.Root.traverse(lookup)
}

// New constructs a new Novis innstance
func New() *Novis {
	return &Novis{
		Root: &Branch{
			branches: make(map[string]*Branch),
		},
	}
}

// Branch is a single url path node
type Branch struct {
	name     string
	path     string
	params   []string
	branches map[string]*Branch
	parent   *Branch
}

// traverse follows the lookup path to the end placing branches
// ontp a receive only channel callers can range over
func (branch *Branch) traverse(lookup string) <-chan *Branch {
	ch := make(chan *Branch)
	ok := false
	route := strings.Split(lookup, ".")
	go func() {
		defer close(ch)
		for _, name := range route {
			if branch == nil {
				break
			}
			branch, ok = branch.Get(name)
			if !ok {
				break
			}
			ch <- branch
		}
	}()
	return ch
}

// Rel returns the branch relative path
func (branch *Branch) Rel() string {
	return branch.path
}

// Returns the full branch path from root to tip
func (branch *Branch) Path() string {
	parts := []string{branch.path}
	b := branch
	for {
		b = b.parent
		if b == nil {
			break
		}
		parts = append([]string{b.path}, parts...)
	}
	return path.Join(parts...)
}

// Get returns a child branch on this branch by name
func (branch *Branch) Get(name string) (b *Branch, ok bool) {
	b, ok = branch.branches[name]
	return
}

// Add adds a new child branch to this branch
func (branch *Branch) Add(route, path string, params ...string) *Branch {
	parent := branch
	parts := strings.Split(route, ".")
	name := parts[len(parts)-1]
	if len(parts) > 1 {
		for b := range branch.traverse(route) {
			parent = b
		}
	}
	nb := NewBranch(name, path, parent, params...)
	parent.branches[name] = nb
	return nb
}

// NewBranch construcs a new Branch instance
func NewBranch(name, path string, parent *Branch, params ...string) *Branch {
	return &Branch{
		name:     name,
		path:     path,
		params:   params,
		branches: make(map[string]*Branch),
		parent:   parent,
	}
}
