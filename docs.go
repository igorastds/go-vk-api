package vk

import ()

// Docs https://new.vk.com/dev/docs
type Docs struct {
	client *VK
}

func (Docs *Docs) GetById(params RequestParams) ([]byte, error) {
	resp, err := Docs.client.CallMethod("docs.getById", params)
	return resp, err
}
