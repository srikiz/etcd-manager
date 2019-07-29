/*
Copyright 2016 The Kubernetes Authors.

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

package do

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/digitalocean/godo"

	"github.com/golang/glog"

	"kope.io/etcd-manager/pkg/volumes"
)

// DOVolumes defines the aws volume implementation
type DOVolumes struct {

}

var _ volumes.Volumes = &DOVolumes{}

// NewDOVolumes returns a new aws volume provider
func NewDOVolumes(clusterName string, volumeTags []string, nameTag string) (*DOVolumes, error) {

	return a, nil
}

func (a *DOVolumes) describeInstance() (*ec2.Instance, error) {

	var instances []*ec2.Instance


	return instances[0], nil
}


func (a *DOVolumes) describeVolumes(request *ec2.DescribeVolumesInput) ([]*volumes.Volume, error) {
	var found []*volumes.Volume
	
	return found, nil
}

func (a *DOVolumes) FindVolumes() ([]*volumes.Volume, error) {
	return a.findVolumes(true)
}

func (a *DOVolumes) findVolumes(filterByAZ bool) ([]*volumes.Volume, error) {
	return nil
}

// FindMountedVolume implements Volumes::FindMountedVolume
func (a *DOVolumes) FindMountedVolume(volume *volumes.Volume) (string, error) {

	return "", nil
}

// assignDevice picks a hopefully unused device and reserves it for the volume attachment
func (a *DOVolumes) assignDevice(volumeID string) (string, error) {
	return "", fmt.Errorf("All devices in use")
}

// AttachVolume attaches the specified volume to this instance, returning the mountpoint & nil if successful
func (a *DOVolumes) AttachVolume(volume *volumes.Volume) error {

}

func (a *DOVolumes) MyIP() (string, error) {
	return nil, nil
}
