package app

import (
	"WhoKnowsMeapp/internal/big_cache"
	"WhoKnowsMeapp/internal/db"
	"WhoKnowsMeapp/internal/rdb"
	"github.com/allegro/bigcache"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type App struct {
	Gorm  *gorm.DB
	Redis *redis.Client
	BigCache *bigcache.BigCache
}

func NewApp() (*App, error) {

	grm, err := db.NewGorm()

	if err != nil {
		return nil, err
	}

	r, err := rdb.NewRedis()

	if err != nil {
		return nil, err
	}

	return &App{
		Gorm:  grm,
		Redis: r,
		BigCache: big_cache.MakeBigCache(),
	}, nil
}
