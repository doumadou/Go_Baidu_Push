package Util

import (
	. "Go_Baidu_Push/Config"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func GenerateSignature(urlString string, para map[string]string) string {
	httpMethod := "POST"
	var arr []string
	//Get all parameters, except sign
	for key, v := range para {
		str := key + "=" + v
		arr = append(arr, str)
	}
	sort.Strings(arr)
	finalstring := strings.Join(arr, "")
	baseString := httpMethod + urlString + finalstring + SECRET_KEY
	encodedString := string(url.Parse(baseString))
	hs := sha1.New()
	hs.Write([]byte(encodedString))
	sign := hex.EncodeToString(hs.Sum(nil))
	return sign
}

func GetURLWithParameters(url string, para map[string]string) (string, error) {
	var result string
	sign := GenerateSignature(para)
	para["sign"] = sign
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	form := url.Values{}
	for k, v := range para {
		form.Add(k, v)
	}
	req.PostForm = from

	resp, err := client.Do(req)
	defer resp.Body.Close()
	bd, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(bd))
	return string(bd), nil
}
