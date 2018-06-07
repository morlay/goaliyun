package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeNqasRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RouterId             string `position:"Query" name:"RouterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	NqaId                string `position:"Query" name:"NqaId"`
	IsDefault            string `position:"Query" name:"IsDefault"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeNqasRequest) Invoke(client goaliyun.Client) (*DescribeNqasResponse, error) {
	resp := &DescribeNqasResponse{}
	err := client.Request("vpc", "DescribeNqas", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeNqasResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Nqas       DescribeNqasNqaList
}

type DescribeNqasNqa struct {
	NqaId         goaliyun.String
	RegionId      goaliyun.String
	Status        goaliyun.String
	RouterId      goaliyun.String
	DestinationIp goaliyun.String
}

type DescribeNqasNqaList []DescribeNqasNqa

func (list *DescribeNqasNqaList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNqasNqa)
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
