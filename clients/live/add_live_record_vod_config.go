package live

import (
	"github.com/morlay/goaliyun"
)

type AddLiveRecordVodConfigRequest struct {
	AppName                    string `position:"Query" name:"AppName"`
	AutoCompose                string `position:"Query" name:"AutoCompose"`
	DomainName                 string `position:"Query" name:"DomainName"`
	CycleDuration              int64  `position:"Query" name:"CycleDuration"`
	OwnerId                    int64  `position:"Query" name:"OwnerId"`
	ComposeVodTranscodeGroupId string `position:"Query" name:"ComposeVodTranscodeGroupId"`
	StreamName                 string `position:"Query" name:"StreamName"`
	VodTranscodeGroupId        string `position:"Query" name:"VodTranscodeGroupId"`
	RegionId                   string `position:"Query" name:"RegionId"`
}

func (req *AddLiveRecordVodConfigRequest) Invoke(client goaliyun.Client) (*AddLiveRecordVodConfigResponse, error) {
	resp := &AddLiveRecordVodConfigResponse{}
	err := client.Request("live", "AddLiveRecordVodConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddLiveRecordVodConfigResponse struct {
	RequestId goaliyun.String
}
