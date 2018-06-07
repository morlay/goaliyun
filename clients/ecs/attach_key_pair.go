package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AttachKeyPairRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	KeyPairName          string `position:"Query" name:"KeyPairName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AttachKeyPairRequest) Invoke(client goaliyun.Client) (*AttachKeyPairResponse, error) {
	resp := &AttachKeyPairResponse{}
	err := client.Request("ecs", "AttachKeyPair", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AttachKeyPairResponse struct {
	RequestId   goaliyun.String
	TotalCount  goaliyun.String
	FailCount   goaliyun.String
	KeyPairName goaliyun.String
	Results     AttachKeyPairResultList
}

type AttachKeyPairResult struct {
	InstanceId goaliyun.String
	Success    goaliyun.String
	Code       goaliyun.String
	Message    goaliyun.String
}

type AttachKeyPairResultList []AttachKeyPairResult

func (list *AttachKeyPairResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AttachKeyPairResult)
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
