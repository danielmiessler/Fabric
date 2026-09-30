package db

// Storage is the contract for a named-entity store. Each implementation
// specifies the names that are valid. It rejects an invalid name with a
// typed error. Callers can map this error to a client error.
type Storage[T any] interface {
	Configure() (err error)
	Get(name string) (ret *T, err error)
	GetNames() (ret []string, err error)
	Delete(name string) (err error)
	// Exists reports false for an invalid name. It cannot show the
	// difference between a rejected name and an absent entry.
	Exists(name string) (ret bool)
	Rename(oldName, newName string) (err error)
	Save(name string, content []byte) (err error)
	Load(name string) (ret []byte, err error)
	ListNames(shellCompleteList bool) (err error)
}
