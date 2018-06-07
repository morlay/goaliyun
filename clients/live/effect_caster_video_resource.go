package live

import (
	"github.com/morlay/goaliyun"
)

type EffectCasterVideoResourceRequest struct {
	ResourceId string `position:"Query" name:"ResourceId"`
	CasterId   string `position:"Query" name:"CasterId"`
	SceneId    string `position:"Query" name:"SceneId"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *EffectCasterVideoResourceRequest) Invoke(client goaliyun.Client) (*EffectCasterVideoResourceResponse, error) {
	resp := &EffectCasterVideoResourceResponse{}
	err := client.Request("live", "EffectCasterVideoResource", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EffectCasterVideoResourceResponse struct {
	RequestId goaliyun.String
}
