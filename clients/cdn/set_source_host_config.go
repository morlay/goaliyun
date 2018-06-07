package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetSourceHostConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	BackSrcDomain string `position:"Query" name:"BackSrcDomain"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetSourceHostConfigRequest) Invoke(client goaliyun.Client) (*SetSourceHostConfigResponse, error) {
	resp := &SetSourceHostConfigResponse{}
	err := client.Request("cdn", "SetSourceHostConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetSourceHostConfigResponse struct {
	RequestId goaliyun.String
}
