package websockets

import "aahframe.work/ws"

// HandleEvents method is to demostrate aah WebSocket events published for each
// WebSocket connection made to aah server and its life cycle.
func HandleEvents(eventName string, ctx *ws.Context) {
	switch eventName {
	case ws.EventOnPreConnect:
		ctx.Log().Infof("I'm %s event", eventName)
	case ws.EventOnPostConnect:
		ctx.Log().Infof("I'm %s event", eventName)
		// aah assigns unique ID for each WebSocket Request.
		// Unique ID generated using MongoDB Object ID algrothim.
		_ = ctx.ReplyText("Greetings from Server!")
		_ = ctx.ReplyText("Just letting you know, your ID: " + ctx.Req.ID)
	case ws.EventOnPostDisconnect:
		ctx.Log().Infof("I'm %s event", eventName)
		ctx.Log().Info("WebSocket client disconnected")
	case ws.EventOnError:
		ctx.Log().Infof("I'm %s event", eventName)
		ctx.Log().Errorf("Ooh no! error occurred: %s", ctx.ErrorReason())
	}
}
