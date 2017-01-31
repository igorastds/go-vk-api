package vk

import ()

// Wall https://new.vk.com/dev/wall
type Wall struct {
	client *VK
}

func (Wall *Wall) GetComments(params RequestParams) ([]byte, error) {
	resp, err := Wall.client.CallMethod("wall.getComments", params)
	return resp, err
}

func (Wall *Wall) GetReposts(params RequestParams) ([]byte, error) {
	resp, err := Wall.client.CallMethod("wall.getReposts", params)
	return resp, err
}

func (Wall *Wall) GetById(params RequestParams) ([]byte, error) {
	resp, err := Wall.client.CallMethod("wall.getById", params)
	return resp, err
}

func (Wall *Wall) Get(params RequestParams) ([]byte, error) {
	resp, err := Wall.client.CallMethod("wall.get", params)
	return resp, err
}
