package main

import "testing"

func TestSum(t *testing.T) {
    total := Sum(17, 17)

    if total != 34 {
        t.Errorf("\nExpected: %d\nGot: %d", 34, total)
    }
}
