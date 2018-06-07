package polardb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeParameterTemplatesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeParameterTemplatesRequest) Invoke(client goaliyun.Client) (*DescribeParameterTemplatesResponse, error) {
	resp := &DescribeParameterTemplatesResponse{}
	err := client.Request("polardb", "DescribeParameterTemplates", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeParameterTemplatesResponse struct {
	RequestId      goaliyun.String
	Engine         goaliyun.String
	DBType         goaliyun.String
	DBVersion      goaliyun.String
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
