package backup

import "github.com/google/uuid"

type chatStorage struct {
	chatLogs map[uuid.UUID][]string
}

func createStorage() *chatStorage {
	return &chatStorage{
		chatLogs: make(map[uuid.UUID][]string),
	}
}

func (s *chatStorage) addChat(session uuid.UUID, content string) {
	if _, ok := s.chatLogs[session]; !ok {
		s.chatLogs[session] = make([]string, 0, 16)
	}
	s.chatLogs[session] = append(s.chatLogs[session], content)
}

func (s *chatStorage) getFullChat(session uuid.UUID) ([]string, int) {
	logs := s.chatLogs[session]
	return logs, len(logs)
}

func (s *chatStorage) getChatFrom(session uuid.UUID, page int) ([]string, int) {
	if logs, ok := s.chatLogs[session]; ok {
		return logs[page:], len(logs)
	}
	return nil, 0
}
