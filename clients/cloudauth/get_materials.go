package cloudauth

import (
	"github.com/morlay/goaliyun"
)

type GetMaterialsRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Biz             string `position:"Query" name:"Biz"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	TicketId        string `position:"Query" name:"TicketId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetMaterialsRequest) Invoke(client goaliyun.Client) (*GetMaterialsResponse, error) {
	resp := &GetMaterialsResponse{}
	err := client.Request("cloudauth", "GetMaterials", "2018-05-04", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetMaterialsResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      GetMaterialsData
}

type GetMaterialsData struct {
	Name                 goaliyun.String
	IdentificationNumber goaliyun.String
	IdCardType           goaliyun.String
	IdCardStartDate      goaliyun.String
	IdCardExpiry         goaliyun.String
	Address              goaliyun.String
	Sex                  goaliyun.String
	IdCardFrontPic       goaliyun.String
	IdCardBackPic        goaliyun.String
	FacePic              goaliyun.String
	EthnicGroup          goaliyun.String
}
