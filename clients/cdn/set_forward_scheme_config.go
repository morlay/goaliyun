package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetForwardSchemeConfigRequest struct {
	SchemeOrigin     string `position:"Query" name:"SchemeOrigin"`
	SchemeOriginPort string `position:"Query" name:"SchemeOriginPort"`
	SecurityToken    string `position:"Query" name:"SecurityToken"`
	Enable           string `position:"Query" name:"Enable"`
	DomainName       string `position:"Query" name:"DomainName"`
	OwnerId          int64  `position:"Query" name:"OwnerId"`
	RegionId         string `position:"Query" name:"RegionId"`
}

func (req *SetForwardSchemeConfigRequest) Invoke(client goaliyun.Client) (*SetForwardSchemeConfigResponse, error) {
	resp := &SetForwardSchemeConfigResponse{}
	err := client.Request("cdn", "SetForwardSchemeConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetForwardSchemeConfigResponse struct {
	RequestId goaliyun.String
}
