package main

import (
	"testing"
	"time"
) 

func TestBirhtParsing(t *testing.T) {
        mem := New("1995-12-29", "80")
        exp := time.Date(1995, 12, 29, 0, 0, 0, 0, time.UTC)
        if mem.Birth != exp {
                t.Errorf("Expected %v, got %v", exp, mem.Birth)
        }
}

func TestCaclcWeek(t *testing.T) {
        mem := New("1995-12-29", "80")
        exp := 80 * 52
        res := mem.CalcWeek()
        if res != exp {
                t.Errorf("Expected %v, got %v", exp, res)
        }
}

func TestCalcElapsedWeeks(t *testing.T) {
        mem := New("1995-12-29", "80")
        exp := 52 * 25 + 4
        now := time.Date(2020, 12, 29, 0, 0, 0, 0, time.UTC)
        res := mem.CalcElapsedWeeks(now)

        if res != exp {
                t.Errorf("Expected %v, got %v", exp, res)
        }



}
