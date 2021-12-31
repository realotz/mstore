package myjwt

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/realotz/mstore/pkg/myctx"
	"strings"
	"time"
)

// Token jwt
type PublicToken struct {
	publicKey  *rsa.PublicKey

}

type PrivateToken struct {
	PublicToken
	privateKey *rsa.PrivateKey
}

func New(privateKeyByte, publicKeyByte []byte) (*PrivateToken, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)
	if err != nil {
		return nil, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
	if err != nil {
		return nil, err
	}
	return &PrivateToken{
		PublicToken: PublicToken{
			publicKey: publicKey,
		},
		privateKey:  privateKey,
	}, nil
}

func NewPublic(publicKeyByte []byte) (*PublicToken, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)
	if err != nil {
		return nil, err
	}
	return &PublicToken{
		publicKey: publicKey,
	}, nil
}

//Decode 解码
func (srv *PublicToken) Decode(tokenStr string) (*myctx.JwtCustomClaims, error) {
	if srv.publicKey == nil {
		return nil, errors.New("private key is nill")
	}
	tokenStr = strings.ReplaceAll(tokenStr, "Bearer ", "")
	t, err := jwt.ParseWithClaims(tokenStr, &myctx.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return srv.publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 解密转换类型并返回
	if claims, ok := t.Claims.(*myctx.JwtCustomClaims); ok && t.Valid {
		return claims, nil
	}
	return nil, err
}

func (srv *PrivateToken) Encode(userIssues string,up myctx.JwtProfile, ut string) (string, error) {
	if srv.privateKey == nil {
		return "", errors.New("private key is nill")
	}
	claims := myctx.JwtCustomClaims{
		JwtProfile: up,
		UserType:   ut,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			NotBefore: time.Now().Unix() - 1000,
			IssuedAt:  time.Now().Unix(),
			Issuer:    userIssues,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tk, err := jwtToken.SignedString(srv.privateKey)
	if err != nil {
		return "", err
	}
	return tk, nil
}
