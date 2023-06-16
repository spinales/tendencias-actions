package main

import (
	"testing"
)

func TestCreatingDatabase(t *testing.T) {
	//db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	t.Log("OK")
}

func TestMigrations(t *testing.T) {
	t.Log("OK")
}

func TestInsertPerson(t *testing.T) {
	t.Log("OK")
}

func TestRetrievePerson(t *testing.T) {
	t.Log("OK")
}

var casesReverse = []struct {
	Value    string
	Expected string
}{
	{"juan", "nauj"},
	{"oso", "oso"},
	{"World", "dlroW"},
	{"Hello", "olleH"},
	{"Geeks", "skeeG"},
	{"Lee", "eeL"},
}

func TestReverse(t *testing.T) {
	for _, r := range casesReverse {
		result := reverse(r.Value)
		if result != r.Expected {
			t.Errorf("Expected: %s Got:%s", r.Expected, result)
		}
	}
}
