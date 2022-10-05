package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown_word")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Add("test", "this is just a test")

		want := "this is just a test"

		assertError(t, err, nil)
		assertDefinition(t, dict, "test", want)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		dict := Dictionary{word: defination}

		err := dict.Add(word, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, defination)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}
		newDifination := "this is the new definition"

		err := dict.Update(word, newDifination)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDifination)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{}
		newDifination := "this is the new definition"

		err := dict.Update(word, newDifination)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is just a definition"
	dict := Dictionary{word: definition}

	dict.Delete(word)
	_, err := dict.Search(word)
	if err != ErrorNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q given", got, want)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}
