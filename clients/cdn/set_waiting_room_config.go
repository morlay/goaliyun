package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetWaitingRoomConfigRequest struct {
	WaitUrl       string `position:"Query" name:"WaitUrl"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	WaitUri       string `position:"Query" name:"WaitUri"`
	MaxQps        int64  `position:"Query" name:"MaxQps"`
	MaxTimeWait   int64  `position:"Query" name:"MaxTimeWait"`
	DomainName    string `position:"Query" name:"DomainName"`
	AllowPct      int64  `position:"Query" name:"AllowPct"`
	GapTime       int64  `position:"Query" name:"GapTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetWaitingRoomConfigRequest) Invoke(client goaliyun.Client) (*SetWaitingRoomConfigResponse, error) {
	resp := &SetWaitingRoomConfigResponse{}
	err := client.Request("cdn", "SetWaitingRoomConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetWaitingRoomConfigResponse struct {
	RequestId goaliyun.String
}
