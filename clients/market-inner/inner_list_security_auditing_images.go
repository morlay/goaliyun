package market_inner

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type InnerListSecurityAuditingImagesRequest struct {
	Channel   int64  `position:"Query" name:"Channel"`
	PageIndex int64  `position:"Query" name:"PageIndex"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *InnerListSecurityAuditingImagesRequest) Invoke(client goaliyun.Client) (*InnerListSecurityAuditingImagesResponse, error) {
	resp := &InnerListSecurityAuditingImagesResponse{}
	err := client.Request("market-inner", "InnerListSecurityAuditingImages", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InnerListSecurityAuditingImagesResponse struct {
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	ImageList  InnerListSecurityAuditingImagesImageList
}

type InnerListSecurityAuditingImagesImage struct {
	RegionNo        goaliyun.String
	ProductCode     goaliyun.String
	ImageNo         goaliyun.String
	ImageVersion    goaliyun.String
	SupplierId      goaliyun.Integer
	ProductName     goaliyun.String
	InstanceId      goaliyun.String
	InstanceAddress goaliyun.String
	OsType          goaliyun.String
	UserName        goaliyun.String
	SupplierName    goaliyun.String
	Password        goaliyun.String
	OsKind          goaliyun.String
	OsBit           goaliyun.Integer
}

type InnerListSecurityAuditingImagesImageList []InnerListSecurityAuditingImagesImage

func (list *InnerListSecurityAuditingImagesImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerListSecurityAuditingImagesImage)
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
