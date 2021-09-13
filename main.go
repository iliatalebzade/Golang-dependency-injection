package main

import "log"

// DATABASE PART
type Database interface {
	GetUser() string
	GetAllUsers() []string
}

type DefaultDatabase struct{}

func (db DefaultDatabase) GetUser() string {
	return "bob"
}

// FAMOUS DATABASE
type FamousDatabase struct{}

func (db FamousDatabase) GetUser() string {
	return "Will Smith"
}

func (db FamousDatabase) GetAllUsers() []string {
	return []string{}
}

// GREETER PART
type Greeter interface {
	Greet(userName string)
}

type NiceGreeter struct{}

func (ng NiceGreeter) Greet(userName string) {
	log.Printf("Hello %s !! Nice to meet you", userName)
}

// PROGRAMM PART
type Program struct {
	Db      Database
	Greeter Greeter
}

func (p Program) Execute() {
	user := p.Db.GetUser()
	p.Greeter.Greet(user)
}

// MAIN PART
func main() {
	db := FamousDatabase{}
	greeter := NiceGreeter{}
	p := Program{
		Db:      db,
		Greeter: greeter,
	}

	p.Execute()
}
