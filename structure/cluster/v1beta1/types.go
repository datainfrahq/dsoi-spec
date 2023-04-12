package v1beta1

import v1 "k8s.io/api/core/v1"

// MyApp defines the structure for an app spec.
type MyApp struct {
	Spec MyAppSpec `yaml:"spec"`
}

// MyAppSpec embeds all the specs.
type MyAppSpec struct {
	DeploymentOrder  []string               `yaml:"deploymentOrder"`
	External         ExternalSpec           `yaml:"external,omitempty"`
	K8sConfigGroup   []K8sConfigGroupSpec   `yaml:"k8sConfigGroup"`
	MyAppConfigGroup []MyAppConfigGroupSpec `yaml:"myAppConfigGroup"`
	Nodes            []NodeSpec             `yaml:"nodes"`
}

// ExternalSpec embeds all the external specs required for the app.
type ExternalSpec struct {
	Zookeeper ZookeeperSpec `yaml:"zookeeper"`
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

// NodeK8sConfigGroupSpec embeds all the k8s specific configuration.
type K8sConfigGroupSpec struct {
	Name               string            `json:"name"`
	Volumes            []v1.Volume       `json:"volumes,omitempty"`
	VolumeMount        []v1.VolumeMount  `json:"volumeMount,omitempty"`
	Image              string            `json:"image"`
	ImagePullPolicy    v1.PullPolicy     `json:"imagePullPolicy,omitempty"`
	ServiceAccountName string            `json:"serviceAccountName,omitempty"`
	Env                []v1.EnvVar       `json:"env,omitempty"`
	Tolerations        []v1.Toleration   `json:"tolerations,omitempty"`
	PodMetadata        Metadata          `json:"podMetadata,omitempty"`
	StorageConfig      []StorageConfig   `json:"storageConfig,omitempty"`
	NodeSelector       map[string]string `json:"nodeSelector,omitempty"`
	Service            *v1.ServiceSpec   `json:"service,omitempty"`
}

type Metadata struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
}

type StorageConfig struct {
	Name      string                       `json:"name"`
	MountPath string                       `json:"mountPath"`
	PvcSpec   v1.PersistentVolumeClaimSpec `json:"spec"`
}

// NodeAppConfigGroupSpec embeds all the app configs specific to each nodetype or common to all nodetypes.
type MyAppConfigGroupSpec struct {
	CommonRuntimeProperties string `yaml:"commonRuntimeProperties"`
	RuntimeProperties       string `yaml:"runtimeProperties"`
	OverrideProperties      string `yaml:"overrideProperties"`
}

// NodeSpec maps configs to a node of a specific nodeType.
type NodeSpec struct {
	Name                 string `json:"name"`
	Kind                 string `json:"kind"`
	NodeType             string `json:"nodeType"`
	Replicas             int    `json:"replicas"`
	K8sConfigGroup       string `json:"k8sConfigGroup"`
	MyAppConfigGroupName string `json:"myAppConfigGroup"`
}
