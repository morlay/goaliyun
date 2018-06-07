package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamRecordIndexFilesRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamRecordIndexFilesRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamRecordIndexFilesResponse, error) {
	resp := &DescribeLiveStreamRecordIndexFilesResponse{}
	err := client.Request("cdn", "DescribeLiveStreamRecordIndexFiles", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamRecordIndexFilesResponse struct {
	RequestId           goaliyun.String
	RecordIndexInfoList DescribeLiveStreamRecordIndexFilesRecordIndexInfoList
}

type DescribeLiveStreamRecordIndexFilesRecordIndexInfo struct {
	RecordId   goaliyun.String
	RecordUrl  goaliyun.String
	DomainName goaliyun.String
	AppName    goaliyun.String
	StreamName goaliyun.String
	OssObject  goaliyun.String
	StartTime  goaliyun.String
	EndTime    goaliyun.String
	Duration   goaliyun.Float
	Height     goaliyun.Integer
	Width      goaliyun.Integer
	CreateTime goaliyun.String
}

type DescribeLiveStreamRecordIndexFilesRecordIndexInfoList []DescribeLiveStreamRecordIndexFilesRecordIndexInfo

func (list *DescribeLiveStreamRecordIndexFilesRecordIndexInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRecordIndexFilesRecordIndexInfo)
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
