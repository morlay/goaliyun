package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstanceHAConfigRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstanceHAConfigRequest) Invoke(client goaliyun.Client) (*DescribeDBInstanceHAConfigResponse, error) {
	resp := &DescribeDBInstanceHAConfigResponse{}
	err := client.Request("rds", "DescribeDBInstanceHAConfig", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstanceHAConfigResponse struct {
	RequestId         goaliyun.String
	DBInstanceId      goaliyun.String
	SyncMode          goaliyun.String
	HAMode            goaliyun.String
	HostInstanceInfos DescribeDBInstanceHAConfigNodeInfoList
}

type DescribeDBInstanceHAConfigNodeInfo struct {
	NodeId       goaliyun.String
	RegionId     goaliyun.String
	LogSyncTime  goaliyun.String
	DataSyncTime goaliyun.String
	NodeType     goaliyun.String
	ZoneId       goaliyun.String
	SyncStatus   goaliyun.String
}

type DescribeDBInstanceHAConfigNodeInfoList []DescribeDBInstanceHAConfigNodeInfo

func (list *DescribeDBInstanceHAConfigNodeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceHAConfigNodeInfo)
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
