package main

import "testing"

// to run test
// 1/ make sure there is a go.mod file, if not, in cli run: go mod init learn/hello-world
// 2/ in cli run: go test

// rules for writing tests:
// 1/ it needs to be in a file with name like xxx_test.go
// 2/ the test function must start with the word Test
// 3/ the test function takes one argument only t *testing.T
// 4/ in order to use the *testing.T type, you need to import "testing"

func TestHello(t *testing.T) {
	// sub tests with t.Run
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Alejandro", "Spanish")
		want := "Hola, Alejandro"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Amelie", "French")
		want := "Bonjour, Amelie"
		assertCorrectMessage(t, got, want)
	})
}

// about testing.TB
// for helper functions, it is a good idea to accept a testing.TB which is an interface that *testing.T and *testing.B both satisfy
// so you can call helper function from a test or a benchmark
func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper, for devs to track problems easier
	// by doing this, when it fails the line number reported will be in our function call rather than inside our test helper
	t.Helper()
	if got != want {
		// &q wraps your values in double quotes
		t.Errorf("got %q want %q", got, want)
	}
}
