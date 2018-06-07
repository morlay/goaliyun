package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type UpdateLiveSnapshotDetectPornConfigRequest struct {
	OssBucket     string                                       `position:"Query" name:"OssBucket"`
	AppName       string                                       `position:"Query" name:"AppName"`
	SecurityToken string                                       `position:"Query" name:"SecurityToken"`
	DomainName    string                                       `position:"Query" name:"DomainName"`
	OssEndpoint   string                                       `position:"Query" name:"OssEndpoint"`
	Interval      int64                                        `position:"Query" name:"Interval"`
	OwnerId       int64                                        `position:"Query" name:"OwnerId"`
	OssObject     string                                       `position:"Query" name:"OssObject"`
	Scenes        *UpdateLiveSnapshotDetectPornConfigSceneList `position:"Query" type:"Repeated" name:"Scene"`
	RegionId      string                                       `position:"Query" name:"RegionId"`
}

func (req *UpdateLiveSnapshotDetectPornConfigRequest) Invoke(client goaliyun.Client) (*UpdateLiveSnapshotDetectPornConfigResponse, error) {
	resp := &UpdateLiveSnapshotDetectPornConfigResponse{}
	err := client.Request("live", "UpdateLiveSnapshotDetectPornConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateLiveSnapshotDetectPornConfigResponse struct {
	RequestId goaliyun.String
}

type UpdateLiveSnapshotDetectPornConfigSceneList []string

func (list *UpdateLiveSnapshotDetectPornConfigSceneList) UnmarshalJSON(data []byte) error {
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
