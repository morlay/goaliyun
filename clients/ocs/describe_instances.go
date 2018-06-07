package ocs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstancesRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	OcsInstanceStatus    string `position:"Query" name:"OcsInstanceStatus"`
	PageNo               int64  `position:"Query" name:"PageNo"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	VpcId                string `position:"Query" name:"VpcId"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	PrivateIps           string `position:"Query" name:"PrivateIps"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstancesRequest) Invoke(client goaliyun.Client) (*DescribeInstancesResponse, error) {
	resp := &DescribeInstancesResponse{}
	err := client.Request("ocs", "DescribeInstances", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstancesResponse struct {
	RequestId               goaliyun.String
	GetOcsInstancesResponse DescribeInstancesGetOcsInstancesResponse
}

type DescribeInstancesGetOcsInstancesResponse struct {
	Total        goaliyun.Integer
	PageNo       goaliyun.Integer
	PageSize     goaliyun.Integer
	OcsInstances DescribeInstancesOcsInstanceList
}

type DescribeInstancesOcsInstance struct {
	OcsInstanceId     goaliyun.String
	OcsInstanceName   goaliyun.String
	Capacity          goaliyun.Integer
	Qps               goaliyun.Integer
	Bandwidth         goaliyun.Integer
	Connections       goaliyun.Integer
	ConnectionDomain  goaliyun.String
	Port              goaliyun.Integer
	UserName          goaliyun.String
	RegionId          goaliyun.String
	OcsInstanceStatus goaliyun.String
	GmtCreated        goaliyun.String
	EndTime           goaliyun.String
	ChargeType        goaliyun.String
	IzId              goaliyun.String
	NetworkType       goaliyun.String
	VpcId             goaliyun.String
	VSwitchId         goaliyun.String
	PrivateIp         goaliyun.String
}

type DescribeInstancesOcsInstanceList []DescribeInstancesOcsInstance

func (list *DescribeInstancesOcsInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesOcsInstance)
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
