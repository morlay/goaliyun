package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamSnapshotInfoRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Limit         int64  `position:"Query" name:"Limit"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamSnapshotInfoRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamSnapshotInfoResponse, error) {
	resp := &DescribeLiveStreamSnapshotInfoResponse{}
	err := client.Request("live", "DescribeLiveStreamSnapshotInfo", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamSnapshotInfoResponse struct {
	RequestId                  goaliyun.String
	NextStartTime              goaliyun.String
	LiveStreamSnapshotInfoList DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfoList
}

type DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfo struct {
	OssEndpoint goaliyun.String
	OssBucket   goaliyun.String
	OssObject   goaliyun.String
	CreateTime  goaliyun.String
}

type DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfoList []DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfo

func (list *DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfo)
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
