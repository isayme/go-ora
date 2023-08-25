package main

import (
	"fmt"
	"time"

	"github.com/isayme/go-ora"
)

func main() {
	ora := ora.New()
	ora.Start()
	defer ora.Stop()

	for i := 0; i < 5; i++ {
		ora.Suffix(fmt.Sprintf("loop %d", i))
		time.Sleep(time.Second)
	}

	ora.Color("bgRed", "white")
	ora.Text("bgRed, fgWhite")
	time.Sleep(time.Second)

	ora.Info("Info")
	ora.Warn("Warn")
	ora.Fail("Fail")
	ora.Succeed("Succeed")
}
