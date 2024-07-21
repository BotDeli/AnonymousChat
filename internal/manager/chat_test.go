package manager_test

import (
	"AnonimousChat/internal/data"
	"AnonimousChat/internal/manager"
	"testing"
	"time"
)

var durationWaiting = 1 * time.Second

func TestIsDontExistsChat(t *testing.T) {
	chatManager := manager.NewChatManager()
	user := data.NewInterlocutorFromTelegram(123456)

	exists := chatManager.IsExistsChat(user)
	if exists {
		t.Error("User dont find chat, already exists chat")
	}
}

func TestStartSearchChatOneUser(t *testing.T) {
	chatManager := manager.NewChatManager()
	user := data.NewInterlocutorFromTelegram(1)

	ff := func(text string) {
		if text == "Собеседник найден!" {
			t.Fatal("Interlocutor finded, have only one user")
		}
	}

	chatManager.StartSearchChat(user, ff)

	time.Sleep(durationWaiting)
}

func TestMultySearchChatOneUser(t *testing.T) {
	chatManager := manager.NewChatManager()
	user := data.NewInterlocutorFromTelegram(1)

	var countCalls int

	ff := func(text string) {
		countCalls++
	}

	for i := 0; i < 10; i++ {
		chatManager.StartSearchChat(user, ff)
	}

	time.Sleep(durationWaiting)

	if countCalls != 1 {
		t.Fail()
	}
}

func TestStartSearchChatTwoUsers(t *testing.T) {
	chatManager := manager.NewChatManager()
	user1 := data.NewInterlocutorFromTelegram(1)
	user2 := data.NewInterlocutorFromTelegram(2)

	var countCalls int

	ff := func(text string) {
		countCalls++
	}

	chatManager.StartSearchChat(user1, ff)
	chatManager.StartSearchChat(user2, ff)

	time.Sleep(durationWaiting)

	if countCalls != 4 {
		t.Fail()
	}
}

func TestMultySearchChatTwoUser(t *testing.T) {
	chatManager := manager.NewChatManager()
	user1 := data.NewInterlocutorFromTelegram(1)
	user2 := data.NewInterlocutorFromTelegram(2)

	var countCalls int

	ff := func(text string) {
		countCalls++
	}

	countSearch := 10

	for i := 0; i < countSearch; i++ {
		chatManager.StartSearchChat(user1, ff)
		chatManager.StartSearchChat(user2, ff)
	}

	time.Sleep(durationWaiting)

	if countCalls != (countSearch-1)*6+4 {
		t.Fail()
	}
}

func TestExistsChatTwoUsers(t *testing.T) {
	chatManager := manager.NewChatManager()
	user1 := data.NewInterlocutorFromTelegram(1)
	user2 := data.NewInterlocutorFromTelegram(2)

	if chatManager.IsExistsChat(user1) || chatManager.IsExistsChat(user2) {
		t.Fatal("Expected dont exists chat")
	}

	ff := func(text string) {}

	chatManager.StartSearchChat(user1, ff)
	chatManager.StartSearchChat(user2, ff)

	if !(chatManager.IsExistsChat(user1) && chatManager.IsExistsChat(user2)) {
		t.Fatal("Expected all exists chats")
	}
}

func TestSearchChatThreeUser(t *testing.T) {
	chatManager := manager.NewChatManager()
	user1 := data.NewInterlocutorFromTelegram(1)
	user2 := data.NewInterlocutorFromTelegram(2)

	var countCalls int

	ff := func(text string) {
		countCalls++
	}

	chatManager.StartSearchChat(user1, ff)
	chatManager.StartSearchChat(user2, ff)

	time.Sleep(durationWaiting)

	user3 := data.NewInterlocutorFromTelegram(3)
	countCalls = 0

	chatManager.StartSearchChat(user3, ff)

	time.Sleep(durationWaiting)

	if countCalls != 1 {
		t.Fail()
	}
}

