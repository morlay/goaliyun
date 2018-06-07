package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetVideoSeekConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetVideoSeekConfigRequest) Invoke(client goaliyun.Client) (*SetVideoSeekConfigResponse, error) {
	resp := &SetVideoSeekConfigResponse{}
	err := client.Request("cdn", "SetVideoSeekConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetVideoSeekConfigResponse struct {
	RequestId goaliyun.String
}
