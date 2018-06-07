package live

import (
	"github.com/morlay/goaliyun"
)

type SetCasterConfigRequest struct {
	SideOutputUrl    string  `position:"Query" name:"SideOutputUrl"`
	CasterId         string  `position:"Query" name:"CasterId"`
	DomainName       string  `position:"Query" name:"DomainName"`
	ProgramEffect    int64   `position:"Query" name:"ProgramEffect"`
	ProgramName      string  `position:"Query" name:"ProgramName"`
	OwnerId          int64   `position:"Query" name:"OwnerId"`
	RecordConfig     string  `position:"Query" name:"RecordConfig"`
	UrgentMaterialId string  `position:"Query" name:"UrgentMaterialId"`
	TranscodeConfig  string  `position:"Query" name:"TranscodeConfig"`
	Delay            float64 `position:"Query" name:"Delay"`
	CasterName       string  `position:"Query" name:"CasterName"`
	CallbackUrl      string  `position:"Query" name:"CallbackUrl"`
	RegionId         string  `position:"Query" name:"RegionId"`
}

func (req *SetCasterConfigRequest) Invoke(client goaliyun.Client) (*SetCasterConfigResponse, error) {
	resp := &SetCasterConfigResponse{}
	err := client.Request("live", "SetCasterConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetCasterConfigResponse struct {
	RequestId goaliyun.String
	CasterId  goaliyun.String
}
