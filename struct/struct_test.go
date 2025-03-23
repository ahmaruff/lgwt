package structs

import "testing"


type Shape interface {
    Area() float64
}

func TestPerimeter(t *testing.T) {
    rect := Rectangle{10.0, 10.0}
    got := Perimeter(rect)
    want := 40.0

    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}


func TestArea(t *testing.T) {


    areaTest := []struct {
        s Shape
        want float64
    }{
        {Rectangle{10.0, 4.0}, 40.0},
        {Circle{10.0}, 314.1592653589793},
    }


    for _, test := range(areaTest) {
        got := test.s.Area()

        if got != test.want {
			t.Errorf("got %g want %g", got, test.want)
        }

    }
}
