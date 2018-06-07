package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveRecordConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveRecordConfigRequest) Invoke(client goaliyun.Client) (*DescribeLiveRecordConfigResponse, error) {
	resp := &DescribeLiveRecordConfigResponse{}
	err := client.Request("cdn", "DescribeLiveRecordConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveRecordConfigResponse struct {
	RequestId         goaliyun.String
	LiveAppRecordList DescribeLiveRecordConfigLiveAppRecordList
}

type DescribeLiveRecordConfigLiveAppRecord struct {
	DomainName      goaliyun.String
	AppName         goaliyun.String
	OssEndpoint     goaliyun.String
	OssBucket       goaliyun.String
	OssObjectPrefix goaliyun.String
	CreateTime      goaliyun.String
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
