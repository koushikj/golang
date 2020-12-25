package main

import "fmt"

type runInterface interface {
	run()
}

func boot (r runInterface)  {
	fmt.Println("im from boot",r)
	switch r.(type) {
	case local:
		fmt.Println("im from boot and the type is local",r.(local).name)
		r.run()
	case remote:
		fmt.Println("im from boot and the type is remote",r.(remote).name)
		r.run()
	}
}
type local struct {
	name string
}

type remote struct {
	name string
}

func (local) run()  {
	fmt.Println("im local method")
}

func (remote) run()  {
	fmt.Println("im remote method")
}

func main()  {
	l1:= local{
		"local instance 1 ",
	}
	l2 := local{
		"local instance 2",
	}
	r1 := remote{
		"remote instance 1",
	}
	r2:= remote{
		"remote instance 2",
	}
	l1.run()
	l2.run()
	r1.run()
	r2.run()

	boot(l1)
	boot(l2)
	boot(r1)
	boot(r2)

}