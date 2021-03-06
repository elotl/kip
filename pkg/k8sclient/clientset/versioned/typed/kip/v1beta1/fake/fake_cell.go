/*
Copyright 2020 Elotl Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/elotl/kip/pkg/apis/kip/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCells implements CellInterface
type FakeCells struct {
	Fake *FakeKiyotV1beta1
}

var cellsResource = schema.GroupVersionResource{Group: "kiyot.elotl.co", Version: "v1beta1", Resource: "cells"}

var cellsKind = schema.GroupVersionKind{Group: "kiyot.elotl.co", Version: "v1beta1", Kind: "Cell"}

// Get takes name of the cell, and returns the corresponding cell object, and an error if there is any.
func (c *FakeCells) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Cell, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(cellsResource, name), &v1beta1.Cell{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Cell), err
}

// List takes label and field selectors, and returns the list of Cells that match those selectors.
func (c *FakeCells) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CellList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(cellsResource, cellsKind, opts), &v1beta1.CellList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CellList{ListMeta: obj.(*v1beta1.CellList).ListMeta}
	for _, item := range obj.(*v1beta1.CellList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cells.
func (c *FakeCells) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(cellsResource, opts))
}

// Create takes the representation of a cell and creates it.  Returns the server's representation of the cell, and an error, if there is any.
func (c *FakeCells) Create(ctx context.Context, cell *v1beta1.Cell, opts v1.CreateOptions) (result *v1beta1.Cell, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(cellsResource, cell), &v1beta1.Cell{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Cell), err
}

// Update takes the representation of a cell and updates it. Returns the server's representation of the cell, and an error, if there is any.
func (c *FakeCells) Update(ctx context.Context, cell *v1beta1.Cell, opts v1.UpdateOptions) (result *v1beta1.Cell, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(cellsResource, cell), &v1beta1.Cell{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Cell), err
}

// Delete takes name of the cell and deletes it. Returns an error if one occurs.
func (c *FakeCells) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(cellsResource, name), &v1beta1.Cell{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCells) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(cellsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.CellList{})
	return err
}

// Patch applies the patch and returns the patched cell.
func (c *FakeCells) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Cell, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(cellsResource, name, pt, data, subresources...), &v1beta1.Cell{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Cell), err
}
