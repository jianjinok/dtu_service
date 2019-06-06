package client

import (
    "log"
    "net"
)

func AddDevice(conn net.Conn, dtuid string){

    _,ok := getDtu(dtuid)
    if ok{
        DelDevice(dtuid)
    }
    reIp := conn.RemoteAddr().String()
    inChan, outChan, ctrlChan := createClient(conn, dtuid)
    addDtu(conn, dtuid, reIp, inChan, outChan, ctrlChan)
}

func DelDevice(dtuid string)(string, bool){
    dtu,ok := getDtu(dtuid)
    if ok{
        delClient(dtu.ctrlChan)
        return "delete device ok", true
    }
    return "not find dtuid", false
}

func DownloadDeviceMsg(dtuid, msg string)(string, bool){
    return sendDtuMsg(dtuid, msg)
}

func GetDeviceStatus(dtuid string)(dtu, bool){
    return getDtu(dtuid)
}

func GetServiceStatus() deviceCtrl{
    return getDtuService()
}

func RUN(){
}

func init(){
    deviceCtrls.dtumap = make(map[string]dtu)
    log.Printf("tcp client running...\n")
}

