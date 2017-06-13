package vk

import ()

// Users https://vk.com/dev/users.search
type Users struct {
	client *VK
}

func (users *Users) Search(params RequestParams) ([]byte, error) {
	resp, err := users.client.CallMethod("users.search", params)
	return resp, err
}
