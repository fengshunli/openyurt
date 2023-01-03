//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1beta1 "k8s.io/kubelet/config/v1beta1"

	config "github.com/openyurtio/openyurt/pkg/util/kubernetes/kubelet/apis/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1beta1.KubeletAnonymousAuthentication)(nil), (*config.KubeletAnonymousAuthentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeletAnonymousAuthentication_To_config_KubeletAnonymousAuthentication(a.(*v1beta1.KubeletAnonymousAuthentication), b.(*config.KubeletAnonymousAuthentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeletAnonymousAuthentication)(nil), (*v1beta1.KubeletAnonymousAuthentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeletAnonymousAuthentication_To_v1beta1_KubeletAnonymousAuthentication(a.(*config.KubeletAnonymousAuthentication), b.(*v1beta1.KubeletAnonymousAuthentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.KubeletAuthentication)(nil), (*config.KubeletAuthentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeletAuthentication_To_config_KubeletAuthentication(a.(*v1beta1.KubeletAuthentication), b.(*config.KubeletAuthentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeletAuthentication)(nil), (*v1beta1.KubeletAuthentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeletAuthentication_To_v1beta1_KubeletAuthentication(a.(*config.KubeletAuthentication), b.(*v1beta1.KubeletAuthentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.KubeletAuthorization)(nil), (*config.KubeletAuthorization)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeletAuthorization_To_config_KubeletAuthorization(a.(*v1beta1.KubeletAuthorization), b.(*config.KubeletAuthorization), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeletAuthorization)(nil), (*v1beta1.KubeletAuthorization)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeletAuthorization_To_v1beta1_KubeletAuthorization(a.(*config.KubeletAuthorization), b.(*v1beta1.KubeletAuthorization), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.KubeletConfiguration)(nil), (*config.KubeletConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeletConfiguration_To_config_KubeletConfiguration(a.(*v1beta1.KubeletConfiguration), b.(*config.KubeletConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeletConfiguration)(nil), (*v1beta1.KubeletConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeletConfiguration_To_v1beta1_KubeletConfiguration(a.(*config.KubeletConfiguration), b.(*v1beta1.KubeletConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.KubeletWebhookAuthentication)(nil), (*config.KubeletWebhookAuthentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeletWebhookAuthentication_To_config_KubeletWebhookAuthentication(a.(*v1beta1.KubeletWebhookAuthentication), b.(*config.KubeletWebhookAuthentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeletWebhookAuthentication)(nil), (*v1beta1.KubeletWebhookAuthentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeletWebhookAuthentication_To_v1beta1_KubeletWebhookAuthentication(a.(*config.KubeletWebhookAuthentication), b.(*v1beta1.KubeletWebhookAuthentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.KubeletWebhookAuthorization)(nil), (*config.KubeletWebhookAuthorization)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeletWebhookAuthorization_To_config_KubeletWebhookAuthorization(a.(*v1beta1.KubeletWebhookAuthorization), b.(*config.KubeletWebhookAuthorization), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeletWebhookAuthorization)(nil), (*v1beta1.KubeletWebhookAuthorization)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeletWebhookAuthorization_To_v1beta1_KubeletWebhookAuthorization(a.(*config.KubeletWebhookAuthorization), b.(*v1beta1.KubeletWebhookAuthorization), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.KubeletX509Authentication)(nil), (*config.KubeletX509Authentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeletX509Authentication_To_config_KubeletX509Authentication(a.(*v1beta1.KubeletX509Authentication), b.(*config.KubeletX509Authentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeletX509Authentication)(nil), (*v1beta1.KubeletX509Authentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeletX509Authentication_To_v1beta1_KubeletX509Authentication(a.(*config.KubeletX509Authentication), b.(*v1beta1.KubeletX509Authentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.SerializedNodeConfigSource)(nil), (*config.SerializedNodeConfigSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SerializedNodeConfigSource_To_config_SerializedNodeConfigSource(a.(*v1beta1.SerializedNodeConfigSource), b.(*config.SerializedNodeConfigSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SerializedNodeConfigSource)(nil), (*v1beta1.SerializedNodeConfigSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SerializedNodeConfigSource_To_v1beta1_SerializedNodeConfigSource(a.(*config.SerializedNodeConfigSource), b.(*v1beta1.SerializedNodeConfigSource), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_KubeletAnonymousAuthentication_To_config_KubeletAnonymousAuthentication(in *v1beta1.KubeletAnonymousAuthentication, out *config.KubeletAnonymousAuthentication, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.Enabled, &out.Enabled, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_KubeletAnonymousAuthentication_To_config_KubeletAnonymousAuthentication is an autogenerated conversion function.
