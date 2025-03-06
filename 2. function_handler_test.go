package main

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

// Home page Handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the Homepage")
}

// About Page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(3, 4)
	fmt.Fprintf(w, "this is the About Page and 3 + 4 is %d", sum)
}

// fucntion that add Values into About Handler
func AddValues(x, y int) int {
	sum := x + y
	return sum
}

// divide Handler
func Divides(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100, 3)
	if err != nil {
		fmt.Fprint(w, "Cannot Divide By 0")

		//stop execution when error
		return
	}

	angka := []float32{100, 3}
	fmt.Fprintf(w, "%f divided by %f is %f", angka[0], angka[1], f)
}

// function divide
func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by Zero")
		return 0, err
	}

	result := x / y
	return result, nil

}

func TestFunctionHandler(t *testing.T) {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divides)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
