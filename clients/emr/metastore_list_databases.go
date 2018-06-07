package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type MetastoreListDatabasesRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *MetastoreListDatabasesRequest) Invoke(client goaliyun.Client) (*MetastoreListDatabasesResponse, error) {
	resp := &MetastoreListDatabasesResponse{}
	err := client.Request("emr", "MetastoreListDatabases", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MetastoreListDatabasesResponse struct {
	RequestId   goaliyun.String
	Description goaliyun.String
	DbNames     MetastoreListDatabasesDbNameList
}

type MetastoreListDatabasesDbNameList []goaliyun.String

func (list *MetastoreListDatabasesDbNameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
