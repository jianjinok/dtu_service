package client

import (
    "log"
    "fmt"
    "net"
    "sync"
    "time"
)

type dtu struct{
    conn net.Conn
    dtuid string
    reIp string
    inChan, outChan, ctrlChan chan string
    lastTime time.Time
    blockNum int
}

type deviceCtrl struct{
    dtumap map[string]dtu
    mutex sync.RWMutex
}

var deviceCtrls deviceCtrl

func addDtu(conn net.Conn, dtuid string, reIp string, inChan, outChan, ctrlChan chan string){

    log.Printf("add device: %s %v\n",dtuid, conn)
    deviceCtrls.mutex.Lock()
    var newDtu = dtu{conn:conn, dtuid:dtuid, reIp:reIp, inChan:inChan, outChan:outChan, ctrlChan:ctrlChan, lastTime:time.Now()}
    deviceCtrls.dtumap[dtuid] = newDtu 
    deviceNum := len(deviceCtrls.dtumap)
    deviceCtrls.mutex.Unlock()
    log.Printf("add device num: %d\n",deviceNum)
}

func delDtu(dtuid string, conn net.Conn){

    log.Printf("del device: %s\n",dtuid)
    deviceCtrls.mutex.Lock()
    dtu,ok := deviceCtrls.dtumap[dtuid] 
    if ok{
        if dtu.conn == conn{
            delete(deviceCtrls.dtumap, dtuid)
        }
    }
    deviceNum := len(deviceCtrls.dtumap)
    deviceCtrls.mutex.Unlock()
    log.Printf("del device num: %d\n",deviceNum)

}

func getDtu(dtuid string)(dtu, bool){

    deviceCtrls.mutex.RLock()
    dtu, ok := deviceCtrls.dtumap[dtuid]
    deviceCtrls.mutex.RUnlock()

    return dtu, ok
}

func getDtuService()deviceCtrl{

    deviceCtrls.mutex.RLock()
    service := deviceCtrls
    deviceCtrls.mutex.RUnlock()
    
    return service
}

func getDtuId(conn net.Conn)(string, bool){

    deviceCtrls.mutex.RLock()
    for dtuid, dtu := range deviceCtrls.dtumap{
        if dtu.conn == conn{
            return dtuid, true
        }
    }
    deviceCtrls.mutex.RUnlock()

    return "", false
}

func getDtuList()[]string{

    var list []string

    deviceCtrls.mutex.RLock()
    for dtuid, _ := range deviceCtrls.dtumap{
        list = append(list,dtuid)
    }
    deviceCtrls.mutex.RUnlock()

    return list
}

func changeBlockNum(dtuid string, num int){

    deviceCtrls.mutex.Lock()
    dtu, ok := deviceCtrls.dtumap[dtuid]
    if ok{
        dtu.blockNum +=num
        deviceCtrls.dtumap[dtuid] = dtu
    }
    deviceCtrls.mutex.Unlock()
}

func sendDtuMsg(dtuid, msg string)(string, bool){


    dtu,ok := getDtu(dtuid)
    if !ok{
        log.Printf("no find dtuid %s\n", dtuid)
        return "no find dtu",false
    }

    var hex []byte
    if _, err := fmt.Sscanf(msg,"%X",&hex); err!=nil{
        log.Printf("DownloadDeviceMsg dtu %s msg %s msg format error\n", dtuid, msg)
        return "msg format err",false
    }

    changeBlockNum(dtuid, 1)
    defer changeBlockNum(dtuid, -1)

    return sendClientMsg(dtu.inChan, dtu.outChan, msg), true
}

