package v1beta1

// MyApp defines the structure for an app spec.
type MyAppAutoScaling struct {
	Spec MyAppAutoScalingSpec `yaml:"spec"`
}

type MyAppAutoScalingSpec struct {
	External ExternalSpec `yaml:"external"`
}

type ExternalSpec struct {
	Spec ExternalConfig `yaml:"spec"`
}

type ExternalConfig struct {
	URL string `yaml:"url"`
}
