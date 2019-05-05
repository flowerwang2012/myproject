package main

import "fmt"

// 工厂模式

type Shape interface {
	draw()
}

//Circle /ˈsəːk(ə)l/ 圆形
type Circle struct{}

func (c *Circle) draw() { //Circle结构体实现了接口，所以实例化结构体
	fmt.Println("draw circle shape...")
}

//Square /skwɛː/ 正方形
type Square struct{}

func (s *Square) draw() {
	fmt.Println("draw square shape...")
}

//Rectangle /ˈrɛktaŋɡ(ə)l/ 长方形
type Rectangle struct{}

func (r *Rectangle) draw() {
	fmt.Println("draw rectangle shape...")
}

//ShapeFactory 对象工厂
type ShapeFactory struct{}

func (f *ShapeFactory) getShape(shape string) Shape {
	switch shape {
	case "circle":
		return new(Circle)
	case "square":
		return new(Square)
	case "rectangle":
		return new(Rectangle)
	}
	return nil
}

func main() {
	f := new(ShapeFactory)
	var s Shape
	s = f.getShape("circle") //s是接口类型的引用，指向实现类的对象
	s.draw()
}
