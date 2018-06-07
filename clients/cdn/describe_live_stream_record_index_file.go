package cdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamRecordIndexFileRequest struct {
	RecordId      string `position:"Query" name:"RecordId"`
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamRecordIndexFileRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamRecordIndexFileResponse, error) {
	resp := &DescribeLiveStreamRecordIndexFileResponse{}
	err := client.Request("cdn", "DescribeLiveStreamRecordIndexFile", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamRecordIndexFileResponse struct {
	RequestId       goaliyun.String
	RecordIndexInfo DescribeLiveStreamRecordIndexFileRecordIndexInfo
}

type DescribeLiveStreamRecordIndexFileRecordIndexInfo struct {
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
