package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ListApAssetRequest struct {
	OrderCol       string `position:"Query" name:"OrderCol"`
	SearchName     string `position:"Query" name:"SearchName"`
	SearchSerialNo string `position:"Query" name:"SearchSerialNo"`
	Length         int64  `position:"Query" name:"Length"`
	PageIndex      int64  `position:"Query" name:"PageIndex"`
	SearchMac      string `position:"Query" name:"SearchMac"`
	OrderDir       string `position:"Query" name:"OrderDir"`
	SearchModel    string `position:"Query" name:"SearchModel"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *ListApAssetRequest) Invoke(client goaliyun.Client) (*ListApAssetResponse, error) {
	resp := &ListApAssetResponse{}
	err := client.Request("cloudwf", "ListApAsset", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListApAssetResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
