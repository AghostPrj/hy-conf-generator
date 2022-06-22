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
	confKeys = [24][]string{
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
		{object.HyServerKey, object.HyServerKey},
		{object.HyAuthStrKey, object.HyAuthStrKey},
		{object.HyServerNameKey, object.HyServerNameKey},
		{object.HyRcvWindowKey, object.HyRcvWindowKey},
		{object.HySocks5PortKey, object.HySocks5PortKey},
		{object.HyHttpPortKey, object.HyHttpPortKey},
		{object.HyProxyUsernameKey, object.HyProxyUsernameKey},
		{object.HyProxyPasswordKey, object.HyProxyPasswordKey},
		{object.HyProxyTimeoutKey, object.HyProxyTimeoutKey},
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
		Protocol:          viper.GetString(object.HyProtocolKey),
		Cert:              viper.GetString(object.HyCertKey),
		Key:               viper.GetString(object.HyKeyKey),
		UpMbps:            viper.GetInt32(object.HyUpMbpsKey),
		DownMbps:          viper.GetInt32(object.HyDownMbpsKey),
		Obfs:              viper.GetString(object.HyObfsKey),
		Alpn:              viper.GetString(object.HyAlpnKey),
		RcvWindowConn:     viper.GetInt64(object.HyRcvWindowConnKey),
		MaxConnClient:     viper.GetInt32(object.HyMaxConnClientKey),
		Resolver:          viper.GetString(object.HyResolverKey),
		ResolvePreference: viper.GetString(object.HyResolvePreferenceKey),

		Server:     viper.GetString(object.HyServerKey),
		AuthStr:    viper.GetString(object.HyAuthStrKey),
		ServerName: viper.GetString(object.HyServerNameKey),

		Auth:   nil,
		Socks5: nil,
		Http:   nil,
	}

	if viper.GetString(object.HyListenKey) != "---" {
		conf.Listen = viper.GetString(object.HyListenKey)
	}

	if viper.GetInt64(object.HyRcvWindowClientKey) != 0 {
		conf.RcvWindowClient = viper.GetInt64(object.HyRcvWindowClientKey)
	}

	if viper.GetInt64(object.HyRcvWindowClientKey) != 0 {
		conf.RcvWindowClient = viper.GetInt64(object.HyRcvWindowClientKey)
	}

	if viper.GetInt64(object.HyMaxConnClientKey) != 0 {
		conf.MaxConnClient = viper.GetInt32(object.HyMaxConnClientKey)
	}

	if viper.GetInt32(object.HySocks5PortKey) > 0 && viper.GetInt32(object.HySocks5PortKey) < 65536 {
		socks5Conf := object.Socks5Conf{
			Listen:   "0.0.0.0:" + viper.GetString(object.HySocks5PortKey),
			Timeout:  viper.GetInt32(object.HyProxyTimeoutKey),
			User:     viper.GetString(object.HyProxyUsernameKey),
			Password: viper.GetString(object.HyProxyPasswordKey),
		}
		conf.Socks5 = &socks5Conf
	}

	if viper.GetInt32(object.HyHttpPortKey) > 0 && viper.GetInt32(object.HyHttpPortKey) < 65536 &&
		viper.GetInt32(object.HyHttpPortKey) != viper.GetInt32(object.HySocks5PortKey) {
		httpConf := object.HttpConf{
			Listen:   "0.0.0.0:" + viper.GetString(object.HyHttpPortKey),
			Timeout:  viper.GetInt32(object.HyProxyTimeoutKey),
			User:     viper.GetString(object.HyProxyUsernameKey),
			Password: viper.GetString(object.HyProxyPasswordKey),
		}
		conf.Http = &httpConf
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
