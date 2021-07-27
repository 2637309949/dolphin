package svc

import (
	"io/ioutil"

	"github.com/medivhzhan/weapp/v2"
	"github.com/spf13/viper"
)

// Weapp defined TODO
type Weapp interface {
	Code2Session(code string) (*weapp.LoginResponse, error)
	GetAccessToken() (*weapp.TokenResponse, error)
	CreateQRCode(path string, width int) ([]byte, error)
}

// GetAccessToken defined TODO
func (svc *SvcHepler) CreateQRCode(path string, width int) ([]byte, error) {
	creator := weapp.QRCodeCreator{
		Path:  path,
		Width: width,
	}
	resp, res, err := creator.Create("access-token")
	if err != nil {
		return []byte{}, nil
	}
	if err := res.GetResponseError(); err != nil {
		return []byte{}, nil
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// Code2Session defined TODO
func (svc *SvcHepler) Code2Session(code string) (*weapp.LoginResponse, error) {
	appid := viper.GetString("weapp.appid")
	secret := viper.GetString("weapp.secret")
	res, err := weapp.Login(appid, secret, code)
	if err != nil {
		return nil, err
	}
	if err := res.GetResponseError(); err != nil {
		return nil, err
	}
	return res, err
}

// GetAccessToken defined TODO
func (svc *SvcHepler) GetAccessToken() (*weapp.TokenResponse, error) {
	appid := viper.GetString("weapp.appid")
	secret := viper.GetString("weapp.secret")
	res, err := weapp.GetAccessToken(appid, secret)
	if err != nil {
		return nil, err
	}
	if err := res.GetResponseError(); err != nil {
		return nil, err
	}
	return res, err
}
