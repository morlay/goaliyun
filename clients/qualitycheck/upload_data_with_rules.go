package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type UploadDataWithRulesRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UploadDataWithRulesRequest) Invoke(client goaliyun.Client) (*UploadDataWithRulesResponse, error) {
	resp := &UploadDataWithRulesResponse{}
	err := client.Request("qualitycheck", "UploadDataWithRules", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UploadDataWithRulesResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
}
