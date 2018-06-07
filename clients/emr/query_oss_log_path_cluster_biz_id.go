package emr

import (
	"github.com/morlay/goaliyun"
)

type QueryOssLogPathClusterBizIdRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *QueryOssLogPathClusterBizIdRequest) Invoke(client goaliyun.Client) (*QueryOssLogPathClusterBizIdResponse, error) {
	resp := &QueryOssLogPathClusterBizIdResponse{}
	err := client.Request("emr", "QueryOssLogPathClusterBizId", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryOssLogPathClusterBizIdResponse struct {
	RequestId              goaliyun.String
	OssLogPathClusterBizId goaliyun.String
}
