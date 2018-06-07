package push

import (
	"github.com/morlay/goaliyun"
)

type UnbindAliasRequest struct {
	AliasName string `position:"Query" name:"AliasName"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	DeviceId  string `position:"Query" name:"DeviceId"`
	UnbindAll string `position:"Query" name:"UnbindAll"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *UnbindAliasRequest) Invoke(client goaliyun.Client) (*UnbindAliasResponse, error) {
	resp := &UnbindAliasResponse{}
	err := client.Request("push", "UnbindAlias", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UnbindAliasResponse struct {
	RequestId goaliyun.String
}
