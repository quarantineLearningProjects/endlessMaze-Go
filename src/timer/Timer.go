package main

import (
	"fmt"
	"errors"
	"time"
)

type timer struct {
	initialTime, currentTime int
	ticker *time.Ticker
}

// Constructor
func NewTimer(duration int) (timer, error) {
	var t timer
	var err error = nil
	// Error checking
	if duration < 0 {
		err = errors.New("Cannot accept negative parameters")
	}
	t = timer {duration, duration, nil}
	return t, err
}

func (t timer) CurrentTime() int {
	return t.currentTime
}

func (t timer) SetCurrentTime(duration int) error {
	if duration < 0 {
		return errors.New("Cannot accept negative parameters")
	} else {
		t.currentTime = duration
		return nil
	}
}

func (t timer) IsZero() bool {
	return t.currentTime == 0
}

func (t timer) IsPositive() bool {
	return t.currentTime > 0
}

func (t timer) AddSeconds(duration int) error {
	if duration < 0 {
		return errors.New("Cannot accept negative parameters")
	} else {
		t.currentTime += duration * 1000
		return nil
	}
}

func (t timer) SubtractSeconds(duration int) error  {
	newTime := t.currentTime - duration * 1000
	if duration < 0 {
		return errors.New("Cannot accept negative parameters")
	} else if newTime < 0 {
		return errors.New("Resulting time would be negative")
	} else {
		t.currentTime = newTime
		return nil
	}
}

func (t timer) AddMilliseconds(duration int) error {
	if duration < 0 {
		return errors.New("Cannot accept negative parameters")
	} else {
		t.currentTime += duration
		return nil
	}
}

func (t timer) SubtractMilliseconds(duration int) error {
	newTime := t.currentTime - duration
	if duration < 0 {
		return errors.New("Cannot accept negative parameters")
	} else if newTime < 0 {
		return errors.New("Resulting time would be negative")
	} else {
		t.currentTime = newTime
		return nil
	}
}

func (t timer) IncrementSecond() {
	t.AddSeconds(1)
}

func (t timer) DecrementSecond() {
	t.SubtractSeconds(1)
}

func (t timer) IncrementMillisecond() {
	t.AddMilliseconds(1)
}

func (t timer) DecrementMillisecond() {
	t.SubtractMilliseconds(1)
}

func (t timer) Minutes() int {
	return t.currentTime / 60000
}

func (t timer) Seconds() int {
	return t.currentTime / 1000
}

func (t timer) Milliseconds() int {
	return t.currentTime % 1000
}

func (t timer) Start() {
	t.ticker = time.NewTicker(time.Millisecond)
	for _ = range t.ticker.C {
		t.DecrementMillisecond()
		if t.IsZero() {
			t.ticker.Stop()
		}
	}
}

func (t timer) Pause() {
	t.ticker.Stop()
}

func main() {
	r, e := NewTimer(1)
	fmt.Println(r.CurrentTime(), e)
	r.SetCurrentTime(0)
	fmt.Println(r.currentTime)
	fmt.Println(r.IsZero())
}