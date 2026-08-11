package repo

type EventType string

// Actions
const (
	EventChatToGoblin EventType = "chat_to_goblin"
	EventChatToTeam   EventType = "chat_to_team"
	EventAskTeam      EventType = "ask_team"
	EventOffer        EventType = "offer"
)

// Status Changes
const (
	EventChatInit       EventType = "init"
	EventChatFromGoblin EventType = "chat_from_goblin"
	EventChatFromTeam   EventType = "chat_from_team"
	EventSentiment      EventType = "Sentiment"
	EventTeamEmotion    EventType = "emotion"
	EventEnd            EventType = "end"
)
