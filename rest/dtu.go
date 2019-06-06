package rest

import (
    "log"
    "fmt"
    "encoding/json"
    "io/ioutil"
    "dtu_service/tcp/client"
    "github.com/ant0ine/go-json-rest/rest"
)

type requestSt struct{
    Id string   `json:"id"`
    Msg string  `json:"msg"`
}

type responseSt struct{
    Id string   `json:"id"`
    Msg string  `json:"msg"`
    Status string   `json:"status"`
    Info string     `json:"info"`
}

func dtu_execrawcmd(w rest.ResponseWriter, req *rest.Request){

    response := responseSt{Status:"ok"}
    var request requestSt

    jsonbytes, _ := ioutil.ReadAll(req.Body)
    json.Unmarshal(jsonbytes, &request)
    log.Println(request)

    response.Id = request.Id
    dtuid := fmt.Sprintf("%X",request.Id)
    recvmsg, ok := client.DownloadDeviceMsg(dtuid, request.Msg)
    if !ok{
        response.Status = "fail"
        response.Info = recvmsg
        recvmsg = ""
    }
    response.Msg = recvmsg

    w.WriteJson(response)
}

func dtu_execclose(w rest.ResponseWriter, req *rest.Request){

    response := responseSt{Status:"ok"}
    var request requestSt

    jsonbytes, _ := ioutil.ReadAll(req.Body)
    json.Unmarshal(jsonbytes, &request)

    response.Id = request.Id
    dtuid := fmt.Sprintf("%X",request.Id)
    info, ok := client.DelDevice(dtuid)
    if !ok{
        response.Status = "fail"
    }
    response.Info = info

    w.WriteJson(response)
}

func dtu_execcmd(w rest.ResponseWriter, req *rest.Request){
    resp := make(map[string]string)
    resp["exec func"] = "dtu_execcmd"
    w.WriteJson(resp)
}

func dtu_getstatus(w rest.ResponseWriter, req *rest.Request){
    resp := make(map[string]string)
    resp["exec func"] = "dtu_getstatus"
    w.WriteJson(resp)
}

