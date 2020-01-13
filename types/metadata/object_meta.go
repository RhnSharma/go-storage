// Code generated by go generate internal/cmd/metadata; DO NOT EDIT.
package metadata

import (
	"github.com/Xuanwo/storage/pkg/storageclass"
)

var _ storageclass.Type

// All available metadata.
const (
	ObjectMetaContentMD5   = "content-md5"
	ObjectMetaContentType  = "content-type"
	ObjectMetaETag         = "etag"
	ObjectMetaStorageClass = "storage-class"
)

// GetContentMD5 will get content-md5 value from metadata.
func (m ObjectMeta) GetContentMD5() (string, bool) {
	v, ok := m.m[ObjectMetaContentMD5]
	if !ok {
		return "", false
	}
	return v.(string), true
}

// MustGetContentMD5 will get content-md5 value from metadata.
func (m ObjectMeta) MustGetContentMD5() string {
	return m.m[ObjectMetaContentMD5].(string)
}

// SetContentMD5 will set content-md5 value into metadata.
func (m ObjectMeta) SetContentMD5(v string) ObjectMeta {
	m.m[ObjectMetaContentMD5] = v
	return m
}

// GetContentType will get content-type value from metadata.
func (m ObjectMeta) GetContentType() (string, bool) {
	v, ok := m.m[ObjectMetaContentType]
	if !ok {
		return "", false
	}
	return v.(string), true
}

// MustGetContentType will get content-type value from metadata.
func (m ObjectMeta) MustGetContentType() string {
	return m.m[ObjectMetaContentType].(string)
}

// SetContentType will set content-type value into metadata.
func (m ObjectMeta) SetContentType(v string) ObjectMeta {
	m.m[ObjectMetaContentType] = v
	return m
}

// GetETag will get etag value from metadata.
func (m ObjectMeta) GetETag() (string, bool) {
	v, ok := m.m[ObjectMetaETag]
	if !ok {
		return "", false
	}
	return v.(string), true
}

// MustGetETag will get etag value from metadata.
func (m ObjectMeta) MustGetETag() string {
	return m.m[ObjectMetaETag].(string)
}

// SetETag will set etag value into metadata.
func (m ObjectMeta) SetETag(v string) ObjectMeta {
	m.m[ObjectMetaETag] = v
	return m
}

// GetStorageClass will get storage-class value from metadata.
func (m ObjectMeta) GetStorageClass() (storageclass.Type, bool) {
	v, ok := m.m[ObjectMetaStorageClass]
	if !ok {
		return "", false
	}
	return v.(storageclass.Type), true
}

// MustGetStorageClass will get storage-class value from metadata.
func (m ObjectMeta) MustGetStorageClass() storageclass.Type {
	return m.m[ObjectMetaStorageClass].(storageclass.Type)
}

// SetStorageClass will set storage-class value into metadata.
func (m ObjectMeta) SetStorageClass(v storageclass.Type) ObjectMeta {
	m.m[ObjectMetaStorageClass] = v
	return m
}
