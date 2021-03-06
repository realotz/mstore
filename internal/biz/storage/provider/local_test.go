package provider

import (
	"context"
	"fmt"
	"github.com/realotz/mstore/pkg/utils"
	"testing"
)

func Test_localProvider_List(t *testing.T) {
	p,err := NewLocalProvider(utils.JsonBytes(map[string]string{"path":"/Users"}))
	if err!=nil{
		t.Error(err)
	}
	got, err := p.List(context.Background(),ListOption{
		Path:     "",
		HideFile: false,
		SortFlag: 0,
	})
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(got)
}

