package emr

import (
	"github.com/morlay/goaliyun"
)

type ResetSgPortRequest struct {
	SourceIp    string `position:"Query" name:"SourceIp"`
	ClusterId   string `position:"Query" name:"ClusterId"`
	OperateType string `position:"Query" name:"OperateType"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *ResetSgPortRequest) Invoke(client goaliyun.Client) (*ResetSgPortResponse, error) {
	resp := &ResetSgPortResponse{}
	err := client.Request("emr", "ResetSgPort", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResetSgPortResponse struct {
	RequestId goaliyun.String
	Success   goaliyun.String
	ErrCode   goaliyun.String
	ErrMsg    goaliyun.String
}
