package cdn

import (
	"github.com/morlay/goaliyun"
)

type UpdateLiveAppSnapshotConfigRequest struct {
	TimeInterval       int64  `position:"Query" name:"TimeInterval"`
	OssBucket          string `position:"Query" name:"OssBucket"`
	AppName            string `position:"Query" name:"AppName"`
	SecurityToken      string `position:"Query" name:"SecurityToken"`
	DomainName         string `position:"Query" name:"DomainName"`
	OssEndpoint        string `position:"Query" name:"OssEndpoint"`
	SequenceOssObject  string `position:"Query" name:"SequenceOssObject"`
	OverwriteOssObject string `position:"Query" name:"OverwriteOssObject"`
	OwnerId            int64  `position:"Query" name:"OwnerId"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *UpdateLiveAppSnapshotConfigRequest) Invoke(client goaliyun.Client) (*UpdateLiveAppSnapshotConfigResponse, error) {
	resp := &UpdateLiveAppSnapshotConfigResponse{}
	err := client.Request("cdn", "UpdateLiveAppSnapshotConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateLiveAppSnapshotConfigResponse struct {
	RequestId goaliyun.String
}
