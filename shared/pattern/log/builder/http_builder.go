// File: http_builder.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Logica de manejo y de logs utilizando el canal de envio HTTP

// Cualquier cambio a la estructura en este archivo, debe ser previa autorizacion del area de arquitectura

package builder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type LogMessage struct {
	Level     string `json:"level"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	Service   string `json:"service"`
	Host      string `json:"host"`
}

type HTTPLogWriter struct {
	endpoint   string
	client     *http.Client
	headers    map[string]string
	logChannel chan LogMessage
	wg         sync.WaitGroup
	quit       chan struct{}
}

type HTTPLogWriterBuilder struct {
	endpoint string
	timeout  time.Duration
	headers  map[string]string
	buffer   int
	service  string
	host     string
}

func NewHTTPLogWriterBuilder() *HTTPLogWriterBuilder {
	return &HTTPLogWriterBuilder{
		timeout: 5 * time.Second, // Timeout por defecto
		headers: make(map[string]string),
		buffer:  100, // Buffer de 100 mensajes
		service: "default-service",
		host:    "localhost",
	}
}

func (objeto *HTTPLogWriterBuilder) SetEndpoint(endpoint string) *HTTPLogWriterBuilder {
	objeto.endpoint = endpoint
	return objeto
}

func (objeto *HTTPLogWriterBuilder) SetTimeout(timeout time.Duration) *HTTPLogWriterBuilder {
	objeto.timeout = timeout
	return objeto
}

func (objeto *HTTPLogWriterBuilder) AddHeader(key, value string) *HTTPLogWriterBuilder {
	objeto.headers[key] = value
	return objeto
}

func (objeto *HTTPLogWriterBuilder) SetBufferSize(size int) *HTTPLogWriterBuilder {
	objeto.buffer = size
	return objeto
}

func (objeto *HTTPLogWriterBuilder) SetServiceName(service string) *HTTPLogWriterBuilder {
	objeto.service = service
	return objeto
}

func (objeto *HTTPLogWriterBuilder) SetHostName(host string) *HTTPLogWriterBuilder {
	objeto.host = host
	return objeto
}

func (objeto *HTTPLogWriterBuilder) Build() *HTTPLogWriter {
	writer := &HTTPLogWriter{
		endpoint:   objeto.endpoint,
		client:     &http.Client{Timeout: objeto.timeout},
		headers:    objeto.headers,
		logChannel: make(chan LogMessage, objeto.buffer),
		quit:       make(chan struct{}),
	}

	// Iniciar la goroutine de procesamiento de logs
	go writer.processLogs()

	return writer
}

func (objLogW *HTTPLogWriter) processLogs() {
	for {
		select {
		case logEntry := <-objLogW.logChannel:
			objLogW.sendLog(logEntry) // Envía el log a la API REST
		case <-objLogW.quit:
			close(objLogW.logChannel)
			return
		}
	}
}

// envía un log en formato JSON a la API REST
func (objLogW *HTTPLogWriter) sendLog(logEntry LogMessage) {
	jsonData, err := json.Marshal(logEntry)
	if err != nil {
		fmt.Println("Error serializando JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", objLogW.endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creando el request:", err)
		return
	}

	// Agregar headers personalizados
	for key, value := range objLogW.headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := objLogW.client.Do(req)
	if err != nil {
		fmt.Println("Error enviando log:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		fmt.Printf("Error en respuesta de API: %d\n", resp.StatusCode)
	}
}

// Implementación de la interfaz Logger
func (objLogW *HTTPLogWriter) Info(msg string) {
	objLogW.logChannel <- LogMessage{
		Level:     "INFO",
		Timestamp: time.Now().Format(time.RFC3339),
		Message:   msg,
		Service:   "default-service",
		Host:      "localhost",
	}
}

func (objLogW *HTTPLogWriter) Debug(msg string) {
	objLogW.logChannel <- LogMessage{
		Level:     "DEBUG",
		Timestamp: time.Now().Format(time.RFC3339),
		Message:   msg,
		Service:   "default-service",
		Host:      "localhost",
	}
}

func (objLogW *HTTPLogWriter) Error(msg string) {
	objLogW.logChannel <- LogMessage{
		Level:     "ERROR",
		Timestamp: time.Now().Format(time.RFC3339),
		Message:   msg,
		Service:   "default-service",
		Host:      "localhost",
	}
}

// Close cierra la goroutine de logs de forma segura
func (objLogW *HTTPLogWriter) Close() {
	close(objLogW.quit)
}
