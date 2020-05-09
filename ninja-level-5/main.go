package main

import "fmt"

type person struct{
	firstName string
	lastName string
	favoritIcecrams []string
}

type vehicle struct{
	doors int
	color string
}

type truck struct{
	vehicle
	fourWheel bool
}

type sedan struct{
	vehicle
	luxury bool
}

func main() {

	exercise1()
	fmt.Println()
	exercise2()
	fmt.Println()
	exercise3()
	fmt.Println()
	exercise4()
}

func exercise1(){

	p1 := person{
		firstName: "mario",
		lastName: "pereira",
		favoritIcecrams: []string{ "chocolat", "banana", "strawberry"},
	}

	p2 := person{
		firstName: "james",
		lastName: "bond",
		favoritIcecrams: []string{ "vanilla", "martini"},
	}

	fmt.Printf("the first name is: %v and the last name is %v\n", p1.firstName, p1.lastName)

	for _, v := range p1.favoritIcecrams{
		fmt.Println("\t", v)
	}

	fmt.Printf("the first name is: %v and the last name is %v\n", p2.firstName, p2.lastName)

	for _, v := range p2.favoritIcecrams{
		fmt.Println("\t",v)
	}
}


func exercise2(){

	p1 := person{
		firstName: "mario",
		lastName: "pereira",
		favoritIcecrams: []string{ "chocolat", "banana", "strawberry"},
	}

	p2 := person{
		firstName: "james",
		lastName: "bond",
		favoritIcecrams: []string{ "vanilla", "martini"},
	}

	myMap := map[string]person{
		p1.lastName : p1,
		p2.lastName : p2,
	}

	for _, v := range myMap {
		fmt.Printf("the first name is: %v and the last name is %v\n", v.firstName, v.lastName)

		for _, u := range v.favoritIcecrams{
			fmt.Println("\t", u)
		}

	}
}

func exercise3(){

	myTruck := truck{
		vehicle: vehicle{
			color: "white",
			doors: 2,
		},
		fourWheel: false,
	}

	mySedan := sedan{
		vehicle: vehicle{
			color: "black",
			doors: 5,
		},
		luxury: true,
	}

	fmt.Println(myTruck)
	fmt.Println(mySedan)

	fmt.Printf("the truck is %v and has %v doors but the fourWhell thing is %v\n", myTruck.color, myTruck.doors, myTruck.fourWheel)

	fmt.Printf("the sedan is %v and has %v doors but the luxury thing is %v\n", mySedan.color, mySedan.doors, mySedan.luxury)
}

func exercise4(){

	s := struct{
		first string
		last string
		hobbies []string
		mapHobbies map[string][]string
	}{
		first: "mario",
		last: "pereira",
		hobbies: []string{"batata", "esbelela", "potato"},
		mapHobbies: map[string][]string{
			"mario": {"esbeler", "potato", "tratata"},
		},
	}

	fmt.Println(s)
	fmt.Println(s.first)
	fmt.Println(s.last)

	for i, v := range s.hobbies{
		fmt.Println(i, v)
	}

	for k, v := range s.mapHobbies{
		fmt.Println(k, v)
	}
}