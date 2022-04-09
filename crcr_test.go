package main

import (
	
	"testing"
	"reflect"
)

func TestCrcr(t *testing.T) {
	assertString := func(t testing.TB, got, excepted string) {
		t.Helper()
		if got != excepted {
			t.Errorf("got %q excepted %q", got , excepted)
		}
	}

	assertCharInArray  := func(t testing.TB, got, excepted map[string]int, isPresent bool ) {
		t.Helper()
		if !reflect.DeepEqual(got,excepted) {
			t.Errorf("got %v excepted %v isPresent %v", got, excepted, isPresent)
		}
	}

	t.Run("TestStart", func(t *testing.T) {
		got := Start()
		excepted := "Start"
		assertString(t, got, excepted)

	})

	t.Run("TestArgs", func (t *testing.T ) {
		got := HandleArgs()
		excepted := "testtext.txt"
		if got != excepted {
			t.Errorf("got %q excepted %q", got, excepted)
		}
	})

	t.Run("TestCharInArray", func (t *testing.T) {
		got, isPresent := CharInArray("z", charsLower)
		excepted := map[string]int{"z":1}
		assertCharInArray(t, got, excepted, isPresent)
	})
}
