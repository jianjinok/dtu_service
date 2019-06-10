package client

import (
    "log"
    "net"
    "fmt"
    "time"
    "dtu_service/config"
    "dtu_service/callback"
)

func closeChan(channel chan string){

    select{
        case <-channel:
            return
        default:
    }
    close(channel)
}

func clientTimeout(timeChan chan string, timeout int64){

    time.Sleep(time.Duration(timeout) * time.Second)
    close(timeChan)
}

func clientSendRecv(dtuid string, sendChan, recvChan chan []byte, msg string)string{

    timeout := make(chan string)
    var data []byte
    var recvmsg string

    _,err := fmt.Sscanf(msg, "%X", &data)
    if err != nil{
        return recvmsg
    
    }

    sendChan <- []byte(data[:])
    go clientTimeout(timeout, config.TcpCmdTimeout)
    for{
        select{
            case data = <- recvChan:
                if ok := callback.UploadDtuMsgHook(dtuid, data); !ok{
                    recvmsg = fmt.Sprintf("%X",data)
                    goto END
                }
            case <-timeout:
                recvmsg = ""
                goto END
        }
    }
END:
    return recvmsg
}

func clientCloseProc(dtuid string, conn net.Conn, inChan, outChan chan string){

    delDtu(dtuid, conn)
    close(outChan)
    conn.Close()
    for{
        select{
            case <- inChan:
            default:
                return
        }
    }
}

func clientgo(conn net.Conn, dtuid string, inChan, outChan chan string, ctrlChan chan string){

    sendChan := make(chan []byte)
    recvChan := make(chan []byte)
    go recvProc(conn, recvChan, ctrlChan)
    go sendProc(conn, sendChan, ctrlChan)
    
    defer log.Printf("close client %s\n", dtuid)
    defer clientCloseProc(dtuid, conn, inChan, outChan)
    
    for{
        select{
        case msg := <- inChan:
            recvMsg := clientSendRecv(dtuid, sendChan, recvChan, msg)
            outChan <- recvMsg
            log.Printf("send dtuid %s msg %s recv msg %s\n", dtuid, msg, recvMsg)
        
        case data := <- recvChan:
            callback.UploadDtuMsg(dtuid, data)

        case _,ok := <- ctrlChan:
            if !ok{
                return
            }
        }
    }
}

func createClient(conn net.Conn, dtuid string)(chan string, chan string, chan string){

    inChan := make(chan string, 1)
    outChan := make(chan string)
    ctrlChan := make(chan string)

    go clientgo(conn, dtuid, inChan, outChan, ctrlChan)
    
    return inChan, outChan, ctrlChan
}

func delClient(ctrlChan chan string){
   closeChan(ctrlChan)
}

func sendClientMsg(inChan, outChan chan string, msg string) string{

    inChan <- msg
    recv := <- outChan

    return recv
}

