// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/danieloliveira079/kubebuilder-example/pkg/apis/operator/v1beta1"
	"github.com/danieloliveira079/kubebuilder-example/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type OperatorV1beta1Interface interface {
	RESTClient() rest.Interface
	MedusasGetter
}

// OperatorV1beta1Client is used to interact with features provided by the operator.octops.io group.
type OperatorV1beta1Client struct {
	restClient rest.Interface
}

func (c *OperatorV1beta1Client) Medusas(namespace string) MedusaInterface {
	return newMedusas(c, namespace)
}

// NewForConfig creates a new OperatorV1beta1Client for the given config.
func NewForConfig(c *rest.Config) (*OperatorV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &OperatorV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new OperatorV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *OperatorV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new OperatorV1beta1Client for the given RESTClient.
func New(c rest.Interface) *OperatorV1beta1Client {
	return &OperatorV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *OperatorV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
