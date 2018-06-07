package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCommandsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Description          string `position:"Query" name:"Description"`
	Type                 string `position:"Query" name:"Type"`
	CommandId            string `position:"Query" name:"CommandId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Name                 string `position:"Query" name:"Name"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeCommandsRequest) Invoke(client goaliyun.Client) (*DescribeCommandsResponse, error) {
	resp := &DescribeCommandsResponse{}
	err := client.Request("ecs", "DescribeCommands", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCommandsResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Commands   DescribeCommandsCommandList
}

type DescribeCommandsCommand struct {
	CommandId      goaliyun.String
	Name           goaliyun.String
	Type           goaliyun.String
	Description    goaliyun.String
	CommandContent goaliyun.String
	WorkingDir     goaliyun.String
	Timeout        goaliyun.Integer
}

type DescribeCommandsCommandList []DescribeCommandsCommand

func (list *DescribeCommandsCommandList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCommandsCommand)
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
