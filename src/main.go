package main

import (
	"flag"
	"fmt"
	"github.com/huhuhudia/URLShortener/src/urlshort"
	"log"
	"net/http"
)

func main(){
	urlmap := map[string]string{
		"/todayisgood":"http://localhost:8080/hello",
	}
	yamlFileName := flag.String("yaml-file", "redirect.yaml", "Yaml file name with redirection URL")
	flag.Parse()
	mux := defaultMux()
	mapHandler := urlshort.NewHttpRedirectHandler(urlshort.NewBaseUrlMapper(urlmap) ,mux)
	yamlUrlMapper, err := urlshort.NewYamlUrlMapper(*yamlFileName)
	if err != nil{
		log.Fatalf("cant create yaml redirect url provider : %v", err)
	}
	yamlHandler := urlshort.NewHttpRedirectHandler(yamlUrlMapper, mapHandler)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello world")
}