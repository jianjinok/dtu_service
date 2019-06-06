package main

import(
    "log"
    "dtu_service/config"
    _ "dtu_service/callback"
    "dtu_service/tcp"
    "dtu_service/rest"
)

func main(){

    config.RUN()

    go tcp.RUN(config.TcpAddr)
    
    rest.RUN(config.RestAddr)
}

func init(){
    log.SetFlags(log.Llongfile | log.LstdFlags)
    log.Println("main running...")
}
