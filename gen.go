package main

// +build ignore

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
)

var VersionString = "unset"

func main() {

	const url = "https://github.com/golang/go/raw/master/CONTRIBUTORS"

	resp, err := http.Get(url)
	die(err)
	defer resp.Body.Close()

	sc := bufio.NewScanner(resp.Body)
	carls := []string{}

	for sc.Scan() {
		if strings.Contains(sc.Text(), "Carl") {
			carls = append(carls, sc.Text())
		}
	}
	die(sc.Err())

	f, err := os.Create("contributors.go")
	die(err)
	defer f.Close()

	packageTemplate.Execute(f, struct {
		Timestamp time.Time
		URL       string
		Carls     []string
	}{
		Timestamp: time.Now(),
		URL:       url,
		Carls:     carls,
	})

	fmt.Println("Version: ", VersionString)
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
	// this file was generated by robots at
	// {{ .Timestamp }}
	// {{ .URL }}
	package contributor

	var Contributors = []string{
		{{- range .Carls }}
			{{ printf "%q" . }},
		{{- end }}
	}
	`))
