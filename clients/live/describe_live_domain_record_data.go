package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveDomainRecordDataRequest struct {
	RecordType string `position:"Query" name:"RecordType"`
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveDomainRecordDataRequest) Invoke(client goaliyun.Client) (*DescribeLiveDomainRecordDataResponse, error) {
	resp := &DescribeLiveDomainRecordDataResponse{}
	err := client.Request("live", "DescribeLiveDomainRecordData", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveDomainRecordDataResponse struct {
	RequestId       goaliyun.String
	RecordDataInfos DescribeLiveDomainRecordDataRecordDataInfoList
}

type DescribeLiveDomainRecordDataRecordDataInfo struct {
	Date   goaliyun.String
	Total  goaliyun.Integer
	Detail DescribeLiveDomainRecordDataDetail
}

type DescribeLiveDomainRecordDataDetail struct {
	MP4 goaliyun.Integer
	FLV goaliyun.Integer
	TS  goaliyun.Integer
}

type DescribeLiveDomainRecordDataRecordDataInfoList []DescribeLiveDomainRecordDataRecordDataInfo

func (list *DescribeLiveDomainRecordDataRecordDataInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveDomainRecordDataRecordDataInfo)
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
