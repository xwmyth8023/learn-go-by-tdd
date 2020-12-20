package MapsTest

import "testing"

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, key, value string) {
	t.Helper()
	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if value != got {
		t.Errorf("got '%s' want '%s'", got, value)
	}
}
func TestSearch(t *testing.T) {
	// dictionary := map[string]string{"test": "this is just a test"}
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("exist word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	// got := Search(dictionary, "test")

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, errorNotFound)
	})
}

// func TestAdd(t *testing.T)  {
// 	dictionary := Dictionary{}
// 	dictionary.Add("test", "this is just a test")
// 	want := "this is just a test"
// 	got, err := dictionary.Search("test")

// 	if err != nil {
// 		t.Fatal("should find added word:", err)
// 	}

// 	if want != got {
// 		t.Errorf("got '%s' want '%s'", got, want)
// 	}
// }
func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		value := "this is just a test"

		err := dictionary.Add(word, value)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, value)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, errorWordExits)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update exist key value", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		newDefinition := "new definition"
		dictionary := Dictionary{key: value}
		dictionary.Update(key, newDefinition)

		assertDefinition(t, dictionary, key, newDefinition)
	})

	t.Run("update new key value", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "new"
		value := "new test"
		err := dictionary.Update(key, value)

		assertError(t, err, ErrWordDoesNotExist)
		// assertDefinition(t, dictionary, key, value)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete exist key value", func(t *testing.T) {
		key := "test"
		dictionary := Dictionary{key: "test definition"}
		dictionary.Delete(key)

		_, err := dictionary.Search(key)
		if err != errorNotFound {
			t.Errorf("Expected '%s' to be deleted", key)
		}
	})

	t.Run("delete not exist key value", func(t *testing.T) {
		dictionary := Dictionary{"test": "test definition"}
		key := "notExist"

		err := dictionary.Delete(key)

		assertError(t, err, ErrWordDoesNotExist)
	})
}
