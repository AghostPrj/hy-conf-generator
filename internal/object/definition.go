/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2022/6/22 12:14
 * @Desc:
 */

package object

const (
	HyListenKey     = "HY_LISTEN"
	HyListenDefault = "0.0.0.0:12345"

	HyProtocolKey     = "HY_PROTOCOL"
	HyProtocolDefault = "udp"

	HyCertKey = "HY_CERT"
	HyKeyKey  = "HY_KEY"

	HyUpMbpsKey     = "HY_UP_MBPS"
	HyUpMbpsDefault = int32(1000)

	HyDownMbpsKey     = "HY_DOWN_MBPS"
	HyDownMbpsDefault = int32(1000)

	HyObfsKey = "HY_OBFS"
	HyAlpnKey = "HY_ALPN"

	HyRcvWindowConnKey     = "HY_RCV_WINDOW_CONN"
	HyRcvWindowConnDefault = int64(15728640)

	HyRcvWindowClientKey     = "HY_RCV_WINDOW_CLIENT"
	HyRcvWindowClientDefault = int64(67108864)

	HyMaxConnClientKey     = "HY_MAX_CONN_CLIENT"
	HyMaxConnClientDefault = int32(4096)

	HyResolverKey     = "HY_RESOLVER"
	HyResolverDefault = "8.8.8.8:53"

	HyResolvePreferenceKey     = "HY_RESOLVE_PREFERENCE"
	HyResolvePreferenceDefault = "4"

	HyAuthModeKey       = "HY_AUTH_MODE"
	HyAuthConfStringKey = "HY_AUTH_CONFIG"
)

type HyConf struct {
	Listen            string      `json:"listen"`
	Protocol          string      `json:"protocol,omitempty"`
	Cert              string      `json:"cert,omitempty"`
	Key               string      `json:"key,omitempty"`
	UpMbps            int32       `json:"up_mbps,omitempty"`
	DownMbps          int32       `json:"down_mbps,omitempty"`
	Obfs              string      `json:"obfs,omitempty"`
	Alpn              string      `json:"alpn,omitempty"`
	RcvWindowConn     int64       `json:"recv_window_conn,omitempty"`
	RcvWindowClient   int64       `json:"recv_window_client,omitempty"`
	MaxConnClient     int32       `json:"max_conn_client,omitempty"`
	Resolver          string      `json:"resolver,omitempty"`
	ResolvePreference string      `json:"resolve_preference,omitempty"`
	Auth              *HyAuthConf `json:"auth,omitempty"`
}

type HyAuthConf struct {
	Mode   string   `json:"mode"`
	Config []string `json:"config,omitempty"`
}
