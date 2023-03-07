package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("nope")

		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("add new key", func(t *testing.T) {
		key := "test"
		val := "this is just a test"
		d := Dictionary{}

		err := d.Add(key, val)

		assertError(t, err, nil)
		assertDefinition(t, d, key, val)
	})

	t.Run("add existing key", func(t *testing.T) {
		key := "test"
		val := "this is just a test"
		d := Dictionary{key: val}

		err := d.Add(key, val)

		assertError(t, err, ErrKeyAlreadyExists)
		assertDefinition(t, d, key, val)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("updates existing word", func(t *testing.T) {
		key := "test"
		val := "this is just a test"

		d := Dictionary{key: val}
		newVal := "something new under the sun"
		err := d.Update(key, newVal)

		assertError(t, err, nil)
		assertDefinition(t, d, key, newVal)
	})

	t.Run("updates non-existent word", func(t *testing.T) {
		key := "test"

		d := Dictionary{}
		newVal := "something new under the sun"
		err := d.Update(key, newVal)

		assertError(t, err, ErrKeyDoesNotExist)
		assertDefinition(t, d, key, newVal)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	d := Dictionary{key: "val"}

	d.Delete(key)

	_, err := d.Search(key)
	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", err)
	}
}

func assertDefinition(t testing.TB, d Dictionary, key, val string) {
	t.Helper()
	d.Add(key, val)
	got, err := d.Search(key)
	want := val

	if err != nil {
		t.Fatal("should find added entry", err)
	}

	assertStrings(t, got, want)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q ", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}
