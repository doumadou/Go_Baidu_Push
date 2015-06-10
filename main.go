package main

import (
	"Go_Baidu_Push/pushManager"
	"log"
)

func main() {
	log.Println("main")
	pushManager.SharedPushManager().PushToAll()
}
