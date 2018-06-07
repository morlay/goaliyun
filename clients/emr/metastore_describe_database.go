package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type MetastoreDescribeDatabaseRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	DbName          string `position:"Query" name:"DbName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *MetastoreDescribeDatabaseRequest) Invoke(client goaliyun.Client) (*MetastoreDescribeDatabaseResponse, error) {
	resp := &MetastoreDescribeDatabaseResponse{}
	err := client.Request("emr", "MetastoreDescribeDatabase", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MetastoreDescribeDatabaseResponse struct {
	RequestId   goaliyun.String
	DbName      goaliyun.String
	Description goaliyun.String
	LocationUri goaliyun.String
	Parameters  MetastoreDescribeDatabaseParameterList
}

type MetastoreDescribeDatabaseParameter struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type MetastoreDescribeDatabaseParameterList []MetastoreDescribeDatabaseParameter

func (list *MetastoreDescribeDatabaseParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MetastoreDescribeDatabaseParameter)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
