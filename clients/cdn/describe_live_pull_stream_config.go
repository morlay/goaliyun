package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLivePullStreamConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLivePullStreamConfigRequest) Invoke(client goaliyun.Client) (*DescribeLivePullStreamConfigResponse, error) {
	resp := &DescribeLivePullStreamConfigResponse{}
	err := client.Request("cdn", "DescribeLivePullStreamConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLivePullStreamConfigResponse struct {
	RequestId         goaliyun.String
	LiveAppRecordList DescribeLivePullStreamConfigLiveAppRecordList
}

type DescribeLivePullStreamConfigLiveAppRecord struct {
	DomainName goaliyun.String
	StreamName goaliyun.String
	SourceUrl  goaliyun.String
	StartTime  goaliyun.String
	EndTime    goaliyun.String
}

type DescribeLivePullStreamConfigLiveAppRecordList []DescribeLivePullStreamConfigLiveAppRecord

func (list *DescribeLivePullStreamConfigLiveAppRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLivePullStreamConfigLiveAppRecord)
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
