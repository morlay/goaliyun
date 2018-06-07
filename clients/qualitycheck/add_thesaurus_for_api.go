package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type AddThesaurusForApiRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *AddThesaurusForApiRequest) Invoke(client goaliyun.Client) (*AddThesaurusForApiResponse, error) {
	resp := &AddThesaurusForApiResponse{}
	err := client.Request("qualitycheck", "AddThesaurusForApi", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddThesaurusForApiResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.Integer
}
