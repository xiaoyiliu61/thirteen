package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"unicode/utf8"
)

func main() {
	http.HandleFunc("/user_info1", UserInfo)

	err := http.ListenAndServe("127.0.0.1:8090", nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
func UserInfo(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	username:=request.Form.Get("username")
	length:=utf8.RuneCountInString(username)
	if (length>8) {
		fmt.Println("名字不能超过8位")
	    writer.Write([]byte("名字不能超过8位"))
		return
	}
	//'^[x{4e00}-\x{9fa5}]+$'

	result,err:=regexp.MatchString(`^[x{4e00}-\x{9fa5}]+$`,username)
	if result{
		fmt.Println(username)
		writer.Write([]byte("欢迎你："+username))
	}else {
		fmt.Println(username)
		writer.Write([]byte("对不起,名字必须是中文"+username))
		return
	}


	 ageStr:=request.Form.Get("age")
	 ageResult,err:=regexp.MatchString(`[0-9]$`,ageStr)
	if ageResult {
		writer.Write([]byte("对不起，你输入的年龄格式不正确"))
		return
	}
	age,err:=strconv.Atoi(ageStr)
	if err != nil {
		log.Fatal(err)
		writer.Write([]byte("对不起，你输入的格式不正确"))
		return
	}

	 fmt.Println("年龄",age)

	 /*ageResult,err:=regexp.MatchString(`^[0-9]+$`,ageStr)
     if err!=nil{
     	writer.Write([]byte("对不起，你输入的格式不正确"))
	 }
	if ageResult {
		fmt.Println("年龄",ageResult)
	}*/
	 mobileStr:=request.Form.Get("mobile")

	 mobileResult,err:=regexp.MatchString(`^(1[3|5|7|8][0-9]\d{8}$)`,mobileStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	if !mobileResult {
		writer.Write([]byte("手机格式不正确，请重新尝试"))
		return
	}
	xueli:=request.Form.Get("xueli")
	xueliSlice:=[]string{"xiaoxue","chuzhong","gaozhong","dazhuan"}
	var xueliResult bool
	for _,value:=range xueliSlice {
	 	if value==xueli {
		  xueliResult=true
		  continue
		}else {
			xueliResult=false
			break
		}
	}
	if !xueliResult {
		fmt.Println("学历不合法")
	    writer.Write([]byte("学历不合法，请重新输入"))
	}
	 userCard:=request.Form.Get("usercard")
	 cardResult,err:=regexp.MatchString(`^(\d{17})([0-9]x)$`,userCard)
	if err!=nil {
		fmt.Println(err)
		return
	}
	if !cardResult {
		writer.Write([]byte("对不起，身份证不符合规则"))
		return
	}
}
