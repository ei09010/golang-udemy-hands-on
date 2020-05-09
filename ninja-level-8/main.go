package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type user struct {
	First string
	Age   int
}

type Profile struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

type newUser struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	exercise1()
	fmt.Println()
	exercise2()
	fmt.Println()
	exercise3()
	fmt.Println()
	exercise4()
	fmt.Println()
	exercise5()
}


func exercise1(){

	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	userByteSlice, err := json.Marshal(users)

	if err != nil{
		fmt.Println("returned error form marshling: ", err)
	}

	fmt.Println(string(userByteSlice))

}

func exercise2(){

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	fmt.Println()

	var profiles []Profile

	err := json.Unmarshal([]byte(s), &profiles)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(profiles)

	for _, person := range profiles{
		fmt.Println("\t", person.First, person.Last, person.Age)
		for _, saying := range person.Sayings{
			fmt.Println("\t","this person has the following saying: ", saying)
		}
	}
}

func exercise3(){

	u1 := newUser{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := newUser{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := newUser{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []newUser{u1, u2, u3}

	myEncoder := json.NewEncoder(os.Stdout)

	err := myEncoder.Encode(users)

	if err != nil {
		fmt.Println(err)
	}

}

func exercise4(){
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)

	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)
}

type ByAge []newUser

func (a ByAge) Len() int { return len(a)}
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }


type ByLast []newUser

func (b ByLast) Len() int { return len(b)}
func (b ByLast) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByLast) Less(i, j int) bool { return b[i].Last < b[j].Last }

func exercise5(){
	u1 := newUser{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := newUser{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := newUser{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []newUser{u1, u2, u3}

	fmt.Println(users)

	sort.Sort(ByAge(users))

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t",v)
		}
	}

	fmt.Println("---------------")

	sort.Sort(ByLast(users))

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t",v)
		}
	}

}