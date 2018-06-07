package drds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRdsListRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeRdsListRequest) Invoke(client goaliyun.Client) (*DescribeRdsListResponse, error) {
	resp := &DescribeRdsListResponse{}
	err := client.Request("drds", "DescribeRdsList", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRdsListResponse struct {
	RequestId goaliyun.String
	Success   bool
	Data      DescribeRdsListRdsInstanceList
}

type DescribeRdsListRdsInstance struct {
	InstanceId       goaliyun.Integer
	InstanceName     goaliyun.String
	ConnectUrl       goaliyun.String
	Port             goaliyun.Integer
	InstanceStatus   goaliyun.String
	DbType           goaliyun.String
	ReadWeight       goaliyun.Integer
	ReadOnlyChildren DescribeRdsListChildList
}

type DescribeRdsListChild struct {
	InstanceId     goaliyun.String
	InstanceName   goaliyun.String
	ConnectUrl     goaliyun.String
	Port           goaliyun.Integer
	InstanceStatus goaliyun.String
	DbType         goaliyun.String
	ReadWeight     goaliyun.Integer
}

type DescribeRdsListRdsInstanceList []DescribeRdsListRdsInstance

func (list *DescribeRdsListRdsInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRdsListRdsInstance)
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

type DescribeRdsListChildList []DescribeRdsListChild

func (list *DescribeRdsListChildList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRdsListChild)
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
