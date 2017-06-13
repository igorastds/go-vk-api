package vk

import ()

// Database https://new.vk.com/dev/photos
type Database struct {
	client *VK
}

func (database *Database) GetChairs(params RequestParams) ([]byte, error) {
	resp, err := database.client.CallMethod("database.getChairs", params)
	return resp, err
}

func (database *Database) GetFaculties(params RequestParams) ([]byte, error) {
	resp, err := database.client.CallMethod("database.getFaculties", params)
	return resp, err
}


func (database *Database) GetUniversities(params RequestParams) ([]byte, error) {
	resp, err := database.client.CallMethod("database.getUniversities", params)
	return resp, err
}

func (database *Database) GetSchools(params RequestParams) ([]byte, error) {
	resp, err := database.client.CallMethod("database.getSchools", params)
	return resp, err
}

func (database *Database) GetCities(params RequestParams) ([]byte, error) {
	resp, err := database.client.CallMethod("database.getCities", params)
	return resp, err
}

func (database *Database) GetRegions(params RequestParams) ([]byte, error) {
	resp, err := database.client.CallMethod("database.getRegions", params)
	return resp, err
}
