package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetRemoveQueryStringConfigRequest struct {
	KeepOssArgs   string `position:"Query" name:"KeepOssArgs"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	AliRemoveArgs string `position:"Query" name:"AliRemoveArgs"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetRemoveQueryStringConfigRequest) Invoke(client goaliyun.Client) (*SetRemoveQueryStringConfigResponse, error) {
	resp := &SetRemoveQueryStringConfigResponse{}
	err := client.Request("cdn", "SetRemoveQueryStringConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetRemoveQueryStringConfigResponse struct {
	RequestId goaliyun.String
}
