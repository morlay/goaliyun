package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteLiveStreamsNotifyUrlConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteLiveStreamsNotifyUrlConfigRequest) Invoke(client goaliyun.Client) (*DeleteLiveStreamsNotifyUrlConfigResponse, error) {
	resp := &DeleteLiveStreamsNotifyUrlConfigResponse{}
	err := client.Request("live", "DeleteLiveStreamsNotifyUrlConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLiveStreamsNotifyUrlConfigResponse struct {
	RequestId goaliyun.String
}
