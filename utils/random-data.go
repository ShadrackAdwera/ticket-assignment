package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "qazwsxedcrfvtgbyhnujmikolp"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func GenerateRandomString() string {
	return RandomString(10)
}

func GenerateTicketTitle() string {
	return RandomString(15)
}

func GenerateTicketDescription() string {
	return RandomString(60)
}

func GetAgentStatus() string {
	status := []string{ACTIVE, INACTIVE}
	return status[rand.Intn(len(status))]
}

func RandomInteger(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
