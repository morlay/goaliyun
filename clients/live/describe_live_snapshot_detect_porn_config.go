package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveSnapshotDetectPornConfigRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNum       int64  `position:"Query" name:"PageNum"`
	Order         string `position:"Query" name:"Order"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveSnapshotDetectPornConfigRequest) Invoke(client goaliyun.Client) (*DescribeLiveSnapshotDetectPornConfigResponse, error) {
	resp := &DescribeLiveSnapshotDetectPornConfigResponse{}
	err := client.Request("live", "DescribeLiveSnapshotDetectPornConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveSnapshotDetectPornConfigResponse struct {
	RequestId                        goaliyun.String
	PageNum                          goaliyun.Integer
	PageSize                         goaliyun.Integer
	Order                            goaliyun.String
	TotalNum                         goaliyun.Integer
	TotalPage                        goaliyun.Integer
	LiveSnapshotDetectPornConfigList DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfigList
}

type DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfig struct {
	DomainName  goaliyun.String
	AppName     goaliyun.String
	OssEndpoint goaliyun.String
	OssBucket   goaliyun.String
	OssObject   goaliyun.String
	Interval    goaliyun.Integer
	Scenes      DescribeLiveSnapshotDetectPornConfigSceneList
}

type DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfigList []DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfig

func (list *DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfig)
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

type DescribeLiveSnapshotDetectPornConfigSceneList []goaliyun.String

func (list *DescribeLiveSnapshotDetectPornConfigSceneList) UnmarshalJSON(data []byte) error {
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
