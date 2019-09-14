package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"

	shell "github.com/memoio/mefs-http-api-go"
)

// 需要提前准备好，已充值测试Token的User
// 否则如果要在Shell里create，则需要写转账逻辑

const UserSK = ""

const UserAddr = ""

func main() {
	//连接已经启动的MEFS,如果非本机，还需设置跨域访问
	sh := shell.NewShell("localhost:5001")

	//通过私钥，启动User
	err := sh.StartUser(UserAddr, shell.SetSecretKey(UserSK))
	if err != nil {
		log.Println(err)
		return
	}

	//设置地址选项
	addrOpt := shell.SetOp("address", UserAddr)

	//在该User的加密存储空间内新建一Bucket
	bucket, err := sh.CreateBucket("bucket01", addrOpt)
	if err != nil {
		log.Println(err)
		return
	}
	//构造随机文件
	r := rand.Int63n(1024 * 1024 * 40)
	data := make([]byte, r)
	fillRandom(data)
	buf := bytes.NewBuffer(data)
	objectName := UserAddr + "_" + strconv.Itoa(int(r))

	// 上传文件
	ob, err := sh.PutObject(buf, objectName, bucket.Buckets[0].BucketName, addrOpt)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(ob)
	// 下载该文件
	readerCloser, err := sh.GetObject(ob.Objects[0].ObjectName, bucket.Buckets[0].BucketName, addrOpt)
	if err != nil {
		log.Println(err)
		return
	}
	tempfile, err := os.Create("/tmp/tempfile")
	if err != nil {
		log.Println(err)
		return
	}
	// 将下载的文件暂存
	_, err = io.Copy(tempfile, readerCloser)
	if err != nil {
		log.Println(err)
		return
	}
	//查看一些信息
	bks, err := sh.ListBuckets(addrOpt)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(bks)
	//查看一些信息
	obs, err := sh.ListObjects(bks.Buckets[0].BucketName, addrOpt)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(obs)
}

func fillRandom(p []byte) {
	for i := 0; i < len(p); i += 7 {
		val := rand.Int63()
		for j := 0; i+j < len(p) && j < 7; j++ {
			p[i+j] = byte(val)
			val >>= 8
		}
	}
}
