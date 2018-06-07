package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteLiveRecordNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteLiveRecordNotifyConfigRequest) Invoke(client goaliyun.Client) (*DeleteLiveRecordNotifyConfigResponse, error) {
	resp := &DeleteLiveRecordNotifyConfigResponse{}
	err := client.Request("live", "DeleteLiveRecordNotifyConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLiveRecordNotifyConfigResponse struct {
	RequestId goaliyun.String
}
