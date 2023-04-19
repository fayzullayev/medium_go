package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type TodoList struct {
	TodoCount int
	Todos     []string
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func write(w http.ResponseWriter, msg string) {
	_, err := w.Write([]byte(msg))

	if err != nil {
		errorCheck(err)
	}
}

func english(w http.ResponseWriter, r *http.Request) {
	write(w, "Hello Internet")
}

func spanish(w http.ResponseWriter, r *http.Request) {
	write(w, "Hola Internet")
}

func interactHandler(w http.ResponseWriter, r *http.Request) {
	todos := getStrings("data.txt")

	tmpl, err := template.ParseFiles("view.html")
	errorCheck(err)

	todoss := TodoList{
		TodoCount: len(todos),
		Todos:     todos,
	}
	err = tmpl.Execute(w, todoss)

}

func newHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("new.html")

	errorCheck(err)

	err = tmpl.Execute(w, nil)
	errorCheck(err)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	todo := r.FormValue("todo")

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("data.txt", options, os.FileMode(0600))
	errorCheck(err)
	_, err = fmt.Fprintln(file, todo)
	errorCheck(err)
	err = file.Close()
	errorCheck(err)

	http.Redirect(w, r, "/interact", http.StatusFound)
}

func getStrings(fileName string) []string {
	var lines []string

	file, err := os.Open(fileName)

	if os.IsNotExist(err) {
		return nil
	}

	errorCheck(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	errorCheck(scanner.Err())

	return lines
}

func main() {

	http.HandleFunc("/hello", english)
	http.HandleFunc("/hola", spanish)
	http.HandleFunc("/interact", interactHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/create", createHandler)

	log.Fatal(http.ListenAndServe(":4003", nil))

}
