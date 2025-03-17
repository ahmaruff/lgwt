package helloworld

import "testing"

func TestHello(t *testing.T){
    t.Run("Saying Hello to people", func(t *testing.T) {
        got := Hello("Chris", "")
        want := "Hello, Chris"

        assertCorrectMessage(t, got, want);
    })

    t.Run("Saying 'Hello, World' when empty string is provided", func(t *testing.T) {
        got := Hello("", "")
        want := "Hello, World"

        assertCorrectMessage(t, got, want);
    })

    t.Run("In Spanish", func(t *testing.T) {
        got := Hello("Esmeralda", "es")
        want := "Hola, Esmeralda"

        assertCorrectMessage(t, got, want);
    })

    t.Run("In French", func(t *testing.T) {
        got := Hello("Malik", "fr")
        want := "Bonjour, Malik"

        assertCorrectMessage(t, got, want);
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()
    if(got != want) {    
        t.Errorf("get %q want %q", got, want)
    }
}
