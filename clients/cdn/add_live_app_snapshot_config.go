package cdn

import (
	"github.com/morlay/goaliyun"
)

type AddLiveAppSnapshotConfigRequest struct {
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

func (req *AddLiveAppSnapshotConfigRequest) Invoke(client goaliyun.Client) (*AddLiveAppSnapshotConfigResponse, error) {
	resp := &AddLiveAppSnapshotConfigResponse{}
	err := client.Request("cdn", "AddLiveAppSnapshotConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddLiveAppSnapshotConfigResponse struct {
	RequestId goaliyun.String
}
