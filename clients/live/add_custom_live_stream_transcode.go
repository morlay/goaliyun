package live

import (
	"github.com/morlay/goaliyun"
)

type AddCustomLiveStreamTranscodeRequest struct {
	App           string `position:"Query" name:"App"`
	Template      string `position:"Query" name:"Template"`
	FPS           int64  `position:"Query" name:"FPS"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	TemplateType  string `position:"Query" name:"TemplateType"`
	Domain        string `position:"Query" name:"Domain"`
	Width         int64  `position:"Query" name:"Width"`
	VideoBitrate  int64  `position:"Query" name:"VideoBitrate"`
	Height        int64  `position:"Query" name:"Height"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *AddCustomLiveStreamTranscodeRequest) Invoke(client goaliyun.Client) (*AddCustomLiveStreamTranscodeResponse, error) {
	resp := &AddCustomLiveStreamTranscodeResponse{}
	err := client.Request("live", "AddCustomLiveStreamTranscode", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddCustomLiveStreamTranscodeResponse struct {
	RequestId goaliyun.String
}
