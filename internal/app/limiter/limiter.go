package limiter

import (
	"context"
	"strings"
	"sync"

	"github.com/emrzvv/tages-test/cfg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type clientCounters struct {
	uploads   int
	downloads int
	lists     int
}

type CounterLimiter struct {
	mutex                 sync.Mutex
	clientCountersMapping map[string]*clientCounters
	config                *cfg.Config
}

func NewCounterLimiter(config *cfg.Config) *CounterLimiter {
	return &CounterLimiter{
		mutex:                 sync.Mutex{},
		clientCountersMapping: make(map[string]*clientCounters),
		config:                config,
	}
}

func (c *CounterLimiter) getClientKeyFromCtx(ctx context.Context) string {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return "unknown"
	}
	return p.Addr.String()
}

func (c *CounterLimiter) getOrCreateCounter(clientKey string) *clientCounters {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	cnt, ok := c.clientCountersMapping[clientKey]
	if !ok {
		cnt = &clientCounters{}
		c.clientCountersMapping[clientKey] = cnt
	}
	return cnt
}

func (c *CounterLimiter) UnaryInterceptor(
	ctx context.Context,
	request any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {
	clientKey := c.getClientKeyFromCtx(ctx)
	counters := c.getOrCreateCounter(clientKey)

	switch {
	case strings.Contains(info.FullMethod, "/GetImagesList"):
		if !c.acquireList(counters) {
			return nil, status.Error(codes.ResourceExhausted, "too many concurrent list requests")
		}
		defer c.releaseList(counters)

	}

	return handler(ctx, request)
}

func (c *CounterLimiter) StreamInterceptor(
	srv any,
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	key := c.getClientKeyFromCtx(ss.Context())
	counters := c.getOrCreateCounter(key)

	switch {
	case strings.HasSuffix(info.FullMethod, "/UploadImage"):
		if !c.acquireUpload(counters) {
			return status.Error(codes.ResourceExhausted, "too many concurrent upload requests")
		}
		defer c.releaseUpload(counters)

	case strings.HasSuffix(info.FullMethod, "/DownloadImage"):
		if !c.acquireDownload(counters) {
			return status.Error(codes.ResourceExhausted, "too many concurrent download requests")
		}
		defer c.releaseDownload(counters)
	}

	return handler(srv, ss)
}

func (c *CounterLimiter) acquireUpload(cnt *clientCounters) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if cnt.uploads >= c.config.Limits["upload"] {
		return false
	}
	cnt.uploads++
	return true
}
func (c *CounterLimiter) releaseUpload(cnt *clientCounters) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	cnt.uploads--
}

func (c *CounterLimiter) acquireDownload(cnt *clientCounters) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if cnt.downloads >= c.config.Limits["download"] {
		return false
	}
	cnt.downloads++
	return true
}
func (c *CounterLimiter) releaseDownload(cnt *clientCounters) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	cnt.downloads--
}

func (c *CounterLimiter) acquireList(cnt *clientCounters) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if cnt.lists >= c.config.Limits["list"] {
		return false
	}
	cnt.lists++
	return true
}
func (c *CounterLimiter) releaseList(cnt *clientCounters) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	cnt.lists--
}
