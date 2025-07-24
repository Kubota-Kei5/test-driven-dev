package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestMoneyMultiplier1(t *testing.T) {
//     assert.Equal(t, MoneyMultiplier(0, 7), 0,  "0 * 7 should equal 0")
// }

// func TestMoneyMultiplier2(t *testing.T) {
//     assert.Equal(t, MoneyMultiplier(2, 5), 10,  "2 * 5 should equal 10")
// }


func TestMoneyMultiplier1(t *testing.T) {
    t.Run("Real Number Multiplication", func(t *testing.T) {
        assert.Equal(t, MoneyMultiplier(2, 5), 10,  "2 * 5 should equal 10")
    })
    t.Run("Zero Multiplication", func(t *testing.T) {
        assert.Equal(t, MoneyMultiplier(0, 7), 0,  "0 * 7 should equal 0")
    })
}
