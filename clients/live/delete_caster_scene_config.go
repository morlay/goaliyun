package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteCasterSceneConfigRequest struct {
	CasterId string `position:"Query" name:"CasterId"`
	SceneId  string `position:"Query" name:"SceneId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	Type     string `position:"Query" name:"Type"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteCasterSceneConfigRequest) Invoke(client goaliyun.Client) (*DeleteCasterSceneConfigResponse, error) {
	resp := &DeleteCasterSceneConfigResponse{}
	err := client.Request("live", "DeleteCasterSceneConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCasterSceneConfigResponse struct {
	RequestId goaliyun.String
}
