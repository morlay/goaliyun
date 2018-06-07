package cdn

import (
	"github.com/morlay/goaliyun"
)

type DeleteLiveAppRecordConfigRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteLiveAppRecordConfigRequest) Invoke(client goaliyun.Client) (*DeleteLiveAppRecordConfigResponse, error) {
	resp := &DeleteLiveAppRecordConfigResponse{}
	err := client.Request("cdn", "DeleteLiveAppRecordConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLiveAppRecordConfigResponse struct {
	RequestId goaliyun.String
}
