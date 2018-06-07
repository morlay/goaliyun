package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeKeyPairsRequest struct {
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Tag2Key              string `position:"Query" name:"Tag.2.Key"`
	Tag5Key              string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	KeyPairFingerPrint   string `position:"Query" name:"KeyPairFingerPrint"`
	Tag3Key              string `position:"Query" name:"Tag.3.Key"`
	KeyPairName          string `position:"Query" name:"KeyPairName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tag5Value            string `position:"Query" name:"Tag.5.Value"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	Tag1Key              string `position:"Query" name:"Tag.1.Key"`
	Tag1Value            string `position:"Query" name:"Tag.1.Value"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	Tag2Value            string `position:"Query" name:"Tag.2.Value"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	Tag4Key              string `position:"Query" name:"Tag.4.Key"`
	Tag3Value            string `position:"Query" name:"Tag.3.Value"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeKeyPairsRequest) Invoke(client goaliyun.Client) (*DescribeKeyPairsResponse, error) {
	resp := &DescribeKeyPairsResponse{}
	err := client.Request("ecs", "DescribeKeyPairs", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeKeyPairsResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	KeyPairs   DescribeKeyPairsKeyPairList
}

type DescribeKeyPairsKeyPair struct {
	KeyPairName        goaliyun.String
	KeyPairFingerPrint goaliyun.String
}

type DescribeKeyPairsKeyPairList []DescribeKeyPairsKeyPair

func (list *DescribeKeyPairsKeyPairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeKeyPairsKeyPair)
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
