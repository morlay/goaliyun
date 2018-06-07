package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteLiveAppSnapshotConfigRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteLiveAppSnapshotConfigRequest) Invoke(client goaliyun.Client) (*DeleteLiveAppSnapshotConfigResponse, error) {
	resp := &DeleteLiveAppSnapshotConfigResponse{}
	err := client.Request("live", "DeleteLiveAppSnapshotConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLiveAppSnapshotConfigResponse struct {
	RequestId goaliyun.String
}
