package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetErrorPageConfigRequest struct {
	PageType      string `position:"Query" name:"PageType"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	CustomPageUrl string `position:"Query" name:"CustomPageUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetErrorPageConfigRequest) Invoke(client goaliyun.Client) (*SetErrorPageConfigResponse, error) {
	resp := &SetErrorPageConfigResponse{}
	err := client.Request("cdn", "SetErrorPageConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetErrorPageConfigResponse struct {
	RequestId goaliyun.String
}
