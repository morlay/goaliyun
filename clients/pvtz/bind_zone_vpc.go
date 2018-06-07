package pvtz

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type BindZoneVpcRequest struct {
	UserClientIp string               `position:"Query" name:"UserClientIp"`
	ZoneId       string               `position:"Query" name:"ZoneId"`
	Lang         string               `position:"Query" name:"Lang"`
	Vpcss        *BindZoneVpcVpcsList `position:"Query" type:"Repeated" name:"Vpcs"`
	RegionId     string               `position:"Query" name:"RegionId"`
}

func (req *BindZoneVpcRequest) Invoke(client goaliyun.Client) (*BindZoneVpcResponse, error) {
	resp := &BindZoneVpcResponse{}
	err := client.Request("pvtz", "BindZoneVpc", "2018-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type BindZoneVpcVpcs struct {
	RegionId string `name:"RegionId"`
	VpcId    string `name:"VpcId"`
}

type BindZoneVpcResponse struct {
	RequestId goaliyun.String
}

type BindZoneVpcVpcsList []BindZoneVpcVpcs

func (list *BindZoneVpcVpcsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]BindZoneVpcVpcs)
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
