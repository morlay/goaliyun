package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCasterScenesRequest struct {
	CasterId string `position:"Query" name:"CasterId"`
	SceneId  string `position:"Query" name:"SceneId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribeCasterScenesRequest) Invoke(client goaliyun.Client) (*DescribeCasterScenesResponse, error) {
	resp := &DescribeCasterScenesResponse{}
	err := client.Request("live", "DescribeCasterScenes", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCasterScenesResponse struct {
	RequestId goaliyun.String
	Total     goaliyun.Integer
	SceneList DescribeCasterScenesSceneList
}

type DescribeCasterScenesScene struct {
	SceneId      goaliyun.String
	SceneName    goaliyun.String
	OutputType   goaliyun.String
	LayoutId     goaliyun.String
	StreamUrl    goaliyun.String
	Status       goaliyun.Integer
	StreamInfos  DescribeCasterScenesStreamInfoList
	ComponentIds DescribeCasterScenesComponentIdList
}

type DescribeCasterScenesStreamInfo struct {
	TranscodeConfig goaliyun.String
	VideoFormat     goaliyun.String
	OutputStreamUrl goaliyun.String
}

type DescribeCasterScenesSceneList []DescribeCasterScenesScene

func (list *DescribeCasterScenesSceneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterScenesScene)
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

type DescribeCasterScenesStreamInfoList []DescribeCasterScenesStreamInfo

func (list *DescribeCasterScenesStreamInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterScenesStreamInfo)
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

type DescribeCasterScenesComponentIdList []goaliyun.String

func (list *DescribeCasterScenesComponentIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
