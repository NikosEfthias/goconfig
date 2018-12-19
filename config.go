package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var _defaults map[string]string

//Load load the initial config
func Load(defaults map[string]string) error { //{{{
	f, err := os.Open("conf.json")
	if nil != err {
		switch {
		case os.IsNotExist(err):
			d, _ := json.MarshalIndent(defaults, "", "	")
			err := ioutil.WriteFile("conf.json", d, 0644)
			if nil != err {
				return err
			}
		default:
			return err
		}
		_defaults = defaults
		return nil
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&_defaults)
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
	return _defaults[k]
} //}}}
