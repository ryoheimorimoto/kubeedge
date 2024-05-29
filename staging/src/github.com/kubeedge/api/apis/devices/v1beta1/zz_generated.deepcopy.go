//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KubeEdge Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBMethodConfig) DeepCopyInto(out *DBMethodConfig) {
	*out = *in
	if in.Influxdb2 != nil {
		in, out := &in.Influxdb2, &out.Influxdb2
		*out = new(DBMethodInfluxdb2)
		(*in).DeepCopyInto(*out)
	}
	if in.Redis != nil {
		in, out := &in.Redis, &out.Redis
		*out = new(DBMethodRedis)
		(*in).DeepCopyInto(*out)
	}
	if in.TDEngine != nil {
		in, out := &in.TDEngine, &out.TDEngine
		*out = new(DBMethodTDEngine)
		(*in).DeepCopyInto(*out)
	}
	if in.Mysql != nil {
		in, out := &in.Mysql, &out.Mysql
		*out = new(DBMethodMySQL)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBMethodConfig.
func (in *DBMethodConfig) DeepCopy() *DBMethodConfig {
	if in == nil {
		return nil
	}
	out := new(DBMethodConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBMethodInfluxdb2) DeepCopyInto(out *DBMethodInfluxdb2) {
	*out = *in
	if in.Influxdb2ClientConfig != nil {
		in, out := &in.Influxdb2ClientConfig, &out.Influxdb2ClientConfig
		*out = new(Influxdb2ClientConfig)
		**out = **in
	}
	if in.Influxdb2DataConfig != nil {
		in, out := &in.Influxdb2DataConfig, &out.Influxdb2DataConfig
		*out = new(Influxdb2DataConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBMethodInfluxdb2.
func (in *DBMethodInfluxdb2) DeepCopy() *DBMethodInfluxdb2 {
	if in == nil {
		return nil
	}
	out := new(DBMethodInfluxdb2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBMethodMySQL) DeepCopyInto(out *DBMethodMySQL) {
	*out = *in
	if in.MySQLClientConfig != nil {
		in, out := &in.MySQLClientConfig, &out.MySQLClientConfig
		*out = new(MySQLClientConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBMethodMySQL.
func (in *DBMethodMySQL) DeepCopy() *DBMethodMySQL {
	if in == nil {
		return nil
	}
	out := new(DBMethodMySQL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBMethodRedis) DeepCopyInto(out *DBMethodRedis) {
	*out = *in
	if in.RedisClientConfig != nil {
		in, out := &in.RedisClientConfig, &out.RedisClientConfig
		*out = new(RedisClientConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBMethodRedis.
func (in *DBMethodRedis) DeepCopy() *DBMethodRedis {
	if in == nil {
		return nil
	}
	out := new(DBMethodRedis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBMethodTDEngine) DeepCopyInto(out *DBMethodTDEngine) {
	*out = *in
	if in.TDEngineClientConfig != nil {
		in, out := &in.TDEngineClientConfig, &out.TDEngineClientConfig
		*out = new(TDEngineClientConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBMethodTDEngine.
func (in *DBMethodTDEngine) DeepCopy() *DBMethodTDEngine {
	if in == nil {
		return nil
	}
	out := new(DBMethodTDEngine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Device) DeepCopyInto(out *Device) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Device.
func (in *Device) DeepCopy() *Device {
	if in == nil {
		return nil
	}
	out := new(Device)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Device) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceList) DeepCopyInto(out *DeviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Device, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceList.
func (in *DeviceList) DeepCopy() *DeviceList {
	if in == nil {
		return nil
	}
	out := new(DeviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceModel) DeepCopyInto(out *DeviceModel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceModel.
func (in *DeviceModel) DeepCopy() *DeviceModel {
	if in == nil {
		return nil
	}
	out := new(DeviceModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceModel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceModelList) DeepCopyInto(out *DeviceModelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeviceModel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceModelList.
func (in *DeviceModelList) DeepCopy() *DeviceModelList {
	if in == nil {
		return nil
	}
	out := new(DeviceModelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceModelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceModelSpec) DeepCopyInto(out *DeviceModelSpec) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]ModelProperty, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceModelSpec.
func (in *DeviceModelSpec) DeepCopy() *DeviceModelSpec {
	if in == nil {
		return nil
	}
	out := new(DeviceModelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceProperty) DeepCopyInto(out *DeviceProperty) {
	*out = *in
	in.Desired.DeepCopyInto(&out.Desired)
	in.Visitors.DeepCopyInto(&out.Visitors)
	if in.PushMethod != nil {
		in, out := &in.PushMethod, &out.PushMethod
		*out = new(PushMethod)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceProperty.
func (in *DeviceProperty) DeepCopy() *DeviceProperty {
	if in == nil {
		return nil
	}
	out := new(DeviceProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSpec) DeepCopyInto(out *DeviceSpec) {
	*out = *in
	if in.DeviceModelRef != nil {
		in, out := &in.DeviceModelRef, &out.DeviceModelRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]DeviceProperty, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Protocol.DeepCopyInto(&out.Protocol)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSpec.
func (in *DeviceSpec) DeepCopy() *DeviceSpec {
	if in == nil {
		return nil
	}
	out := new(DeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceStatus) DeepCopyInto(out *DeviceStatus) {
	*out = *in
	if in.Twins != nil {
		in, out := &in.Twins, &out.Twins
		*out = make([]Twin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceStatus.
func (in *DeviceStatus) DeepCopy() *DeviceStatus {
	if in == nil {
		return nil
	}
	out := new(DeviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Influxdb2ClientConfig) DeepCopyInto(out *Influxdb2ClientConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Influxdb2ClientConfig.
func (in *Influxdb2ClientConfig) DeepCopy() *Influxdb2ClientConfig {
	if in == nil {
		return nil
	}
	out := new(Influxdb2ClientConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Influxdb2DataConfig) DeepCopyInto(out *Influxdb2DataConfig) {
	*out = *in
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Influxdb2DataConfig.
func (in *Influxdb2DataConfig) DeepCopy() *Influxdb2DataConfig {
	if in == nil {
		return nil
	}
	out := new(Influxdb2DataConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelProperty) DeepCopyInto(out *ModelProperty) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelProperty.
func (in *ModelProperty) DeepCopy() *ModelProperty {
	if in == nil {
		return nil
	}
	out := new(ModelProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLClientConfig) DeepCopyInto(out *MySQLClientConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLClientConfig.
func (in *MySQLClientConfig) DeepCopy() *MySQLClientConfig {
	if in == nil {
		return nil
	}
	out := new(MySQLClientConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtocolConfig) DeepCopyInto(out *ProtocolConfig) {
	*out = *in
	if in.ConfigData != nil {
		in, out := &in.ConfigData, &out.ConfigData
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtocolConfig.
func (in *ProtocolConfig) DeepCopy() *ProtocolConfig {
	if in == nil {
		return nil
	}
	out := new(ProtocolConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushMethod) DeepCopyInto(out *PushMethod) {
	*out = *in
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = new(PushMethodHTTP)
		**out = **in
	}
	if in.MQTT != nil {
		in, out := &in.MQTT, &out.MQTT
		*out = new(PushMethodMQTT)
		**out = **in
	}
	if in.OTEL != nil {
		in, out := &in.OTEL, &out.OTEL
		*out = new(PushMethodOTEL)
		**out = **in
	}
	if in.DBMethod != nil {
		in, out := &in.DBMethod, &out.DBMethod
		*out = new(DBMethodConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushMethod.
func (in *PushMethod) DeepCopy() *PushMethod {
	if in == nil {
		return nil
	}
	out := new(PushMethod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushMethodHTTP) DeepCopyInto(out *PushMethodHTTP) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushMethodHTTP.
func (in *PushMethodHTTP) DeepCopy() *PushMethodHTTP {
	if in == nil {
		return nil
	}
	out := new(PushMethodHTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushMethodMQTT) DeepCopyInto(out *PushMethodMQTT) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushMethodMQTT.
func (in *PushMethodMQTT) DeepCopy() *PushMethodMQTT {
	if in == nil {
		return nil
	}
	out := new(PushMethodMQTT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushMethodOTEL) DeepCopyInto(out *PushMethodOTEL) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushMethodOTEL.
func (in *PushMethodOTEL) DeepCopy() *PushMethodOTEL {
	if in == nil {
		return nil
	}
	out := new(PushMethodOTEL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisClientConfig) DeepCopyInto(out *RedisClientConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisClientConfig.
func (in *RedisClientConfig) DeepCopy() *RedisClientConfig {
	if in == nil {
		return nil
	}
	out := new(RedisClientConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TDEngineClientConfig) DeepCopyInto(out *TDEngineClientConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TDEngineClientConfig.
func (in *TDEngineClientConfig) DeepCopy() *TDEngineClientConfig {
	if in == nil {
		return nil
	}
	out := new(TDEngineClientConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Twin) DeepCopyInto(out *Twin) {
	*out = *in
	in.Reported.DeepCopyInto(&out.Reported)
	in.ObservedDesired.DeepCopyInto(&out.ObservedDesired)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Twin.
func (in *Twin) DeepCopy() *Twin {
	if in == nil {
		return nil
	}
	out := new(Twin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinProperty) DeepCopyInto(out *TwinProperty) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinProperty.
func (in *TwinProperty) DeepCopy() *TwinProperty {
	if in == nil {
		return nil
	}
	out := new(TwinProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VisitorConfig) DeepCopyInto(out *VisitorConfig) {
	*out = *in
	if in.ConfigData != nil {
		in, out := &in.ConfigData, &out.ConfigData
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VisitorConfig.
func (in *VisitorConfig) DeepCopy() *VisitorConfig {
	if in == nil {
		return nil
	}
	out := new(VisitorConfig)
	in.DeepCopyInto(out)
	return out
}
