package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//生成手机验证码
func GeneratePhoneCode(codeLen uint) string {

	//默认长度为6
	if codeLen <= 0 {
		codeLen = 6
	}

	var codeList []string
	randNew := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := uint(0); i < codeLen; i++ {
		randValue := randNew.Intn(10)
		randString := fmt.Sprintf("%d", randValue)
		codeList = append(codeList, randString)
	}
	codeValue := strings.Join(codeList, "")
	return codeValue

}
