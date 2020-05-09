package main

import "fmt"

func main() {

	exercise1()
	fmt.Println()
	exercise2()
	fmt.Println()
	exercise5()
	fmt.Println()
	exercise7()
	fmt.Println()
	exercise8()
	fmt.Println()
	exercise10()

}

func exercise1(){

	x :=  [5]int{42, 43, 54, 77, 123}

	for i,v := range x{
		fmt.Println(i,v)
	}
	fmt.Printf("%T\n",x)
}


func exercise2(){

	x :=  []int{42, 43, 54, 77, 123, 12, 5, 6, 88, 9, 90}

	for i,v := range x{
		fmt.Println(i,v)
	}
	fmt.Printf("%T\n",x)
}

func exercise3(){

	x :=  []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Println(x)
}

func exercise4(){

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x, 52)

	fmt.Println(x)

	x = append(x, 53, 54, 55)

	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}

	x = append(x, y...)

	fmt.Println(x)
}

func exercise5(){

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := append(x[:3],x[6:]...)

	fmt.Println(y)

}

func exercise6(){

	x := make([]string, 50, 500)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// put a value into each of the 50 positions in the length of the slice
	// we are putting values into the 50 positions that are the length of the slice
	for i := 0; i < 50; i++ {
		x[i] = fmt.Sprintf("Position %d", i)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// append values to the slice grows the length of the slice
	// the underlying array "capacity" of 500 is the same
	x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

func exercise7(){

	tempSlice1 := []string{"James", "Bond", "Shaken, not stirred"}
	tempslice2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	var x = [][]string{tempSlice1, tempslice2}

	fmt.Println(x)

	for i, v := range x{
		fmt.Println("slice number: ",i)
		for j, u := range v{
			fmt.Println(j,u)
		}
	}
}

func exercise8(){

	m := map[string][]string{
		"bond_james": []string{"Shaken", "not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr": []string{"Being evil", "Ice cream", "Sunsets"},
	}

	m["pereira_mario"] = []string{"beer", "tech", "the good life"}

	for _, f := range m{
		fmt.Println(f)
		for j,v := range f{
			fmt.Println("\t",j, v)
		}
	}
}
func exercise10(){

	m := map[string][]string{
		"bond_james": []string{"Shaken", "not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr": []string{"Being evil", "Ice cream", "Sunsets"},
	}

	m["pereira_mario"] = []string{"beer", "tech", "the good life"}


	delete(m, "no_dr")

	for k, f := range m{
		fmt.Println(k)
		for j,v := range f{
			fmt.Println("\t",j, v)
		}
	}
}
