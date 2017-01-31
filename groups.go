package vk

import ()

// Groups https://new.vk.com/dev/groups
type Groups struct {
	client *VK
}

func (Groups *Groups) Search(params RequestParams) ([]byte, error) {
	resp, err := Groups.client.CallMethod("groups.search", params)
	return resp, err
}

func (Groups *Groups) Get(params RequestParams) ([]byte, error) {
	resp, err := Groups.client.CallMethod("groups.get", params)
	return resp, err
}

func (Groups *Groups) GetById(params RequestParams) ([]byte, error) {
	resp, err := Groups.client.CallMethod("groups.getById", params)
	return resp, err
}
