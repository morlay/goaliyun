package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveSnapshotConfigRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNum       int64  `position:"Query" name:"PageNum"`
	StreamName    string `position:"Query" name:"StreamName"`
	Order         string `position:"Query" name:"Order"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveSnapshotConfigRequest) Invoke(client goaliyun.Client) (*DescribeLiveSnapshotConfigResponse, error) {
	resp := &DescribeLiveSnapshotConfigResponse{}
	err := client.Request("cdn", "DescribeLiveSnapshotConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveSnapshotConfigResponse struct {
	RequestId                    goaliyun.String
	PageNum                      goaliyun.Integer
	PageSize                     goaliyun.Integer
	Order                        goaliyun.String
	TotalNum                     goaliyun.Integer
	TotalPage                    goaliyun.Integer
	LiveStreamSnapshotConfigList DescribeLiveSnapshotConfigLiveStreamSnapshotConfigList
}

type DescribeLiveSnapshotConfigLiveStreamSnapshotConfig struct {
	DomainName         goaliyun.String
	AppName            goaliyun.String
	TimeInterval       goaliyun.Integer
	OssEndpoint        goaliyun.String
	OssBucket          goaliyun.String
	OverwriteOssObject goaliyun.String
	SequenceOssObject  goaliyun.String
	CreateTime         goaliyun.String
}

type DescribeLiveSnapshotConfigLiveStreamSnapshotConfigList []DescribeLiveSnapshotConfigLiveStreamSnapshotConfig

func (list *DescribeLiveSnapshotConfigLiveStreamSnapshotConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveSnapshotConfigLiveStreamSnapshotConfig)
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
