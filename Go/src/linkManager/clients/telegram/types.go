package telegram

const(
	getUpdatesMethod = "getUpdates"
	sendMessageMethod = "sendMessage"

)

type Update struct {
	ID      int    `json:"update_id"`
	Message string `json:"message"`
}

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}
