package link_test

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"

	"github.com/sagernet/sing-box/common/link"
)

func TestShadowSocks(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		link string
		want link.ShadowSocks
	}{
		{
			link: "ss://YWVzLTEyOC1nY206dGVzdA@192.168.100.1:8888#Example1",
			want: link.ShadowSocks{
				Address:  "192.168.100.1",
				Port:     8888,
				Ps:       "Example1",
				Method:   "aes-128-gcm",
				Password: "test",
			},
		},
		{
			link: "ss://cmM0LW1kNTpwYXNzd2Q@192.168.100.1:8888/?plugin=obfs-local%3Bobfs%3Dhttp%3Bobfs-host=abc.com#Example2",
			want: link.ShadowSocks{
				Address:    "192.168.100.1",
				Port:       8888,
				Ps:         "Example2",
				Method:     "rc4-md5",
				Password:   "passwd",
				Plugin:     "obfs-local",
				PluginOpts: "obfs=http;obfs-host=abc.com",
			},
		},
		{
			link: "ss://2022-blake3-aes-256-gcm:YctPZ6U7xPPcU%2Bgp3u%2B0tx%2FtRizJN9K8y%2BuKlW2qjlI%3D@192.168.100.1:8888#Example3",
			want: link.ShadowSocks{
				Address:  "192.168.100.1",
				Port:     8888,
				Ps:       "Example3",
				Method:   "2022-blake3-aes-256-gcm",
				Password: "YctPZ6U7xPPcU gp3u 0tx/tRizJN9K8y uKlW2qjlI=",
			},
		},
		{
			link: "ss://2022-blake3-aes-256-gcm:YctPZ6U7xPPcU%2Bgp3u%2B0tx%2FtRizJN9K8y%2BuKlW2qjlI%3D@192.168.100.1:8888/?plugin=v2ray-plugin%3Bserver&unsupported-arguments=should-be-ignored#Example3",
			want: link.ShadowSocks{
				Address:    "192.168.100.1",
				Port:       8888,
				Ps:         "Example3",
				Method:     "2022-blake3-aes-256-gcm",
				Password:   "YctPZ6U7xPPcU gp3u 0tx/tRizJN9K8y uKlW2qjlI=",
				Plugin:     "v2ray-plugin",
				PluginOpts: "server",
			},
		},
	}
	for _, tc := range testCases {
		u, err := url.Parse(tc.link)
		if err != nil {
			t.Fatal(err)
		}
		link, err := link.ParseShadowSocks(u)
		if err != nil {
			t.Error(err)
			return
		}
		if *link != tc.want {
			t.Errorf("want %v, got %v", tc.want, link)
		}
	}
}

func TestShadowSocksConvert(t *testing.T) {
	t.Parallel()
	testCases := []*link.ShadowSocks{
		{
			Ps:         "remarks",
			Method:     "aes-128-gcm",
			Password:   "password-密码",
			Address:    "192.168.1.1",
			Port:       443,
			Plugin:     "v2ray-plugin",
			PluginOpts: "server;tls;host=example.com",
		},
	}
	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprint("#", i), func(t *testing.T) {
			t.Parallel()
			uri, err := tc.URL()
			if err != nil {
				t.Fatal(err)
			}
			u, err := url.Parse(uri)
			if err != nil {
				t.Fatal(err)
			}
			link, err := link.ParseShadowSocks(u)
			if err != nil {
				t.Error(err)
				return
			}
			if !reflect.DeepEqual(link, tc) {
				t.Errorf("want %#v, got %#v", tc, link)
			}
		})
	}
}
