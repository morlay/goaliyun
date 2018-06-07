package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveRecordVodConfigsRequest struct {
	AppName    string `position:"Query" name:"AppName"`
	DomainName string `position:"Query" name:"DomainName"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	PageNum    int64  `position:"Query" name:"PageNum"`
	StreamName string `position:"Query" name:"StreamName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveRecordVodConfigsRequest) Invoke(client goaliyun.Client) (*DescribeLiveRecordVodConfigsResponse, error) {
	resp := &DescribeLiveRecordVodConfigsResponse{}
	err := client.Request("live", "DescribeLiveRecordVodConfigs", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveRecordVodConfigsResponse struct {
	RequestId            goaliyun.String
	PageNum              goaliyun.Integer
	PageSize             goaliyun.Integer
	Total                goaliyun.String
	LiveRecordVodConfigs DescribeLiveRecordVodConfigsLiveRecordVodConfigList
}

type DescribeLiveRecordVodConfigsLiveRecordVodConfig struct {
	CreateTime                 goaliyun.String
	DomainName                 goaliyun.String
	AppName                    goaliyun.String
	StreamName                 goaliyun.String
	VodTranscodeGroupId        goaliyun.String
	CycleDuration              goaliyun.Integer
	AutoCompose                goaliyun.String
	ComposeVodTranscodeGroupId goaliyun.String
}

type DescribeLiveRecordVodConfigsLiveRecordVodConfigList []DescribeLiveRecordVodConfigsLiveRecordVodConfig

func (list *DescribeLiveRecordVodConfigsLiveRecordVodConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveRecordVodConfigsLiveRecordVodConfig)
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
