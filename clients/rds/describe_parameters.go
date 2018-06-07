package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeParametersRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeParametersRequest) Invoke(client goaliyun.Client) (*DescribeParametersResponse, error) {
	resp := &DescribeParametersResponse{}
	err := client.Request("rds", "DescribeParameters", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeParametersResponse struct {
	RequestId         goaliyun.String
	Engine            goaliyun.String
	EngineVersion     goaliyun.String
	ConfigParameters  DescribeParametersDBInstanceParameterList
	RunningParameters DescribeParametersDBInstanceParameterList
}

type DescribeParametersDBInstanceParameter struct {
	ParameterName        goaliyun.String
	ParameterValue       goaliyun.String
	ParameterDescription goaliyun.String
}

type DescribeParametersDBInstanceParameterList []DescribeParametersDBInstanceParameter

func (list *DescribeParametersDBInstanceParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeParametersDBInstanceParameter)
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
