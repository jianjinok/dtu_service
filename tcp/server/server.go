package server

import (
    "log"
    "net"
    "time"
    "dtu_service/tcp/connect"
    "dtu_service/config"
)

func server(addr string){
    netListen, err := net.Listen("tcp", addr)
    if err != nil{
        log.Println(err)
        return
    }

    defer netListen.Close()
    for{
        conn,err := netListen.Accept()
        if err != nil{
            log.Println("accept error!")
            continue
        }
        go connect.AddConnect(conn)
    }
}

func RUN(addr string){

    log.Printf("start tcp server %s\n", addr)

    for{
        server(addr)
        time.Sleep(time.Duration(config.TcpServerTimeout) * time.Second)
    }
}

func init(){
    log.Printf("tcp server running...\n")
}

