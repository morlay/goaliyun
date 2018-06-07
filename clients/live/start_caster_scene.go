package live

import (
	"github.com/morlay/goaliyun"
)

type StartCasterSceneRequest struct {
	CasterId string `position:"Query" name:"CasterId"`
	SceneId  string `position:"Query" name:"SceneId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *StartCasterSceneRequest) Invoke(client goaliyun.Client) (*StartCasterSceneResponse, error) {
	resp := &StartCasterSceneResponse{}
	err := client.Request("live", "StartCasterScene", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StartCasterSceneResponse struct {
	RequestId goaliyun.String
	StreamUrl goaliyun.String
}
