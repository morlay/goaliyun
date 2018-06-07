package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DetachKeyPairRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	KeyPairName          string `position:"Query" name:"KeyPairName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DetachKeyPairRequest) Invoke(client goaliyun.Client) (*DetachKeyPairResponse, error) {
	resp := &DetachKeyPairResponse{}
	err := client.Request("ecs", "DetachKeyPair", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DetachKeyPairResponse struct {
	RequestId   goaliyun.String
	TotalCount  goaliyun.String
	FailCount   goaliyun.String
	KeyPairName goaliyun.String
	Results     DetachKeyPairResultList
}

type DetachKeyPairResult struct {
	InstanceId goaliyun.String
	Success    goaliyun.String
	Code       goaliyun.String
	Message    goaliyun.String
}

type DetachKeyPairResultList []DetachKeyPairResult

func (list *DetachKeyPairResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetachKeyPairResult)
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
