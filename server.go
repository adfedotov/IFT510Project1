package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	// Password Generator
	http.HandleFunc("/generate_password/", GeneratePasswordHandler)

	// Static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Server started at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GeneratePasswordHandler(w http.ResponseWriter, r *http.Request) {
	length := 8 // default
	if r.URL.Query().Has("length") {
		length, _ = strconv.Atoi(r.URL.Query().Get("length"))
	}

	if length < 5 || length > 1024 {
		fmt.Fprintf(w, "length parameter must be between 5 and 1024 inclusive")
		return
	}

	useUpper := true
	if r.URL.Query().Has("upper") {
		if r.URL.Query().Get("upper") == "false" {
			useUpper = false
		}
	}

	useSpecial := true
	if r.URL.Query().Has("special") {
		if r.URL.Query().Get("special") == "false" {
			useSpecial = false
		}
	}

	useDigits := true
	if r.URL.Query().Has("digits") {
		if r.URL.Query().Get("digits") == "false" {
			useDigits = false
		}
	}

	collection := "abcdefghijklmnopqrstuvwxyz"

	digit := "0123456789"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special := "~!@#$^&*()_+[]?"

	arr := make([]byte, length)

	if useUpper {
		collection = collection + uppercase
	}
	if useSpecial {
		collection = collection + special
	}
	if useDigits {
		collection = collection + digit
	}

	for i := 0; i < length; i++ {
		rnd := rand.Intn(len(collection))
		arr[i] = collection[rnd]
	}

	password := string(arr)

	fmt.Fprintf(w, password)
}
