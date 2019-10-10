// Code generated by go generate internal/cmd/metdata_gen; DO NOT EDIT.
package types

// GetChecksum will get checksum value from metadata.
func (m Metadata) GetChecksum() (string, bool) {
	v, ok := m[Checksum]
	if !ok {
		return "", false
	}
	return v.(string), true
}

// SetChecksum will set checksum value into metadata.
func (m Metadata) SetChecksum(v string) {
	m[Checksum] = v
}

// GetLocation will get location value from metadata.
func (m Metadata) GetLocation() (string, bool) {
	v, ok := m[Location]
	if !ok {
		return "", false
	}
	return v.(string), true
}

// SetLocation will set location value into metadata.
func (m Metadata) SetLocation(v string) {
	m[Location] = v
}

// GetSize will get size value from metadata.
func (m Metadata) GetSize() (int64, bool) {
	v, ok := m[Size]
	if !ok {
		return 0, false
	}
	return v.(int64), true
}

// SetSize will set size value into metadata.
func (m Metadata) SetSize(v int64) {
	m[Size] = v
}

// GetStorageClass will get storage_class value from metadata.
func (m Metadata) GetStorageClass() (string, bool) {
	v, ok := m[StorageClass]
	if !ok {
		return "", false
	}
	return v.(string), true
}

// SetStorageClass will set storage_class value into metadata.
func (m Metadata) SetStorageClass(v string) {
	m[StorageClass] = v
}

// GetType will get type value from metadata.
func (m Metadata) GetType() (string, bool) {
	v, ok := m[Type]
	if !ok {
		return "", false
	}
	return v.(string), true
}

// SetType will set type value into metadata.
func (m Metadata) SetType(v string) {
	m[Type] = v
}
