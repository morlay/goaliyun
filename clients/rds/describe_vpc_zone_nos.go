package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeVpcZoneNosRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Region               string `position:"Query" name:"Region"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeVpcZoneNosRequest) Invoke(client goaliyun.Client) (*DescribeVpcZoneNosResponse, error) {
	resp := &DescribeVpcZoneNosResponse{}
	err := client.Request("rds", "DescribeVpcZoneNos", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVpcZoneNosResponse struct {
	RequestId goaliyun.String
	Items     DescribeVpcZoneNosVpcZoneIdList
}

type DescribeVpcZoneNosVpcZoneId struct {
	ZoneId    goaliyun.String
	Region    goaliyun.String
	SubDomain goaliyun.String
}

type DescribeVpcZoneNosVpcZoneIdList []DescribeVpcZoneNosVpcZoneId

func (list *DescribeVpcZoneNosVpcZoneIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpcZoneNosVpcZoneId)
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
