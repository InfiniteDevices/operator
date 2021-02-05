// +build !ignore_autogenerated

/*
Copyright 2019 infinimesh, inc.

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
// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Platform) DeepCopyInto(out *Platform) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Platform.
func (in *Platform) DeepCopy() *Platform {
	if in == nil {
		return nil
	}
	out := new(Platform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Platform) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformApiserver) DeepCopyInto(out *PlatformApiserver) {
	*out = *in
	in.GRPC.DeepCopyInto(&out.GRPC)
	in.Restful.DeepCopyInto(&out.Restful)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformApiserver.
func (in *PlatformApiserver) DeepCopy() *PlatformApiserver {
	if in == nil {
		return nil
	}
	out := new(PlatformApiserver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformApp) DeepCopyInto(out *PlatformApp) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = make([]extensionsv1beta1.IngressTLS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformApp.
func (in *PlatformApp) DeepCopy() *PlatformApp {
	if in == nil {
		return nil
	}
	out := new(PlatformApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformDgraph) DeepCopyInto(out *PlatformDgraph) {
	*out = *in
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformDgraph.
func (in *PlatformDgraph) DeepCopy() *PlatformDgraph {
	if in == nil {
		return nil
	}
	out := new(PlatformDgraph)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformInfinimeshDefaultStorage) DeepCopyInto(out *PlatformInfinimeshDefaultStorage) {
	*out = *in
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformDgraph.
func (in *PlatformInfinimeshDefaultStorage) DeepCopy() *PlatformInfinimeshDefaultStorage {
	if in == nil {
		return nil
	}
	out := new(PlatformInfinimeshDefaultStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformGRPCApiserver) DeepCopyInto(out *PlatformGRPCApiserver) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = make([]extensionsv1beta1.IngressTLS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformGRPCApiserver.
func (in *PlatformGRPCApiserver) DeepCopy() *PlatformGRPCApiserver {
	if in == nil {
		return nil
	}
	out := new(PlatformGRPCApiserver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformKafka) DeepCopyInto(out *PlatformKafka) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformKafka.
func (in *PlatformKafka) DeepCopy() *PlatformKafka {
	if in == nil {
		return nil
	}
	out := new(PlatformKafka)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformList) DeepCopyInto(out *PlatformList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Platform, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformList.
func (in *PlatformList) DeepCopy() *PlatformList {
	if in == nil {
		return nil
	}
	out := new(PlatformList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlatformList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformMQTTBroker) DeepCopyInto(out *PlatformMQTTBroker) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformMQTTBroker.
func (in *PlatformMQTTBroker) DeepCopy() *PlatformMQTTBroker {
	if in == nil {
		return nil
	}
	out := new(PlatformMQTTBroker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
// func (in *PlatformInfinimeshDefaultStorage) DeepCopyInto(out *PlatformInfinimeshDefaultStorage) {
// 	*out = *in
// 	return
// }

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformStorage.
// func (in *PlatformInfinimeshDefaultStorage) DeepCopy() *PlatformInfinimeshDefaultStorage {
// 	if in == nil {
// 		return nil
// 	}
// 	out := new(PlatformInfinimeshDefaultStorage)
// 	in.DeepCopyInto(out)
// 	return out
// }

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformRestfulApiserver) DeepCopyInto(out *PlatformRestfulApiserver) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = make([]extensionsv1beta1.IngressTLS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformRestfulApiserver.
func (in *PlatformRestfulApiserver) DeepCopy() *PlatformRestfulApiserver {
	if in == nil {
		return nil
	}
	out := new(PlatformRestfulApiserver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformSpec) DeepCopyInto(out *PlatformSpec) {
	*out = *in
	out.MQTT = in.MQTT
	in.DGraph.DeepCopyInto(&out.DGraph)
	out.Kafka = in.Kafka
	in.Apiserver.DeepCopyInto(&out.Apiserver)
	in.App.DeepCopyInto(&out.App)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformSpec.
func (in *PlatformSpec) DeepCopy() *PlatformSpec {
	if in == nil {
		return nil
	}
	out := new(PlatformSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformStatus) DeepCopyInto(out *PlatformStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformStatus.
func (in *PlatformStatus) DeepCopy() *PlatformStatus {
	if in == nil {
		return nil
	}
	out := new(PlatformStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformTimescaleDB) DeepCopyInto(out *PlatformTimescaleDB) {
	*out = *in
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformTimescaleDB.
func (in *PlatformTimescaleDB) DeepCopy() *PlatformTimescaleDB {
	if in == nil {
		return nil
	}
	out := new(PlatformTimescaleDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformTimeseries) DeepCopyInto(out *PlatformTimeseries) {
	*out = *in
	if in.TimescaleDB != nil {
		in, out := &in.TimescaleDB, &out.TimescaleDB
		*out = new(PlatformTimescaleDB)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformTimeseries.
func (in *PlatformTimeseries) DeepCopy() *PlatformTimeseries {
	if in == nil {
		return nil
	}
	out := new(PlatformTimeseries)
	in.DeepCopyInto(out)
	return out
}
