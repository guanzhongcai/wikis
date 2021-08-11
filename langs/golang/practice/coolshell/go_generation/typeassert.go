package go_generation

import "fmt"

// Container is a generic container, accepting anything.
type Container []interface{}

func (c *Container) Put(elem interface{}) {
	*c = append(*c, elem)
}

func (c *Container) Get() interface{} {
	elem := (*c)[0]
	*c = (*c)[1:]
	return elem
}

func Usage()  {
	intContainer := &Container{}
	intContainer.Put(7)
	intContainer.Put(42)

	elem, ok := intContainer.Get().(int)
	if !ok {
		fmt.Println("unable to read an int from intContainer")
	} else {
		fmt.Printf("type assert example: %d (%T)\n", elem, elem)
	}
}

// output
// type assert example: 7 (int)