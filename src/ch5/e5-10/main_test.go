package main

import (
	"testing"
)

func TestName(t *testing.T) {
	courses := topoSort(prereqs)

	index := make(map[string]int, len(courses))
	for i, name := range courses {
		if index[name] != 0 {
			t.Errorf("Duplicate of class: %s\n", name)
		}
		index[name] = i + 1
	}

	for class, class_prereqs := range prereqs {
		if index[class] == 0 {
			t.Errorf("Missing class: %s\n", class)
		}
		for _, preq := range class_prereqs {
			if index[class] < index[preq] {
				t.Errorf("The class (%s) cannot be taken without taking the class (%s)\n", class, preq)
			}
			if index[preq] == 0 {
				t.Errorf("Missing class: %s\n", preq)
			}
		}
	}
}
