package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetPublicAccessUrlsRequest struct {
	DomainType string                          `position:"Query" name:"DomainType"`
	LibraryId  string                          `position:"Query" name:"LibraryId"`
	PhotoIds   *GetPublicAccessUrlsPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName  string                          `position:"Query" name:"StoreName"`
	ZoomType   string                          `position:"Query" name:"ZoomType"`
	RegionId   string                          `position:"Query" name:"RegionId"`
}

func (req *GetPublicAccessUrlsRequest) Invoke(client goaliyun.Client) (*GetPublicAccessUrlsResponse, error) {
	resp := &GetPublicAccessUrlsResponse{}
	err := client.Request("cloudphoto", "GetPublicAccessUrls", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetPublicAccessUrlsResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   GetPublicAccessUrlsResultList
}

type GetPublicAccessUrlsResult struct {
	Code       goaliyun.String
	Message    goaliyun.String
	PhotoId    goaliyun.Integer
	PhotoIdStr goaliyun.String
	AccessUrl  goaliyun.String
}

type GetPublicAccessUrlsPhotoIdList []int64

func (list *GetPublicAccessUrlsPhotoIdList) UnmarshalJSON(data []byte) error {
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

type GetPublicAccessUrlsResultList []GetPublicAccessUrlsResult

func (list *GetPublicAccessUrlsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPublicAccessUrlsResult)
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
