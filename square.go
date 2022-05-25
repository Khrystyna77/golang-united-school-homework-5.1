package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (receiver *Square) End() Point {
	// implement m
	v := receiver.start
	var v2 uint = receiver.a
	endx := v.x + int(v2)
	endy := v.y + int(v2)
	Output := Point{x: endx, y: endy}
	//fmt.Println(Output)
	return Output

}

func (receiver *Square) Area() uint {
	// implement me
	var uinta uint = receiver.a
	uintAr := uinta * uinta
	return uintAr
}

func (receiver *Square) Perimeter() uint {
	// implement me
	var pers uint = receiver.a
	persq := 2 * (pers + pers)
	//fmt.Println(persq)
	return persq

}
