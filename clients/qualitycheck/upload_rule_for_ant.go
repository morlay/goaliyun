package qualitycheck

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type UploadRuleForAntRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UploadRuleForAntRequest) Invoke(client goaliyun.Client) (*UploadRuleForAntResponse, error) {
	resp := &UploadRuleForAntResponse{}
	err := client.Request("qualitycheck", "UploadRuleForAnt", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UploadRuleForAntResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      UploadRuleForAntDatumList
}

type UploadRuleForAntDatumList []goaliyun.String

func (list *UploadRuleForAntDatumList) UnmarshalJSON(data []byte) error {
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
