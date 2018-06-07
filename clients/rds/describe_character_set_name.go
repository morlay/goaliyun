package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCharacterSetNameRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Engine               string `position:"Query" name:"Engine"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeCharacterSetNameRequest) Invoke(client goaliyun.Client) (*DescribeCharacterSetNameResponse, error) {
	resp := &DescribeCharacterSetNameResponse{}
	err := client.Request("rds", "DescribeCharacterSetName", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCharacterSetNameResponse struct {
	RequestId             goaliyun.String
	Engine                goaliyun.String
	CharacterSetNameItems DescribeCharacterSetNameCharacterSetNameItemList
}

type DescribeCharacterSetNameCharacterSetNameItemList []goaliyun.String

func (list *DescribeCharacterSetNameCharacterSetNameItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
