package vk

import ()

// likes https://new.vk.com/dev/likes
type Likes struct {
	client *VK
}

func (Likes *Likes) GetList(params RequestParams) ([]byte, error) {
	resp, err := Likes.client.CallMethod("likes.getList", params)
	return resp, err
}
