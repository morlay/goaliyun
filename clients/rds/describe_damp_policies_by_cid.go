package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDampPoliciesByCidRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDampPoliciesByCidRequest) Invoke(client goaliyun.Client) (*DescribeDampPoliciesByCidResponse, error) {
	resp := &DescribeDampPoliciesByCidResponse{}
	err := client.Request("rds", "DescribeDampPoliciesByCid", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDampPoliciesByCidResponse struct {
	RequestId goaliyun.String
	Policies  DescribeDampPoliciesByCidPolicyList
}

type DescribeDampPoliciesByCidPolicy struct {
	PolicyName goaliyun.String
	Comment    goaliyun.String
}

type DescribeDampPoliciesByCidPolicyList []DescribeDampPoliciesByCidPolicy

func (list *DescribeDampPoliciesByCidPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDampPoliciesByCidPolicy)
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
