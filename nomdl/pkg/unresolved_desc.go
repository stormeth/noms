package pkg

import "github.com/attic-labs/noms/types"

const UnresolvedKind = 100

// UnresolvedDesc represents a named reference to a type.
type UnresolvedDesc struct {
	Namespace string
	Name      string
}

func (desc UnresolvedDesc) Kind() types.NomsKind {
	return UnresolvedKind
}

func (desc UnresolvedDesc) Equals(other types.TypeDesc) bool {
	if other.Kind() != UnresolvedKind {
		return false
	}
	d2 := other.(UnresolvedDesc)
	return d2.Namespace == desc.Namespace && d2.Name == desc.Namespace
}

func makeUnresolvedType(namespace, name string) *types.Type {
	return &types.Type{Desc: UnresolvedDesc{namespace, name}}
}
