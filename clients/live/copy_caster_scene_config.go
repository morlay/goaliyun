package live

import (
	"github.com/morlay/goaliyun"
)

type CopyCasterSceneConfigRequest struct {
	FromSceneId string `position:"Query" name:"FromSceneId"`
	CasterId    string `position:"Query" name:"CasterId"`
	OwnerId     int64  `position:"Query" name:"OwnerId"`
	ToSceneId   string `position:"Query" name:"ToSceneId"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *CopyCasterSceneConfigRequest) Invoke(client goaliyun.Client) (*CopyCasterSceneConfigResponse, error) {
	resp := &CopyCasterSceneConfigResponse{}
	err := client.Request("live", "CopyCasterSceneConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CopyCasterSceneConfigResponse struct {
	RequestId goaliyun.String
}
