package cdn

import (
	"github.com/morlay/goaliyun"
)

type DeleteHttpHeaderConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigID      string `position:"Query" name:"ConfigID"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteHttpHeaderConfigRequest) Invoke(client goaliyun.Client) (*DeleteHttpHeaderConfigResponse, error) {
	resp := &DeleteHttpHeaderConfigResponse{}
	err := client.Request("cdn", "DeleteHttpHeaderConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteHttpHeaderConfigResponse struct {
	RequestId goaliyun.String
}
