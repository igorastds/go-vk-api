package vk

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"strings"
)

// Messages https://new.vk.com/dev/messages
type Messages struct {
	client *VK
}

func (messages *Messages) Delete(params RequestParams, multiple bool) (int, error) {
	body, err := messages.client.CallMethod("messages.delete", params)

	type ResponseFlag struct {
		Response int64 `json:"response"`
	}

	type ResponseFlagMultiple struct {
		Response map[string]int `json:"response"`
	}

	deleted := 0
	deleted_unmarshal_cause := false
	if !multiple {
		var resp ResponseFlag
		err = json.Unmarshal(body, &resp)
		if err != nil {
			//return 0, errors.New(fmt.Sprintf("cannot unmarshal delete response: %v, body: %s", err, string(body)))
			deleted_unmarshal_cause = true
			goto delete_multi
		}
		if resp.Response != 1 {
			return 0, errors.New("no success deleting message")
		}
		goto delete_done

	}

delete_multi:
	if multiple {

		var resp ResponseFlagMultiple
		err = json.Unmarshal(body, &resp)
		if err != nil {
			return 0, errors.New(fmt.Sprintf("cannot unmarshal delete response: %v, body: %s", err, string(body)))
		}

		if deleted_unmarshal_cause {
			log.Warnf("delete: response was in multi message format for single message request: %s", string(body))
		}

		failedKeys := make([]string, 0)
		for k, v := range resp.Response {
			if v == 0 {
				failedKeys = append(failedKeys, k)
			} else {
				deleted++
			}
		}

		if len(failedKeys) > 0 {
			return deleted, errors.New(fmt.Sprintf("failed to delete following messages: %s", strings.Join(failedKeys, ",")))
		}

	}

delete_done:
	return deleted, nil
}

// Send https://new.vk.com/dev/messages.send
func (messages *Messages) Send(params RequestParams) (int64, error) {
	resp, err := messages.client.CallMethod("messages.send", params)
	if err != nil {
		return 0, err
	}

	type JSONBody struct {
		MessageID int64 `json:"response"`
	}

	var body JSONBody

	if err := json.Unmarshal(resp, &body); err != nil {
		return 0, err
	}

	return body.MessageID, nil
}

func (messages *Messages) GetHistory(params RequestParams) ([]byte, error) {
	resp, err := messages.client.CallMethod("messages.getHistory", params)
	return resp, err
}

func (messages *Messages) GetDialogs(params RequestParams) ([]byte, error) {
	resp, err := messages.client.CallMethod("messages.getDialogs", params)
	return resp, err

	/*	if err != nil {
		return nil, err
	}*/
}
