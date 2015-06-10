package util

import (
	"crypto/md5"
	"encoding/hex"
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
