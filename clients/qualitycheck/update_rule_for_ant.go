package qualitycheck

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type UpdateRuleForAntRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UpdateRuleForAntRequest) Invoke(client goaliyun.Client) (*UpdateRuleForAntResponse, error) {
	resp := &UpdateRuleForAntResponse{}
	err := client.Request("qualitycheck", "UpdateRuleForAnt", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateRuleForAntResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      UpdateRuleForAntDatumList
}

type UpdateRuleForAntDatumList []goaliyun.String

func (list *UpdateRuleForAntDatumList) UnmarshalJSON(data []byte) error {
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
