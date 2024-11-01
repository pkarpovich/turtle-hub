package bot

type Openai struct{}

func NewOpenai() *Openai {
	return &Openai{}
}

func (l *Openai) ShouldHandle(msg Message) bool {
	return true
}

func (l *Openai) OnMessage(msg Message) Response {
	return Response{
		Reaction: &ResponseReaction{
			MessageID: msg.ID,
			Emoji:     "👀",
		},
		ChatID: msg.ChatID,
	}
}
