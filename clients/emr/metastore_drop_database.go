package emr

import (
	"github.com/morlay/goaliyun"
)

type MetastoreDropDatabaseRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	DbName          string `position:"Query" name:"DbName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *MetastoreDropDatabaseRequest) Invoke(client goaliyun.Client) (*MetastoreDropDatabaseResponse, error) {
	resp := &MetastoreDropDatabaseResponse{}
	err := client.Request("emr", "MetastoreDropDatabase", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MetastoreDropDatabaseResponse struct {
	RequestId goaliyun.String
}
