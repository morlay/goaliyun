package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteLiveDetectNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteLiveDetectNotifyConfigRequest) Invoke(client goaliyun.Client) (*DeleteLiveDetectNotifyConfigResponse, error) {
	resp := &DeleteLiveDetectNotifyConfigResponse{}
	err := client.Request("live", "DeleteLiveDetectNotifyConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLiveDetectNotifyConfigResponse struct {
	RequestId goaliyun.String
}
