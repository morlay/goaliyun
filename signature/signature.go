package signature

import (
	"crypto/hmac"
	crypto_rand "crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/morlay/goaliyun/transform"
)

type Credential struct {
	AccessKey string
	SecretKey string
}

func (c *Credential) sign(stringToSignature string, suffix string) string {
	key := []byte(c.SecretKey + suffix)
	hmac := hmac.New(sha1.New, key)
	hmac.Write([]byte(stringToSignature))
	signedBytes := hmac.Sum(nil)
	signedString := base64.StdEncoding.EncodeToString(signedBytes)
	return signedString
}

func NewAutoSignTransport(credential *Credential) transform.Transport {
	return func(rt http.RoundTripper) http.RoundTripper {
		return &AutoSignRoundTripper{
			Credential:       credential,
			NextRoundTripper: rt,
		}
	}
}

type AutoSignRoundTripper struct {
	*Credential
	NextRoundTripper http.RoundTripper
}

var _ interface {
	http.RoundTripper
} = (*AutoSignRoundTripper)(nil)

func (t *AutoSignRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	t.PatchSignForRpc(request)
	return t.NextRoundTripper.RoundTrip(request)
}

func (t *AutoSignRoundTripper) PatchSignForRpc(request *http.Request) {
	signMethod := "HMAC-SHA1"

	query := request.URL.Query()

	query.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	query.Set("SignatureMethod", signMethod)
	query.Set("SignatureType", "")
	query.Set("SignatureVersion", "1.0")
	query.Set("SignatureNonce", randStr())
	query.Set("AccessKeyId", t.AccessKey)

	query.Del("Signature")

	canonicalizedQueryString := percentReplace(query.Encode())
	stringToSign := request.Method + "&%2F&" + url.QueryEscape(canonicalizedQueryString)

	query.Set("Signature", t.sign(stringToSign, "&"))

	request.URL.RawQuery = query.Encode()
}

func percentReplace(str string) string {
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	return str
}

const dictionary = "_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func randStr() string {
	b := make([]byte, 32)
	l := len(dictionary)

	_, err := crypto_rand.Read(b)

	if err != nil {
		// fail back to insecure rand
		rand.Seed(time.Now().UnixNano())
		for i := range b {
			b[i] = dictionary[rand.Int()%l]
		}
	} else {
		for i, v := range b {
			b[i] = dictionary[v%byte(l)]
		}
	}

	return string(b)
}
