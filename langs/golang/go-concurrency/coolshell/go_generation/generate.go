package go_generation

import "fmt"

//go:generate ./gen.sh ./template/container.tmp.go gen uint32 container
func generateUin32Example()  {
	var u int32 = 42
	c := NewUint32Container()
	c.Put(u)
	v := c.Get()
	fmt.Printf("generateUin32Example: %d (%T)\n", v, v)
}

//go:generate ./gen.sh ./template/container.tmp.go gen string container
func generateStringExample()  {
	var s string = "Hello"
	c := NewStringContainer()
	c.Put(s)
	v := c.Get()
	fmt.Printf("generateStringExample: %s (%T)\n", v, v)
}
