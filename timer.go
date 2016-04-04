package timer

import "time"

const (
	StateIdle = iota
	StateActive
	StateExpired
)

type Timer struct {
	C         chan time.Time
	duration  time.Duration
	state     int
	fn        func()
	startedAt time.Time
	t         *time.Timer
}

func AfterFunc(d time.Duration, f func()) *Timer {
	t := new(Timer)
	t.duration = d
	t.fn = func() {
		t.state = StateExpired
		f()
	}
	return t
}

func NewTimer(d time.Duration) *Timer {
	c := make(chan time.Time, 1)
	t := new(Timer)
	t.C = c
	t.duration = d
	t.fn = func() {
		t.state = StateExpired
		t.C <- time.Now()
	}
	return t
}

func (t *Timer) Pause() bool {
	if t.state != StateActive {
		return false
	}
	if !t.t.Stop() {
		t.state = StateExpired
		return false
	}
	t.state = StateIdle
	dur := time.Now().Sub(t.startedAt)
	t.duration = t.duration - dur
	return true
}

func (t *Timer) Start() bool {
	if t.state != StateIdle {
		return false
	}
	t.startedAt = time.Now()
	t.state = StateActive
	t.t = time.AfterFunc(t.duration, t.fn)
	return true
}

func (t *Timer) Stop() bool {
	if t.state != StateActive {
		return false
	}
	t.startedAt = time.Now()
	t.state = StateExpired
	t.t.Stop()
	return true
}
