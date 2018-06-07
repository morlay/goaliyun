package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DeleteAlbumsRequest struct {
	LibraryId string                   `position:"Query" name:"LibraryId"`
	AlbumIds  *DeleteAlbumsAlbumIdList `position:"Query" type:"Repeated" name:"AlbumId"`
	StoreName string                   `position:"Query" name:"StoreName"`
	RegionId  string                   `position:"Query" name:"RegionId"`
}

func (req *DeleteAlbumsRequest) Invoke(client goaliyun.Client) (*DeleteAlbumsResponse, error) {
	resp := &DeleteAlbumsResponse{}
	err := client.Request("cloudphoto", "DeleteAlbums", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteAlbumsResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   DeleteAlbumsResultList
}

type DeleteAlbumsResult struct {
	Id      goaliyun.Integer
	IdStr   goaliyun.String
	Code    goaliyun.String
	Message goaliyun.String
}

type DeleteAlbumsAlbumIdList []int64

func (list *DeleteAlbumsAlbumIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type DeleteAlbumsResultList []DeleteAlbumsResult

func (list *DeleteAlbumsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteAlbumsResult)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
