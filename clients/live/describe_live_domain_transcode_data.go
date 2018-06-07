package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveDomainTranscodeDataRequest struct {
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveDomainTranscodeDataRequest) Invoke(client goaliyun.Client) (*DescribeLiveDomainTranscodeDataResponse, error) {
	resp := &DescribeLiveDomainTranscodeDataResponse{}
	err := client.Request("live", "DescribeLiveDomainTranscodeData", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveDomainTranscodeDataResponse struct {
	RequestId          goaliyun.String
	TranscodeDataInfos DescribeLiveDomainTranscodeDataTranscodeDataInfoList
}

type DescribeLiveDomainTranscodeDataTranscodeDataInfo struct {
	Date   goaliyun.String
	Total  goaliyun.Integer
	Detail goaliyun.String
}

type DescribeLiveDomainTranscodeDataTranscodeDataInfoList []DescribeLiveDomainTranscodeDataTranscodeDataInfo

func (list *DescribeLiveDomainTranscodeDataTranscodeDataInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveDomainTranscodeDataTranscodeDataInfo)
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
