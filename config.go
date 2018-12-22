package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

var _defaults = map[string]string{}
var _l sync.Mutex

//Load load the initial config
func Load(defaults map[string]string) error { //{{{
	f, err := os.Open("conf.json")
	if nil != err {
		switch {
		case os.IsNotExist(err):
			err := _write_config()
			if nil != err {
				return err
			}
		default:
			return err
		}
		_l.Lock()
		_defaults = defaults
		_l.Unlock()
		return nil
	}
	defer f.Close()
	_l.Lock()
	err = json.NewDecoder(f).Decode(&_defaults)
	_l.Unlock()
	if nil != err {
		return err
	}
	return nil
} //}}}
//Get config
func Get(k string) string { //{{{
	if d := os.Getenv(k); d != "" {
		return d
	}
	_l.Lock()
	defer _l.Unlock()
	return _defaults[k]
} //}}}
func _write_config() error { //{{{
	_l.Lock()
	defer _l.Unlock()
	d, err := json.MarshalIndent(_defaults, "", " ")
	if nil != err {
		return err
	}
	return ioutil.WriteFile("conf.json", d, 0644)
} //}}}
func Set(k, v string) { //{{{
	Load(_defaults)
	_l.Lock()
	_defaults[k] = v
	_l.Unlock()
	_write_config()
} //}}}
func Get_or_set(k, v string) string { //{{{
	val := Get(k)
	if "" != val {
		return val
	}
	go Set(k, v)
	return v
} //}}}
