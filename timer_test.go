package timer

import (
	"testing"
	"time"
)

func TestStartStop(t *testing.T) {
	var now = time.Now()
	var timer = AfterFunc(time.Second, func() {
		println(time.Now().Sub(now))
		if time.Now().Sub(now) < time.Second*2 {
			t.Fatal("Early func")
		}
	})
	timer.Start()
	time.Sleep(time.Microsecond * 500)
	timer.Pause()
	time.Sleep(time.Second)
	timer.Start()
	time.Sleep(time.Second)
}
