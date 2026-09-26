package adapters

import (
	"encoding/json"
	"net/http"

	"github.com/leandrojacome/concurrent-rate-limiter-go/application"
	"github.com/leandrojacome/concurrent-rate-limiter-go/domain"
)

type RateLimitHandler struct{ limiter application.Limiter }

func NewRateLimitHandler(limiter application.Limiter) *RateLimitHandler {
	return &RateLimitHandler{limiter: limiter}
}
func (h *RateLimitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key, err := domain.NewClientKey(r.Header.Get("X-Client-ID"))
	if err != nil {
		http.Error(w, "invalid X-Client-ID", http.StatusBadRequest)
		return
	}
	decision := h.limiter.Allow(r.Context(), key)
	w.Header().Set("Content-Type", "application/json")
	if !decision.Allowed {
		w.WriteHeader(http.StatusTooManyRequests)
	}
	_ = json.NewEncoder(w).Encode(decision)
}
