package infra

import (
	"fmt"
	"log"
	"notchman8600/authentication-provider/interfaces/database"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type RedisHandler struct {
	Conn *redis.Client
}

const Nil = redis.Nil

const (
	tokenKey      = "TOKEN_"
	tokenDuration = time.Hour * 24
)

//180秒話さなかったら警告
//テストでは10秒

func makeUniqueId() (uid_str string, err error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	uid_str = uid.String()
	return
}

func NewRedisHandler() database.RedisHandler {
	redisPath, err := NewEnvHandler().ReadEnv("REDIS_HOST")
	if err != nil {
		log.Println(errors.WithStack(err))
		return nil
	}
	redisHandler := new(RedisHandler)
	redisHandler.Conn, err = New(redisPath)
	if err != nil {
		log.Println(errors.WithStack(err))
		return nil
	}
	return redisHandler
}

func (handler *RedisHandler) AddValue(target string, duration time.Duration) (err error) {
	err = handler.Conn.Get(target).Err()
	if err == redis.Nil {
		log.Printf("%s does not exist. creating now...\n", target)

		//とりあえず永続化
		err = handler.Conn.Set(target, 1, duration).Err()
		if err != nil {
			return errors.Wrap(err, "Failed to set client")
		}
	} else if err != nil {
		return errors.Wrapf(err, "Failed to get %s", target)

	} else {
		_, err := handler.Conn.Incr(target).Result()
		if err != nil {
			return errors.Wrapf(err, "Failed to incr %s", target)
		}
	}
	return nil
}

func (handler *RedisHandler) DeclValue(target string) (int, error) {

	currentNum, err := handler.Conn.Decr(target).Result()
	if err != nil {
		return 0, errors.Wrap(err, "Failed to decr value")
	}
	return int(currentNum), nil
}

func (handler *RedisHandler) Count(target string) (value int, err error) {
	//Todo: ここの決め打ちを解消する
	err = handler.AddValue(target, time.Hour*1)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to add connection")
	}

	value, err = handler.DeclValue(target)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to decl connection")
	}
	return
}

func (handler *RedisHandler) Get(key string) (message string, err error) {
	//キーの存在可否の修正
	err = handler.Conn.Get(key).Err()
	if err == redis.Nil {
		fmt.Println("Key does not exist")
		return
	} else if err != nil {
		log.Println(err)
		return
	} else {
		message, err = handler.Conn.Get(key).Result()

		if err != nil {
			log.Println(errors.WithStack(err))
		}
	}
	return
}

func (handler *RedisHandler) Set(value string) (key string, err error) {
	key, err = makeUniqueId()
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	_, err = handler.Conn.Set(key, value, 0).Result()
	if err != nil {
		err = errors.WithStack(err)
		key = ""
		return
	}
	return
}

func (handler *RedisHandler) SetWithKey(key string, value string) (err error) {

	_, err = handler.Conn.Set(key, value, 0).Result()
	if err != nil {
		log.Println(err)
	}
	return
}
func (handler *RedisHandler) SetWithTime(key string, value string, duration time.Duration) (err error) {

	_, err = handler.Conn.Set(key, value, duration).Result()
	if err != nil {
		log.Println(err)
	}
	return
}
func (handler *RedisHandler) Update(key string, value string) (err error) {

	//キーの存在可否の修正
	err = handler.Conn.Get(key).Err()
	if err == redis.Nil {
		fmt.Println("Redis item is null")
		err = handler.SetWithKey(key, value)
		return
	} else if err != nil {
		log.Println(err)
		return
	}

	err = handler.Conn.Del(key).Err()
	if err != nil {
		fmt.Println("Redis item is failed to delete")
		err = errors.WithStack(err)
		return
	}
	return handler.SetWithKey(key, value)
	// _, err = handler.Conn.Get(key).Result()
	// if err == redis.Nil {
	// 	fmt.Println("key does not exist")
	// 	return
	// } else if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// err = handler.Conn.Del(key).Err()
	// if err != nil {
	// 	err = errors.WithStack(err)
	// 	return
	// }
	// return handler.SetWithKey(key, value)
}

func (handler *RedisHandler) SADD(key string, value string, time time.Duration) (flag int64, err error) {
	flag, err = handler.Conn.SAdd(key, value).Result()
	if flag == 0 {
		log.Println(value + " was already a member of the set!")
	}
	return
}

func (handler *RedisHandler) SREM(key string, value string) (err error) {
	err = handler.Conn.SRem(key, value).Err()
	if err != nil {
		log.Println(err)
	}
	return
}

func (handler *RedisHandler) SMEMBERS(key string) (lists []string, err error) {
	lists, err = handler.Conn.SMembers(key).Result()

	if err != nil {
		log.Println(err)
	}
	return
}

func New(dsn string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: "",
		DB:       0,
	})
	if err := client.Ping().Err(); err != nil {
		return nil, errors.WithStack(errors.Wrapf(err, "failed to ping redis server"))
	}
	return client, nil
}

func SetToken(cli *redis.Client, token string, userId int) error {
	if err := cli.Set(tokenKey+token, userId, tokenDuration).Err(); err != nil {
		return errors.WithStack(errors.Wrapf(err, "failed to set value"))

	}
	return nil
}

func GetIDByToken(cli *redis.Client, token string) (int, error) {
	v, err := cli.Get(tokenKey + token).Result()
	if err != nil {
		return 0, errors.WithStack(errors.Wrapf(err, "failed to get id from redis by token"))
	}
	id, err := strconv.Atoi(v)
	if err != nil {
		return 0, errors.WithStack(errors.Wrapf(err, "failed to convert string to inte"))
	}
	return id, nil
}
