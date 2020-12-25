package main

import "fmt"

func main(){
	fmt.Println("hello")

	fmt.Println(`
	strcut and embeded stuct
	`)

	type employee struct {
		name string
		id int
	}
	e1:= employee{
		name: "koushik",
		id : 12,
	}
	e2:= employee{
		name: "jayakumar",
		id : 121,
	}
	fmt.Println(e1,e2)
	fmt.Println(e1.id,e2.name)

	type manager struct {
		employee
		teamMembers int
	}
	m1:= manager{
		employee: employee{
			name: "mgrname",
			id:   10,
		},
		teamMembers:1,
	}
	fmt.Println(m1)

	type person1 struct {
		first string
		last string
		favourites []string
	}
	p3:=person1{
		first: "James",
		last: "Bond",
		favourites: []string{"shooting","driving"},
	}
	p4:=person1{
		first: "David",
		last:"Goggins",
		favourites: []string{"running","excercise"},
	}
	fmt.Println(p3,p4)
	fmt.Println(p3.favourites,p4.favourites)

	x23 := map[string]person1{
		p3.last:p3,
		p4.last:p4,
	}
	fmt.Println(x23)
	fmt.Println(x23["Bond"])
	fmt.Println(x23["David"])
	fmt.Println(x23["Goggins"])
	for i,v := range x23{
		fmt.Println("key:",i,"::",v.first,v.last)
		for j,v1 := range v.favourites{
			fmt.Print("\t")
			fmt.Println("key:",j,"value:",v1)

		}
	}

	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheeler bool
	}
	type sedan struct {
		vehicle
		luxury bool
	}
	t1:= truck{
		vehicle:vehicle{
			doors: 0,
			color: "white",
		},
		fourWheeler:true,
	}
	s1:= sedan{
		vehicle:vehicle{
			doors: 1,
			color: "blue",
		},
	}

	s1.color="pink"
	s1.luxury=true
	s1.vehicle.doors=10

	fmt.Println(t1,s1)
	fmt.Println(t1.color,t1.doors,t1.fourWheeler,t1.vehicle.doors)
	fmt.Println(s1.color,s1.doors,s1.luxury,s1.vehicle.doors)

	//anonymous struts
	s2:= struct {
		name string
		age int
		hobbies []string
		marks map[string]int
	}{
		name: "ron",
		age: 10,
		hobbies: []string{"playing","reading"},
		marks: map[string]int{"science":100,"english":90,"maths":80},
	}

	fmt.Println(s2.name,s2.age,s2.hobbies,s2.marks)
	for i,v := range s2.hobbies{
		fmt.Printf("\t hobby :%v of %v is :%v \n",i,s2.name,v)
	}
	for i,v:= range s2.marks{
		fmt.Println(i,"=",v)
	}

}
