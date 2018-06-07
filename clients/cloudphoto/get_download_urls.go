package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetDownloadUrlsRequest struct {
	LibraryId string                      `position:"Query" name:"LibraryId"`
	PhotoIds  *GetDownloadUrlsPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                      `position:"Query" name:"StoreName"`
	RegionId  string                      `position:"Query" name:"RegionId"`
}

func (req *GetDownloadUrlsRequest) Invoke(client goaliyun.Client) (*GetDownloadUrlsResponse, error) {
	resp := &GetDownloadUrlsResponse{}
	err := client.Request("cloudphoto", "GetDownloadUrls", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetDownloadUrlsResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   GetDownloadUrlsResultList
}

type GetDownloadUrlsResult struct {
	Code        goaliyun.String
	Message     goaliyun.String
	PhotoId     goaliyun.Integer
	PhotoIdStr  goaliyun.String
	DownloadUrl goaliyun.String
}

type GetDownloadUrlsPhotoIdList []int64

func (list *GetDownloadUrlsPhotoIdList) UnmarshalJSON(data []byte) error {
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

type GetDownloadUrlsResultList []GetDownloadUrlsResult

func (list *GetDownloadUrlsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetDownloadUrlsResult)
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
