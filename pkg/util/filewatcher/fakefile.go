package filewatcher

type FakeWatcher struct {
	FakeContents string
	FakeVersion  int
}

func (f *FakeWatcher) Contents() string {
	return f.FakeContents
}
func (f *FakeWatcher) Version() int {
	return f.FakeVersion
}
