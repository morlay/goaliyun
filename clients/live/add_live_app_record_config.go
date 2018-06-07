package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AddLiveAppRecordConfigRequest struct {
	OssBucket     string                                  `position:"Query" name:"OssBucket"`
	DomainName    string                                  `position:"Query" name:"DomainName"`
	OssEndpoint   string                                  `position:"Query" name:"OssEndpoint"`
	EndTime       string                                  `position:"Query" name:"EndTime"`
	StartTime     string                                  `position:"Query" name:"StartTime"`
	OwnerId       int64                                   `position:"Query" name:"OwnerId"`
	AppName       string                                  `position:"Query" name:"AppName"`
	SecurityToken string                                  `position:"Query" name:"SecurityToken"`
	RecordFormats *AddLiveAppRecordConfigRecordFormatList `position:"Query" type:"Repeated" name:"RecordFormat"`
	OnDemand      int64                                   `position:"Query" name:"OnDemand"`
	StreamName    string                                  `position:"Query" name:"StreamName"`
	RegionId      string                                  `position:"Query" name:"RegionId"`
}

func (req *AddLiveAppRecordConfigRequest) Invoke(client goaliyun.Client) (*AddLiveAppRecordConfigResponse, error) {
	resp := &AddLiveAppRecordConfigResponse{}
	err := client.Request("live", "AddLiveAppRecordConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddLiveAppRecordConfigRecordFormat struct {
	Format               string `name:"Format"`
	OssObjectPrefix      string `name:"OssObjectPrefix"`
	SliceOssObjectPrefix string `name:"SliceOssObjectPrefix"`
	CycleDuration        int64  `name:"CycleDuration"`
}

type AddLiveAppRecordConfigResponse struct {
	RequestId goaliyun.String
}

type AddLiveAppRecordConfigRecordFormatList []AddLiveAppRecordConfigRecordFormat

func (list *AddLiveAppRecordConfigRecordFormatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddLiveAppRecordConfigRecordFormat)
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
