package websockets

import (
	"html"

	"aahframe.work/ws"
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
func (w *SimpleChatWebSocket) SimpleChat() {
	for {
		text, err := w.ReadText()
		if err != nil {
			if ws.IsDisconnected(err) {
				// WebSocket client is gone, exit here
				return
			}

			w.Log().Error(err)
			continue // we are moving on to next WS frame/msg
		}

		w.Log().Info("Received from client: ", text)

		// if no error, echo back the Text to WebSocket client
		if err = w.ReplyText(html.UnescapeString(text)); err != nil {
			w.Log().Error(err)
		}
	}
}
