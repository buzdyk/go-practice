package main

import (
	"html/template"
	"net/http"
	"strings"
)

func main() {
	story, err := ParseScript()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		arc := strings.Trim(r.RequestURI, "/")
		script, exists := story[arc]

		if exists == false {
			script = story["intro"]
		}

		t, _ := template.New("arc").Parse(`
			<h1>{{.Title}}</h1>
			<p>{{.Story}}</p>
			{{range $option := .Options}}
				<a href="/{{$option.Arc}}">{{$option.Text}}</p>
			{{end}}
		`)

		t.Execute(w, script)
	})

	http.ListenAndServe(":8080", nil)
}
