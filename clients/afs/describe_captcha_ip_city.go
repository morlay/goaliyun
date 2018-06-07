package afs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCaptchaIpCityRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	ConfigName      string `position:"Query" name:"ConfigName"`
	Time            string `position:"Query" name:"Time"`
	Type            string `position:"Query" name:"Type"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeCaptchaIpCityRequest) Invoke(client goaliyun.Client) (*DescribeCaptchaIpCityResponse, error) {
	resp := &DescribeCaptchaIpCityResponse{}
	err := client.Request("afs", "DescribeCaptchaIpCity", "2018-01-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCaptchaIpCityResponse struct {
	RequestId     goaliyun.String
	BizCode       goaliyun.String
	HasData       bool
	CaptchaCities DescribeCaptchaIpCityCaptchaCitieList
	CaptchaIps    DescribeCaptchaIpCityCaptchaIpList
}

type DescribeCaptchaIpCityCaptchaCitie struct {
	Location goaliyun.String
	Lat      goaliyun.String
	Lng      goaliyun.String
	Pv       goaliyun.Integer
}

type DescribeCaptchaIpCityCaptchaIp struct {
	Ip    goaliyun.String
	Value goaliyun.Integer
}

type DescribeCaptchaIpCityCaptchaCitieList []DescribeCaptchaIpCityCaptchaCitie

func (list *DescribeCaptchaIpCityCaptchaCitieList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCaptchaIpCityCaptchaCitie)
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

type DescribeCaptchaIpCityCaptchaIpList []DescribeCaptchaIpCityCaptchaIp

func (list *DescribeCaptchaIpCityCaptchaIpList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCaptchaIpCityCaptchaIp)
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
