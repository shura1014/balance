package balance

import (
	"testing"
)

type TestNode struct {
}

func (node TestNode) IsValid() bool {
	return true
}

func TestGetBalance(t *testing.T) {
	balance := GetBalance[TestNode]("random")
	t.Log(balance.Name())
}
