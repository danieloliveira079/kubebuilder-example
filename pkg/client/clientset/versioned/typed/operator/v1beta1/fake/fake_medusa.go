// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/danieloliveira079/kubebuilder-example/pkg/apis/operator/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMedusas implements MedusaInterface
type FakeMedusas struct {
	Fake *FakeOperatorV1beta1
	ns   string
}

var medusasResource = schema.GroupVersionResource{Group: "operator.octops.io", Version: "v1beta1", Resource: "medusas"}

var medusasKind = schema.GroupVersionKind{Group: "operator.octops.io", Version: "v1beta1", Kind: "Medusa"}

// Get takes name of the medusa, and returns the corresponding medusa object, and an error if there is any.
func (c *FakeMedusas) Get(name string, options v1.GetOptions) (result *v1beta1.Medusa, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(medusasResource, c.ns, name), &v1beta1.Medusa{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Medusa), err
}

// List takes label and field selectors, and returns the list of Medusas that match those selectors.
func (c *FakeMedusas) List(opts v1.ListOptions) (result *v1beta1.MedusaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(medusasResource, medusasKind, c.ns, opts), &v1beta1.MedusaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.MedusaList{}
	for _, item := range obj.(*v1beta1.MedusaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested medusas.
func (c *FakeMedusas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(medusasResource, c.ns, opts))

}

// Create takes the representation of a medusa and creates it.  Returns the server's representation of the medusa, and an error, if there is any.
func (c *FakeMedusas) Create(medusa *v1beta1.Medusa) (result *v1beta1.Medusa, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(medusasResource, c.ns, medusa), &v1beta1.Medusa{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Medusa), err
}

// Update takes the representation of a medusa and updates it. Returns the server's representation of the medusa, and an error, if there is any.
func (c *FakeMedusas) Update(medusa *v1beta1.Medusa) (result *v1beta1.Medusa, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(medusasResource, c.ns, medusa), &v1beta1.Medusa{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Medusa), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMedusas) UpdateStatus(medusa *v1beta1.Medusa) (*v1beta1.Medusa, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(medusasResource, "status", c.ns, medusa), &v1beta1.Medusa{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Medusa), err
}

// Delete takes name of the medusa and deletes it. Returns an error if one occurs.
func (c *FakeMedusas) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(medusasResource, c.ns, name), &v1beta1.Medusa{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMedusas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(medusasResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.MedusaList{})
	return err
}

// Patch applies the patch and returns the patched medusa.
func (c *FakeMedusas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Medusa, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(medusasResource, c.ns, name, data, subresources...), &v1beta1.Medusa{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Medusa), err
}
