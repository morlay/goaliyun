package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type UpdateDomainToDomainGroupRequest struct {
	DataSource    int64                                    `position:"Query" name:"DataSource"`
	UserClientIp  string                                   `position:"Query" name:"UserClientIp"`
	FileToUpload  string                                   `position:"Body" name:"FileToUpload"`
	DomainNames   *UpdateDomainToDomainGroupDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	Replace       string                                   `position:"Query" name:"Replace"`
	Lang          string                                   `position:"Query" name:"Lang"`
	DomainGroupId int64                                    `position:"Query" name:"DomainGroupId"`
	RegionId      string                                   `position:"Query" name:"RegionId"`
}

func (req *UpdateDomainToDomainGroupRequest) Invoke(client goaliyun.Client) (*UpdateDomainToDomainGroupResponse, error) {
	resp := &UpdateDomainToDomainGroupResponse{}
	err := client.Request("domain", "UpdateDomainToDomainGroup", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateDomainToDomainGroupResponse struct {
	RequestId goaliyun.String
}

type UpdateDomainToDomainGroupDomainNameList []string

func (list *UpdateDomainToDomainGroupDomainNameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
