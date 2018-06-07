package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetHttpErrorPageConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageUrl       string `position:"Query" name:"PageUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	ErrorCode     string `position:"Query" name:"ErrorCode"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetHttpErrorPageConfigRequest) Invoke(client goaliyun.Client) (*SetHttpErrorPageConfigResponse, error) {
	resp := &SetHttpErrorPageConfigResponse{}
	err := client.Request("cdn", "SetHttpErrorPageConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetHttpErrorPageConfigResponse struct {
	RequestId goaliyun.String
}
