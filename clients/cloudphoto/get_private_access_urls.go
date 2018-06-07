package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetPrivateAccessUrlsRequest struct {
	LibraryId string                           `position:"Query" name:"LibraryId"`
	PhotoIds  *GetPrivateAccessUrlsPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                           `position:"Query" name:"StoreName"`
	ZoomType  string                           `position:"Query" name:"ZoomType"`
	RegionId  string                           `position:"Query" name:"RegionId"`
}

func (req *GetPrivateAccessUrlsRequest) Invoke(client goaliyun.Client) (*GetPrivateAccessUrlsResponse, error) {
	resp := &GetPrivateAccessUrlsResponse{}
	err := client.Request("cloudphoto", "GetPrivateAccessUrls", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetPrivateAccessUrlsResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   GetPrivateAccessUrlsResultList
}

type GetPrivateAccessUrlsResult struct {
	Code       goaliyun.String
	Message    goaliyun.String
	PhotoId    goaliyun.Integer
	PhotoIdStr goaliyun.String
	AccessUrl  goaliyun.String
}

type GetPrivateAccessUrlsPhotoIdList []int64

func (list *GetPrivateAccessUrlsPhotoIdList) UnmarshalJSON(data []byte) error {
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

type GetPrivateAccessUrlsResultList []GetPrivateAccessUrlsResult

func (list *GetPrivateAccessUrlsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPrivateAccessUrlsResult)
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
