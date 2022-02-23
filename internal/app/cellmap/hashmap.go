package cellmap

import (
	"evo/internal/app/simulation"
)

type CellMap struct {
	m map[**simulation.Cell]*simulation.Cell
}

// New instantiates a hash map.
func New() *CellMap {
	return &CellMap{m: make(map[**simulation.Cell]*simulation.Cell)}
}

// Put inserts element into the map.
func (cm *CellMap) Put(value *simulation.Cell) {
	cm.m[&value] = value
}

// func (cm *CellMap) Values() []*simulation.Cell {
// 	values := make([]*simulation.Cell, cm.Size())
// 	for _, value := range cm.m {
// 		values = append(values, value)
// 	}
// 	return values
// }

// // Values returns all values (random order).
func (m *CellMap) Values() []*simulation.Cell {
	values := make([]*simulation.Cell, m.Size())
	count := 0
	for _, value := range m.m {
		values[count] = value
		count++
	}
	return values
}

func (m *CellMap) Size() int {
	return len(m.m)
}

func (cm *CellMap) Get(key **simulation.Cell) (value *simulation.Cell, found bool) {
	value, found = cm.m[key]
	return
}

func (cm *CellMap) GetM() map[**simulation.Cell]*simulation.Cell {
	return cm.m
}

func (cm *CellMap) Remove(key **simulation.Cell) {
	delete(cm.m, key)
	// value, _ := cm.Get(key)

	// fmt.Println("DELETE", value)
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

// // Get searches the element in the map by key and returns its value or nil if key is not found in map.
// // Second return parameter is true if key was found, otherwise false.
// func (m *Map) Get(key interface{}) (value interface{}, found bool) {
// 	value, found = m.m[key]
// 	return
// }

// // Remove removes the element from the map by key.

// // Empty returns true if map does not contain any elements
// func (m *Map) Empty() bool {
// 	return m.Size() == 0
// }

// // Size returns number of elements in the map.

// // Keys returns all keys (random order).
//

// // Values returns all values (random order).
// func (m *Map) Values() []interface{} {
// 	values := make([]interface{}, m.Size())
// 	count := 0
// 	for _, value := range m.m {
// 		values[count] = value
// 		count++
// 	}
// 	return values
// }

// // Clear removes all elements from the map.
// func (m *Map) Clear() {
// 	m.m = make(map[interface{}]interface{})
// }

// // String returns a string representation of container
// func (m *Map) String() string {
// 	str := "HashMap\n"
// 	str += fmt.Sprintf("%v", m.m)
// 	return str
// }