func Convert_v1beta1_KubeletAnonymousAuthentication_To_config_KubeletAnonymousAuthentication(in *v1beta1.KubeletAnonymousAuthentication, out *config.KubeletAnonymousAuthentication, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeletAnonymousAuthentication_To_config_KubeletAnonymousAuthentication(in, out, s)
}

func autoConvert_config_KubeletAnonymousAuthentication_To_v1beta1_KubeletAnonymousAuthentication(in *config.KubeletAnonymousAuthentication, out *v1beta1.KubeletAnonymousAuthentication, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.Enabled, &out.Enabled, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_KubeletAnonymousAuthentication_To_v1beta1_KubeletAnonymousAuthentication is an autogenerated conversion function.
func Convert_config_KubeletAnonymousAuthentication_To_v1beta1_KubeletAnonymousAuthentication(in *config.KubeletAnonymousAuthentication, out *v1beta1.KubeletAnonymousAuthentication, s conversion.Scope) error {
	return autoConvert_config_KubeletAnonymousAuthentication_To_v1beta1_KubeletAnonymousAuthentication(in, out, s)
}

func autoConvert_v1beta1_KubeletAuthentication_To_config_KubeletAuthentication(in *v1beta1.KubeletAuthentication, out *config.KubeletAuthentication, s conversion.Scope) error {
	if err := Convert_v1beta1_KubeletX509Authentication_To_config_KubeletX509Authentication(&in.X509, &out.X509, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_KubeletWebhookAuthentication_To_config_KubeletWebhookAuthentication(&in.Webhook, &out.Webhook, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_KubeletAnonymousAuthentication_To_config_KubeletAnonymousAuthentication(&in.Anonymous, &out.Anonymous, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_KubeletAuthentication_To_config_KubeletAuthentication is an autogenerated conversion function.
func Convert_v1beta1_KubeletAuthentication_To_config_KubeletAuthentication(in *v1beta1.KubeletAuthentication, out *config.KubeletAuthentication, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeletAuthentication_To_config_KubeletAuthentication(in, out, s)
}

func autoConvert_config_KubeletAuthentication_To_v1beta1_KubeletAuthentication(in *config.KubeletAuthentication, out *v1beta1.KubeletAuthentication, s conversion.Scope) error {
	if err := Convert_config_KubeletX509Authentication_To_v1beta1_KubeletX509Authentication(&in.X509, &out.X509, s); err != nil {
		return err
	}
	if err := Convert_config_KubeletWebhookAuthentication_To_v1beta1_KubeletWebhookAuthentication(&in.Webhook, &out.Webhook, s); err != nil {
		return err
	}
	if err := Convert_config_KubeletAnonymousAuthentication_To_v1beta1_KubeletAnonymousAuthentication(&in.Anonymous, &out.Anonymous, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_KubeletAuthentication_To_v1beta1_KubeletAuthentication is an autogenerated conversion function.
func Convert_config_KubeletAuthentication_To_v1beta1_KubeletAuthentication(in *config.KubeletAuthentication, out *v1beta1.KubeletAuthentication, s conversion.Scope) error {
	return autoConvert_config_KubeletAuthentication_To_v1beta1_KubeletAuthentication(in, out, s)
}

func autoConvert_v1beta1_KubeletAuthorization_To_config_KubeletAuthorization(in *v1beta1.KubeletAuthorization, out *config.KubeletAuthorization, s conversion.Scope) error {
	out.Mode = config.KubeletAuthorizationMode(in.Mode)
	if err := Convert_v1beta1_KubeletWebhookAuthorization_To_config_KubeletWebhookAuthorization(&in.Webhook, &out.Webhook, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_KubeletAuthorization_To_config_KubeletAuthorization is an autogenerated conversion function.
func Convert_v1beta1_KubeletAuthorization_To_config_KubeletAuthorization(in *v1beta1.KubeletAuthorization, out *config.KubeletAuthorization, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeletAuthorization_To_config_KubeletAuthorization(in, out, s)
}

func autoConvert_config_KubeletAuthorization_To_v1beta1_KubeletAuthorization(in *config.KubeletAuthorization, out *v1beta1.KubeletAuthorization, s conversion.Scope) error {
	out.Mode = v1beta1.KubeletAuthorizationMode(in.Mode)
	if err := Convert_config_KubeletWebhookAuthorization_To_v1beta1_KubeletWebhookAuthorization(&in.Webhook, &out.Webhook, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_KubeletAuthorization_To_v1beta1_KubeletAuthorization is an autogenerated conversion function.
func Convert_config_KubeletAuthorization_To_v1beta1_KubeletAuthorization(in *config.KubeletAuthorization, out *v1beta1.KubeletAuthorization, s conversion.Scope) error {
	return autoConvert_config_KubeletAuthorization_To_v1beta1_KubeletAuthorization(in, out, s)
}

func autoConvert_v1beta1_KubeletConfiguration_To_config_KubeletConfiguration(in *v1beta1.KubeletConfiguration, out *config.KubeletConfiguration, s conversion.Scope) error {
	out.StaticPodPath = in.StaticPodPath
	out.SyncFrequency = in.SyncFrequency
	out.FileCheckFrequency = in.FileCheckFrequency
	out.HTTPCheckFrequency = in.HTTPCheckFrequency
	out.StaticPodURL = in.StaticPodURL
	out.StaticPodURLHeader = *(*map[string][]string)(unsafe.Pointer(&in.StaticPodURLHeader))
	out.Address = in.Address
	out.Port = in.Port
	out.ReadOnlyPort = in.ReadOnlyPort
	out.TLSCertFile = in.TLSCertFile
	out.TLSPrivateKeyFile = in.TLSPrivateKeyFile
	out.TLSCipherSuites = *(*[]string)(unsafe.Pointer(&in.TLSCipherSuites))
	out.TLSMinVersion = in.TLSMinVersion
	out.RotateCertificates = in.RotateCertificates
	out.ServerTLSBootstrap = in.ServerTLSBootstrap
	if err := Convert_v1beta1_KubeletAuthentication_To_config_KubeletAuthentication(&in.Authentication, &out.Authentication, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_KubeletAuthorization_To_config_KubeletAuthorization(&in.Authorization, &out.Authorization, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int32_To_int32(&in.RegistryPullQPS, &out.RegistryPullQPS, s); err != nil {
		return err
	}
	out.RegistryBurst = in.RegistryBurst
	if err := v1.Convert_Pointer_int32_To_int32(&in.EventRecordQPS, &out.EventRecordQPS, s); err != nil {
		return err
	}
	out.EventBurst = in.EventBurst
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableDebuggingHandlers, &out.EnableDebuggingHandlers, s); err != nil {
		return err
	}
	out.EnableContentionProfiling = in.EnableContentionProfiling
	if err := v1.Convert_Pointer_int32_To_int32(&in.HealthzPort, &out.HealthzPort, s); err != nil {
		return err
	}
	out.HealthzBindAddress = in.HealthzBindAddress
	if err := v1.Convert_Pointer_int32_To_int32(&in.OOMScoreAdj, &out.OOMScoreAdj, s); err != nil {
		return err
	}
	out.ClusterDomain = in.ClusterDomain
	out.ClusterDNS = *(*[]string)(unsafe.Pointer(&in.ClusterDNS))
	out.StreamingConnectionIdleTimeout = in.StreamingConnectionIdleTimeout
	out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
	out.NodeStatusReportFrequency = in.NodeStatusReportFrequency
	out.NodeLeaseDurationSeconds = in.NodeLeaseDurationSeconds
	out.ImageMinimumGCAge = in.ImageMinimumGCAge
	if err := v1.Convert_Pointer_int32_To_int32(&in.ImageGCHighThresholdPercent, &out.ImageGCHighThresholdPercent, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int32_To_int32(&in.ImageGCLowThresholdPercent, &out.ImageGCLowThresholdPercent, s); err != nil {
		return err
	}
	out.VolumeStatsAggPeriod = in.VolumeStatsAggPeriod
	out.KubeletCgroups = in.KubeletCgroups
	out.SystemCgroups = in.SystemCgroups
	out.CgroupRoot = in.CgroupRoot
	if err := v1.Convert_Pointer_bool_To_bool(&in.CgroupsPerQOS, &out.CgroupsPerQOS, s); err != nil {
		return err
	}
	out.CgroupDriver = in.CgroupDriver
	out.CPUManagerPolicy = in.CPUManagerPolicy
	out.CPUManagerReconcilePeriod = in.CPUManagerReconcilePeriod
	out.TopologyManagerPolicy = in.TopologyManagerPolicy
	out.QOSReserved = *(*map[string]string)(unsafe.Pointer(&in.QOSReserved))
	out.RuntimeRequestTimeout = in.RuntimeRequestTimeout
	out.HairpinMode = in.HairpinMode
	out.MaxPods = in.MaxPods
	out.PodCIDR = in.PodCIDR
	if err := v1.Convert_Pointer_int64_To_int64(&in.PodPidsLimit, &out.PodPidsLimit, s); err != nil {
		return err
	}
	out.ResolverConfig = in.ResolverConfig
	if err := v1.Convert_Pointer_bool_To_bool(&in.CPUCFSQuota, &out.CPUCFSQuota, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_v1_Duration_To_v1_Duration(&in.CPUCFSQuotaPeriod, &out.CPUCFSQuotaPeriod, s); err != nil {
		return err
	}
	out.MaxOpenFiles = in.MaxOpenFiles
	out.ContentType = in.ContentType
	if err := v1.Convert_Pointer_int32_To_int32(&in.KubeAPIQPS, &out.KubeAPIQPS, s); err != nil {
		return err
	}
	out.KubeAPIBurst = in.KubeAPIBurst
	if err := v1.Convert_Pointer_bool_To_bool(&in.SerializeImagePulls, &out.SerializeImagePulls, s); err != nil {
		return err
	}
	out.EvictionHard = *(*map[string]string)(unsafe.Pointer(&in.EvictionHard))
	out.EvictionSoft = *(*map[string]string)(unsafe.Pointer(&in.EvictionSoft))
	out.EvictionSoftGracePeriod = *(*map[string]string)(unsafe.Pointer(&in.EvictionSoftGracePeriod))
	out.EvictionPressureTransitionPeriod = in.EvictionPressureTransitionPeriod
	out.EvictionMaxPodGracePeriod = in.EvictionMaxPodGracePeriod
	out.EvictionMinimumReclaim = *(*map[string]string)(unsafe.Pointer(&in.EvictionMinimumReclaim))
	out.PodsPerCore = in.PodsPerCore
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableControllerAttachDetach, &out.EnableControllerAttachDetach, s); err != nil {
		return err
	}
	out.ProtectKernelDefaults = in.ProtectKernelDefaults
	if err := v1.Convert_Pointer_bool_To_bool(&in.MakeIPTablesUtilChains, &out.MakeIPTablesUtilChains, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int32_To_int32(&in.IPTablesMasqueradeBit, &out.IPTablesMasqueradeBit, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int32_To_int32(&in.IPTablesDropBit, &out.IPTablesDropBit, s); err != nil {
		return err
	}
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	if err := v1.Convert_Pointer_bool_To_bool(&in.FailSwapOn, &out.FailSwapOn, s); err != nil {
		return err
	}
	out.ContainerLogMaxSize = in.ContainerLogMaxSize
	if err := v1.Convert_Pointer_int32_To_int32(&in.ContainerLogMaxFiles, &out.ContainerLogMaxFiles, s); err != nil {
		return err
	}
	out.ConfigMapAndSecretChangeDetectionStrategy = config.ResourceChangeDetectionStrategy(in.ConfigMapAndSecretChangeDetectionStrategy)
	out.SystemReserved = *(*map[string]string)(unsafe.Pointer(&in.SystemReserved))
	out.KubeReserved = *(*map[string]string)(unsafe.Pointer(&in.KubeReserved))
	out.ReservedSystemCPUs = in.ReservedSystemCPUs
	out.ShowHiddenMetricsForVersion = in.ShowHiddenMetricsForVersion
	out.SystemReservedCgroup = in.SystemReservedCgroup
	out.KubeReservedCgroup = in.KubeReservedCgroup
	out.EnforceNodeAllocatable = *(*[]string)(unsafe.Pointer(&in.EnforceNodeAllocatable))
	out.AllowedUnsafeSysctls = *(*[]string)(unsafe.Pointer(&in.AllowedUnsafeSysctls))
	return nil
}

// Convert_v1beta1_KubeletConfiguration_To_config_KubeletConfiguration is an autogenerated conversion function.
func Convert_v1beta1_KubeletConfiguration_To_config_KubeletConfiguration(in *v1beta1.KubeletConfiguration, out *config.KubeletConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeletConfiguration_To_config_KubeletConfiguration(in, out, s)
}

func autoConvert_config_KubeletConfiguration_To_v1beta1_KubeletConfiguration(in *config.KubeletConfiguration, out *v1beta1.KubeletConfiguration, s conversion.Scope) error {
	out.StaticPodPath = in.StaticPodPath
	out.SyncFrequency = in.SyncFrequency
	out.FileCheckFrequency = in.FileCheckFrequency
	out.HTTPCheckFrequency = in.HTTPCheckFrequency
	out.StaticPodURL = in.StaticPodURL
	out.StaticPodURLHeader = *(*map[string][]string)(unsafe.Pointer(&in.StaticPodURLHeader))
	out.Address = in.Address
	out.Port = in.Port
	out.ReadOnlyPort = in.ReadOnlyPort
	out.TLSCertFile = in.TLSCertFile
	out.TLSPrivateKeyFile = in.TLSPrivateKeyFile
	out.TLSCipherSuites = *(*[]string)(unsafe.Pointer(&in.TLSCipherSuites))
	out.TLSMinVersion = in.TLSMinVersion
	out.RotateCertificates = in.RotateCertificates
	out.ServerTLSBootstrap = in.ServerTLSBootstrap
	if err := Convert_config_KubeletAuthentication_To_v1beta1_KubeletAuthentication(&in.Authentication, &out.Authentication, s); err != nil {
		return err
	}
	if err := Convert_config_KubeletAuthorization_To_v1beta1_KubeletAuthorization(&in.Authorization, &out.Authorization, s); err != nil {
		return err
	}
	if err := v1.Convert_int32_To_Pointer_int32(&in.RegistryPullQPS, &out.RegistryPullQPS, s); err != nil {
		return err
	}
	out.RegistryBurst = in.RegistryBurst
	if err := v1.Convert_int32_To_Pointer_int32(&in.EventRecordQPS, &out.EventRecordQPS, s); err != nil {
		return err
	}
	out.EventBurst = in.EventBurst
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableDebuggingHandlers, &out.EnableDebuggingHandlers, s); err != nil {
		return err
	}
	out.EnableContentionProfiling = in.EnableContentionProfiling
	if err := v1.Convert_int32_To_Pointer_int32(&in.HealthzPort, &out.HealthzPort, s); err != nil {
		return err
	}
	out.HealthzBindAddress = in.HealthzBindAddress
	if err := v1.Convert_int32_To_Pointer_int32(&in.OOMScoreAdj, &out.OOMScoreAdj, s); err != nil {
		return err
	}
	out.ClusterDomain = in.ClusterDomain
	out.ClusterDNS = *(*[]string)(unsafe.Pointer(&in.ClusterDNS))
	out.StreamingConnectionIdleTimeout = in.StreamingConnectionIdleTimeout
	out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
	out.NodeStatusReportFrequency = in.NodeStatusReportFrequency
	out.NodeLeaseDurationSeconds = in.NodeLeaseDurationSeconds
	out.ImageMinimumGCAge = in.ImageMinimumGCAge
	if err := v1.Convert_int32_To_Pointer_int32(&in.ImageGCHighThresholdPercent, &out.ImageGCHighThresholdPercent, s); err != nil {
		return err
	}
	if err := v1.Convert_int32_To_Pointer_int32(&in.ImageGCLowThresholdPercent, &out.ImageGCLowThresholdPercent, s); err != nil {
		return err
	}
	out.VolumeStatsAggPeriod = in.VolumeStatsAggPeriod
	out.KubeletCgroups = in.KubeletCgroups
	out.SystemCgroups = in.SystemCgroups
	out.CgroupRoot = in.CgroupRoot
	if err := v1.Convert_bool_To_Pointer_bool(&in.CgroupsPerQOS, &out.CgroupsPerQOS, s); err != nil {
		return err
	}
	out.CgroupDriver = in.CgroupDriver
	out.CPUManagerPolicy = in.CPUManagerPolicy
	out.CPUManagerReconcilePeriod = in.CPUManagerReconcilePeriod
	out.TopologyManagerPolicy = in.TopologyManagerPolicy
	out.QOSReserved = *(*map[string]string)(unsafe.Pointer(&in.QOSReserved))
	out.RuntimeRequestTimeout = in.RuntimeRequestTimeout
	out.HairpinMode = in.HairpinMode
	out.MaxPods = in.MaxPods
	out.PodCIDR = in.PodCIDR
	if err := v1.Convert_int64_To_Pointer_int64(&in.PodPidsLimit, &out.PodPidsLimit, s); err != nil {
		return err
	}
	out.ResolverConfig = in.ResolverConfig
	if err := v1.Convert_bool_To_Pointer_bool(&in.CPUCFSQuota, &out.CPUCFSQuota, s); err != nil {
		return err
	}
	if err := v1.Convert_v1_Duration_To_Pointer_v1_Duration(&in.CPUCFSQuotaPeriod, &out.CPUCFSQuotaPeriod, s); err != nil {
		return err
	}
	out.MaxOpenFiles = in.MaxOpenFiles
	out.ContentType = in.ContentType
	if err := v1.Convert_int32_To_Pointer_int32(&in.KubeAPIQPS, &out.KubeAPIQPS, s); err != nil {
		return err
	}
	out.KubeAPIBurst = in.KubeAPIBurst
	if err := v1.Convert_bool_To_Pointer_bool(&in.SerializeImagePulls, &out.SerializeImagePulls, s); err != nil {
		return err
	}
	out.EvictionHard = *(*map[string]string)(unsafe.Pointer(&in.EvictionHard))
	out.EvictionSoft = *(*map[string]string)(unsafe.Pointer(&in.EvictionSoft))
	out.EvictionSoftGracePeriod = *(*map[string]string)(unsafe.Pointer(&in.EvictionSoftGracePeriod))
	out.EvictionPressureTransitionPeriod = in.EvictionPressureTransitionPeriod
	out.EvictionMaxPodGracePeriod = in.EvictionMaxPodGracePeriod
	out.EvictionMinimumReclaim = *(*map[string]string)(unsafe.Pointer(&in.EvictionMinimumReclaim))
	out.PodsPerCore = in.PodsPerCore
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableControllerAttachDetach, &out.EnableControllerAttachDetach, s); err != nil {
		return err
	}
	out.ProtectKernelDefaults = in.ProtectKernelDefaults
	if err := v1.Convert_bool_To_Pointer_bool(&in.MakeIPTablesUtilChains, &out.MakeIPTablesUtilChains, s); err != nil {
		return err
	}
	if err := v1.Convert_int32_To_Pointer_int32(&in.IPTablesMasqueradeBit, &out.IPTablesMasqueradeBit, s); err != nil {
		return err
	}
	if err := v1.Convert_int32_To_Pointer_int32(&in.IPTablesDropBit, &out.IPTablesDropBit, s); err != nil {
		return err
	}
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	if err := v1.Convert_bool_To_Pointer_bool(&in.FailSwapOn, &out.FailSwapOn, s); err != nil {
		return err
	}
	out.ContainerLogMaxSize = in.ContainerLogMaxSize
	if err := v1.Convert_int32_To_Pointer_int32(&in.ContainerLogMaxFiles, &out.ContainerLogMaxFiles, s); err != nil {
		return err
	}
	out.ConfigMapAndSecretChangeDetectionStrategy = v1beta1.ResourceChangeDetectionStrategy(in.ConfigMapAndSecretChangeDetectionStrategy)
	out.AllowedUnsafeSysctls = *(*[]string)(unsafe.Pointer(&in.AllowedUnsafeSysctls))
	out.SystemReserved = *(*map[string]string)(unsafe.Pointer(&in.SystemReserved))
	out.KubeReserved = *(*map[string]string)(unsafe.Pointer(&in.KubeReserved))
	out.SystemReservedCgroup = in.SystemReservedCgroup
	out.KubeReservedCgroup = in.KubeReservedCgroup
	out.EnforceNodeAllocatable = *(*[]string)(unsafe.Pointer(&in.EnforceNodeAllocatable))
	out.ReservedSystemCPUs = in.ReservedSystemCPUs
	out.ShowHiddenMetricsForVersion = in.ShowHiddenMetricsForVersion
	return nil
}

// Convert_config_KubeletConfiguration_To_v1beta1_KubeletConfiguration is an autogenerated conversion function.
func Convert_config_KubeletConfiguration_To_v1beta1_KubeletConfiguration(in *config.KubeletConfiguration, out *v1beta1.KubeletConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeletConfiguration_To_v1beta1_KubeletConfiguration(in, out, s)
}

func autoConvert_v1beta1_KubeletWebhookAuthentication_To_config_KubeletWebhookAuthentication(in *v1beta1.KubeletWebhookAuthentication, out *config.KubeletWebhookAuthentication, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.Enabled, &out.Enabled, s); err != nil {
		return err
	}
	out.CacheTTL = in.CacheTTL
	return nil
}

// Convert_v1beta1_KubeletWebhookAuthentication_To_config_KubeletWebhookAuthentication is an autogenerated conversion function.
func Convert_v1beta1_KubeletWebhookAuthentication_To_config_KubeletWebhookAuthentication(in *v1beta1.KubeletWebhookAuthentication, out *config.KubeletWebhookAuthentication, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeletWebhookAuthentication_To_config_KubeletWebhookAuthentication(in, out, s)
}

func autoConvert_config_KubeletWebhookAuthentication_To_v1beta1_KubeletWebhookAuthentication(in *config.KubeletWebhookAuthentication, out *v1beta1.KubeletWebhookAuthentication, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.Enabled, &out.Enabled, s); err != nil {
		return err
	}
	out.CacheTTL = in.CacheTTL
	return nil
}

// Convert_config_KubeletWebhookAuthentication_To_v1beta1_KubeletWebhookAuthentication is an autogenerated conversion function.
func Convert_config_KubeletWebhookAuthentication_To_v1beta1_KubeletWebhookAuthentication(in *config.KubeletWebhookAuthentication, out *v1beta1.KubeletWebhookAuthentication, s conversion.Scope) error {
	return autoConvert_config_KubeletWebhookAuthentication_To_v1beta1_KubeletWebhookAuthentication(in, out, s)
}

func autoConvert_v1beta1_KubeletWebhookAuthorization_To_config_KubeletWebhookAuthorization(in *v1beta1.KubeletWebhookAuthorization, out *config.KubeletWebhookAuthorization, s conversion.Scope) error {
	out.CacheAuthorizedTTL = in.CacheAuthorizedTTL
	out.CacheUnauthorizedTTL = in.CacheUnauthorizedTTL
	return nil
}

// Convert_v1beta1_KubeletWebhookAuthorization_To_config_KubeletWebhookAuthorization is an autogenerated conversion function.
func Convert_v1beta1_KubeletWebhookAuthorization_To_config_KubeletWebhookAuthorization(in *v1beta1.KubeletWebhookAuthorization, out *config.KubeletWebhookAuthorization, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeletWebhookAuthorization_To_config_KubeletWebhookAuthorization(in, out, s)
}

func autoConvert_config_KubeletWebhookAuthorization_To_v1beta1_KubeletWebhookAuthorization(in *config.KubeletWebhookAuthorization, out *v1beta1.KubeletWebhookAuthorization, s conversion.Scope) error {
	out.CacheAuthorizedTTL = in.CacheAuthorizedTTL
	out.CacheUnauthorizedTTL = in.CacheUnauthorizedTTL
	return nil
}

// Convert_config_KubeletWebhookAuthorization_To_v1beta1_KubeletWebhookAuthorization is an autogenerated conversion function.
func Convert_config_KubeletWebhookAuthorization_To_v1beta1_KubeletWebhookAuthorization(in *config.KubeletWebhookAuthorization, out *v1beta1.KubeletWebhookAuthorization, s conversion.Scope) error {
	return autoConvert_config_KubeletWebhookAuthorization_To_v1beta1_KubeletWebhookAuthorization(in, out, s)
}

func autoConvert_v1beta1_KubeletX509Authentication_To_config_KubeletX509Authentication(in *v1beta1.KubeletX509Authentication, out *config.KubeletX509Authentication, s conversion.Scope) error {
	out.ClientCAFile = in.ClientCAFile
	return nil
}

// Convert_v1beta1_KubeletX509Authentication_To_config_KubeletX509Authentication is an autogenerated conversion function.
func Convert_v1beta1_KubeletX509Authentication_To_config_KubeletX509Authentication(in *v1beta1.KubeletX509Authentication, out *config.KubeletX509Authentication, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeletX509Authentication_To_config_KubeletX509Authentication(in, out, s)
}

func autoConvert_config_KubeletX509Authentication_To_v1beta1_KubeletX509Authentication(in *config.KubeletX509Authentication, out *v1beta1.KubeletX509Authentication, s conversion.Scope) error {
	out.ClientCAFile = in.ClientCAFile
	return nil
}

// Convert_config_KubeletX509Authentication_To_v1beta1_KubeletX509Authentication is an autogenerated conversion function.
func Convert_config_KubeletX509Authentication_To_v1beta1_KubeletX509Authentication(in *config.KubeletX509Authentication, out *v1beta1.KubeletX509Authentication, s conversion.Scope) error {
	return autoConvert_config_KubeletX509Authentication_To_v1beta1_KubeletX509Authentication(in, out, s)
}

func autoConvert_v1beta1_SerializedNodeConfigSource_To_config_SerializedNodeConfigSource(in *v1beta1.SerializedNodeConfigSource, out *config.SerializedNodeConfigSource, s conversion.Scope) error {
	out.Source = in.Source
	return nil
}

// Convert_v1beta1_SerializedNodeConfigSource_To_config_SerializedNodeConfigSource is an autogenerated conversion function.
func Convert_v1beta1_SerializedNodeConfigSource_To_config_SerializedNodeConfigSource(in *v1beta1.SerializedNodeConfigSource, out *config.SerializedNodeConfigSource, s conversion.Scope) error {
	return autoConvert_v1beta1_SerializedNodeConfigSource_To_config_SerializedNodeConfigSource(in, out, s)
}

func autoConvert_config_SerializedNodeConfigSource_To_v1beta1_SerializedNodeConfigSource(in *config.SerializedNodeConfigSource, out *v1beta1.SerializedNodeConfigSource, s conversion.Scope) error {
	out.Source = in.Source
	return nil
}

// Convert_config_SerializedNodeConfigSource_To_v1beta1_SerializedNodeConfigSource is an autogenerated conversion function.
func Convert_config_SerializedNodeConfigSource_To_v1beta1_SerializedNodeConfigSource(in *config.SerializedNodeConfigSource, out *v1beta1.SerializedNodeConfigSource, s conversion.Scope) error {
	return autoConvert_config_SerializedNodeConfigSource_To_v1beta1_SerializedNodeConfigSource(in, out, s)
}