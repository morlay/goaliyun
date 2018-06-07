package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteLiveSnapshotDetectPornConfigRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteLiveSnapshotDetectPornConfigRequest) Invoke(client goaliyun.Client) (*DeleteLiveSnapshotDetectPornConfigResponse, error) {
	resp := &DeleteLiveSnapshotDetectPornConfigResponse{}
	err := client.Request("live", "DeleteLiveSnapshotDetectPornConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLiveSnapshotDetectPornConfigResponse struct {
	RequestId goaliyun.String
}
