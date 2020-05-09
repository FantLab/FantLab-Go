package timeid

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/binary"
	"sync"
	"time"
)

var mutex sync.Mutex

func Generate(t time.Time) string {
	bytes := [20]byte{}

	var err error

	mutex.Lock()
	_, err = rand.Read(bytes[4:])
	mutex.Unlock()

	if err != nil {
		panic(err)
	}

	binary.BigEndian.PutUint32(bytes[:4], uint32(t.Unix()-14e8))

	return base32.StdEncoding.EncodeToString(bytes[:])
}

func GenerateNow() string {
	return Generate(time.Now())
}
