package market_inner

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type InnerReviewHistoryImageRequest struct {
	ProductCode  string `position:"Query" name:"ProductCode"`
	ReviewType   string `position:"Query" name:"ReviewType"`
	ImageVersion string `position:"Query" name:"ImageVersion"`
	ImageNo      string `position:"Query" name:"ImageNo"`
	PageIndex    int64  `position:"Query" name:"PageIndex"`
	RegionNo     string `position:"Query" name:"RegionNo"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *InnerReviewHistoryImageRequest) Invoke(client goaliyun.Client) (*InnerReviewHistoryImageResponse, error) {
	resp := &InnerReviewHistoryImageResponse{}
	err := client.Request("market-inner", "InnerReviewHistoryImage", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InnerReviewHistoryImageResponse struct {
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	IsSuccess  bool
	ImageList  InnerReviewHistoryImageImageList
}

type InnerReviewHistoryImageImage struct {
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
	InstanceStatus  goaliyun.String
	ImageCreateTime goaliyun.String
}

type InnerReviewHistoryImageImageList []InnerReviewHistoryImageImage

func (list *InnerReviewHistoryImageImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerReviewHistoryImageImage)
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
