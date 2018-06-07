package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DeleteFacesRequest struct {
	LibraryId string                 `position:"Query" name:"LibraryId"`
	StoreName string                 `position:"Query" name:"StoreName"`
	FaceIds   *DeleteFacesFaceIdList `position:"Query" type:"Repeated" name:"FaceId"`
	RegionId  string                 `position:"Query" name:"RegionId"`
}

func (req *DeleteFacesRequest) Invoke(client goaliyun.Client) (*DeleteFacesResponse, error) {
	resp := &DeleteFacesResponse{}
	err := client.Request("cloudphoto", "DeleteFaces", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteFacesResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   DeleteFacesResultList
}

type DeleteFacesResult struct {
	Id      goaliyun.Integer
	IdStr   goaliyun.String
	Code    goaliyun.String
	Message goaliyun.String
}

type DeleteFacesFaceIdList []int64

func (list *DeleteFacesFaceIdList) UnmarshalJSON(data []byte) error {
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

type DeleteFacesResultList []DeleteFacesResult

func (list *DeleteFacesResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteFacesResult)
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
