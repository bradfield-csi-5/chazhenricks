package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("find a match", func(t *testing.T) {

		dictionary := Dictionary{"test": "this is just a test"}
		got, gotErr := dictionary.Search("test")
		want := "this is just a test"
		assertNoErrors(t, gotErr)
		assertStringsEqual(t, got, want)

	})

	t.Run("return no match", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got, gotErr := dictionary.Search("fart")
		want := ""

		assertErrors(t, gotErr, NotFoundError)
		assertStringsEqual(t, got, want)

	})

	t.Run("add item to dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is another test")

		got, gotErr := dictionary.Search("test")
		want := "this is another test"

		assertNoErrors(t, gotErr)
		assertStringsEqual(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "word"
		definition := "this is a word"
		dictionary := Dictionary{word: definition}
		addErr := dictionary.Add(word, "another def for word")

		assertErrors(t, addErr, WordExistsError)

	})

	t.Run("non- existing word", func(t *testing.T) {
		word := "test"
		def := "this is a test"
		dictionary := Dictionary{}

		updateErr := dictionary.Update(word, def)
		assertErrors(t, updateErr, WordDoesNotExistError)

	})

	t.Run("delete existing word", func(t *testing.T) {
		word := "test"
		def := "def"
		dictionary := Dictionary{word: def}

		deleteErr := dictionary.Delete(word)
		assertNoErrors(t, deleteErr)

    _, gotErr := dictionary.Search(word)
    assertErrors(t, gotErr, NotFoundError)
	})

  t.Run("error on delete non-existing word", func(t *testing.T){
    
		word := "test"
    dict := Dictionary{}
    deleteErr := dict.Delete(word)

    assertErrors(t, deleteErr, DeleteNonExistError)
  })

}
func assertNoErrors(t testing.TB, gotErr error) {
	if gotErr != nil {
		t.Fatalf("dictionary.Search: didnt want an err and I got one: %v", gotErr)
	}
}

func assertErrors(t testing.TB, gotErr, wantErr error) {

	if gotErr == nil {
		t.Fatal("Was expecting an error and didnt get one")
	}

	if gotErr != wantErr {
		t.Errorf("expected: %q - got: %q", wantErr, gotErr)
	}
}

func assertStringsEqual(t testing.TB, got, want string) {

	if got != want {
		t.Errorf("got:'%s' - want:'%s'", got, want)
	}
}
