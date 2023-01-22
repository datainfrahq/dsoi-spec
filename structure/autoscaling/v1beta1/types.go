package v1beta1

// MyApp defines the structure for an app spec.
type MyAppAutoScaling struct {
	Spec MyAppAutoScalingSpec `yaml:"spec"`
}

type MyAppAutoScalingSpec struct {
	External    ExternalSpec      `yaml:"external"`
	Autoscaling []AutoscalingSpec `yaml:"autoscaling"`
}

type AutoscalingSpec struct {
	Name string                `yaml:"name"`
	Spec AutoscalingConfigSpec `yaml:"spec"`
}

type AutoscalingConfigSpec struct {
	Enable     bool   `yaml:"enable"`
	MinReplica int    `yaml:"minReplica"`
	MaxReplica int    `yaml:"maxReplica"`
	Threshold  int    `yaml:"threshold"`
	CoolDown   string `yaml:"coolDown"`
}

type ExternalSpec struct {
	Spec ExternalConfig `yaml:"spec"`
}

type ExternalConfig struct {
	URL string `yaml:"url"`
}
