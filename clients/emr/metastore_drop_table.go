package emr

import (
	"github.com/morlay/goaliyun"
)

type MetastoreDropTableRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	DbName          string `position:"Query" name:"DbName"`
	TableName       string `position:"Query" name:"TableName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *MetastoreDropTableRequest) Invoke(client goaliyun.Client) (*MetastoreDropTableResponse, error) {
	resp := &MetastoreDropTableResponse{}
	err := client.Request("emr", "MetastoreDropTable", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MetastoreDropTableResponse struct {
	RequestId goaliyun.String
}
