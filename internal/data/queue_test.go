package data_test

import (
	"AnonimousChat/internal/data"
	"testing"
)

func TestInterlocutorQueuePush(t *testing.T) {
	queue := data.NewInterlocutorQueue()

	queue.Push(nil)
	queue.Push(nil)
	queue.Push(&data.Interlocutor{})

	queue.Push(&data.Interlocutor{
		Tag: "tg:123",
	})

	queue.Push(&data.Interlocutor{
		Source: "tg",
		ID:     123,
	})

	queue.Push(&data.Interlocutor{
		Tag:              "app:9999999",
		Source:           "app",
		ID:               9999999,
		CountConnections: 1,
		SumDonation:      1,
	})
}

func TestInterlocutorQueuePopOneNil(t *testing.T) {
	queue := data.NewInterlocutorQueue()

	user := queue.Pop()
	if user != nil {
		t.Errorf("Expected user = nil, got: %v", user)
	}
}

func TestInterlocutorQueuePopMoreNull(t *testing.T) {
	queue := data.NewInterlocutorQueue()

	for i := 0; i < 10; i++ {
		queue.Push(nil)
	}

	for i := 0; i < 10; i++ {
		user := queue.Pop()
		if user != nil {
			t.Errorf("Expected user = nil, got: %v", user)
		}
	}
}

func TestInterlocutorQueuePopEqualsPointers(t *testing.T) {
	queue := data.NewInterlocutorQueue()

	testCases := []*data.Interlocutor{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}

	for i := 0; i < len(testCases); i++ {
		queue.Push(testCases[i])
	}

	for i := 0; i < len(testCases); i++ {
		user := queue.Pop()
		if user != testCases[i] {
			t.Errorf("Expected %p pointer, got %p pointer", testCases[i], user)
		}
	}

	queue.Push(testCases[1])
	if queue.Pop() != testCases[1] {
		t.Fail()
	}
}
