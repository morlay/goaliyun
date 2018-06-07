package push

import (
	"github.com/morlay/goaliyun"
)

type BindAliasRequest struct {
	AliasName string `position:"Query" name:"AliasName"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	DeviceId  string `position:"Query" name:"DeviceId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *BindAliasRequest) Invoke(client goaliyun.Client) (*BindAliasResponse, error) {
	resp := &BindAliasResponse{}
	err := client.Request("push", "BindAlias", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type BindAliasResponse struct {
	RequestId goaliyun.String
}
