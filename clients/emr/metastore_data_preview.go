package emr

import (
	"github.com/morlay/goaliyun"
)

type MetastoreDataPreviewRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	DbName          string `position:"Query" name:"DbName"`
	TableName       string `position:"Query" name:"TableName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *MetastoreDataPreviewRequest) Invoke(client goaliyun.Client) (*MetastoreDataPreviewResponse, error) {
	resp := &MetastoreDataPreviewResponse{}
	err := client.Request("emr", "MetastoreDataPreview", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MetastoreDataPreviewResponse struct {
	RequestId goaliyun.String
	Samples   goaliyun.String
}
