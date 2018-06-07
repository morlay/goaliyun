package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetThumbnailsRequest struct {
	LibraryId string                    `position:"Query" name:"LibraryId"`
	PhotoIds  *GetThumbnailsPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                    `position:"Query" name:"StoreName"`
	ZoomType  string                    `position:"Query" name:"ZoomType"`
	RegionId  string                    `position:"Query" name:"RegionId"`
}

func (req *GetThumbnailsRequest) Invoke(client goaliyun.Client) (*GetThumbnailsResponse, error) {
	resp := &GetThumbnailsResponse{}
	err := client.Request("cloudphoto", "GetThumbnails", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetThumbnailsResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   GetThumbnailsResultList
}

type GetThumbnailsResult struct {
	Code         goaliyun.String
	Message      goaliyun.String
	PhotoId      goaliyun.Integer
	PhotoIdStr   goaliyun.String
	ThumbnailUrl goaliyun.String
}

type GetThumbnailsPhotoIdList []int64

func (list *GetThumbnailsPhotoIdList) UnmarshalJSON(data []byte) error {
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

type GetThumbnailsResultList []GetThumbnailsResult

func (list *GetThumbnailsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetThumbnailsResult)
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
