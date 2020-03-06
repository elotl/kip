/*
Copyright 2019 Elotl Inc.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	v1beta1 "github.com/elotl/cloud-instance-provider/pkg/apis/kip/v1beta1"
	scheme "github.com/elotl/cloud-instance-provider/pkg/k8sclient/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CellsGetter has a method to return a CellInterface.
// A group's client should implement this interface.
type CellsGetter interface {
	Cells() CellInterface
}

// CellInterface has methods to work with Cell resources.
type CellInterface interface {
	Create(*v1beta1.Cell) (*v1beta1.Cell, error)
	Update(*v1beta1.Cell) (*v1beta1.Cell, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.Cell, error)
	List(opts v1.ListOptions) (*v1beta1.CellList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Cell, err error)
	CellExpansion
}

// cells implements CellInterface
type cells struct {
	client rest.Interface
}

// newCells returns a Cells
func newCells(c *KiyotV1beta1Client) *cells {
	return &cells{
		client: c.RESTClient(),
	}
}

// Get takes name of the cell, and returns the corresponding cell object, and an error if there is any.
func (c *cells) Get(name string, options v1.GetOptions) (result *v1beta1.Cell, err error) {
	result = &v1beta1.Cell{}
	err = c.client.Get().
		Resource("cells").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Cells that match those selectors.
func (c *cells) List(opts v1.ListOptions) (result *v1beta1.CellList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.CellList{}
	err = c.client.Get().
		Resource("cells").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cells.
func (c *cells) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("cells").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cell and creates it.  Returns the server's representation of the cell, and an error, if there is any.
func (c *cells) Create(cell *v1beta1.Cell) (result *v1beta1.Cell, err error) {
	result = &v1beta1.Cell{}
	err = c.client.Post().
		Resource("cells").
		Body(cell).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cell and updates it. Returns the server's representation of the cell, and an error, if there is any.
func (c *cells) Update(cell *v1beta1.Cell) (result *v1beta1.Cell, err error) {
	result = &v1beta1.Cell{}
	err = c.client.Put().
		Resource("cells").
		Name(cell.Name).
		Body(cell).
		Do().
		Into(result)
	return
}

// Delete takes name of the cell and deletes it. Returns an error if one occurs.
func (c *cells) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("cells").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cells) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("cells").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cell.
func (c *cells) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Cell, err error) {
	result = &v1beta1.Cell{}
	err = c.client.Patch(pt).
		Resource("cells").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}