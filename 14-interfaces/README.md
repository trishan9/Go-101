# Interfaces

- An interface in Go is a set of method signatures without any implementation.

### Interfaces are implemented implicitly

- If a type (struct) implements all the methods defined in the interface, it is said to implement that interface. A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Interfaces provide several benefits in software development:

- **Decoupling**: allows for loose coupling between components, enabling easier maintenance and extensibility.
- **Abstraction**: provides a way to define abstract behavior that can be implemented by different types.
- **Polymorphism**: enables polymorphic behavior, allowing different types to be used interchangeably.
- **Modularity and Separation of Concerns**: Interfaces help in breaking down complex systems into smaller, more manageable modules.

### Simple Example

In the following example, `Shape` must be able to return its area. Both `Square` and `Circle` fulfill the interface.

```go
type Shape interface {
    Area() float64
}

type Square struct {
    side float64
}

func (s Square) Area() float64 {
    return s.side * s.side
}

type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}
```

Now, we can write a function that takes a `Shape` interface and calculates the area:

```go
func printArea(s Shape) {
    fmt.Println("Area:", s.Area())
}
```

We can call `printArea` with instances of `Square` or `Circle` because they both implement the `Shape` interface:

```go
sq := Square{side: 5}
c := Circle{radius: 3}

printArea(sq) // Output: Area: 25
printArea(c)  // Output: Area: 28.274333882308138
```

## Type Assertions in Go

Type assertion in Go is a way to extract the underlying value of an interface. Go interfaces work by storing a pair: the value and the type of the value. Type assertion is a way to retrieve the dynamic value from the interface. They are typically used when you have an interface value and you want to access methods or fields specific to the concrete type.

Syntax: `value, ok := interfaceValue.(ConcreteType)`
The `ok` value is a boolean that reports whether the assertion succeeded or not.

### Example

```go
var i interface{} = "hello"

s := i.(string)
fmt.Println(s)

s, ok := i.(string)
fmt.Println(s, ok)

f, ok := i.(float64)
fmt.Println(f, ok)

f = i.(float64) // panic
fmt.Println(f)
```

### Another Example

```go
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	var shape Shape
	shape = Circle{Radius: 5}

	circle, ok := shape.(Circle)
	if ok {
		fmt.Println("Circle Area:", circle.Area()) // Outputs: Circle Area: 78.5
	} else {
		fmt.Println("Not a Circle")
	}
}

```
