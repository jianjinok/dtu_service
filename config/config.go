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
    TcpAliveTimeout int64
    TcpBufSize int
)

func RUN(){
    flag.Parse()

    log.Printf("tcp server addr: %s", TcpAddr)
    log.Printf("rest sever addr: %s", RestAddr)
    log.Printf("tcp restart timout %ds", TcpServerTimeout)
    log.Printf("tcp connect timout %ds", TcpConnectTimeout)
    log.Printf("tcp send cmd timout %ds", TcpCmdTimeout)
    log.Printf("tcp client alive timout %ds", TcpAliveTimeout)
    log.Printf("tcp buff size %d byte", TcpBufSize)
}

func init(){

    flag.StringVar(&TcpAddr, "T",":3334", "tcp server addr")
    flag.StringVar(&RestAddr, "R",":3333", "rest server addr")
    flag.Int64Var(&TcpServerTimeout, "S", 1, "tcp server restart timeout")
    flag.Int64Var(&TcpConnectTimeout, "C", 30, "tcp connect timeout")
    flag.IntVar(&TcpBufSize, "B", 1024, "tcp Buf size")
    flag.Int64Var(&TcpCmdTimeout, "M", 10, "tcp send cmd timeout")
    flag.Int64Var(&TcpAliveTimeout, "A", 66, "tcp keep alive timeout")
}

