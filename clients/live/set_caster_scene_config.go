package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SetCasterSceneConfigRequest struct {
	ComponentIds *SetCasterSceneConfigComponentIdList `position:"Query" type:"Repeated" name:"ComponentId"`
	CasterId     string                               `position:"Query" name:"CasterId"`
	SceneId      string                               `position:"Query" name:"SceneId"`
	OwnerId      int64                                `position:"Query" name:"OwnerId"`
	LayoutId     string                               `position:"Query" name:"LayoutId"`
	RegionId     string                               `position:"Query" name:"RegionId"`
}

func (req *SetCasterSceneConfigRequest) Invoke(client goaliyun.Client) (*SetCasterSceneConfigResponse, error) {
	resp := &SetCasterSceneConfigResponse{}
	err := client.Request("live", "SetCasterSceneConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetCasterSceneConfigResponse struct {
	RequestId goaliyun.String
}

type SetCasterSceneConfigComponentIdList []string

func (list *SetCasterSceneConfigComponentIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
