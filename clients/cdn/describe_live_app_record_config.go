package cdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeLiveAppRecordConfigRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveAppRecordConfigRequest) Invoke(client goaliyun.Client) (*DescribeLiveAppRecordConfigResponse, error) {
	resp := &DescribeLiveAppRecordConfigResponse{}
	err := client.Request("cdn", "DescribeLiveAppRecordConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveAppRecordConfigResponse struct {
	RequestId     goaliyun.String
	LiveAppRecord DescribeLiveAppRecordConfigLiveAppRecord
}

type DescribeLiveAppRecordConfigLiveAppRecord struct {
	DomainName      goaliyun.String
	AppName         goaliyun.String
	OssEndpoint     goaliyun.String
	OssBucket       goaliyun.String
	OssObjectPrefix goaliyun.String
	CreateTime      goaliyun.String
}
