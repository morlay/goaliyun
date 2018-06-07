package live

import (
	"github.com/morlay/goaliyun"
)

type StopCasterSceneRequest struct {
	CasterId string `position:"Query" name:"CasterId"`
	SceneId  string `position:"Query" name:"SceneId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *StopCasterSceneRequest) Invoke(client goaliyun.Client) (*StopCasterSceneResponse, error) {
	resp := &StopCasterSceneResponse{}
	err := client.Request("live", "StopCasterScene", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StopCasterSceneResponse struct {
	RequestId goaliyun.String
}
