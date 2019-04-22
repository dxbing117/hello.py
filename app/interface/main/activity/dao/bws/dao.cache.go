// Code generated by $GOPATH/src/go-common/app/tool/cache/gen. DO NOT EDIT.

/*
  Package bws is a generated cache proxy package.
  It is generated from:
  type _cache interface {
		//cache: -sync=true
		UsersMid(c context.Context, key int64) (*bwsmdl.Users, error)
		//cache: -sync=true
		UsersKey(c context.Context, key string) (*bwsmdl.Users, error)
		//cache: -sync=true
		Points(c context.Context, bid int64) (*bwsmdl.Points, error)
		//cache: -sync=true
		Achievements(c context.Context, bid int64) (*bwsmdl.Achievements, error)
		//cache: -sync=true
		UserAchieves(c context.Context, bid int64, key string) ([]*bwsmdl.UserAchieve, error)
		//cache: -sync=true
		UserPoints(c context.Context, bid int64, key string) ([]*bwsmdl.UserPoint, error)
		//cache: -sync=true
		AchieveCounts(c context.Context, bid int64, day string) ([]*bwsmdl.CountAchieves, error)
	}
*/

package bws

import (
	"context"

	bwsmdl "go-common/app/interface/main/activity/model/bws"
	"go-common/library/stat/prom"
)

var _ _cache

// UsersMid get data from cache if miss will call source method, then add to cache.
func (d *Dao) UsersMid(c context.Context, id int64) (res *bwsmdl.Users, err error) {
	addCache := true
	res, err = d.CacheUsersMid(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	if res != nil {
		prom.CacheHit.Incr("UsersMid")
		return
	}
	prom.CacheMiss.Incr("UsersMid")
	res, err = d.RawUsersMid(c, id)
	if err != nil {
		return
	}
	miss := res
	if !addCache {
		return
	}
	d.AddCacheUsersMid(c, id, miss)
	return
}

// UsersKey get data from cache if miss will call source method, then add to cache.
func (d *Dao) UsersKey(c context.Context, id string) (res *bwsmdl.Users, err error) {
	addCache := true
	res, err = d.CacheUsersKey(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	if res != nil {
		prom.CacheHit.Incr("UsersKey")
		return
	}
	prom.CacheMiss.Incr("UsersKey")
	res, err = d.RawUsersKey(c, id)
	if err != nil {
		return
	}
	miss := res
	if !addCache {
		return
	}
	d.AddCacheUsersKey(c, id, miss)
	return
}

// Points get data from cache if miss will call source method, then add to cache.
func (d *Dao) Points(c context.Context, id int64) (res *bwsmdl.Points, err error) {
	addCache := true
	res, err = d.CachePoints(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	if res != nil {
		prom.CacheHit.Incr("Points")
		return
	}
	prom.CacheMiss.Incr("Points")
	res, err = d.RawPoints(c, id)
	if err != nil {
		return
	}
	miss := res
	if !addCache {
		return
	}
	d.AddCachePoints(c, id, miss)
	return
}

// Achievements get data from cache if miss will call source method, then add to cache.
func (d *Dao) Achievements(c context.Context, id int64) (res *bwsmdl.Achievements, err error) {
	addCache := true
	res, err = d.CacheAchievements(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	if res != nil {
		prom.CacheHit.Incr("Achievements")
		return
	}
	prom.CacheMiss.Incr("Achievements")
	res, err = d.RawAchievements(c, id)
	if err != nil {
		return
	}
	miss := res
	if !addCache {
		return
	}
	d.AddCacheAchievements(c, id, miss)
	return
}

// UserAchieves get data from cache if miss will call source method, then add to cache.
func (d *Dao) UserAchieves(c context.Context, id int64, key string) (res []*bwsmdl.UserAchieve, err error) {
	addCache := true
	res, err = d.CacheUserAchieves(c, id, key)
	if err != nil {
		addCache = false
		err = nil
	}
	if len(res) != 0 {
		prom.CacheHit.Incr("UserAchieves")
		return
	}
	prom.CacheMiss.Incr("UserAchieves")
	res, err = d.RawUserAchieves(c, id, key)
	if err != nil {
		return
	}
	miss := res
	if !addCache {
		return
	}
	d.AddCacheUserAchieves(c, id, miss, key)
	return
}

// UserPoints get data from cache if miss will call source method, then add to cache.
func (d *Dao) UserPoints(c context.Context, id int64, key string) (res []*bwsmdl.UserPoint, err error) {
	addCache := true
	res, err = d.CacheUserPoints(c, id, key)
	if err != nil {
		addCache = false
		err = nil
	}
	if len(res) != 0 {
		prom.CacheHit.Incr("UserPoints")
		return
	}
	prom.CacheMiss.Incr("UserPoints")
	res, err = d.RawUserPoints(c, id, key)
	if err != nil {
		return
	}
	miss := res
	if !addCache {
		return
	}
	d.AddCacheUserPoints(c, id, miss, key)
	return
}

// AchieveCounts get data from cache if miss will call source method, then add to cache.
func (d *Dao) AchieveCounts(c context.Context, id int64, day string) (res []*bwsmdl.CountAchieves, err error) {
	addCache := true
	res, err = d.CacheAchieveCounts(c, id, day)
	if err != nil {
		addCache = false
		err = nil
	}
	if len(res) != 0 {
		prom.CacheHit.Incr("AchieveCounts")
		return
	}
	prom.CacheMiss.Incr("AchieveCounts")
	res, err = d.RawAchieveCounts(c, id, day)
	if err != nil {
		return
	}
	miss := res
	if !addCache {
		return
	}
	d.AddCacheAchieveCounts(c, id, miss, day)
	return
}