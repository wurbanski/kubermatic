package fake

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/kubermatic/api/extensions"
	uapi "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/rest/fake"
)

func objBody(object interface{}) io.ReadCloser {
	output, err := json.MarshalIndent(object, "", "")
	if err != nil {
		panic(err)
	}
	return ioutil.NopCloser(bytes.NewReader([]byte(output)))
}

// ClientsetWithExtensions returns a fake extensions.Clientset interface which should only be used for testing
func ClientsetWithExtensions() extensions.Clientset {
	fakeClient := &fake.RESTClient{
		APIRegistry:          api.Registry,
		NegotiatedSerializer: api.Codecs,
		Resp: &http.Response{
			StatusCode: 200,
			Body:       objBody(&uapi.APIVersions{Versions: []string{"version1", api.Registry.GroupOrDie(api.GroupName).GroupVersion.String()}}),
		},
		Client: fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
			header := http.Header{}
			header.Set("Content-Type", runtime.ContentTypeJSON)
			return &http.Response{StatusCode: 200, Header: header, Body: objBody(&uapi.APIVersions{Versions: []string{"version1", api.Registry.GroupOrDie(api.GroupName).GroupVersion.String()}})}, nil
		}),
	}

	return &extensions.WrappedClientset{
		Client: fakeClient,
	}
}
