package csb

import (
	"github.com/morlay/goaliyun"
)

type UpdateServiceQPSRequest struct {
	Qps       string `position:"Query" name:"Qps"`
	ServiceId int64  `position:"Query" name:"ServiceId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *UpdateServiceQPSRequest) Invoke(client goaliyun.Client) (*UpdateServiceQPSResponse, error) {
	resp := &UpdateServiceQPSResponse{}
	err := client.Request("csb", "UpdateServiceQPS", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateServiceQPSResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}
