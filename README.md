# medium

## Content
- [Solid Design Principle](#Solid-Design-Principle)
- [Object Oriented Programming](#Object-Oriented-Programming)
# Solid-Design-Principle
- The acronym SOLID was derived by Robert C.Martin also known as Uncle Bob. It defines the principle used in object-oriented design. These principles serve as a guide for software developers to build applications that are extensible and maintainable. Although Golang is partly object-oriented, as such it does not implement principles like inheritance, these SOLID principles can still be utilized with Go to build maintainable applications.

- The SOLID means:

- S: Single Responsibility Principle (SRP)
- Single Responsibility Principle (SRP): This states that a type should have one primary responsibility, as a result, it should only have one reason to change which should be the primary responsibility. For instance, let’s assume you have a chef, his or her responsibility is to prepare meals, and delicacies for you. For no reason should the chef start styling your hair or decide to start doing your laundry. This is exactly what the SRP is trying to explain. This SRP principle extends to another principle called SOC (separation of concerns). This simply means different problems or different concerns/functions that a system solves should reside in a different construct or package. To further understand these principles I have a written a simple example of a Journal that shows how the SRP is implemented

- O: Open-Closed Principle(OCP)

- L: Liskov Substitution Principle(LSP)

- I:Interface Segregation Principle(ISP)

- D: Dependency Inversion Principle(DIP)

# Object-Oriented-Programming
## ENCAPSULATION
- Encapsulation keeps data safe from external interferences. In oop languages, this is done at object level by the use of access modifiers. But in Golang, this is done at package level.
- Go does not have access modifiers such as public, private or protected. Instead, there are exported and unexported fields. This difference comes from the first letters of the variables. Lowercase means unexported, capital means exported. You can think of this difference as private or public.
```go
type Customer struct {
	id   int
	name string
}

func (c *Customer) GetID() int {
	return c.id
}

func (c *Customer) GetName() string {
	return c.name
}
```
- Another difference we see is that there are no classes but we have structures instead. The reason is that Go came out as an alternative for C and C++ and structs are mainly used in those languages.
- These are package level visibilities, so variables starting with lowercase letters are not visible from outside the package they are in.

## INHERITANCE
- In OOP, computer programs are designed in such a way where everything is an object that interacts with one another. Inheritance is an integral part of OOP languages which lets the properties of one class to be inherited by the other. It basically helps in reusing the code and establish a relationship between different classes.
```go
type Vehicle struct {
	Seats int
	Color string
}

type Car struct {
	Vehicle
}

type MotorCycle struct {
	Base Vehicle
}
```
- In Go, we can use compositions to get similar results.
- In the above example, car struct is being embedded by another structure anonymously.
- This means that we have direct access to the fields. This method is similar to what we are used to on the OOP side.

## CONSTRUCTORS

- Constructors are special functions for reliably creating multiple instances of similar objects in a class-based oop language.
- There is no specific definition for initialization in GO, but we can do the same by writing a function.
```go
type Customer struct {
	Name string
	Age  int
}

func (c *Customer) NewCustomer(name string, age int) {
	c.Name = name
	c.Age = age
}
```
- The NewCustomer function above does the same job as the constructors in oop languages.
- As a small note here, there is a concept of Default or Empty values in Go. Everything is initialized to something.
```go
type Customer struct {
   Name string // " "
   Age  int    // 0
}
```
- This error, which is called “NullReferenceException” in C#, “NullPointerException” in Java and where no value is assigned while initializing, does not appear in GO.

## POLYMORPHISM

- Golang supports polymorphism through interfaces only. A type implements an interface if it provides definitions for all the methods declared in the interface.
```go
type Instrument interface {
	Play()
}

func PlayInstrument(i Instrument) {
	i.Play()
}

type Guitar struct {
	Type string
}

func (g Guitar) Play() {
	fmt.Println("Guitar Sounds")
}
```
- In the above example, Guitar implements the Instrument interface by providing definition for Play method which is the only method declared in the interface.

## ENUMS

- Enum is a powerful data type for Enumerated Data in oop languages. On the Go side, we can only use constants and assign incrementing values by the use of iota. Although the benefit of using Enum is that we don’t need to give the order explicitly.
```go
const (
	Guitar = iota // 0
	Violin        // 1
	Piano         // 2
	Drums         // 3
)
```

## GENERICS

- Generics in Go is a new concept came in with the 1.18 version. It is called type parameters and we are be able to get the types of the parameters of the functions from an interface that we have to define which can contain multiple acceptable types. In this way, we can get rid of parameter type dependency and create reusable functions. We can also use this to create more advanced type parameters as well.
```go
type Instrument interface {
	Play()
}

func PlayInstrument[I Instrument](i I) {
	i.Play()
}

type Guitar struct {
	Sound string
}

func (g Guitar) Play() {
	fmt.Println(g.Sound)
}
```
- In the above example, the PlayInstrument function accepts any value that fulfills the contract of “Instrument” interface.

## CONCLUSION
- Golang promotes an interesting take on object-oriented programming with the use of interfaces although it is not meant to be an OOP language. The implementation of object oriented programming in Go is incredibly flexible, direct and lightweight.
- The usage of OOP in Go might feel strange, but getting all the advantages of Go without the resource-hungry frameworks is definitely a big plus for many developers.