package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamRecordContentRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamRecordContentRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamRecordContentResponse, error) {
	resp := &DescribeLiveStreamRecordContentResponse{}
	err := client.Request("cdn", "DescribeLiveStreamRecordContent", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamRecordContentResponse struct {
	RequestId             goaliyun.String
	RecordContentInfoList DescribeLiveStreamRecordContentRecordContentInfoList
}

type DescribeLiveStreamRecordContentRecordContentInfo struct {
	OssEndpoint     goaliyun.String
	OssBucket       goaliyun.String
	OssObjectPrefix goaliyun.String
	StartTime       goaliyun.String
	EndTime         goaliyun.String
	Duration        goaliyun.Float
}

type DescribeLiveStreamRecordContentRecordContentInfoList []DescribeLiveStreamRecordContentRecordContentInfo

func (list *DescribeLiveStreamRecordContentRecordContentInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRecordContentRecordContentInfo)
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