func TestExistsChatThreeUsers(t *testing.T) {
	chatManager := manager.NewChatManager()
	user1 := data.NewInterlocutorFromTelegram(1)
	user2 := data.NewInterlocutorFromTelegram(2)

	ff := func(text string) {}

	chatManager.StartSearchChat(user1, ff)
	chatManager.StartSearchChat(user2, ff)

	time.Sleep(durationWaiting)

	user3 := data.NewInterlocutorFromTelegram(3)

	chatManager.StartSearchChat(user3, ff)

	if !(chatManager.IsExistsChat(user1) && chatManager.IsExistsChat(user2)) {
		t.Fatal("Expected all exists chats")
	}

	if chatManager.IsExistsChat(user3) {
		t.Fatal("Expected dont exists chat")
	}
}

func TestCloseChat(t *testing.T) {
	chatManager := manager.NewChatManager()
	user1 := data.NewInterlocutorFromTelegram(1)
	user2 := data.NewInterlocutorFromTelegram(2)

	var countCalls int

	ff := func(text string) {
		countCalls++
	}

	chatManager.StartSearchChat(user1, ff)
	chatManager.StartSearchChat(user2, ff)

	time.Sleep(durationWaiting)

	chatManager.CloseChat(user1)

	if countCalls != 6 {
		t.Fail()
	}
}

func TestSendMessageTwoUsers(t *testing.T) {
	chatManager := manager.NewChatManager()
	user1 := data.NewInterlocutorFromTelegram(1)
	user2 := data.NewInterlocutorFromTelegram(2)

	var lastMessage1 string
	ff1 := func(text string) {
		lastMessage1 = text
	}

	var lastMessage2 string
	ff2 := func(text string) {
		lastMessage2 = text
	}

	chatManager.StartSearchChat(user1, ff1)
	chatManager.StartSearchChat(user2, ff2)

	time.Sleep(durationWaiting)

	expectedStartMessage := "Собеседник найден!"

	if lastMessage1 != expectedStartMessage || lastMessage2 != expectedStartMessage {
		t.Errorf("Dont correct response message for search chat, expected: \"%s\", got: \"%s\" and \"%s\"", expectedStartMessage, lastMessage1, lastMessage2)
	}

	currentText := "Привет"
	chatManager.SendMessageToChat(user1, currentText)
	chatManager.SendMessageToChat(user2, currentText)

	if lastMessage1 != currentText || lastMessage2 != currentText {
		t.Errorf("Dont correct response message, expected: \"%s\", got \"%s\" and \"%s\"", currentText, lastMessage1, lastMessage2)
	}

	currentText = "Как дела?"
	chatManager.SendMessageToChat(user1, currentText)
	if lastMessage2 != currentText {
		t.Errorf("Expected message: \"%s\", got \"%s\"", currentText, lastMessage2)
	}

	if lastMessage1 == currentText {
		t.Errorf("Dont expected change last message for user1")
	}

	currentText = "Все хорошо)"
	chatManager.SendMessageToChat(user2, currentText)
	if lastMessage1 != currentText {
		t.Errorf("Expected message: \"%s\", got \"%s\"", currentText, lastMessage1)
	}

	if lastMessage2 == currentText {
		t.Errorf("Dont expected change last message for user2")
	}
}

func TestSendMessagesFromOneUser(t *testing.T) {
	chatManager := manager.NewChatManager()
	user1 := data.NewInterlocutorFromTelegram(1)
	user2 := data.NewInterlocutorFromTelegram(2)

	var lastMessage1 string
	ff1 := func(text string) {
		lastMessage1 = text
	}

	var lastMessage2 string
	ff2 := func(text string) {
		lastMessage2 = text
	}

	chatManager.StartSearchChat(user1, ff1)
	chatManager.StartSearchChat(user2, ff2)

	time.Sleep(durationWaiting)

	messages := []string{
		"Привет",
		"Как твои дела?",
		"Ты чего меня игноришь :/",
		"Да ну тебя...",
		"Пока!",
	}

	expectedStartMessage := "Собеседник найден!"

	for i := 0; i < len(messages); i++ {
		chatManager.SendMessageToChat(user1, messages[i])
		if lastMessage2 != messages[i] {
			t.Fatalf("Expected change message to \"%s\", got \"%s\"", messages[i], lastMessage2)
		}

		if lastMessage1 != expectedStartMessage {
			t.Fatal("Last message changed for user1")
		}
	}
}
