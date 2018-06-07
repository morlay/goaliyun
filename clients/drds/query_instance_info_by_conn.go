package drds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryInstanceInfoByConnRequest struct {
	Port     int64  `position:"Query" name:"Port"`
	Host     string `position:"Query" name:"Host"`
	UserName string `position:"Query" name:"UserName"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *QueryInstanceInfoByConnRequest) Invoke(client goaliyun.Client) (*QueryInstanceInfoByConnResponse, error) {
	resp := &QueryInstanceInfoByConnResponse{}
	err := client.Request("drds", "QueryInstanceInfoByConn", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryInstanceInfoByConnResponse struct {
	RequestId goaliyun.String
	Success   bool
	Data      QueryInstanceInfoByConnData
}

type QueryInstanceInfoByConnData struct {
	DrdsInstanceId     goaliyun.String
	Type               goaliyun.String
	RegionId           goaliyun.String
	ZoneId             goaliyun.String
	Description        goaliyun.String
	NetworkType        goaliyun.String
	Status             goaliyun.String
	CreateTime         goaliyun.Integer
	Version            goaliyun.Integer
	Specification      goaliyun.String
	SpecTypeId         goaliyun.String
	SpecTypeName       goaliyun.String
	VpcCloudInstanceId goaliyun.String
	Vips               QueryInstanceInfoByConnVipList
}

type QueryInstanceInfoByConnVip struct {
	IP        goaliyun.String
	Port      goaliyun.String
	Type      goaliyun.String
	VpcId     goaliyun.String
	VswitchId goaliyun.String
}

type QueryInstanceInfoByConnVipList []QueryInstanceInfoByConnVip

func (list *QueryInstanceInfoByConnVipList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryInstanceInfoByConnVip)
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
