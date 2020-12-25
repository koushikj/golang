package main

import (
	"fmt"
)

// package level scope
var a int
var b string
var c bool
var d float32


func main()  {
	//short hand variable - declare and assign
	fmt.Println(`   short hand variable - declare and assign`)

	x:=42
	y:="James Bond"
	z:=true
	fmt.Println(x,y,z)
	fmt.Printf("%T %v\n",x,x)
	fmt.Printf("%T %v\n",y,y)
	fmt.Printf("%T %v\n",z,z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(`   package level scope`)

	fmt.Printf("%T %v\n",a,a)
	fmt.Printf("%T %v\n",b,b)
	fmt.Printf("%T %v\n",c,c)
	fmt.Printf("%T %v\n",d,d)

	fmt.Println(`   create custom type`)

	type myint int
	var a1 myint
	fmt.Printf("%T %v\n",a1,a1)
	a1=29
	fmt.Printf("%T %v\n",a1,a1)
	var a2 = int(a1)
	fmt.Printf("%T %v\n",a2,a2)

	fmt.Println("\n\n\t----- numerals -----\n")

	x1 := 191
	fmt.Printf("%v in decimal %d\n",x1,x1)
	fmt.Printf("%v in binary %b\n",x1,x1)
	fmt.Printf("%v in hexa decimal %x\n",x1,x1)
	fmt.Printf("%v in hexa decimal %#x\n",x1,x1)

	fmt.Println("\n\n\t----- cosntants -----\n")

	const(
		x2=3;
		x3 float32 =1;
		x4 bool = x2==x3;
	)
	fmt.Printf("%d %T \n",x2,x2)
	fmt.Printf("%v %T \n",x3,x3)
	fmt.Printf("%v %T \n",x4,x4)



	fmt.Println("\n\n\t----- bit shifting -----\n")

	x5:=9
	fmt.Printf("%v in decimal %d \n",x5,x5)
	fmt.Printf("%v in binary %b \n",x5,x5)
	fmt.Printf("%v in hex %#x \n",x5,x5)

	fmt.Printf("after left shift %v by 1 bit \n",x5)
	x6:=x5<<1
	fmt.Printf("%v in decimal %d \n",x6,x6)
	fmt.Printf("%v in binary %b \n",x6,x6)
	fmt.Printf("%v in hex %#x \n",x6,x6)

	fmt.Printf("after right shift %v by 1 bit \n",x5)
	x7:=x5>>1
	fmt.Printf("%v in decimal %d \n",x7,x7)
	fmt.Printf("%v in binary %b \n",x7,x7)
	fmt.Printf("%v in hex %#x \n",x7,x7)

	fmt.Println("\n\n\t----- iota -----\n")
	const(
		x8=2020
		a = iota + x8
		b= iota + x8
		c= iota + x8
		d= iota + x8
	)
	fmt.Printf("%v \t %v \t %v \t %v \t\n",a,b,c,d)

	fmt.Println("\n\n\t----- for loop -----\n")
	for{
		fmt.Println("abc")
		break
	}
	for x9:=1;x9<10;{
		x9++
		//fmt.Println(x9)
	}
	for x10:=33;x10<129;x10++{
		//fmt.Printf("%d  %#U \n",x10,x10)
	}
	for x11:=10;x11<100;x11++{
		//fmt.Printf("%d mod 4 = %d\n",x11,x11%4)
	}


	fmt.Println("\n\n\t----- array -----\n")
	x15 := [5]int{10,11,12,13,14}
	x16 := [2]string{"abc","def"}
	fmt.Println(x15)
	fmt.Println(x16)
	for i,v := range x15{
		fmt.Println(i,v)
	}
	fmt.Printf("%T %v\n", x15, x15)
	fmt.Printf("%T %v\n", x16, x16)

	x17:= []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(x17[:3])
	fmt.Println(x17[3:6])
	fmt.Println(x17[6:])

	fmt.Println("\n\n\t----- slice -----\n")
	x12:= []int{10,21,212,122,121}
	x13:= []int{0,1,2,3,4}
	x14:= []int{}
	fmt.Println(x12)
	fmt.Println(x13)
	fmt.Println(x14)
	for i,v := range x12{
		fmt.Println(i,v)
	}
	fmt.Printf("%T %v\n",x14,x14)
	fmt.Println(len(x12))
	fmt.Println(x12[2:])
	fmt.Println(append(x12, x13...))
	x14=append(x14, 1,2,3,4,5)
	fmt.Println(x14)
	fmt.Printf("%s %d %d %d\n","capacity length and content of x14",cap(x14),len(x14),x14)
	fmt.Printf("%s %d %d %d\n","capacity length and content of x12",cap(x12),len(x12),x12)
	x14=append(x14[:2],x14[3:]...)
	fmt.Println(x14)

	fmt.Println("\n\n\t----- slice using make -----\n")
	x18 := make([]int,5,10)
	x18[2]=10
	x18[4]=12
	//x18[5]=10
	fmt.Printf("%s %d %d %d\n","slice x18 : capacity, size, content =",cap(x18),len(x18),x18)

	x19:= [][]string{{"abc","def"},{"wee","fef"},{"wpwp"},{"mwe","aaqw","asqee","qpwp"}}
	fmt.Println("slice of slice of string::",x19)
	for i,v := range x19{
		fmt.Print("\t\n",i)
		for j,v1 := range v{
			fmt.Println("\t\t",j,v1)
		}
	}

	fmt.Println("\n\n\t----- map -----\n")
	x20 := map[string]int{
		"abc":12,
		"def":12,
	}
	fmt.Println(x20)
	x21 := map[string]string{
		"abc":"value1",
		"def":"value2",
	}
	fmt.Println(x21,len(x21))

	x22 := map[int][]string{
		1:{"abc","def"},
		3:{"xyz","qwe"},
	}
	x22[10]=[]string{"ppqp","qwqp","qskwow"}
	_,ok := x22[1]
	if ok {
		delete(x22,1)
	}

	fmt.Println(x22)
	for k,v := range x22{
		fmt.Println("the key is::",k)
		for _,v1:=range v{
			fmt.Println("\t",v1)
		}
	}

}
