package service

import (
    "net/http"
    "os"
    "fmt"
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

    formatter := render.New(render.Options{
        Directory:  "templates",
        Extensions: []string{".html"},
        IndentJSON: true,
    })

    n := negroni.Classic()
    mx := mux.NewRouter()
    myMiddleWare := NewMiddleware()
    initRoutes(mx, formatter) 
    n.Use(myMiddleWare)
    n.UseHandler(mx)

    return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
    webRoot := os.Getenv("WEBROOT")
    if len(webRoot) == 0 {
        if root, err := os.Getwd(); err != nil {
            panic("Could not retrive working directory")
        } else {
            webRoot = root
            //fmt.Println(root)
        }
    }
    fmt.Println(webRoot)
    mx.HandleFunc("/login", loginHandler(formatter))
    mx.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")
    mx.PathPrefix("/api/").Handler(http.HandlerFunc(myError))
    mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))
}