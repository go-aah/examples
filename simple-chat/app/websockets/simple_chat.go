package websockets

import (
	"html"

	"aahframework.org/ws.v0"
)

// SimpleChatWebSocket hold simple chat implementation.
type SimpleChatWebSocket struct {
	*ws.Context
}

// SimpleChat method recevies chat text from chat client and echo's
// back to them and logs on server side.
//
// Route: /simple-chat
// Method: WS
func (s *SimpleChatWebSocket) SimpleChat() {
	for {
		text, err := s.ReadText()
		if err != nil {
			if ws.IsDisconnected(err) {
				// WebSocket client is gone, exit here
				return
			}

			s.Log().Error(err)
			continue // we are moving on to next WS frame/msg
		}

		s.Log().Info("Received from client: ", text)

		// if no error, echo back the Text to WebSocket client
		if err = s.ReplyText(html.UnescapeString(text)); err != nil {
			s.Log().Error(err)
		}
	}
}
