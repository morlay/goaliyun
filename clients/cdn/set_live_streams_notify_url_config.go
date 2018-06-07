package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetLiveStreamsNotifyUrlConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	NotifyUrl     string `position:"Query" name:"NotifyUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetLiveStreamsNotifyUrlConfigRequest) Invoke(client goaliyun.Client) (*SetLiveStreamsNotifyUrlConfigResponse, error) {
	resp := &SetLiveStreamsNotifyUrlConfigResponse{}
	err := client.Request("cdn", "SetLiveStreamsNotifyUrlConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetLiveStreamsNotifyUrlConfigResponse struct {
	RequestId goaliyun.String
}
