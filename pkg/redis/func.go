package rds

import (
	"context"
	"errors"
	"net"
	"strconv"

	"github.com/go-redis/redis/v8"
)

//Flushkey method to Flush Key on redis
func Flushkey(ctx context.Context, rds *Rds) (string, error) {
	redisaddress := make([]string, 0)
	redisaddress = append(redisaddress, rds.Data.Address+":"+strconv.Itoa(rds.Data.Port))

	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:            redisaddress,
		MasterName:       rds.Data.Master,
		Password:         rds.Data.Password,
		SentinelPassword: rds.Data.SentinelPassword,
	})
	defer rdb.Close()
	respond, err := rdb.FlushAll(ctx).Result()
	if err != nil {
		return "", err
	}

	return respond, nil
}

//Flushkey method to Flush Key on redis
func FlushKeyAsync(ctx context.Context, rds *Rds) (string, error) {
	redisaddress := make([]string, 0)
	redisaddress = append(redisaddress, rds.Data.Address+":"+strconv.Itoa(rds.Data.Port))

	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:            redisaddress,
		MasterName:       rds.Data.Master,
		Password:         rds.Data.Password,
		SentinelPassword: rds.Data.SentinelPassword,
	})
	defer rdb.Close()
	respond, err := rdb.FlushAllAsync(ctx).Result()
	if err != nil {
		return "", err
	}

	return respond, nil
}

// FlushDB method to flush the specific DB redis
func FlushDB(ctx context.Context, rds *Rds) (string, error) {
	redisaddress := make([]string, 0)
	redisaddress = append(redisaddress, rds.Data.Address+":"+strconv.Itoa(rds.Data.Port))

	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:            redisaddress,
		MasterName:       rds.Data.Master,
		Password:         rds.Data.Password,
		SentinelPassword: rds.Data.SentinelPassword,
		DB:               rds.Data.Database,
	})
	defer rdb.Close()
	respond, err := rdb.FlushDB(ctx).Result()
	if err != nil {
		return "", err
	}

	return respond, nil
}

// FlushDB method to flush the specific DB redis
func FlushDbAsync(ctx context.Context, rds *Rds) (string, error) {
	redisaddress := make([]string, 0)
	redisaddress = append(redisaddress, rds.Data.Address+":"+strconv.Itoa(rds.Data.Port))

	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:            redisaddress,
		MasterName:       rds.Data.Master,
		Password:         rds.Data.Password,
		SentinelPassword: rds.Data.SentinelPassword,
		DB:               rds.Data.Database,
	})
	defer rdb.Close()
	respond, err := rdb.FlushDBAsync(ctx).Result()
	if err != nil {
		return "", err
	}

	return respond, nil
}

//Getkey method to GET KEY redis
func Getkey(ctx context.Context, rds *Rds) (string, error) {
	redisaddress := make([]string, 0)
	redisaddress = append(redisaddress, rds.Data.Address+":"+strconv.Itoa(rds.Data.Port))

	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:            redisaddress,
		MasterName:       rds.Data.Master,
		DB:               rds.Data.Database,
		Password:         rds.Data.Password,
		SentinelPassword: rds.Data.SentinelPassword,
	})
	defer rdb.Close()
	respond, err := rdb.Get(ctx, rds.Data.Key).Result()
	if err != nil {
		return "", err
	}

	return respond, nil
}

// Allkeys method to GET ALL KEYS on redis
func Allkeys(ctx context.Context, rds *Rds) ([]string, error) {
	redisaddress := make([]string, 0)
	redisaddress = append(redisaddress, rds.Data.Address+":"+strconv.Itoa(rds.Data.Port))

	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:            redisaddress,
		MasterName:       rds.Data.Master,
		DB:               rds.Data.Database,
		Password:         rds.Data.Password,
		SentinelPassword: rds.Data.SentinelPassword,
	})
	defer rdb.Close()
	respond, err := rdb.Keys(ctx, rds.Data.Key).Result()
	if err != nil {
		return nil, err
	}

	return respond, nil
}

//Setkeys method to SET KEY on redis
func Setkeys(ctx context.Context, rds *Rds) (string, error) {
	redisaddress := make([]string, 0)
	redisaddress = append(redisaddress, rds.Data.Address+":"+strconv.Itoa(rds.Data.Port))

	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:            redisaddress,
		MasterName:       rds.Data.Master,
		DB:               rds.Data.Database,
		Password:         rds.Data.Password,
		SentinelPassword: rds.Data.SentinelPassword,
	})
	defer rdb.Close()
	resp, err := rdb.Set(ctx, rds.Data.Key, rds.Data.Value, 0).Result()
	if err != nil {
		return "", err
	}

	return resp, nil
}

// GetMasterName function is function that you use to retrieve the master address of sentinel
func GetMasterName(ctx context.Context, rds *Rds) (string, error) {
	redisaddress := make([]string, 0)
	redisaddress = append(redisaddress, rds.Data.Address+":"+strconv.Itoa(rds.Data.Port))

	rdb := redis.NewSentinelClient(&redis.Options{
		Addr:     rds.Data.Address + ":" + strconv.Itoa(rds.Data.Port),
		Password: rds.Data.SentinelPassword,
	})
	defer rdb.Close()
	resp, err := rdb.GetMasterAddrByName(ctx, rds.Data.Master).Result()

	if err != nil {
		return "", err
	}
	// ip is always at index zero
	hostname, err := net.LookupAddr(resp[0])
	if err != nil {
		return "", err
	}
	return hostname[0], nil
}

//Delkey method to DEL KEY redis
func Delkey(ctx context.Context, rds *Rds) error {
	redisaddress := make([]string, 0)
	redisaddress = append(redisaddress, rds.Data.Address+":"+strconv.Itoa(rds.Data.Port))

	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:            redisaddress,
		MasterName:       rds.Data.Master,
		DB:               rds.Data.Database,
		Password:         rds.Data.Password,
		SentinelPassword: rds.Data.SentinelPassword,
	})
	defer rdb.Close()
	resp, _ := rdb.Del(ctx, rds.Data.Key).Result()
	if resp != 1 {
		return errors.New("Key not deleted / doesn't exist")
	}

	return nil
}

//Prefixdelkey method to DEL KEY by PREFIX redis
func Prefixdelkey(ctx context.Context, rds *Rds) (int32, error) {
	var deletedKey int32
	redisaddress := make([]string, 0)
	redisaddress = append(redisaddress, rds.Data.Address+":"+strconv.Itoa(rds.Data.Port))

	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:            redisaddress,
		MasterName:       rds.Data.Master,
		DB:               rds.Data.Database,
		Password:         rds.Data.Password,
		SentinelPassword: rds.Data.SentinelPassword,
	})
	defer rdb.Close()
	iter := rdb.Scan(ctx, 0, rds.Data.Key, 0).Iterator()

	for iter.Next(ctx) {
		err := rdb.Del(ctx, iter.Val()).Err()
		if err != nil {
			return -1, err
		}
		deletedKey += 1
	}
	if err := iter.Err(); err != nil {
		return -1, err
	}
	return deletedKey, nil
}
