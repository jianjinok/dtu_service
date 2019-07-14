package rest

import (
    "log"
    "runtime/pprof"
    "github.com/ant0ine/go-json-rest/rest"
    "dtu_service/tcp/client"
)

func service_getruntine(w rest.ResponseWriter, req *rest.Request) {

    p := pprof.Lookup("goroutine")
    log.Println(p)
    w.WriteJson(p)
}

func service_getdtus(w rest.ResponseWriter, req *rest.Request){
    resp := make(map[string]string)
    resp["exec func"] = "service_getdtus"
    w.WriteJson(resp)
}

func service_getonlines(w rest.ResponseWriter, req *rest.Request){
    resp := make(map[string][]string)
    dtulist := client.GetDtuList()
    resp["dtus"] = dtulist

    w.WriteJson(resp)
}

func service_getstatus(w rest.ResponseWriter, req *rest.Request){
    resp := make(map[string]string)
    resp["exec func"] = "service_getstatus"
    w.WriteJson(resp)
}

