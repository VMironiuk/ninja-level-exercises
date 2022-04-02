package main

import "fmt"

type person struct {
	fname   string
	lname   string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fname, "is walking.")
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

func (t truck) transportationDevice() string {
	return "I'm a truck and I transport heavy weight things"
}

type sedan struct {
	vehicle
	luxury bool
}

func (s sedan) transportationDevice() string {
	return "I'm a sedan and I am very pretty"
}

type transportation interface {
	transportationDevice() string
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

type gator int

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

type flamingo bool

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful")
}

type swampCreature interface {
	greeting()
}

func bayou(sc swampCreature) {
	sc.greeting()
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(sl)
	for i := range sl {
		fmt.Println(i)
	}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(m)
	for k := range m {
		fmt.Println(k)
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	p1 := person{
		fname: "James",
		lname: "Bond",
		favFood: []string{
			"Pizza",
			"Pasta",
			"Beef",
			"Pork",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.fname)
	fmt.Println(p1.favFood)
	for _, v := range p1.favFood {
		fmt.Println(v)
	}
	s := p1.walk()
	fmt.Printf("%v", s)

	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: false,
	}
	sed := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "silver",
		},
		luxury: true,
	}
	fmt.Println(t)
	fmt.Println(sed)
	fmt.Println(t.color)
	fmt.Println(sed.color)

	fmt.Println(t.transportationDevice())
	fmt.Println(sed.transportationDevice())

	report(t)
	report(sed)

	var g1 gator
	g1 = 42
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = int(g1)
	fmt.Println(x)

	g1.greeting()

	var f flamingo
	bayou(g1)
	bayou(f)

	str := `i'm sorry dave i can't do that`
	fmt.Println(str)
	fmt.Println([]byte(str))
	fmt.Println(string([]byte(str)))
	fmt.Println(str[:14])
	fmt.Println(str[10:22])
	fmt.Println(str[17:])
	for _, v := range str {
		fmt.Printf("%c\n", v)
	}
}
