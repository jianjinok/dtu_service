package config

import(
    "log"
    "flag"
)

var(
    TcpAddr string
    RestAddr string
    TcpServerTimeout int64
    TcpConnectTimeout int64
    TcpCmdTimeout int64
    TcpBufSize int
)

func RUN(){
    flag.Parse()

    log.Printf("tcp server addr: %s", TcpAddr)
    log.Printf("rest sever addr: %s", RestAddr)
    log.Printf("tcp restart timout %ds", TcpServerTimeout)
    log.Printf("tcp connect timout %ds", TcpConnectTimeout)
    log.Printf("tcp send cmd timout %ds", TcpCmdTimeout)
    log.Printf("tcp buff size %d byte", TcpBufSize)
}

func init(){

    flag.StringVar(&TcpAddr, "T",":3334", "using for tcp server default(:3334)")
    flag.StringVar(&RestAddr, "R",":3333", "using for rest server default(:3333)")
    flag.Int64Var(&TcpServerTimeout, "S", 1, "tcp server restart timeout default(1)s")
    flag.Int64Var(&TcpConnectTimeout, "C", 30, "tcp connect timeout default(30)s")
    flag.IntVar(&TcpBufSize, "B", 1024, "tcp Buf size default(1024)s")
    flag.Int64Var(&TcpCmdTimeout, "M", 10, "tcp send cmd timeout default(10)s")
}

