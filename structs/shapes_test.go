package main

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {

    checkArea := func(t testing.TB, shape Shape, want float64)  {
        t.Helper()
        got := shape.Area()

        if got != want {
            t.Errorf("got %g want %g", got, want)
        }
    }

    t.Run("rectangles", func(t *testing.T) {
        rectangle := Rectangle{12.0, 6.0}
        checkArea(t, rectangle, 72.0)
    })

    t.Run("circles", func(t *testing.T) {
        circle := Circle{10}
        checkArea(t, circle, 314.1592653589793)
    })
}

type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width   float64
    Height  float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}