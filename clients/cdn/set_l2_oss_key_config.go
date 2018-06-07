package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetL2OssKeyConfigRequest struct {
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	DomainName     string `position:"Query" name:"DomainName"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	PrivateOssAuth string `position:"Query" name:"PrivateOssAuth"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *SetL2OssKeyConfigRequest) Invoke(client goaliyun.Client) (*SetL2OssKeyConfigResponse, error) {
	resp := &SetL2OssKeyConfigResponse{}
	err := client.Request("cdn", "SetL2OssKeyConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetL2OssKeyConfigResponse struct {
	RequestId goaliyun.String
}
