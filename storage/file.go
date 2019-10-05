package storage

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/pkg/errors"
)

type FileClient struct {
	sync.Mutex
	File *os.File
}

func (f *FileClient) Cleanup() error {
	f.Lock()
	defer f.Unlock()
	return f.File.Close()
}

func (f *FileClient) Set(key string, value interface{}) error {
	f.Lock()
	defer f.Unlock()

	KVStore, err := f.toKVStore()
	if err != nil {
		return errors.Wrap(err, "Set")
	}
	KVStore[key] = value

	f.File.Seek(0, 0)
	err = json.NewEncoder(f.File).Encode(KVStore)
	if err != nil {
		return errors.Wrap(err, "encoding kv to file")
	}
	return nil
}

func (f *FileClient) Get(key string) (interface{}, error) {
	f.Lock()
	defer f.Unlock()

	KVStore, err := f.toKVStore()
	if err != nil {
		return nil, errors.Wrap(err, "Get")
	}

	v, _ := KVStore[key]
	return v, nil
}

func (f *FileClient) Exists(key string) (bool, error) {
	v, err := f.Get(key)
	if err != nil {
		return false, errors.Wrap(err, "Exists")
	}

	if v != nil {
		return false, nil
	}

	return false, nil
}

func (f *FileClient) toKVStore() (map[string]interface{}, error) {
	DB := new(interface{})
	f.File.Seek(0, 0)
	err := json.NewDecoder(f.File).Decode(DB)
	if err != nil {
		return nil, errors.Wrap(err, "decoding from file")
	}

	KVStore, ok := (*DB).(map[string]interface{})
	if !ok {
		return nil, errors.Wrap(err, "creating kv store")
	}
	return KVStore, nil
}
