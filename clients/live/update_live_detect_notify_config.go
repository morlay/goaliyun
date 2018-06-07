package live

import (
	"github.com/morlay/goaliyun"
)

type UpdateLiveDetectNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	NotifyUrl     string `position:"Query" name:"NotifyUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *UpdateLiveDetectNotifyConfigRequest) Invoke(client goaliyun.Client) (*UpdateLiveDetectNotifyConfigResponse, error) {
	resp := &UpdateLiveDetectNotifyConfigResponse{}
	err := client.Request("live", "UpdateLiveDetectNotifyConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateLiveDetectNotifyConfigResponse struct {
	RequestId goaliyun.String
}
