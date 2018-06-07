package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetIgnoreQueryStringConfigRequest struct {
	KeepOssArgs   string `position:"Query" name:"KeepOssArgs"`
	HashKeyArgs   string `position:"Query" name:"HashKeyArgs"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetIgnoreQueryStringConfigRequest) Invoke(client goaliyun.Client) (*SetIgnoreQueryStringConfigResponse, error) {
	resp := &SetIgnoreQueryStringConfigResponse{}
	err := client.Request("cdn", "SetIgnoreQueryStringConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetIgnoreQueryStringConfigResponse struct {
	RequestId goaliyun.String
}
