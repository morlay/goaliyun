package signature

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tt := assert.New(t)

	rt := &Credential{}
	rt.AccessKey = "testid"
	rt.SecretKey = "testsecret"

	str := "GET&%2F&AccessKeyId%3Dtestid%26Action%3DDescribeRegions%26Format%3DXML%26RegionId%3Dregion1%26SignatureMethod%3DHMAC-SHA1%26SignatureNonce%3DNwDAxvLU6tFE0DVb%26SignatureVersion%3D1.0%26TimeStamp%3D2012-12-26T10%253A33%253A56Z%26Version%3D2014-05-26"

	tt.Equal("K9fCVP6Jrklpd3rLYKh1pfrrFNo=", rt.sign(str, "&"))
}
