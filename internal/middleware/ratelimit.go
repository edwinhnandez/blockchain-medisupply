package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// IPRateLimiter gestiona limitadores de tasa por IP
type IPRateLimiter struct {
	ips             map[string]*rate.Limiter
	mu              *sync.RWMutex
	r               rate.Limit
	b               int
	cleanupInterval time.Duration
}

// NewIPRateLimiter crea un nuevo limitador de tasa por IP
func NewIPRateLimiter(requestsPerWindow int, windowSeconds int) *IPRateLimiter {
	limiter := &IPRateLimiter{
		ips:             make(map[string]*rate.Limiter),
		mu:              &sync.RWMutex{},
		r:               rate.Limit(float64(requestsPerWindow) / float64(windowSeconds)),
		b:               requestsPerWindow,
		cleanupInterval: time.Minute * 5,
	}

	// Iniciar limpieza periódica
	go limiter.cleanupRoutine()

	return limiter
}

// AddIP crea un nuevo limitador para una IP si no existe
func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)
	i.ips[ip] = limiter

	return limiter
}

// GetLimiter retorna el limitador para una IP
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.RLock()
	limiter, exists := i.ips[ip]
	i.mu.RUnlock()

	if !exists {
		return i.AddIP(ip)
	}

	return limiter
}

// cleanupRoutine limpia limitadores no utilizados
func (i *IPRateLimiter) cleanupRoutine() {
	ticker := time.NewTicker(i.cleanupInterval)
	defer ticker.Stop()

	for range ticker.C {
		i.mu.Lock()
		// En una implementación de producción, mantendríamos track del último uso
		// Por ahora, mantenemos todos los limitadores
		i.mu.Unlock()
	}
}

// RateLimitMiddleware crea un middleware de rate limiting
func RateLimitMiddleware(requestsPerWindow, windowSeconds int) gin.HandlerFunc {
	limiter := NewIPRateLimiter(requestsPerWindow, windowSeconds)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := limiter.GetLimiter(ip)

		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Demasiadas peticiones. Por favor, intente más tarde.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
