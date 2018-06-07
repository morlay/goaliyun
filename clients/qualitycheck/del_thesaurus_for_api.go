package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type DelThesaurusForApiRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DelThesaurusForApiRequest) Invoke(client goaliyun.Client) (*DelThesaurusForApiResponse, error) {
	resp := &DelThesaurusForApiResponse{}
	err := client.Request("qualitycheck", "DelThesaurusForApi", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DelThesaurusForApiResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
}
