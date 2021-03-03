package main

import (
	"study6/manage/routers"
)

func main() {
	//upload.LoadTest()
	r := routers.InitRouter()
	r.Run()
}
