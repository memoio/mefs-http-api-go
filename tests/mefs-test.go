package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	shell "github.com/memoio/mefs-http-api-go"
)

const (
	TESTBUCKET     = "b0"
	RandomDataSize = 1024 * 1024 * 100
)

//测试时空值的计算，流程：没过一段时间，user发送一段数据，keeper进行一次计算，将实际值与理论值对比
//目前使用iptb进行测试，传入进行测试的user和keeper端口
func ReslultSumaryTest(userPort, keeperPort string) {
	theory := int64(0)                //时空值的理论值
	testLastSize := int64(0)          //上一次put文件的大小
	testLastTime := time.Now().Unix() //上一次测试的时间
	userSh := shell.NewShell("localhost:" + userPort)
	keeperSh := shell.NewShell("localhost:" + keeperPort)

	userSh.CreateBucket(TESTBUCKET) //先建bucket

	for i := 1; ; i++ { //每过30分钟进行一次测试
		fmt.Println("======================")
		testThisTime := time.Now().Unix() //本次测试时间
		fmt.Println("本次测试时间：", time.Unix(testThisTime, 0).In(time.Local))
		testThisSize := rand.Int63n(RandomDataSize)
		fmt.Println("传入文件大小:", testThisSize)
		data := make([]byte, testThisSize)
		fillRandom(data)
		buf := bytes.NewBuffer(data)
		objectName := "test_" + strconv.Itoa(int(testThisSize))
		_, err := userSh.PutObject(buf, objectName, TESTBUCKET)
		if err != nil {
			fmt.Println("PutObject err!", err)
		}
		theory += (testThisTime - testLastTime) * testLastSize
		actual := keeperSh.ResultSummary() //时空支付的实际值
		fmt.Println("实际值：", actual)
		fmt.Println("理论值", theory/3)
		testLastTime = testThisTime
		testLastSize += testThisSize
		fmt.Println("======================\n")
		time.Sleep(30 * time.Minute)
	}

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
