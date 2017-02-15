# k8s-core-member
a handy utility to check whether user is a core member of kubernetes. user whose name is listed in any OWNERS file is considered as a core member.

# How to Build
1. `go get -u github.com/fatih/color`
2. `go get -u gopkg.in/yaml.v2`
3. `go build main.go`

# How to use
1. edit config.yaml, add users you want to check
2. `git clone https://github.com/kubernetes/kubernetes.git`
3. `go run main.go /path/to/kubernetes/codebase/`
