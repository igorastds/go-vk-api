package vk

import ()

// Photos https://new.vk.com/dev/photos
type Photos struct {
	client *VK
}

func (Photos *Photos) GetComments(params RequestParams) ([]byte, error) {
	resp, err := Photos.client.CallMethod("photos.getComments", params)
	return resp, err
}

func (Photos *Photos) GetById(params RequestParams) ([]byte, error) {
	resp, err := Photos.client.CallMethod("photos.getById", params)
	return resp, err
}
