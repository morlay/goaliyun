package cdn

import (
	"github.com/morlay/goaliyun"
)

type ModifyUserCustomLogConfigRequest struct {
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigId      string `position:"Query" name:"ConfigId"`
	Tag           string `position:"Query" name:"Tag"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *ModifyUserCustomLogConfigRequest) Invoke(client goaliyun.Client) (*ModifyUserCustomLogConfigResponse, error) {
	resp := &ModifyUserCustomLogConfigResponse{}
	err := client.Request("cdn", "ModifyUserCustomLogConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyUserCustomLogConfigResponse struct {
	RequestId goaliyun.String
}
