package routes

import (
	"fmt"
	"log"
	"net/http"
	middle "test/internal/delivery/http/routes/middleware"
	"test/internal/domain"
	"test/internal/usecase"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "status"},
	)

	cpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_usage",
		Help: "CPU usage",
	})

	memoryUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memory_usage",
		Help: "Memory usage",
	})
)

func initMet() {
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(memoryUsage)
	prometheus.MustRegister(httpRequestsTotal)
}

// @title Anti Brute Force Service
// @description The service is designed to detect and prevent brute-force attacks on web resources.
// @version 1.0
// @host localhost:5080
// @BasePath /
func InitRoutes(useCase usecase.UseCase, config *domain.ConfigSetting, secretKey string) {
	e := echo.New()

	// Запуск сбора метрик
	go collectMetrics()
	initMet()

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	middleWare := middle.NewmiddleWare(secretKey)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Ping server
	// @Summary ping server
	// @Router /ping [get]
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})

	e.Use(metricsMiddleware)

	antiBruteForce := e.Group("/api/antiBruteForce")
	admin := e.Group("/api/admin")
	admin.Use(middleWare.CheckBearerToken)

	AntiBrouteForceRoutes(antiBruteForce, useCase)
	AdminRoutes(admin, useCase)

	e.Logger.Fatal(e.Start(":" + viper.GetString("PORT")))
}

func metricsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		httpRequestsTotal.With(prometheus.Labels{"method": c.Request().Method, "status": fmt.Sprintf("%d", http.StatusOK)}).Inc()
		err := next(c)
		return err
	}
}

func collectMetrics() {
	for {
		cpuPercent, err := cpu.Percent(time.Second, false)
		if err != nil {
			log.Println("Failed to get CPU usage:", err)
			continue
		}

		memInfo, err := mem.VirtualMemory()
		if err != nil {
			log.Println("Failed to get memory usage:", err)
			continue
		}

		cpuUsage.Set(cpuPercent[0])
		memoryUsage.Set(memInfo.UsedPercent)

		time.Sleep(time.Second * 5)
	}
}