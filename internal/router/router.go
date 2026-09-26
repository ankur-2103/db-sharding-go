package router

import (
	"db-sharding/internal/shard"
	"errors"
)

type RoutingStrategy interface {
	GetShard(userID int, numberOfShards int) int
}

type ModuloStrategy struct{}

func (m ModuloStrategy) GetShard(userID int, numberOfShards int) int {
	return userID % numberOfShards
}

type ShardRouter struct {
	shards   map[int]shard.Shard
	strategy RoutingStrategy
}

func NewShardRouter(shards map[int]shard.Shard, strategy RoutingStrategy) *ShardRouter {
	newShard := ShardRouter{
		shards:   shards,
		strategy: strategy,
	}
	return &newShard
}

func (s *ShardRouter) Route(userID int) (shard.Shard, error) {
	if userID < 0 {
		return shard.Shard{}, ErrInvalidUserID
	}

	if len(s.shards) == 0 {
		return shard.Shard{}, ErrNoShards
	}

	shardID := s.strategy.GetShard(userID, len(s.shards))
	selectedShard, exists := s.shards[shardID]

	if !exists {
		return shard.Shard{}, ErrInvalidShard
	}

	return selectedShard, nil
}

var (
	ErrInvalidUserID = errors.New("Invalid user")
	ErrNoShards      = errors.New("No shards available")
	ErrInvalidShard  = errors.New("Invalid shard")
)
