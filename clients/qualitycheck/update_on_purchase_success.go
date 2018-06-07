package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type UpdateOnPurchaseSuccessRequest struct {
	Data     string `position:"Query" name:"Data"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UpdateOnPurchaseSuccessRequest) Invoke(client goaliyun.Client) (*UpdateOnPurchaseSuccessResponse, error) {
	resp := &UpdateOnPurchaseSuccessResponse{}
	err := client.Request("qualitycheck", "UpdateOnPurchaseSuccess", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateOnPurchaseSuccessResponse struct {
	RequestId goaliyun.String
	Data      goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
}
