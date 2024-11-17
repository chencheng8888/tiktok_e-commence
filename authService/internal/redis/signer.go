package redis

import (
	"errors"
	"fmt"
	"github.com/chencheng8888/tiktok_e-commence/authService/internal/conf"
	token2 "github.com/chencheng8888/tiktok_e-commence/authService/internal/token"
	"github.com/go-redis/redis"
)

var (
	ErrSetKey = errors.New("redis key set failed")
)

type TokenGenerater interface {
	GenerateJwtToken(userID int32, jwtSecret string) (string, error)
}

type Signer struct {
	client *redis.Client
	t      TokenGenerater
	cf     *conf.Token
}

func NewSigner(client *redis.Client, t TokenGenerater, cf *conf.Token) *Signer {
	return &Signer{client: client, t: t, cf: cf}
}

func (s *Signer) SignToken(userID int32) (string, error) {
	key := GenerateKey(userID)
	//redis中没有该用户的token,则创建
	jwtToken, err := s.t.GenerateJwtToken(userID, s.cf.Secret)

	// 如果签发token失败,重试3遍
	if errors.Is(err, token2.ErrSignString) {
		for i := 0; i < 2; i++ {
			jwtToken, err = s.t.GenerateJwtToken(userID, s.cf.Secret)
			if err == nil {
				break
			}
		}
	}
	//重试三遍后仍失败就返回错误
	if err != nil {
		return "", err
	}

	err = s.client.Set(key, jwtToken, s.cf.Expiration.AsDuration()).Err()
	if err != nil {
		return "", fmt.Errorf("%w:%s", ErrSetKey, err.Error())
	}
	return jwtToken, nil
}
