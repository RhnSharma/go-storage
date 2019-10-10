// Code generated by go generate via internal/cmd/meta_gen; DO NOT EDIT.
package qingstor

import (
	"github.com/Xuanwo/storage"
	"github.com/Xuanwo/storage/types"
)

// CapabilityRead    = true
// CapabilityWrite   = true
// CapabilityFile    = true
// CapabilityStream  = false
// CapabilitySegment = true
const capability = types.Capability(83)

// Capability implements Storager.Capability().
func (c *Client) Capability() types.Capability {
	return capability
}

var allowdPairs = map[string]map[string]struct{}{
	storage.ActionAbortSegment:    {},
	storage.ActionCompleteSegment: {},
	storage.ActionCopy:            {},
	storage.ActionCreateDir: {
		"location": struct{}{},
	},
	storage.ActionDelete:      {},
	storage.ActionInitSegment: {},
	storage.ActionListDir:     {},
	storage.ActionMove:        {},
	storage.ActionReadFile:    {},
	storage.ActionReadSegment: {},
	storage.ActionReadStream:  {},
	storage.ActionStat:        {},
	storage.ActionWriteFile: {
		"checksum":      struct{}{},
		"storage_class": struct{}{},
	},
	storage.ActionWriteSegment: {},
	storage.ActionWriteStream:  {},
}

// IsPairAvailable implements Storager.IsPairAvailable().
func (c *Client) IsPairAvailable(action, pair string) bool {
	if _, ok := allowdPairs[action]; !ok {
		return false
	}
	if _, ok := allowdPairs[action][pair]; !ok {
		return false
	}
	return true
}

type pairAbortSegment struct {
}

func parsePairAbortSegment(opts ...*types.Pair) *pairAbortSegment {
	result := &pairAbortSegment{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowdPairs[storage.ActionAbortSegment]; !ok {
			continue
		}
		if _, ok := allowdPairs[storage.ActionAbortSegment][v.Key]; !ok {
			continue
		}
		values[v.Key] = v
	}
	return result
}

type pairCompleteSegment struct {
}

func parsePairCompleteSegment(opts ...*types.Pair) *pairCompleteSegment {
	result := &pairCompleteSegment{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowdPairs[storage.ActionCompleteSegment]; !ok {
			continue
		}
		if _, ok := allowdPairs[storage.ActionCompleteSegment][v.Key]; !ok {
			continue
		}
		values[v.Key] = v
	}
	return result
}

type pairCopy struct {
}

func parsePairCopy(opts ...*types.Pair) *pairCopy {
	result := &pairCopy{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowdPairs[storage.ActionCopy]; !ok {
			continue
		}
		if _, ok := allowdPairs[storage.ActionCopy][v.Key]; !ok {
			continue
		}
		values[v.Key] = v
	}
	return result
}

type pairCreateDir struct {
	HasLocation bool
	Location    string
}

func parsePairCreateDir(opts ...*types.Pair) *pairCreateDir {
	result := &pairCreateDir{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowdPairs[storage.ActionCreateDir]; !ok {
			continue
		}
		if _, ok := allowdPairs[storage.ActionCreateDir][v.Key]; !ok {
			continue
		}
		values[v.Key] = v
	}
	if v, ok := values["location"]; !ok {
		result.HasLocation = true
		result.Location = v.(string)
	}
	return result
}

type pairDelete struct {
}

func parsePairDelete(opts ...*types.Pair) *pairDelete {
	result := &pairDelete{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowdPairs[storage.ActionDelete]; !ok {
			continue
		}
		if _, ok := allowdPairs[storage.ActionDelete][v.Key]; !ok {
			continue
		}
		values[v.Key] = v
	}
	return result
}

type pairInitSegment struct {
}

