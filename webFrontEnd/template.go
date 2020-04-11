package main

import (
	"log"
	"net/http"
	"sync"
	"text/template"
)

type templateHandler struct {
	one      sync.Once
	fileName string
	tem      *template.Template
}

func (t *templateHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	t.one.Do(func() {
		t.tem = template.Must(template.ParseFiles(t.fileName))
	})
	if err := t.tem.Execute(rw, nil); err != nil {
		log.Println(err)
	}
}
