package main

import (
	"bytes"
	"fmt"
	"os"
	"path"

	shell "github.com/memoio/mefs-http-api-go"
)

func main() {
	sh := shell.NewShell("localhost:5001")
	bucketName := "test"
	// 文件路径
	p := path.Join(os.Getenv("HOME"), "testfile")
	file, err := os.Open(p)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}

	objectName := path.Base(file.Name())
	// 上传文件
	_, err = sh.PutObject(file, objectName, bucketName)
	if err != nil {
		fmt.Println("put object err: ", err)
		return
	}

	// 文件列表
	obs, err := sh.ListObjects(bucketName)
	if err != nil {
		fmt.Println("list object err: ", err)
		return
	}
	fmt.Println(obs)

	// 下载文件
	ob, err := sh.GetObject(bucketName, objectName)
	if err != nil {
		fmt.Println("get object err: ", err)
		return
	}
	obuf := new(bytes.Buffer)
	obuf.ReadFrom(ob)

	fmt.Println(ob)
	return
}
