package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/upload",uploadFile)

	http.ListenAndServe("127.0.0.1:9013",nil)
}
func uploadFile(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("method",request.Method)

	request.ParseMultipartForm( 32<<10)
	file,header,err:=request.FormFile("uploadFile")
    if err!=nil{
    	log.Fatal(err)
		return
	}
	fileName:=header.Filename
	
	fmt.Println("文件名：",fileName)
    fileSize:=header.Size
    fmt.Println("文件大小：",fileSize)

    defer file.Close()

   // data:=make([]byte,100)
    //file.Read(data)
   // strings.ReplaceAll(header.Filename,"\\","/")
    fileNameSlice:=strings.Split(header.Filename,"\\")
    newFile,err:=os.OpenFile("./"+ fileNameSlice[len(fileNameSlice)-1],os.O_CREATE|os.O_RDONLY|os.O_WRONLY,0666)
	if err!=nil{
		log.Fatal(err)
		return
	}
	defer newFile.Close()

    io.Copy(newFile,file)
    writer.Write([]byte("恭喜文件上传成功"))
}
