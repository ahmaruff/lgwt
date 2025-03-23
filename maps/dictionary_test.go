package maps

import "testing"

func TestDictionary(t *testing.T) {
    dictionary := Dictionary{"test":"this is just a test"}


    t.Run("known word", func(t *testing.T) {
        got,_ := dictionary.Search("test")

        want := "this is just a test"

        assertStrings(t, got, want)
    })

    t.Run("unknown word", func(t *testing.T) {
        _,err := dictionary.Search("bejo")
        want := ErrKeyNotFound 

        if err == nil {
            t.Fatal("expected to get an error.")
        }

        assertError(t, err, want)

    })
} 


func assertStrings(t *testing.T, got string, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

func assertError(t *testing.T, got error, want error) {
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
