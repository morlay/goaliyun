package phoenixsp_inner

import (
	"github.com/morlay/goaliyun"
)

type GetValidProdSubListRequest struct {
	Filter   string `position:"Query" name:"Filter"`
	AliUid   int64  `position:"Query" name:"AliUid"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetValidProdSubListRequest) Invoke(client goaliyun.Client) (*GetValidProdSubListResponse, error) {
	resp := &GetValidProdSubListResponse{}
	err := client.Request("phoenixsp-inner", "GetValidProdSubList", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetValidProdSubListResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
	Success   bool
	Data      goaliyun.String
}
