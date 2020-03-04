// +build tools

// This package imports things required by build scripts, to force
// `dep` to see them as dependencies
package tools

import _ "k8s.io/code-generator"
