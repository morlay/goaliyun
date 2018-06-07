package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamDomainAppInfoRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	AppDomain     string `position:"Query" name:"AppDomain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamDomainAppInfoRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamDomainAppInfoResponse, error) {
	resp := &DescribeLiveStreamDomainAppInfoResponse{}
	err := client.Request("cdn", "DescribeLiveStreamDomainAppInfo", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamDomainAppInfoResponse struct {
	RequestId     goaliyun.String
	DomainAppList DescribeLiveStreamDomainAppInfoDomainAppInfoList
}

type DescribeLiveStreamDomainAppInfoDomainAppInfo struct {
	AppDomain    goaliyun.String
	AppId        goaliyun.String
	AppKey       goaliyun.String
	AppOssBucket goaliyun.String
	AppOssHost   goaliyun.String
	AppOwnerId   goaliyun.String
	AppSecret    goaliyun.String
	UpdateTime   goaliyun.String
}

type DescribeLiveStreamDomainAppInfoDomainAppInfoList []DescribeLiveStreamDomainAppInfoDomainAppInfo

func (list *DescribeLiveStreamDomainAppInfoDomainAppInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamDomainAppInfoDomainAppInfo)
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
