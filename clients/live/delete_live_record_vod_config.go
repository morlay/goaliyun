package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteLiveRecordVodConfigRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteLiveRecordVodConfigRequest) Invoke(client goaliyun.Client) (*DeleteLiveRecordVodConfigResponse, error) {
	resp := &DeleteLiveRecordVodConfigResponse{}
	err := client.Request("live", "DeleteLiveRecordVodConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLiveRecordVodConfigResponse struct {
	RequestId goaliyun.String
}
