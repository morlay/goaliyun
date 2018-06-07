package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyRouteTableAttributesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	KbpsBandwidth        string `position:"Query" name:"KbpsBandwidth"`
	RouteTableName       string `position:"Query" name:"RouteTableName"`
	ResourceUid          int64  `position:"Query" name:"ResourceUid"`
	ResourceBid          string `position:"Query" name:"ResourceBid"`
	RouteTableId         string `position:"Query" name:"RouteTableId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyRouteTableAttributesRequest) Invoke(client goaliyun.Client) (*ModifyRouteTableAttributesResponse, error) {
	resp := &ModifyRouteTableAttributesResponse{}
	err := client.Request("vpc", "ModifyRouteTableAttributes", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyRouteTableAttributesResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
	Success   bool
}
