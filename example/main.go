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

	ora.Info("Info")
	ora.Warn("Warn")
	ora.Fail("Fail")
	ora.Succeed("Succeed")

	ora.Start()
	for i := 0; i < 100; i += 10 {
		ora.Suffix(fmt.Sprintf("Progress %d%%", i))
		time.Sleep(time.Millisecond * 500)
	}
	ora.Succeed("Progress 100%")
}
