package main

type Engineer struct {
	Name   string
	Salary float64
}

type Manager struct {
	Name   string
	Salary float64
}

type Technician struct {
	Name   string
	Salary float64
}

type Employee interface {
	GetInfo() (string, float64)
}
