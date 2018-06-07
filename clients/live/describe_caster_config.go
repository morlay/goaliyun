package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCasterConfigRequest struct {
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribeCasterConfigRequest) Invoke(client goaliyun.Client) (*DescribeCasterConfigResponse, error) {
	resp := &DescribeCasterConfigResponse{}
	err := client.Request("live", "DescribeCasterConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCasterConfigResponse struct {
	RequestId        goaliyun.String
	CasterId         goaliyun.String
	CasterName       goaliyun.String
	DomainName       goaliyun.String
	Delay            goaliyun.Float
	UrgentMaterialId goaliyun.String
	SideOutputUrl    goaliyun.String
	CallbackUrl      goaliyun.String
	ProgramName      goaliyun.String
	ProgramEffect    goaliyun.Integer
	TranscodeConfig  DescribeCasterConfigTranscodeConfig
	RecordConfig     DescribeCasterConfigRecordConfig
}

type DescribeCasterConfigTranscodeConfig struct {
	CasterTemplate  goaliyun.String
	LiveTemplateIds DescribeCasterConfigLiveTemplateIdList
}

type DescribeCasterConfigRecordConfig struct {
	OssEndpoint  goaliyun.String
	OssBucket    goaliyun.String
	RecordFormat DescribeCasterConfigRecordFormatItemList
}

type DescribeCasterConfigRecordFormatItem struct {
	Format               goaliyun.String
	OssObjectPrefix      goaliyun.String
	SliceOssObjectPrefix goaliyun.String
	CycleDuration        goaliyun.Integer
}

type DescribeCasterConfigLiveTemplateIdList []goaliyun.String

func (list *DescribeCasterConfigLiveTemplateIdList) UnmarshalJSON(data []byte) error {
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

type DescribeCasterConfigRecordFormatItemList []DescribeCasterConfigRecordFormatItem

func (list *DescribeCasterConfigRecordFormatItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterConfigRecordFormatItem)
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
