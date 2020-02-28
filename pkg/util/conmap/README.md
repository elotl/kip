This came about because I wanted to have concurrent maps but felt bad
about casting out of them.

To re-generate the gen-conmaps.go with updates from conmaps.go:

```bash
go get github.com/justnoise/genny
cd $GOPATH/src/github.com/justnoise/genny
go install
cd $GOPATH/src/github.com/elotl/cloud-instance-provider/pkg/util/conmap
go generate
```
