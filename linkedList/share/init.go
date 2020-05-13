package share

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

var Debug bool

func DebugPrint(args ...interface{}) {
	if !Debug {
		return
	}
	output := append([]interface{}{">>>"}, args...)
	fmt.Println(output...)
}
