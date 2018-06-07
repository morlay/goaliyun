package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteLiveMixNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteLiveMixNotifyConfigRequest) Invoke(client goaliyun.Client) (*DeleteLiveMixNotifyConfigResponse, error) {
	resp := &DeleteLiveMixNotifyConfigResponse{}
	err := client.Request("live", "DeleteLiveMixNotifyConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLiveMixNotifyConfigResponse struct {
	RequestId goaliyun.String
}
