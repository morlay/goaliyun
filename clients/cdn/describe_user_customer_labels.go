package cdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeUserCustomerLabelsRequest struct {
	Uid           int64  `position:"Query" name:"Uid"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeUserCustomerLabelsRequest) Invoke(client goaliyun.Client) (*DescribeUserCustomerLabelsResponse, error) {
	resp := &DescribeUserCustomerLabelsResponse{}
	err := client.Request("cdn", "DescribeUserCustomerLabels", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeUserCustomerLabelsResponse struct {
	RequestId   goaliyun.String
	IsInnerUser bool
}
