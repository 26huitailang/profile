package auth

//
//import (
//	"reflect"
//	"testing"
//	"time"
//)
//
//func TestGenJWT(t *testing.T) {
//	type args struct {
//		name    string
//		isAdmin bool
//		key     []byte
//		exp     time.Duration
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    string
//		wantErr bool
//	}{
//		{name: "token expire", args: {"admin", true, "secret"}}
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := GenJWT(tt.args.name, tt.args.isAdmin, tt.args.key)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GenJWT() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("GenJWT() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestGetUserInfo(t *testing.T) {
//	type args struct {
//		c v4.Context
//	}
//	tests := []struct {
//		name string
//		args args
//		want JwtCustomClaims
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := GetUserInfo(tt.args.c); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetUserInfo() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
