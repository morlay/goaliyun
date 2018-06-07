package imm

import (
	"github.com/morlay/goaliyun"
)

type ListFaceGroupPhotosRequest struct {
	MaxKeys  int64  `position:"Query" name:"MaxKeys"`
	Marker   string `position:"Query" name:"Marker"`
	GroupId  string `position:"Query" name:"GroupId"`
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListFaceGroupPhotosRequest) Invoke(client goaliyun.Client) (*ListFaceGroupPhotosResponse, error) {
	resp := &ListFaceGroupPhotosResponse{}
	err := client.Request("imm", "ListFaceGroupPhotos", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListFaceGroupPhotosResponse struct {
	RequestId  goaliyun.String
	Photos     goaliyun.String
	NextMarker goaliyun.String
}
