package afs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribePersonMachineListRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribePersonMachineListRequest) Invoke(client goaliyun.Client) (*DescribePersonMachineListResponse, error) {
	resp := &DescribePersonMachineListResponse{}
	err := client.Request("afs", "DescribePersonMachineList", "2018-01-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribePersonMachineListResponse struct {
	RequestId        goaliyun.String
	BizCode          goaliyun.String
	PersonMachineRes DescribePersonMachineListPersonMachineRes
}

type DescribePersonMachineListPersonMachineRes struct {
	HasConfiguration goaliyun.String
	PersonMachines   DescribePersonMachineListPersonMachineList
}

type DescribePersonMachineListPersonMachine struct {
	ConfigurationName   goaliyun.String
	Appkey              goaliyun.String
	ConfigurationMethod goaliyun.String
	ApplyType           goaliyun.String
	Scene               goaliyun.String
	LastUpdate          goaliyun.String
}

type DescribePersonMachineListPersonMachineList []DescribePersonMachineListPersonMachine

func (list *DescribePersonMachineListPersonMachineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePersonMachineListPersonMachine)
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
