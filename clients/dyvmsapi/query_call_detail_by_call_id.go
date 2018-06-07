package dyvmsapi

import (
	"github.com/morlay/goaliyun"
)

type QueryCallDetailByCallIdRequest struct {
	CallId               string `position:"Query" name:"CallId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	QueryDate            int64  `position:"Query" name:"QueryDate"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ProdId               int64  `position:"Query" name:"ProdId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryCallDetailByCallIdRequest) Invoke(client goaliyun.Client) (*QueryCallDetailByCallIdResponse, error) {
	resp := &QueryCallDetailByCallIdResponse{}
	err := client.Request("dyvmsapi", "QueryCallDetailByCallId", "2017-05-25", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryCallDetailByCallIdResponse struct {
	RequestId goaliyun.String
	Data      goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
}
