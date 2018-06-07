package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type RegisterNoticeRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *RegisterNoticeRequest) Invoke(client goaliyun.Client) (*RegisterNoticeResponse, error) {
	resp := &RegisterNoticeResponse{}
	err := client.Request("qualitycheck", "RegisterNotice", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RegisterNoticeResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
}
