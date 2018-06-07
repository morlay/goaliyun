package imm

import (
	"github.com/morlay/goaliyun"
)

type GetFaceGroupPhotosRequest struct {
	MaxKeys  int64  `position:"Query" name:"MaxKeys"`
	Marker   string `position:"Query" name:"Marker"`
	GroupId  string `position:"Query" name:"GroupId"`
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetFaceGroupPhotosRequest) Invoke(client goaliyun.Client) (*GetFaceGroupPhotosResponse, error) {
	resp := &GetFaceGroupPhotosResponse{}
	err := client.Request("imm", "GetFaceGroupPhotos", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetFaceGroupPhotosResponse struct {
	RequestId  goaliyun.String
	Photos     goaliyun.String
	NextMarker goaliyun.String
}
