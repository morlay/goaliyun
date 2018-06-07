package market

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLicenseRequest struct {
	LicenseCode string `position:"Query" name:"LicenseCode"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *DescribeLicenseRequest) Invoke(client goaliyun.Client) (*DescribeLicenseResponse, error) {
	resp := &DescribeLicenseResponse{}
	err := client.Request("market", "DescribeLicense", "2015-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLicenseResponse struct {
	RequestId goaliyun.String
	License   DescribeLicenseLicense
}

type DescribeLicenseLicense struct {
	LicenseStatus goaliyun.String
	LicenseCode   goaliyun.String
	InstanceId    goaliyun.String
	CreateTime    goaliyun.String
	ExpiredTime   goaliyun.String
	ActivateTime  goaliyun.String
	ProductSkuId  goaliyun.String
	ProductCode   goaliyun.String
	ProductName   goaliyun.String
	ExtendArray   DescribeLicenseLicenseAttributeList
	ExtendInfo    DescribeLicenseExtendInfo
}

type DescribeLicenseLicenseAttribute struct {
	Code  goaliyun.String
	Value goaliyun.String
}

type DescribeLicenseExtendInfo struct {
	AliUid          goaliyun.Integer
	Email           goaliyun.String
	Mobile          goaliyun.String
	AccountQuantity goaliyun.Integer
}

type DescribeLicenseLicenseAttributeList []DescribeLicenseLicenseAttribute

func (list *DescribeLicenseLicenseAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLicenseLicenseAttribute)
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
