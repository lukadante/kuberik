// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kuberik/kuberik/pkg/apis/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEvents implements EventInterface
type FakeEvents struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var eventsResource = schema.GroupVersionResource{Group: "core.kuberik.io", Version: "v1alpha1", Resource: "events"}

var eventsKind = schema.GroupVersionKind{Group: "core.kuberik.io", Version: "v1alpha1", Kind: "Event"}

// Get takes name of the event, and returns the corresponding event object, and an error if there is any.
func (c *FakeEvents) Get(name string, options v1.GetOptions) (result *v1alpha1.Event, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eventsResource, c.ns, name), &v1alpha1.Event{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Event), err
}

// List takes label and field selectors, and returns the list of Events that match those selectors.
func (c *FakeEvents) List(opts v1.ListOptions) (result *v1alpha1.EventList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eventsResource, eventsKind, c.ns, opts), &v1alpha1.EventList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EventList{ListMeta: obj.(*v1alpha1.EventList).ListMeta}
	for _, item := range obj.(*v1alpha1.EventList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested events.
func (c *FakeEvents) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eventsResource, c.ns, opts))

}

// Create takes the representation of a event and creates it.  Returns the server's representation of the event, and an error, if there is any.
func (c *FakeEvents) Create(event *v1alpha1.Event) (result *v1alpha1.Event, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eventsResource, c.ns, event), &v1alpha1.Event{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Event), err
}

// Update takes the representation of a event and updates it. Returns the server's representation of the event, and an error, if there is any.
func (c *FakeEvents) Update(event *v1alpha1.Event) (result *v1alpha1.Event, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eventsResource, c.ns, event), &v1alpha1.Event{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Event), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEvents) UpdateStatus(event *v1alpha1.Event) (*v1alpha1.Event, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(eventsResource, "status", c.ns, event), &v1alpha1.Event{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Event), err
}

// Delete takes name of the event and deletes it. Returns an error if one occurs.
func (c *FakeEvents) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eventsResource, c.ns, name), &v1alpha1.Event{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEvents) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eventsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EventList{})
	return err
}

// Patch applies the patch and returns the patched event.
func (c *FakeEvents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Event, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eventsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Event{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Event), err
}
