package handler

import (
	"context"
	"cron-job/model"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Handler struct {
	client   *http.Client
	cfg      *model.JobConfig
	initOnce sync.Once // Đảm bảo hàm init chỉ được gọi 1 lần dù có nhiều goroutine gọi
}

func NewHandler(cfg *model.JobConfig) *Handler {
	return &Handler{
		cfg: cfg,
	}
}

func (h *Handler) init() {

	if h.client == nil {
		h.client = &http.Client{
			// No follow redirect
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}

	if h.cfg.HandlerConfig.TimeoutSeconds == 0 {
		h.cfg.HandlerConfig.TimeoutSeconds = 10
	}
}

func (h *Handler) Run() {
	fmt.Println("vao day")

	h.initOnce.Do(h.init)
	cfg := h.cfg.HandlerConfig

	timestamp := time.Now()
	timeout := time.Duration(cfg.TimeoutSeconds) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, cfg.Method, cfg.URL, strings.NewReader(cfg.Body))
	if err != nil {
		return
	}

	for k, v := range cfg.Headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("x-time", fmt.Sprintf("%d", timestamp.Unix()))

	_, err = h.client.Do(req)
	if err != nil {
		return
	}
}
