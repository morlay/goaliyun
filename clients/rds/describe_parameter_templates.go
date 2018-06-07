package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeParameterTemplatesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Engine               string `position:"Query" name:"Engine"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EngineVersion        string `position:"Query" name:"EngineVersion"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeParameterTemplatesRequest) Invoke(client goaliyun.Client) (*DescribeParameterTemplatesResponse, error) {
	resp := &DescribeParameterTemplatesResponse{}
	err := client.Request("rds", "DescribeParameterTemplates", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeParameterTemplatesResponse struct {
	RequestId      goaliyun.String
	Engine         goaliyun.String
	EngineVersion  goaliyun.String
	ParameterCount goaliyun.String
	Parameters     DescribeParameterTemplatesTemplateRecordList
}

type DescribeParameterTemplatesTemplateRecord struct {
	ParameterName        goaliyun.String
	ParameterValue       goaliyun.String
	ForceModify          goaliyun.String
	ForceRestart         goaliyun.String
	CheckingCode         goaliyun.String
	ParameterDescription goaliyun.String
}

type DescribeParameterTemplatesTemplateRecordList []DescribeParameterTemplatesTemplateRecord

func (list *DescribeParameterTemplatesTemplateRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeParameterTemplatesTemplateRecord)
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
