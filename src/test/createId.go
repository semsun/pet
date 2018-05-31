package main

import (
	"fmt"
	"io"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

func GetMd5String(s string) string {
	h := md5.New();
	h.Write([]byte(s));
	return hex.EncodeToString(h.Sum(nil));
}

func UniqueId() string {
	b := make([]byte, 48);

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "";
	}

	fmt.Printf("Middle: %s\n", base64.URLEncoding.EncodeToString(b));

	return GetMd5String(base64.URLEncoding.EncodeToString(b));
}

func main() {
	fmt.Printf("Md5:%s\n", GetMd5String("asdfe"));
	fmt.Printf("UniqueId:%s\n", UniqueId());
}