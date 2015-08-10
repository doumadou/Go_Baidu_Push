package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/url"
	"sort"
	"strings"
)

func ToMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	str := hex.EncodeToString(h.Sum(nil))
	return str
}

func GenerateSignature(method string, urlStr string, secretkey string, parameters map[string]string) string {
	var arr []string
	//Get all parameters, except sign
	for key, _ := range parameters {
		str := key + "=" + parameters[key]
		arr = append(arr, str)
	}
	sort.Strings(arr)

	arr = append(arr, secretkey)

	finalstring := method + urlStr + strings.Join(arr, "")
	log.Println(finalstring)
	return ToMd5(url.QueryEscape(finalstring))
}

func BuildMessage(title string, messages string, parameters map[string]interface{}, deviceType string) string {
	var finalMsg string
	switch {
	case deviceType == "4":
		finalMsg = BuildIOSMessage(messages, parameters)
	case deviceType == "3":
		finalMsg = BuildAndroidMessage(title, messages, parameters)
	}
	return finalMsg
}

func BuildAndroidMessage(title string, message string, parameters map[string]interface{}) string {
	dic := make(map[string]interface{})
	dic["description"] = message
	dic["title"] = title
	dic["notification_basic_style"] = "7"
	if parameters != nil {
		dic["custom_content"] = parameters
	}


	jsonString, err := json.Marshal(&dic)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	log.Println(string(jsonString))

	return string(jsonString)
}

func BuildIOSMessage(message string, parameters map[string]interface{}) string {
	dic := make(map[string]string)
	//aps := make(map[string]string)
	//aps["alert"] = message
	//aps["badge"] = "1"
	//aps["sound"] = "Default"
	//dic["aps"] = aps
	//if parameters != nil {
	//	for k, v := range parameters {
	//		dic[k] = v
	//	}
	//}

	jsonString, err := json.Marshal(&dic)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	log.Println(string(jsonString))
	return string(jsonString)
}
