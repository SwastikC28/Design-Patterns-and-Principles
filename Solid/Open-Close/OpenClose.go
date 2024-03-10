package main

type Item interface {
	GetId() uint
	GetName() string
}

// Implements Items
type Book struct {
	id   uint
	name string
}

func (book *Book) GetId() uint {
	return book.id
}

func (book *Book) GetName() string {
	return book.name
}

func NewBook(name string) *Book {
	return &Book{
		id:   1,
		name: name,
	}
}

// Implements Items
type DVD struct {
	id   uint
	name string
}

func (dvd *DVD) GetId() uint {
	return dvd.id
}

func (dvd *DVD) GetName() string {
	return dvd.name
}

func NewDVD(name string) *DVD {
	return &DVD{
		id:   1,
		name: name,
	}
}

type User interface {
	GetName() string
	GetType() string
}

type Student struct {
	name       string
	typeOfUser string
}

func (student *Student) GetName() string {
	return student.name
}

func (student *Student) GetType() string {
	return student.typeOfUser
}

func NewStudent(name string) *Student {
	return &Student{
		name:       name,
		typeOfUser: "Student",
	}
}

type Faculty struct {
	name       string
	typeOfUser string
}

func (faculty *Faculty) GetName() string {
	return faculty.name
}

func (faculty *Faculty) GetType() string {
	return faculty.typeOfUser
}

func NewFaculty(name string) *Faculty {
	return &Faculty{
		name:       name,
		typeOfUser: "Faculty",
	}
}

type Library struct {
	Inventory []Item
	Users     []User
}

func (library *Library) AddInInventory(item Item) {
	library.Inventory = append(library.Inventory, item)
}

func (library *Library) AddUser(user User) {
	library.Users = append(library.Users, user)
}

func NewLibrary() *Library {
	return &Library{
		Inventory: make([]Item, 0),
		Users:     make([]User, 0),
	}
}

func main() {
	library := NewLibrary()

	book1 := NewBook("The chronicles of narnia")
	dvd1 := NewDVD("Encyclopedia")

	library.AddInInventory(book1)
	library.AddInInventory(dvd1)

	student1 := NewStudent("Swastik")
	faculty1 := NewFaculty("Teacher")

	library.AddUser(student1)
	library.AddUser(faculty1)
}
