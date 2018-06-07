package emr

import (
	"github.com/morlay/goaliyun"
)

type DecryptBizIdRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	Type            string `position:"Query" name:"Type"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DecryptBizIdRequest) Invoke(client goaliyun.Client) (*DecryptBizIdResponse, error) {
	resp := &DecryptBizIdResponse{}
	err := client.Request("emr", "DecryptBizId", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DecryptBizIdResponse struct {
	RequestId goaliyun.String
	BizId     goaliyun.String
}
