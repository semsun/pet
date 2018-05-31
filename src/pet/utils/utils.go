package utils

import (
	"fmt"
	"time"
	"encoding/pem"
	"crypto/x509"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"bytes"
	"io"
)

func GetCertificateOwner(creatorByte []byte) string {
	certStart := bytes.IndexAny(creatorByte, "-----");
	if certStart == -1 {
			fmt.Printf("No certificate found. [%d, %s]\n", certStart, string(creatorByte));
			return "";
	}
	fmt.Printf("Certificate Start %d, %s.\n", certStart, string(creatorByte));

	certText := creatorByte[certStart:];
	bl, _ := pem.Decode(certText);
	if bl == nil {
			fmt.Errorf("Could not decode the PEM structure\n");
			return "";
	}

	fmt.Println(string(certText));
	cert, err := x509.ParseCertificate(bl.Bytes);
	if err != nil {
			fmt.Errorf("ParseCertificate failed");
			return "";
	}
	fmt.Println(cert);
	uname:=cert.Subject.CommonName;
	fmt.Println("Name:"+uname);

	return uname;
}

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

func CurTimestamp() int64 {
	return time.Now().UTC().UnixNano();
}