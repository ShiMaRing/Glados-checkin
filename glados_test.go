package main

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestBuildCookies(t *testing.T) {
	configInit()
	cookie := viper.GetString("cookie") //Gets the configured cookie
	cookies := buildCookies(cookie)
	fmt.Println(len(cookies))
	for _, v := range cookies {
		fmt.Println(v.Name)
	}

}
