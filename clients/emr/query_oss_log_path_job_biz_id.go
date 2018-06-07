package emr

import (
	"github.com/morlay/goaliyun"
)

type QueryOssLogPathJobBizIdRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	WorkNodeExecId  string `position:"Query" name:"WorkNodeExecId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *QueryOssLogPathJobBizIdRequest) Invoke(client goaliyun.Client) (*QueryOssLogPathJobBizIdResponse, error) {
	resp := &QueryOssLogPathJobBizIdResponse{}
	err := client.Request("emr", "QueryOssLogPathJobBizId", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryOssLogPathJobBizIdResponse struct {
	RequestId          goaliyun.String
	OssLogPathJobBizId goaliyun.String
}
