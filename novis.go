package novis

import (
	"path"
	"strings"
	"sync"
)

// Used for singleton implementations
var novis = New()

// Add adds a new branch to the root returning the new branch
func Add(name, path string) *Branch {
	return novis.Add(name, path)
}

// Rev reversees the dotted name lookup path and returns the
// url path for the name, if param values are provided they will
// be used in place of any params in the path
func Rev(name string, values ...string) string {
	return novis.Rev(name, values...)
}

// GetBranch returns a branch by look up name
// For example "foo.bar" would return the "bar" branch if
// the branch exists
func GetBranch(name string) *Branch {
	return novis.GetBranch(name)
}

// Novis is the root all branches grow from
type Novis struct {
	Root *Branch
}

// GetBranch returns a branch by look up name
// For example "foo.bar" would return the "bar" branch if
// the branch exists
func (novis *Novis) GetBranch(name string) *Branch {
	var ok bool
	route := strings.Split(name, ".")
	last := route[len(route)-1]
	branch := novis.Root
	for i := 0; i < len(route); i++ {
		branch, ok = branch.Get(route[i])
		if !ok {
			break
		}
		if branch.name == last {
			break
		}
	}
	return branch
}

// Add adds a new branch to the root returning the new branch
func (novis *Novis) Add(name, path string) *Branch {
	return novis.Root.Add(name, path)
}

// Rev reversees the dotted name lookup path and returns the
// url path for the name, if param values are provided they will
// be used in place of any params in the path
func (novis *Novis) Rev(name string, values ...string) string {
	var ok bool
	var parts []string
	var params []string
	route := strings.Split(name, ".")
	branch := novis.Root              // Starting Node
	for i := 0; i < len(route); i++ { // Traverse the branches
		branch, ok = branch.Get(route[i])
		if !ok {
			break
		}
		parts = append(parts, branch.path)
		params = append(params, branch.params...)
	}
	p := path.Join(parts...) // join the parts together
	for i := 0; i < len(values) && i < len(params); i++ {
		p = strings.Replace(p, params[i], values[i], -1)
	}
	return p
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
	lock     sync.Mutex
	name     string
	path     string
	params   []string
	branches map[string]*Branch
	parent   *Branch
}

// Rel returns the branch relative path
func (branch *Branch) Rel() string {
	return branch.path
}

// Returns the full branch path from root to tip
func (branch *Branch) Path() string {
	parts := []string{branch.path}
	for {
		b := branch.parent
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
func (branch *Branch) Add(name, path string, params ...string) *Branch {
	b := NewBranch(name, path, branch, params...)
	branch.lock.Lock()
	branch.branches[name] = b
	branch.lock.Unlock()
	return b
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
