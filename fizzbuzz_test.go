package main

import "testing"

func TestFizzBuzz(t *testing.T) {

    // test case for multiple 3
    t.Run("Fizz", func(t *testing.T) {
        expected := "Fizz"
        result := FizzBuzz(3)
        if result != expected {
            t.Errorf("Expected %s, but got %s", expected, result)
        }
    })

    // test case for multiple 5
    t.Run("Buzz", func(t *testing.T) {
        expected := "Buzz"
        result := FizzBuzz(5)
        if result != expected {
            t.Errorf("Expected %s, but got %s", expected, result)
        }
    })

    // test case for multiple 3 & 5
    t.Run("FizzBuzz", func(t *testing.T) {
        expected := "FizzBuzz"
        result := FizzBuzz(15)
        if result != expected {
            t.Errorf("Expected %s, but got %s", expected, result)
        }
    })

    // Test case for default number
    t.Run("Default", func(t *testing.T) {
        expected := "7"
        result := FizzBuzz(7)
        if result != expected {
            t.Errorf("Expected %s, but got %s", expected, result)
        } else {
			t.Logf("Test case Success")
		}
    })
}
