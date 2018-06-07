package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type StartCasterRequest struct {
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *StartCasterRequest) Invoke(client goaliyun.Client) (*StartCasterResponse, error) {
	resp := &StartCasterResponse{}
	err := client.Request("live", "StartCaster", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StartCasterResponse struct {
	RequestId     goaliyun.String
	PvwSceneInfos StartCasterSceneInfoList
	PgmSceneInfos StartCasterSceneInfo1List
}

type StartCasterSceneInfo struct {
	SceneId   goaliyun.String
	StreamUrl goaliyun.String
}

type StartCasterSceneInfo1 struct {
	SceneId     goaliyun.String
	StreamUrl   goaliyun.String
	StreamInfos StartCasterStreamInfoList
}

type StartCasterStreamInfo struct {
	TranscodeConfig goaliyun.String
	VideoFormat     goaliyun.String
	OutputStreamUrl goaliyun.String
}

type StartCasterSceneInfoList []StartCasterSceneInfo

func (list *StartCasterSceneInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartCasterSceneInfo)
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

type StartCasterSceneInfo1List []StartCasterSceneInfo1

func (list *StartCasterSceneInfo1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartCasterSceneInfo1)
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

type StartCasterStreamInfoList []StartCasterStreamInfo

func (list *StartCasterStreamInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartCasterStreamInfo)
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
