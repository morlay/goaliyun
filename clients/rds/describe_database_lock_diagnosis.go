package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDatabaseLockDiagnosisRequest struct {
	EndTime      string `position:"Query" name:"EndTime"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	StartTime    string `position:"Query" name:"StartTime"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeDatabaseLockDiagnosisRequest) Invoke(client goaliyun.Client) (*DescribeDatabaseLockDiagnosisResponse, error) {
	resp := &DescribeDatabaseLockDiagnosisResponse{}
	err := client.Request("rds", "DescribeDatabaseLockDiagnosis", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDatabaseLockDiagnosisResponse struct {
	RequestId    goaliyun.String
	DeadLockList DescribeDatabaseLockDiagnosisDeadLockListList
}

type DescribeDatabaseLockDiagnosisDeadLockListList []goaliyun.String

func (list *DescribeDatabaseLockDiagnosisDeadLockListList) UnmarshalJSON(data []byte) error {
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
