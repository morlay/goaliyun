package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeResourceDiagnosisRequest struct {
	EndTime      string `position:"Query" name:"EndTime"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	StartTime    string `position:"Query" name:"StartTime"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeResourceDiagnosisRequest) Invoke(client goaliyun.Client) (*DescribeResourceDiagnosisResponse, error) {
	resp := &DescribeResourceDiagnosisResponse{}
	err := client.Request("rds", "DescribeResourceDiagnosis", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeResourceDiagnosisResponse struct {
	RequestId  goaliyun.String
	StartTime  goaliyun.String
	EndTime    goaliyun.String
	CPU        DescribeResourceDiagnosisCPUList
	Memory     DescribeResourceDiagnosisMemoryList
	Storage    DescribeResourceDiagnosisStorageList
	IOPS       DescribeResourceDiagnosisIOPList
	Connection DescribeResourceDiagnosisConnectionList
}

type DescribeResourceDiagnosisCPUList []goaliyun.String

func (list *DescribeResourceDiagnosisCPUList) UnmarshalJSON(data []byte) error {
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

type DescribeResourceDiagnosisMemoryList []goaliyun.String

func (list *DescribeResourceDiagnosisMemoryList) UnmarshalJSON(data []byte) error {
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

type DescribeResourceDiagnosisStorageList []goaliyun.String

func (list *DescribeResourceDiagnosisStorageList) UnmarshalJSON(data []byte) error {
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

type DescribeResourceDiagnosisIOPList []goaliyun.String

func (list *DescribeResourceDiagnosisIOPList) UnmarshalJSON(data []byte) error {
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

type DescribeResourceDiagnosisConnectionList []goaliyun.String

func (list *DescribeResourceDiagnosisConnectionList) UnmarshalJSON(data []byte) error {
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
