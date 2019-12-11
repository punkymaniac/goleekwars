package apileek

import (
	"strconv"
	"encoding/json"
)

type AiFolderService struct {
	apiService
}

type id struct {
    Id uint64 `json:"id"`
}

// Change a AI folder location
func (s *AiFolderService) ChangeFolder(
    folderId uint64, // Folder id to move
    dstFolderId uint64, // Folder id to move the folder in
) (error) {
    data := "folder_id=" + strconv.FormatUint(folderId, 10)
    resp, body, err := s.apiRequest("POST", s.url + "change_folder/", &data)
    if err != nil {
        return err
    }

    if resp.StatusCode != 201 {
        return newApiError(resp, body)
    }

    return nil
}

// Delete a AI folder
func (s *AiFolderService) Delete(
    folderId uint64, // Folder id to delete
) (error) {
    data := "folder_id=" + strconv.FormatUint(folderId, 10)
    resp, body, err := s.apiRequest("POST", s.url + "delete/", &data)
    if err != nil {
        return err
    }

    if resp.StatusCode != 200 {
        return newApiError(resp, body)
    }

    return nil
}

// Create a AI folder
func (s *AiFolderService) New(
    folderId uint64, // Parent folder id
) (uint64, error) {
    data := "folder_id=" + strconv.FormatUint(folderId, 10)
    resp, body, err := s.apiRequest("POST", s.url + "new/", &data)
    if err != nil {
        return 0, err
    }

    if resp.StatusCode != 200 {
        return 0, newApiError(resp, body)
    }

    var obj = id{}
    err = json.Unmarshal([]byte(body), &obj)
    if err != nil {
        return 0, err
    }

    return obj.Id, nil
}

// Rename a AI folder
func (s *AiFolderService) Rename(
    folderId uint64, // Folder id to rename
    name string, // The new name of the folder
) (error) {
    data := "folder_id=" + strconv.FormatUint(folderId, 10) + "&new_name=" + name
    resp, body, err := s.apiRequest("POST", s.url + "rename/", &data)
    if err != nil {
        return err
    }

    if resp.StatusCode != 200 {
        return newApiError(resp, body)
    }

    return nil
}

