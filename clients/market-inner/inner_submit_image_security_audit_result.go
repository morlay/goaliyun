package market_inner

import (
	"github.com/morlay/goaliyun"
)

type InnerSubmitImageSecurityAuditResultRequest struct {
	ProductCode  string `position:"Query" name:"ProductCode"`
	Pass         int64  `position:"Query" name:"Pass"`
	ImageVersion string `position:"Query" name:"ImageVersion"`
	Channel      string `position:"Query" name:"Channel"`
	ImageNo      string `position:"Query" name:"ImageNo"`
	Remark       string `position:"Query" name:"Remark"`
	Operator     string `position:"Query" name:"Operator"`
	RegionNo     string `position:"Query" name:"RegionNo"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *InnerSubmitImageSecurityAuditResultRequest) Invoke(client goaliyun.Client) (*InnerSubmitImageSecurityAuditResultResponse, error) {
	resp := &InnerSubmitImageSecurityAuditResultResponse{}
	err := client.Request("market-inner", "InnerSubmitImageSecurityAuditResult", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InnerSubmitImageSecurityAuditResultResponse struct {
	RequestId goaliyun.String
	Success   bool
}
