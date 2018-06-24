package ingress

type Upstreams struct {
	Upstreams []Upstream `json:"upstreams,omitempty"`
}

//Upstream holds the information extracted from Ingress that has passed for any reconcileation
type Upstream struct {
	Key  string `json:"key,omitempty"`
	Addr string `json:"addr,omitempty"`
}
