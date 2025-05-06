package uuid

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

/*
  @Author : lanyulei
  @Desc :
*/

func Get() (uid string) {
	uuidWithHyphen := uuid.New()
	uid = strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return
}

// GetWorkOrderUuid WF + 时间戳 + 随机数
func GetWorkOrderUuid() (uid string) {
	timestamp := time.Now().Format("20060102150405")
	random := generateRandomString(4)
	return "WF" + timestamp + random
}

func generateRandomString(length int) string {
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rng.Intn(len(letterBytes))]
	}
	return string(b)
}
