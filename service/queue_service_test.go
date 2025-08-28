package service

import (
	"testing"
)

func TestCreateNewQueue(t *testing.T) {
	newQueue, err := CreateNewQueue("张三", "166 8888 6666", 3)
	if err != nil {
		t.Errorf("CreateNewQueue err: %v", err)
	}
	t.Logf("%v", newQueue)
}
