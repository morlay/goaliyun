package qualitycheck

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type UploadRuleRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UploadRuleRequest) Invoke(client goaliyun.Client) (*UploadRuleResponse, error) {
	resp := &UploadRuleResponse{}
	err := client.Request("qualitycheck", "UploadRule", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UploadRuleResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      UploadRuleDatumList
}

type UploadRuleDatumList []goaliyun.String

func (list *UploadRuleDatumList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
