CLUSTER_PATH ?= structure/cluster/v1beta1
AUTOSCALING_PATH = "structure/autoscaling/v1beta1"

ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: test
test: #
	cd $(CLUSTER_PATH) && go test . -v
	cd $(AUTOSCALING_PATH) && go test . -v
     
