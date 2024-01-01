package main

import (
	f "fmt"
)

func (e *Engineer) GetInfo() (string, float64) {
	return e.Name, e.Salary
}

func (m *Manager) GetInfo() (string, float64) {
	return m.Name, m.Salary
}

func (t *Technician) GetInfo() (string, float64) {
	return t.Name, t.Salary
}

func main() {
	eng := &Engineer{"Danny", 115000}
	man := &Manager{"Jamie", 150000}
	tech := &Technician{"Ron", 50000}
	f.Println(eng.GetInfo())
	f.Println(man.GetInfo())
	f.Println(tech.GetInfo())
}
