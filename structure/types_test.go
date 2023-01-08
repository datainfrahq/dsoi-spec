package v1beta1

import (
	"io/ioutil"
	"testing"

	"sigs.k8s.io/yaml"
)

func TestSpec(t *testing.T) {
	yfile, err := ioutil.ReadFile("myApp.yaml")

	if err != nil {
		t.Fatal(err)
	}

	var spec MyApp

	t.Logf(string(yfile))

	t.Logf("%+v", spec)
	err = yaml.Unmarshal(yfile, &spec)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(spec)
}
