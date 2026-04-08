package telemetry

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type TelemetryData struct {
	Route      string    `json:"route"`
	APIVersion string    `json:"apiVersion"`
	Timestamp  time.Time `json:"timestamp"`
}

type telemetryService struct {
	enabled bool
}

func (t *telemetryService) TelemetryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !t.enabled {
			c.Next()
			return
		}
		route := c.FullPath()
		go SendTelemetry(route)
		c.Next()
	}
}

type TelemetryService interface {
	TelemetryMiddleware() gin.HandlerFunc
}

func SendTelemetry(route string) {
	if route == "/" {
		return
	}

	telemetry := TelemetryData{
		Route:      route,
		APIVersion: "evo-go",
		Timestamp:  time.Now(),
	}

	url := "https://log.evolution-api.com/telemetry"

	data, err := json.Marshal(telemetry)
	if err != nil {
		log.Println("Erro ao serializar telemetria:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		// Silence DNS lookup errors or connection refused as they are common in restricted environments
		if !strings.Contains(err.Error(), "no such host") && !strings.Contains(err.Error(), "connection refused") {
			log.Println("Erro ao enviar telemetria:", err)
		}
		return
	}
	defer resp.Body.Close()
}

func NewTelemetryService(enabled bool) TelemetryService {
	return &telemetryService{
		enabled: enabled,
	}
}
