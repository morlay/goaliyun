package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetOssLogConfigRequest struct {
	Bucket        string `position:"Query" name:"Bucket"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	Prefix        string `position:"Query" name:"Prefix"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetOssLogConfigRequest) Invoke(client goaliyun.Client) (*SetOssLogConfigResponse, error) {
	resp := &SetOssLogConfigResponse{}
	err := client.Request("cdn", "SetOssLogConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetOssLogConfigResponse struct {
	RequestId goaliyun.String
}
