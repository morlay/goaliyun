package ram

import (
	"github.com/morlay/goaliyun"
)

type GetSecurityPreferenceRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetSecurityPreferenceRequest) Invoke(client goaliyun.Client) (*GetSecurityPreferenceResponse, error) {
	resp := &GetSecurityPreferenceResponse{}
	err := client.Request("ram", "GetSecurityPreference", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetSecurityPreferenceResponse struct {
	RequestId          goaliyun.String
	SecurityPreference GetSecurityPreferenceSecurityPreference
}

type GetSecurityPreferenceSecurityPreference struct {
	LoginProfilePreference GetSecurityPreferenceLoginProfilePreference
	AccessKeyPreference    GetSecurityPreferenceAccessKeyPreference
	PublicKeyPreference    GetSecurityPreferencePublicKeyPreference
	MFAPreference          GetSecurityPreferenceMFAPreference
}

type GetSecurityPreferenceLoginProfilePreference struct {
	EnableSaveMFATicket       bool
	AllowUserToChangePassword bool
}

type GetSecurityPreferenceAccessKeyPreference struct {
	AllowUserToManageAccessKeys bool
}

type GetSecurityPreferencePublicKeyPreference struct {
	AllowUserToManagePublicKeys bool
}

type GetSecurityPreferenceMFAPreference struct {
	AllowUserToManageMFADevices bool
}
