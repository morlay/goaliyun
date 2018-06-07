package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetFramedPhotoUrlsRequest struct {
	FrameId   string                         `position:"Query" name:"FrameId"`
	LibraryId string                         `position:"Query" name:"LibraryId"`
	PhotoIds  *GetFramedPhotoUrlsPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                         `position:"Query" name:"StoreName"`
	RegionId  string                         `position:"Query" name:"RegionId"`
}

func (req *GetFramedPhotoUrlsRequest) Invoke(client goaliyun.Client) (*GetFramedPhotoUrlsResponse, error) {
	resp := &GetFramedPhotoUrlsResponse{}
	err := client.Request("cloudphoto", "GetFramedPhotoUrls", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetFramedPhotoUrlsResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   GetFramedPhotoUrlsResultList
}

type GetFramedPhotoUrlsResult struct {
	Code           goaliyun.String
	Message        goaliyun.String
	PhotoId        goaliyun.Integer
	PhotoIdStr     goaliyun.String
	FramedPhotoUrl goaliyun.String
}

type GetFramedPhotoUrlsPhotoIdList []int64

func (list *GetFramedPhotoUrlsPhotoIdList) UnmarshalJSON(data []byte) error {
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

type GetFramedPhotoUrlsResultList []GetFramedPhotoUrlsResult

func (list *GetFramedPhotoUrlsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetFramedPhotoUrlsResult)
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
