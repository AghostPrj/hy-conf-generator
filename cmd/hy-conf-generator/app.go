/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2022/6/22 12:13
 * @Desc:
 */

package main

import (
	"encoding/json"
	"fmt"
	"github.com/AghostPrj/hy-conf-generator/internal/object"
	"github.com/spf13/viper"
	"strings"
)

var (
	confKeys = [15][]string{
		{object.HyListenKey, object.HyListenKey},
		{object.HyProtocolKey, object.HyProtocolKey},
		{object.HyCertKey, object.HyCertKey},
		{object.HyUpMbpsKey, object.HyUpMbpsKey},
		{object.HyDownMbpsKey, object.HyDownMbpsKey},
		{object.HyObfsKey, object.HyObfsKey},
		{object.HyRcvWindowConnKey, object.HyRcvWindowConnKey},
		{object.HyAlpnKey, object.HyAlpnKey},
		{object.HyKeyKey, object.HyKeyKey},
		{object.HyRcvWindowClientKey, object.HyRcvWindowClientKey},
		{object.HyMaxConnClientKey, object.HyMaxConnClientKey},
		{object.HyResolverKey, object.HyResolverKey},
		{object.HyResolvePreferenceKey, object.HyResolvePreferenceKey},
		{object.HyAuthModeKey, object.HyAuthModeKey},
		{object.HyAuthConfStringKey, object.HyAuthConfStringKey},
	}
)

func main() {
	for _, s := range confKeys {
		err := viper.BindEnv(s...)
		if err != nil {
			panic("viper bind env error")
		}
	}

	viper.SetDefault(object.HyListenKey, object.HyListenDefault)
	viper.SetDefault(object.HyProtocolKey, object.HyProtocolDefault)
	viper.SetDefault(object.HyUpMbpsKey, object.HyUpMbpsDefault)
	viper.SetDefault(object.HyDownMbpsKey, object.HyDownMbpsDefault)
	viper.SetDefault(object.HyRcvWindowConnKey, object.HyRcvWindowConnDefault)
	viper.SetDefault(object.HyRcvWindowClientKey, object.HyRcvWindowClientDefault)
	viper.SetDefault(object.HyMaxConnClientKey, object.HyMaxConnClientDefault)
	viper.SetDefault(object.HyResolverKey, object.HyResolverDefault)
	viper.SetDefault(object.HyResolvePreferenceKey, object.HyResolvePreferenceDefault)

	viper.AllowEmptyEnv(false)
	viper.AutomaticEnv()

	conf := object.HyConf{
		Listen:            viper.GetString(object.HyListenKey),
		Protocol:          viper.GetString(object.HyProtocolKey),
		Cert:              viper.GetString(object.HyCertKey),
		Key:               viper.GetString(object.HyKeyKey),
		UpMbps:            viper.GetInt32(object.HyUpMbpsKey),
		DownMbps:          viper.GetInt32(object.HyDownMbpsKey),
		Obfs:              viper.GetString(object.HyObfsKey),
		Alpn:              viper.GetString(object.HyAlpnKey),
		RcvWindowConn:     viper.GetInt64(object.HyRcvWindowConnKey),
		RcvWindowClient:   viper.GetInt64(object.HyRcvWindowClientKey),
		MaxConnClient:     viper.GetInt32(object.HyMaxConnClientKey),
		Resolver:          viper.GetString(object.HyResolverKey),
		ResolvePreference: viper.GetString(object.HyResolvePreferenceKey),
		Auth:              nil,
	}

	if len(viper.GetString(object.HyAuthModeKey)) > 0 {

		authConf := object.HyAuthConf{
			Mode:   viper.GetString(object.HyAuthModeKey),
			Config: nil,
		}

		if len(viper.GetString(object.HyAuthConfStringKey)) > 0 {
			configs := strings.Split(viper.GetString(object.HyAuthConfStringKey), ",")
			authConf.Config = configs
		}

		conf.Auth = &authConf

	}

	bytes, err := json.Marshal(&conf)
	if err != nil {
		panic("generate config json error")
	}

	fmt.Printf("%s\n", string(bytes))

}
