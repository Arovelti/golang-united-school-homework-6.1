package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity == len(b.shapes) {
		return errors.New("capacity is full, sorry")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errors.New("out of range")
	}

	for index, val := range b.shapes {
		if i == index {
			return val, nil
		}
	}
	return nil, errors.New("there is no Shape with such index")
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if s, err := b.GetByIndex(i); err != nil {
		return nil, err
	} else {
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return s, nil
	}
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	_, err := b.GetByIndex(i)
	if err != nil {
		return nil, err
	}

	newShapes := make([]Shape, len(b.shapes))
	copy(newShapes, b.shapes)
	b.shapes[i] = shape

	return newShapes[i], nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sumPerimeter float64
	for _, v := range b.shapes {
		sumPerimeter += v.CalcPerimeter()
	}
	return sumPerimeter
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sumArea float64
	for _, v := range b.shapes {
		sumArea += v.CalcArea()
	}
	return sumArea
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var exist bool
	for i := 0; i < len(b.shapes); i++ {
		if _, ok := b.shapes[i].(*Circle); ok {
			exist = true
			b.ExtractByIndex(i)
			i--
		}
	}
	if !exist {
		return errors.New("no Circles has found")
	}
	return nil
}
