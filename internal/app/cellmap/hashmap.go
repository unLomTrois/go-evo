package cellmap

import (
	"evo/internal/app/simulation"
)

type CellMap struct {
	m map[**simulation.Cell]*simulation.Cell
}

func New() *CellMap {
	return &CellMap{m: make(map[**simulation.Cell]*simulation.Cell)}
}

func (m *CellMap) Put(value *simulation.Cell) {
	m.m[&value] = value
}

func (m *CellMap) Values() []*simulation.Cell {
	values := make([]*simulation.Cell, m.Size())
	count := 0
	for _, value := range m.m {
		values[count] = value
		count++
	}
	return values
}

func (m *CellMap) Keys() []**simulation.Cell {
	keys := make([]**simulation.Cell, m.Size())
	count := 0
	for key := range m.m {
		keys[count] = key
		count++
	}
	return keys
}

func (m *CellMap) Size() int {
	return len(m.m)
}

func (m *CellMap) Get(key **simulation.Cell) (value *simulation.Cell, found bool) {
	value, found = m.m[key]
	return
}

func (m *CellMap) GetM() map[**simulation.Cell]*simulation.Cell {
	return m.m
}

func (m *CellMap) Remove(key **simulation.Cell) {
	if !m.IsEmpty() {
		delete(m.m, key)
	}
}

func (m *CellMap) IsEmpty() bool {
	return m.Size() == 0
}

func (m *CellMap) Clear() {
	m.m = make(map[**simulation.Cell]*simulation.Cell)
}
