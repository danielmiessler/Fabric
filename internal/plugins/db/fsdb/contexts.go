package fsdb

import "fmt"

type ContextsEntity struct {
	*StorageEntity
}

// Get Load a context from file
func (o *ContextsEntity) Get(name string) (ret *Context, err error) {
	var content []byte
	if content, err = o.Load(name); err != nil {
		return ret, err
	}

	ret = &Context{Name: name, Content: string(content)}
	return ret, err
}

func (o *ContextsEntity) PrintContext(name string) (err error) {
	var context *Context
	if context, err = o.Get(name); err != nil {
		return err
	}
	fmt.Println(context.Content)
	return err
}

type Context struct {
	Name    string
	Content string
}
