package emr

import (
	"github.com/morlay/goaliyun"
)

type ResetSoftwarePasswordRequest struct {
	Password  string `position:"Query" name:"Password"`
	ClusterId string `position:"Query" name:"ClusterId"`
	Username  string `position:"Query" name:"Username"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ResetSoftwarePasswordRequest) Invoke(client goaliyun.Client) (*ResetSoftwarePasswordResponse, error) {
	resp := &ResetSoftwarePasswordResponse{}
	err := client.Request("emr", "ResetSoftwarePassword", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResetSoftwarePasswordResponse struct {
	RequestId goaliyun.String
	Success   goaliyun.String
	ErrCode   goaliyun.String
	ErrMsg    goaliyun.String
}
