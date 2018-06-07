package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type MergeFacesRequest struct {
	LibraryId    string                `position:"Query" name:"LibraryId"`
	TargetFaceId int64                 `position:"Query" name:"TargetFaceId"`
	StoreName    string                `position:"Query" name:"StoreName"`
	FaceIds      *MergeFacesFaceIdList `position:"Query" type:"Repeated" name:"FaceId"`
	RegionId     string                `position:"Query" name:"RegionId"`
}

func (req *MergeFacesRequest) Invoke(client goaliyun.Client) (*MergeFacesResponse, error) {
	resp := &MergeFacesResponse{}
	err := client.Request("cloudphoto", "MergeFaces", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MergeFacesResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   MergeFacesResultList
}

type MergeFacesResult struct {
	Id      goaliyun.Integer
	IdStr   goaliyun.String
	Code    goaliyun.String
	Message goaliyun.String
}

type MergeFacesFaceIdList []int64

func (list *MergeFacesFaceIdList) UnmarshalJSON(data []byte) error {
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

type MergeFacesResultList []MergeFacesResult

func (list *MergeFacesResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MergeFacesResult)
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
