package main

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if v, exists := pathsToUrls[r.RequestURI]; exists == true {
			http.Redirect(w, r, v, 302)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	urls, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}
	m := buildUrlMap(urls)
	return MapHandler(m, fallback), nil
}

func parseYaml(yml []byte) ([]url, error) {
	var urls []url
	if err := yaml.Unmarshal(yml, &urls); err != nil {
		return nil, err
	}
	return urls, nil
}

func buildUrlMap(urls []url) map[string]string {
	r := make(map[string]string)
	for _, v := range urls {
		r[v.Path] = v.Url
	}
	return r
}

type url struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}
