package kloudless

import "time"

type File struct {
	ID       string    `json:"id"`
	RawID    string    `json:"raw_id"`
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Parent   struct {
		ID string `json:"id"`
	} `json:"parent"`
	CanCreateFolders   bool `json:"can_create_folders"`
	CanUploadFiles     bool `json:"can_upload_files"`
	CanListRecursively bool `json:"can_list_recursively"`
	AccID              int  `json:"account"`
}
