package cdn

import (
	"github.com/morlay/goaliyun"
)

type CreateLiveStreamRecordIndexFilesRequest struct {
	OssBucket     string `position:"Query" name:"OssBucket"`
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OssEndpoint   string `position:"Query" name:"OssEndpoint"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	OssObject     string `position:"Query" name:"OssObject"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *CreateLiveStreamRecordIndexFilesRequest) Invoke(client goaliyun.Client) (*CreateLiveStreamRecordIndexFilesResponse, error) {
	resp := &CreateLiveStreamRecordIndexFilesResponse{}
	err := client.Request("cdn", "CreateLiveStreamRecordIndexFiles", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateLiveStreamRecordIndexFilesResponse struct {
	RequestId  goaliyun.String
	RecordInfo CreateLiveStreamRecordIndexFilesRecordInfo
}

type CreateLiveStreamRecordIndexFilesRecordInfo struct {
	RecordId   goaliyun.String
	RecordUrl  goaliyun.String
	Duration   goaliyun.Float
	Height     goaliyun.Integer
	Width      goaliyun.Integer
	CreateTime goaliyun.String
}
