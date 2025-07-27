package db

type Storage[T any] interface {
	Configure() (err error)
	Get(name string) (ret *T, err error)
	GetNames() (ret []string, err error)
	Delete(name string) (err error)
	Exists(name string) (ret bool)
	Rename(oldName, newName string) (err error)
	Save(name string, content []byte) (err error)
	Load(name string) (ret []byte, err error)
	ListNames(shellCompleteList bool) (err error)
}
