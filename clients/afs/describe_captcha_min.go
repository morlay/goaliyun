package afs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCaptchaMinRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	ConfigName      string `position:"Query" name:"ConfigName"`
	Time            string `position:"Query" name:"Time"`
	Type            string `position:"Query" name:"Type"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeCaptchaMinRequest) Invoke(client goaliyun.Client) (*DescribeCaptchaMinResponse, error) {
	resp := &DescribeCaptchaMinResponse{}
	err := client.Request("afs", "DescribeCaptchaMin", "2018-01-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCaptchaMinResponse struct {
	RequestId   goaliyun.String
	BizCode     goaliyun.String
	HasData     bool
	CaptchaMins DescribeCaptchaMinCaptchaMinList
}

type DescribeCaptchaMinCaptchaMin struct {
	Time         goaliyun.String
	Pass         goaliyun.String
	Interception goaliyun.String
}

type DescribeCaptchaMinCaptchaMinList []DescribeCaptchaMinCaptchaMin

func (list *DescribeCaptchaMinCaptchaMinList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCaptchaMinCaptchaMin)
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
