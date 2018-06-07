package alidns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDnsProductAttributesRequest struct {
	VersionCode string `position:"Query" name:"VersionCode"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *DescribeDnsProductAttributesRequest) Invoke(client goaliyun.Client) (*DescribeDnsProductAttributesResponse, error) {
	resp := &DescribeDnsProductAttributesResponse{}
	err := client.Request("alidns", "DescribeDnsProductAttributes", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDnsProductAttributesResponse struct {
	RequestId   goaliyun.String
	TtlMinValue goaliyun.String
	TtlMaxValue goaliyun.String
	RecordLines DescribeDnsProductAttributesRecordLineList
	RecordTypes DescribeDnsProductAttributesRecordTypeList
}

type DescribeDnsProductAttributesRecordLine struct {
	LineCode   goaliyun.String
	FatherCode goaliyun.String
	LineName   goaliyun.String
}

type DescribeDnsProductAttributesRecordLineList []DescribeDnsProductAttributesRecordLine

func (list *DescribeDnsProductAttributesRecordLineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDnsProductAttributesRecordLine)
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

type DescribeDnsProductAttributesRecordTypeList []goaliyun.String

func (list *DescribeDnsProductAttributesRecordTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
