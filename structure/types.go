package v1beta1

import v1 "k8s.io/api/core/v1"

// MyApp defines the structure for an app spec.
type MyApp struct {
	Spec MyAppSpec `yaml:"spec"`
}

// MyAppSpec embeds all the specs.
type MyAppSpec struct {
	External           ExternalSpec                       `yaml:"external"`
	NodeK8sConfigGroup map[string]NodeK8sConfigGroupsSpec `yaml:"nodeK8sConfigGroup"`
}

// ExternalSpec embeds all the external specs required for the app.
type ExternalSpec struct {
	Zookeeper ZookeeperSpec `yaml:"zookeeper"`
	Metadata  MetadataSpec  `yaml:"metadata"`
}

// ZookeeperSpec embeds all the spec, configuration required.
type ZookeeperSpec struct {
	Spec ZkConfig `yaml:"spec"`
}

// ZkConfig provides configs required by zookeeper and app.
type ZkConfig struct {
	ZkRuntimeProperties string `yaml:"zkRuntimeProperties"`
	ZkConnections       string `yaml:"zkConnections"`
}

// MetadataSpec embeds all the spec, configuration required.
type MetadataSpec struct {
	Spec MetaConfig `yaml:"spec"`
}

// MetadataConfig embeds all the spec, configuration required.
type MetaConfig struct {
	MetaRuntimeProperties string `yaml:"metaRuntimeProperties"`
	MetaConnections       string `yaml:"metaConnections"`
}

type NodeK8sConfigGroupsSpec struct {
	Resources v1.ResourceRequirements `yaml:"resources,omitempty"`
}
