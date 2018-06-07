package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type UploadDataRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UploadDataRequest) Invoke(client goaliyun.Client) (*UploadDataResponse, error) {
	resp := &UploadDataResponse{}
	err := client.Request("qualitycheck", "UploadData", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UploadDataResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
}
