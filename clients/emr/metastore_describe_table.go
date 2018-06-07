package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type MetastoreDescribeTableRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	DbName          string `position:"Query" name:"DbName"`
	TableName       string `position:"Query" name:"TableName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *MetastoreDescribeTableRequest) Invoke(client goaliyun.Client) (*MetastoreDescribeTableResponse, error) {
	resp := &MetastoreDescribeTableResponse{}
	err := client.Request("emr", "MetastoreDescribeTable", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MetastoreDescribeTableResponse struct {
	RequestId             goaliyun.String
	CreateTime            goaliyun.Integer
	LastAccessTime        goaliyun.Integer
	LocationUri           goaliyun.String
	InputFormat           goaliyun.String
	OutputFormat          goaliyun.String
	Compressed            bool
	SerializationLib      goaliyun.String
	TableName             goaliyun.String
	DbName                goaliyun.String
	Owner                 goaliyun.String
	TableType             goaliyun.String
	Columns               MetastoreDescribeTableColumnList
	SerdeParameters       MetastoreDescribeTableSerdeParameterList
	StorageDescParameters MetastoreDescribeTableStorageDescParameterList
}

type MetastoreDescribeTableColumn struct {
	Name    goaliyun.String
	Type    goaliyun.String
	Comment goaliyun.String
}

type MetastoreDescribeTableSerdeParameter struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type MetastoreDescribeTableStorageDescParameter struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type MetastoreDescribeTableColumnList []MetastoreDescribeTableColumn

func (list *MetastoreDescribeTableColumnList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MetastoreDescribeTableColumn)
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

type MetastoreDescribeTableSerdeParameterList []MetastoreDescribeTableSerdeParameter

func (list *MetastoreDescribeTableSerdeParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MetastoreDescribeTableSerdeParameter)
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

type MetastoreDescribeTableStorageDescParameterList []MetastoreDescribeTableStorageDescParameter

func (list *MetastoreDescribeTableStorageDescParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MetastoreDescribeTableStorageDescParameter)
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
