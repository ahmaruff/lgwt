package main

import "testing"

func TestHello(t *testing.T){
    t.Run("Saying Hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"

        if got != want {
            t.Errorf("get %q want %q", got, want)
        }
    })

    t.Run("Saying 'Hello, World' when empty string is provided", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"

        if got != want {
            t.Errorf("get %q want %q", got, want)
        }
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()
    if(got != want) {    
        t.Errorf("get %q want %q", got, want)
    }
}
