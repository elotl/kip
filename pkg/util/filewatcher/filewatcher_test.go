package filewatcher

import (
	"testing"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestStatPeriodRefresh(t *testing.T) {
	t.Parallel()
	f, closer := util.MakeTempFile("filewatcher")
	defer closer()

	c1 := "contents1\n"
	f.Write([]byte(c1))
	fw := New(f.Name())
	assert.Equal(t, c1, fw.Contents())
	c2 := "more data\n"
	// Filesystem time is only accurate to 1s
	time.Sleep(1250 * time.Millisecond)
	f.Write([]byte(c2))
	assert.Equal(t, c1, fw.Contents())
	fw.CheckPeriod = 1 * time.Nanosecond
	assert.Equal(t, c1+c2, fw.Contents())
}
