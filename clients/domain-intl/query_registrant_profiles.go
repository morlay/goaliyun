package domain_intl

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryRegistrantProfilesRequest struct {
	RegistrantOrganization   string `position:"Query" name:"RegistrantOrganization"`
	UserClientIp             string `position:"Query" name:"UserClientIp"`
	RegistrantProfileId      int64  `position:"Query" name:"RegistrantProfileId"`
	PageSize                 int64  `position:"Query" name:"PageSize"`
	Lang                     string `position:"Query" name:"Lang"`
	PageNum                  int64  `position:"Query" name:"PageNum"`
	DefaultRegistrantProfile string `position:"Query" name:"DefaultRegistrantProfile"`
	RegionId                 string `position:"Query" name:"RegionId"`
}

func (req *QueryRegistrantProfilesRequest) Invoke(client goaliyun.Client) (*QueryRegistrantProfilesResponse, error) {
	resp := &QueryRegistrantProfilesResponse{}
	err := client.Request("domain-intl", "QueryRegistrantProfiles", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryRegistrantProfilesResponse struct {
	RequestId          goaliyun.String
	TotalItemNum       goaliyun.Integer
	CurrentPageNum     goaliyun.Integer
	TotalPageNum       goaliyun.Integer
	PageSize           goaliyun.Integer
	PrePage            bool
	NextPage           bool
	RegistrantProfiles QueryRegistrantProfilesRegistrantProfileList
}

type QueryRegistrantProfilesRegistrantProfile struct {
	RegistrantProfileId      goaliyun.Integer
	CreateTime               goaliyun.String
	UpdateTime               goaliyun.String
	DefaultRegistrantProfile bool
	RegistrantName           goaliyun.String
	RegistrantOrganization   goaliyun.String
	Country                  goaliyun.String
	Province                 goaliyun.String
	City                     goaliyun.String
	Address                  goaliyun.String
	Email                    goaliyun.String
	PostalCode               goaliyun.String
	TelArea                  goaliyun.String
	Telephone                goaliyun.String
	TelExt                   goaliyun.String
	EmailVerificationStatus  goaliyun.Integer
}

type QueryRegistrantProfilesRegistrantProfileList []QueryRegistrantProfilesRegistrantProfile

func (list *QueryRegistrantProfilesRegistrantProfileList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryRegistrantProfilesRegistrantProfile)
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
