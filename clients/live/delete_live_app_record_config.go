package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteLiveAppRecordConfigRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteLiveAppRecordConfigRequest) Invoke(client goaliyun.Client) (*DeleteLiveAppRecordConfigResponse, error) {
	resp := &DeleteLiveAppRecordConfigResponse{}
	err := client.Request("live", "DeleteLiveAppRecordConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLiveAppRecordConfigResponse struct {
	RequestId goaliyun.String
}
