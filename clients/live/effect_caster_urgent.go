package live

import (
	"github.com/morlay/goaliyun"
)

type EffectCasterUrgentRequest struct {
	CasterId string `position:"Query" name:"CasterId"`
	SceneId  string `position:"Query" name:"SceneId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *EffectCasterUrgentRequest) Invoke(client goaliyun.Client) (*EffectCasterUrgentResponse, error) {
	resp := &EffectCasterUrgentResponse{}
	err := client.Request("live", "EffectCasterUrgent", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EffectCasterUrgentResponse struct {
	RequestId goaliyun.String
}
