package rest

import (
    "log"
    "net/http"
    "github.com/ant0ine/go-json-rest/rest"
)

var routes = [] *rest.Route{
    rest.Post("/dtu/execrawcmd", dtu_execrawcmd),
    rest.Post("/dtu/execclose", dtu_execclose),

    rest.Get("/dtu/execcmd", dtu_execcmd),
    rest.Get("/dtu/getstatus", dtu_getstatus),

    rest.Get("/service/getdtus", service_getdtus),
    rest.Get("/service/getonlines", service_getonlines),
    rest.Get("/service/getstatus", service_getstatus),
    rest.Get("/service/getruntine", service_getruntine),
}

func restserver(addr string){
    
    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    router, err := rest.MakeRouter(routes...)
    if err != nil{
        log.Fatal(err)
    }
    api.SetApp(router)
    http.Handle("/", http.StripPrefix("", api.MakeHandler()))
    log.Fatal(http.ListenAndServe(addr,nil))
}

func RUN(addr string){
    log.Printf("start rest server %s", addr)
    restserver(addr)
}

func init(){
    log.Printf("rest server running...\n")
}

