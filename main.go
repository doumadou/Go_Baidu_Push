package main

import (
	"Go_Baidu_Push/config"
	"log"
)

func main() {
	if config.APPID == "Your app id here" && config.SECRET_KEY == "Your secret key" && config.API_KEY == "Your api key" && config.USERAGENT == "BCCS_SDK/3.0 (GNU/Linux 3.13.0-32-generic x86_64) GO/1.4.2 (Baidu Push Server SDK V3.0.0 )" {
		log.Fatal("Fill with you own APPID, SECRET_KEY, API_KEY, USERAGENT in config.go, enroll your own app in the push console page if not yet have one!")
	}
}