func parsePairInitSegment(opts ...*types.Pair) *pairInitSegment {
	result := &pairInitSegment{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowdPairs[storage.ActionInitSegment]; !ok {
			continue
		}
		if _, ok := allowdPairs[storage.ActionInitSegment][v.Key]; !ok {
			continue
		}
		values[v.Key] = v
	}
	return result
}

type pairListDir struct {
}

func parsePairListDir(opts ...*types.Pair) *pairListDir {
	result := &pairListDir{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowdPairs[storage.ActionListDir]; !ok {
			continue
		}
		if _, ok := allowdPairs[storage.ActionListDir][v.Key]; !ok {
			continue
		}
		values[v.Key] = v
	}
	return result
}

type pairMove struct {
}

func parsePairMove(opts ...*types.Pair) *pairMove {
	result := &pairMove{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowdPairs[storage.ActionMove]; !ok {
			continue
		}
		if _, ok := allowdPairs[storage.ActionMove][v.Key]; !ok {
			continue
		}
		values[v.Key] = v
	}
	return result
}

type pairReadFile struct {
}

func parsePairReadFile(opts ...*types.Pair) *pairReadFile {
	result := &pairReadFile{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowdPairs[storage.ActionReadFile]; !ok {
			continue
		}
		if _, ok := allowdPairs[storage.ActionReadFile][v.Key]; !ok {
			continue
		}
		values[v.Key] = v
	}
	return result
}

type pairReadSegment struct {
}

func parsePairReadSegment(opts ...*types.Pair) *pairReadSegment {
	result := &pairReadSegment{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowdPairs[storage.ActionReadSegment]; !ok {
			continue
		}
		if _, ok := allowdPairs[storage.ActionReadSegment][v.Key]; !ok {
			continue
		}
		values[v.Key] = v
	}
	return result
}

type pairReadStream struct {
}

func parsePairReadStream(opts ...*types.Pair) *pairReadStream {
	result := &pairReadStream{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowdPairs[storage.ActionReadStream]; !ok {
			continue
		}
		if _, ok := allowdPairs[storage.ActionReadStream][v.Key]; !ok {
			continue
		}
		values[v.Key] = v
	}
	return result
}

type pairStat struct {
}

func parsePairStat(opts ...*types.Pair) *pairStat {
	result := &pairStat{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowdPairs[storage.ActionStat]; !ok {
			continue
		}
		if _, ok := allowdPairs[storage.ActionStat][v.Key]; !ok {
			continue
		}
		values[v.Key] = v
	}
	return result
}

type pairWriteFile struct {
	HasChecksum     bool
	Checksum        string
	HasStorageClass bool
	StorageClass    string
}

func parsePairWriteFile(opts ...*types.Pair) *pairWriteFile {
	result := &pairWriteFile{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowdPairs[storage.ActionWriteFile]; !ok {
			continue
		}
		if _, ok := allowdPairs[storage.ActionWriteFile][v.Key]; !ok {
			continue
		}
		values[v.Key] = v
	}
	if v, ok := values["checksum"]; !ok {
		result.HasChecksum = true
		result.Checksum = v.(string)
	}
	if v, ok := values["storage_class"]; !ok {
		result.HasStorageClass = true
		result.StorageClass = v.(string)
	}
	return result
}

type pairWriteSegment struct {
}

func parsePairWriteSegment(opts ...*types.Pair) *pairWriteSegment {
	result := &pairWriteSegment{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowdPairs[storage.ActionWriteSegment]; !ok {
			continue
		}
		if _, ok := allowdPairs[storage.ActionWriteSegment][v.Key]; !ok {
			continue
		}
		values[v.Key] = v
	}
	return result
}

type pairWriteStream struct {
}

func parsePairWriteStream(opts ...*types.Pair) *pairWriteStream {
	result := &pairWriteStream{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowdPairs[storage.ActionWriteStream]; !ok {
			continue
		}
		if _, ok := allowdPairs[storage.ActionWriteStream][v.Key]; !ok {
			continue
		}
		values[v.Key] = v
	}
	return result
}
