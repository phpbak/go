package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

func main() {
	rlt, err := doHttpGetRequest("https://restapi.amap.com/v3/weather/weatherInfo?city=420100&key=78278e172402b74f438f48b13c60f21c")
	  if err != nil {
	    fmt.Println("net req error")
	  } else {
	    fmt.Println(rlt)
	  }
}

// 自定义http get请求函数
// 入参：请求url
// 返回值：rlt，天气数据。err，错误信息
// 网络请求
func doHttpGetRequest(url string) (rlt string, err error) {
 
  // http.Get在net/http中，所以要import "net/http"
  resp, err := http.Get(url)
 
  if err != nil {
    return "", err
  } else {
    // 使用efer resp.Body.Close()。当doHttpGetRequest成功return之后，执行此行语句。多用于句柄关闭
    defer resp.Body.Close()
 
    // io流数据读取。需要引用io/ioutil
    body, err := ioutil.ReadAll(resp.Body)
 
    if err != nil {
      return "", err
    } else {
      return string(body), err
    }
 
  }
 
}

