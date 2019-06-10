package client

import(
    "log"
    "net"
    "time"
    "dtu_service/config"
)

func recvProc(conn net.Conn, recvChan chan []byte, ctrlChan chan string){

    defer log.Println("close recv goruntine")
    data := make([]byte, config.TcpBufSize)
    for{
        conn.SetReadDeadline(time.Now().Add(time.Duration(config.TcpAliveTimeout)*time.Second))
        size,err := conn.Read(data)
        if err != nil{
            delClient(ctrlChan)
            return
        }
        //log.Printf("recv %X", data[:size])
        recvChan <- data[:size]
    }
}

