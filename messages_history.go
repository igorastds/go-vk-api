package vk

import (
	"encoding/json"
	"errors"
)

type Msg struct {
	Id           int         `json:"id"`
	Body         string      `json:"body"`
	User_Id      int         `json:"user_id"`
	From_Id      int         `json:"from_id"`
	Date         int         `json:"date"`
	Read_State   int         `json:"read_state"`
	Out          int         `json:"out"`
	Random_Id    int         `json:"random_id"`
	Fwd_Messages interface{} `json:"fwd_messages"`
}
type Response struct {
	Count int64
	Items []Msg
}
type JSONBody struct {
	Response Response `json:"response"`
}

func (m *Messages) deprecatedGetHistory(params RequestParams) (JSONBody, error) {
	resp, err := m.client.CallMethod("messages.getHistory", params)

	var body JSONBody

	if err != nil {
		return body, err
	}

	if err := json.Unmarshal(resp, &body); err != nil {
		return body, err
	}
	if body.Response.Count < 1 {
		return body, errors.New("zero count")
	}
	return body, nil
}

func (m *Messages) deprecatedDeleteMessage(dialogId int, messageId int) (err error) {
	err = nil
	return
}
