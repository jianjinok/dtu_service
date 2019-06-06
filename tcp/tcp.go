package tcp

import(
    "log"
   "dtu_service/tcp/connect"
   "dtu_service/tcp/client"
   "dtu_service/tcp/server"
)

func RUN(addr string){

    client.RUN()
    connect.RUN()
    server.RUN(addr)
}

func init(){
    log.Println("tcp running...")
}

