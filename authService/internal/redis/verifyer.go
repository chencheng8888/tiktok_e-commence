package redis

import (
	"errors"
	"github.com/chencheng8888/tiktok_e-commence/authService/internal/conf"
	"github.com/go-redis/redis"
	"time"
)

var (
	ErrTokenNotFound = errors.New("don't find the token")
	ErrTokenInvalid  = errors.New("the token is invalid")
)

type TokenVerifyer interface {
	VerifyJwtToken(tokenString, jwtSecret string) (int32, error)
}

type Verifyer struct {
	client *redis.Client
	t      TokenVerifyer
	cf     *conf.Token
}

func NewVerifier(client *redis.Client, t TokenVerifyer, cf *conf.Token) *Verifyer {
	return &Verifyer{client: client, t: t, cf: cf}
}

// VerifyToken 检查token是否合法
// 如果不合法就返回false
// 如果合法,异步为token续期
func (v *Verifyer) VerifyToken(tokenString string) (bool, error) {
	userID, err := v.t.VerifyJwtToken(tokenString, v.cf.Secret)
	if err != nil {
		return false, err
	}

	storedToken, err := v.client.Get(GenerateKey(userID)).Result()
	// 如果该token未找到
	if errors.Is(err, redis.Nil) {
		return false, ErrTokenNotFound
	}
	// 如果传入的token与查到的token不一致也报错
	if storedToken != tokenString {
		return false, ErrTokenInvalid
	}

	// 异步去为token续期,根据剩余时间的多少
	go func() {
		// 添加重试机制
		for i := 0; i < 2; i++ {
			err := v.renewalToken(userID, storedToken)
			if err == nil {
				break
			}
		}
	}()
	return true, nil
}

// 为token续期逻辑
func (v *Verifyer) renewalToken(userID int32, token string) error {
	key := GenerateKey(userID)

	// verify的时候存在,但renewal的时候不存在,说明原本就是快要过期的状态,这时应该直接续期
	if !v.checkKeyExist(key) {
		err := v.client.Set(key, token, v.cf.Expiration.AsDuration()).Err()
		return err
	}

	//获取userID的剩余过期时间
	ttl, err := v.getTTLOfkey(key)
	if err != nil {
		return err
	}

	// 检验剩余时间的程度
	// 根据不同的程度,来增加剩余过期时间
	rich, newTTL := v.checkTimeIfRich(ttl)
	if !rich {
		err = v.client.Set(key, token, newTTL).Err()
		return err
	}
	return nil
}

// 检查key是否存在
func (v *Verifyer) checkKeyExist(key string) bool {
	exists, err := v.client.Exists(key).Result()
	if err != nil {
		return false
	}
	if exists > 0 {
		return true
	}
	return false
}

// 获取key的TTL
func (v *Verifyer) getTTLOfkey(key string) (time.Duration, error) {
	return v.client.TTL(key).Result()
}

// 检查剩余过期时间是否充裕
// 如果充裕返回true
// 不充裕返回false和应该设置的过期时间
func (v *Verifyer) checkTimeIfRich(ttl time.Duration) (bool, time.Duration) {
	expire := v.cf.Expiration.AsDuration()
	if ttl < expire/4 {
		return false, ttl + 3*expire/4
	} else if ttl < expire/2 {
		return false, ttl + expire/2
	} else if ttl < 3*expire/4 {
		return false, ttl + expire/4
	}
	return true, ttl
}
