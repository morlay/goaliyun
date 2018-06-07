package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AddLiveSnapshotDetectPornConfigRequest struct {
	OssBucket     string                                    `position:"Query" name:"OssBucket"`
	AppName       string                                    `position:"Query" name:"AppName"`
	SecurityToken string                                    `position:"Query" name:"SecurityToken"`
	DomainName    string                                    `position:"Query" name:"DomainName"`
	OssEndpoint   string                                    `position:"Query" name:"OssEndpoint"`
	Interval      int64                                     `position:"Query" name:"Interval"`
	OwnerId       int64                                     `position:"Query" name:"OwnerId"`
	OssObject     string                                    `position:"Query" name:"OssObject"`
	Scenes        *AddLiveSnapshotDetectPornConfigSceneList `position:"Query" type:"Repeated" name:"Scene"`
	RegionId      string                                    `position:"Query" name:"RegionId"`
}

func (req *AddLiveSnapshotDetectPornConfigRequest) Invoke(client goaliyun.Client) (*AddLiveSnapshotDetectPornConfigResponse, error) {
	resp := &AddLiveSnapshotDetectPornConfigResponse{}
	err := client.Request("live", "AddLiveSnapshotDetectPornConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddLiveSnapshotDetectPornConfigResponse struct {
	RequestId goaliyun.String
}

type AddLiveSnapshotDetectPornConfigSceneList []string

func (list *AddLiveSnapshotDetectPornConfigSceneList) UnmarshalJSON(data []byte) error {
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
