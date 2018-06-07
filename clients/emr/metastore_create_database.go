package emr

import (
	"github.com/morlay/goaliyun"
)

type MetastoreCreateDatabaseRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	DbName          string `position:"Query" name:"DbName"`
	Description     string `position:"Query" name:"Description"`
	LocationUri     string `position:"Query" name:"LocationUri"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *MetastoreCreateDatabaseRequest) Invoke(client goaliyun.Client) (*MetastoreCreateDatabaseResponse, error) {
	resp := &MetastoreCreateDatabaseResponse{}
	err := client.Request("emr", "MetastoreCreateDatabase", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MetastoreCreateDatabaseResponse struct {
	RequestId goaliyun.String
}
