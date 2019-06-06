package client

import(
    "log"
    "net"
)

func sendProc(conn net.Conn, sendChan chan []byte, ctrlChan chan string){

    defer log.Println("close send goruntine")
    for{
        select{
            case data, ok := <- sendChan:
                if ok{
                    _, err := conn.Write(data)
                        if err != nil{
                            delClient(ctrlChan)
                            return
                        }
                    //log.Printf("send %X", data[:])
                }
            case _,ok := <- ctrlChan:
               if !ok{
                    return
               }
        }
    }
}

