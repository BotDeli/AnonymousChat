package manager

import (
	"AnonimousChat/internal/data"
	"sync"
)

type ChatManager interface {
	IsExistsChat(user *data.Interlocutor) bool
	StartSearchChat(user *data.Interlocutor, sendMessage func(text string))
	SendMessageToChat(user *data.Interlocutor, text string)
	CloseChat(user *data.Interlocutor)
}

type chatManager struct {
	activeChats      map[*data.Interlocutor]*data.Interlocutor
	activeSearchChat map[*data.Interlocutor]bool

	queue                data.InterlocutorQueue
	ffSendMessageForUser map[*data.Interlocutor]func(text string)

	mu sync.Mutex
}

func NewChatManager() ChatManager {
	return &chatManager{
		activeChats:          make(map[*data.Interlocutor]*data.Interlocutor),
		activeSearchChat:     make(map[*data.Interlocutor]bool),
		queue:                data.NewInterlocutorQueue(),
		ffSendMessageForUser: make(map[*data.Interlocutor]func(text string)),
	}
}

func (cm *chatManager) IsExistsChat(user *data.Interlocutor) bool {
	cm.mu.Lock()
	exists := cm.activeChats[user] != nil
	cm.mu.Unlock()

	return exists
}

func (cm *chatManager) StartSearchChat(user *data.Interlocutor, sendMessage func(text string)) {
	if cm.activeSearchChat[user] {
		return
	}

	if cm.IsExistsChat(user) {
		cm.CloseChat(user)
	}

	sendMessage("Поиск собеседника начат...")

	cm.mu.Lock()

	cm.activeSearchChat[user] = true
	cm.ffSendMessageForUser[user] = sendMessage

	interlocutor := cm.queue.Pop()
	if interlocutor == nil {
		cm.queue.Push(user)
		cm.mu.Unlock()
		return
	}

	cm.activeChats[user] = interlocutor
	cm.activeChats[interlocutor] = user

	sendMessage("Собеседник найден!")
	cm.ffSendMessageForUser[interlocutor]("Собеседник найден!")

	cm.activeSearchChat[user] = false
	cm.activeSearchChat[interlocutor] = false

	cm.mu.Unlock()
}

func (cm *chatManager) SendMessageToChat(user *data.Interlocutor, text string) {
	cm.mu.Lock()

	interlocutor := cm.activeChats[user]
	if interlocutor == nil {
		cm.mu.Unlock()
		return
	}

	cm.ffSendMessageForUser[interlocutor](text)

	cm.mu.Unlock()
}

func (cm *chatManager) CloseChat(user *data.Interlocutor) {
	cm.mu.Lock()

	interlocutor := cm.activeChats[user]
	if interlocutor == nil {
		cm.mu.Unlock()
		return
	}

	cm.ffSendMessageForUser[user]("Вы завершили диалог.")
	cm.ffSendMessageForUser[interlocutor]("Собеседник завершил диалог.")

	delete(cm.activeChats, user)
	delete(cm.activeChats, interlocutor)

	delete(cm.ffSendMessageForUser, user)
	delete(cm.ffSendMessageForUser, interlocutor)

	cm.mu.Unlock()
}
