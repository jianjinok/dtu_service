package callback

import(
    "fmt"
    "log"
)

func DtuIsAvalid(dtuid string) bool{

    log.Printf("dtu %s is avalid\n", dtuid)
    return true
}

func UploadDtuMsg(dtuid string, msg string)bool{

    log.Printf("upload dtu: %s msg: %s\n", dtuid, msg)
    return true

}


func UploadDtuMsgProc(dtuid string, data []byte) bool{

    if data[0] == 0xFF{
        log.Printf("%s keep alive msg\n", dtuid)
        return true
    }
    if data[5] == 0xC0{
        msg := fmt.Sprintf("%X",data)
        UploadDtuMsg(dtuid, msg)
        return true
    }

    return false
}

func UploadServerMsg(msg string) bool{

    log.Printf("upload service msg %s\n", msg)
    return true
}

func init(){
    log.Println("callback running...")
}


