// Code generated by protoc-gen-goext. DO NOT EDIT.

package k8s

func (m *ListVersionsResponse) SetAvailableVersions(v []*AvailableVersions) {
	m.AvailableVersions = v
}

func (m *AvailableVersions) SetReleaseChannel(v ReleaseChannel) {
	m.ReleaseChannel = v
}

func (m *AvailableVersions) SetVersions(v []string) {
	m.Versions = v
}
