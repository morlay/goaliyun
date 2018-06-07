package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveRecordConfigRequest struct {
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

func (req *DescribeLiveRecordConfigRequest) Invoke(client goaliyun.Client) (*DescribeLiveRecordConfigResponse, error) {
	resp := &DescribeLiveRecordConfigResponse{}
	err := client.Request("live", "DescribeLiveRecordConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveRecordConfigResponse struct {
	RequestId         goaliyun.String
	PageNum           goaliyun.Integer
	PageSize          goaliyun.Integer
	Order             goaliyun.String
	TotalNum          goaliyun.Integer
	TotalPage         goaliyun.Integer
	LiveAppRecordList DescribeLiveRecordConfigLiveAppRecordList
}

type DescribeLiveRecordConfigLiveAppRecord struct {
	DomainName       goaliyun.String
	AppName          goaliyun.String
	StreamName       goaliyun.String
	OssEndpoint      goaliyun.String
	OssBucket        goaliyun.String
	CreateTime       goaliyun.String
	StartTime        goaliyun.String
	EndTime          goaliyun.String
	OnDemond         goaliyun.Integer
	RecordFormatList DescribeLiveRecordConfigRecordFormatList
}

type DescribeLiveRecordConfigRecordFormat struct {
	Format               goaliyun.String
	OssObjectPrefix      goaliyun.String
	SliceOssObjectPrefix goaliyun.String
	CycleDuration        goaliyun.Integer
}

type DescribeLiveRecordConfigLiveAppRecordList []DescribeLiveRecordConfigLiveAppRecord

func (list *DescribeLiveRecordConfigLiveAppRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveRecordConfigLiveAppRecord)
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

type DescribeLiveRecordConfigRecordFormatList []DescribeLiveRecordConfigRecordFormat

func (list *DescribeLiveRecordConfigRecordFormatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveRecordConfigRecordFormat)
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
