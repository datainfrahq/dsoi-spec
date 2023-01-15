package v1beta1

import v1 "k8s.io/api/core/v1"

// MyApp defines the structure for an app spec.
type MyApp struct {
	Spec MyAppSpec `yaml:"spec"`
}

// MyAppSpec embeds all the specs.
type MyAppSpec struct {
	External           ExternalSpec                      `yaml:"external"`
	NodeK8sConfigGroup map[string]NodeK8sConfigGroupSpec `yaml:"nodeK8sConfigGroup"`
	NodeAppConfigGroup map[string]NodeAppConfigGroupSpec `yaml:"nodeAppConfigGroup"`
	Nodes              map[string]NodeSpec               `yaml:"nodeSpec"`
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

// NodeK8sConfigGroupSpec embeds all the k8s specific configuration.
type NodeK8sConfigGroupSpec struct {
	Resources v1.ResourceRequirements `yaml:"resources,omitempty"`
}

// NodeAppConfigGroupSpec embeds all the app configs specific to each nodetype or common to all nodetypes.
type NodeAppConfigGroupSpec struct {
	CommonRuntimeProperties string `yaml:"commonRuntimeProperties"`
	RuntimeProperties       string `yaml:"runtimeProperties"`
	OverrideProperties      string `yaml:"overrideProperties"`
}

// NodeSpec maps configs to a node of a specific nodeType.
type NodeSpec struct {
	Kind               string   `yaml:"kind"`
	NodeType           string   `yaml:"nodeType"`
	Replicas           int      `yaml:"int"`
	NodeK8sConfigGroup string   `yaml:"nodeK8sConfigGroup"`
	NodeAppConfigGroup []string `yaml:"nodeAppConfigGroup"`
}
