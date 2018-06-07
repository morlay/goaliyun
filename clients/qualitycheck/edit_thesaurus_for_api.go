package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type EditThesaurusForApiRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *EditThesaurusForApiRequest) Invoke(client goaliyun.Client) (*EditThesaurusForApiResponse, error) {
	resp := &EditThesaurusForApiResponse{}
	err := client.Request("qualitycheck", "EditThesaurusForApi", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EditThesaurusForApiResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.Integer
}
