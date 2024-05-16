package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	c1 := sig(2 * time.Second)
	c2 := sig(1 * time.Second)

	resultCh := Or(c1, c2)

	start := time.Now()

	<-resultCh

	if time.Since(start).Seconds() > 2 {
		t.Errorf("Expected less than or equal to 1 second delay")
	}
}
