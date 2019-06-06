package connect

import (
    "log"
    "fmt"
    "net"
    "time"
    "dtu_service/callback"
    "dtu_service/config"
    "dtu_service/tcp/client"
)

func connCheckDevice(dataHex string) bool{
    return callback.DtuIsAvalid(dataHex)
}

func connAddDevice(conn net.Conn, idHex string){
    var expire time.Time
    conn.SetReadDeadline(expire)
    client.AddDevice(conn, idHex)
}

func connProc(conn net.Conn){

    data := make([]byte, config.TcpBufSize)
    conn.SetReadDeadline(time.Now().Add(time.Duration(config.TcpConnectTimeout)*time.Second))
    len,err := conn.Read(data)
    if err != nil{
        log.Println(err)
        conn.Close()
        return
    }
    strData := fmt.Sprintf("%X",data[:len])
    log.Printf("recv %s\n",strData)
    if(connCheckDevice(strData)){
        connAddDevice(conn, strData)
        return
    }
    conn.Close()
}

func AddConnect(conn net.Conn){
    readdr := conn.RemoteAddr()
    log.Printf("add tcp conn %v addr %v", conn, readdr)
    connProc(conn)
}

func RUN(){
}

func init(){

    log.Println("tcp connect running...")
}
