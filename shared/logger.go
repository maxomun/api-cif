// File: main.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Plantilla de servicios API

package shared

import (
	"context"
	"errors"
	"strings"
	"sync"
	"time"

	logFactory "api-cif/shared/pattern/log"

	servicebus "github.com/Azure/azure-service-bus-go"
)

// Estructura para el hook
type BatchHook struct {
	cola      *servicebus.Queue
	nombreApi string
	mensaje   chan string
	stopChan  chan struct{}
	wg        sync.WaitGroup
}

// Crea una nueva instancia del hook
func NewBatchHook(conexion, nombrecola, nombreApi string, objLogger logFactory.Logger) (*BatchHook, error) {

	objServicio, objErrorServicioBus := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(conexion))

	if objErrorServicioBus != nil {
		return nil, errors.New(objErrorServicioBus.Error())
	}

	objCola, objErrorCola := objServicio.NewQueue(nombrecola)

	if objErrorCola != nil {
		return nil, errors.New(objErrorCola.Error())
	}

	objHook := &BatchHook{
		cola:      objCola,
		nombreApi: nombreApi,
		mensaje:   make(chan string, 1000), // Canal con capacidad para manejar mensajes
		stopChan:  make(chan struct{}),
	}

	// Iniciar el procesador de mensajes en un goroutine
	objHook.wg.Add(1)
	go objHook.procesar(objLogger)

	return objHook, nil
}

// Agrega el log al canal de mensajes
func (objHook *BatchHook) Run(objLogger logFactory.Logger, mensaje string) {

	logMensaje := mensaje

	// Agregar el mensaje al canal de forma no bloqueante
	select {
	case objHook.mensaje <- logMensaje:
	default:
		// Si el canal está lleno, descartamos el mensaje
		logFactory.LogMensaje(objLogger, logMensaje, logFactory.Info)
	}
}

// Procesa y envía los mensajes acumulados en segundo plano
func (objHook *BatchHook) procesar(objLogger logFactory.Logger) {

	defer objHook.wg.Done()

	var batch []string

	for {
		select {
		case mensaje := <-objHook.mensaje:
			batch = append(batch, mensaje)
		case <-objHook.stopChan:
			// Enviar cualquier mensaje acumulado al cerrar
			if len(batch) > 0 {
				objHook.enviar(objLogger, batch)
			}
			return
		}
	}
}

// sendBatch envía un lote de mensajes a Azure Service Bus
func (objHook *BatchHook) enviar(objLogger logFactory.Logger, batch []string) {

	objContexto, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	mensajeLote := strings.Join(batch, "\n")
	mensaje := servicebus.NewMessageFromString(mensajeLote)

	if objErrorEnviarLote := objHook.cola.Send(objContexto, mensaje); objErrorEnviarLote != nil {

		logFactory.LogMensaje(objLogger, objErrorEnviarLote.Error(), logFactory.Info)
	} else {
		//fmt.Printf("Enviados %d mensajes a Service Bus\n", len(batch))
	}
}

// Cerrar detiene el procesamiento y envía los mensajes restantes de forma asincrónica
func (objHook *BatchHook) Cerrar(objLogger logFactory.Logger) {

	close(objHook.stopChan)

	// Iniciar el envío final de mensajes en un goroutine separado
	go func() {
		objHook.wg.Wait() // Esperar a que el procesador de mensajes termine
		logFactory.LogMensaje(objLogger, "Todos los mensajes se enviaron correctamente.", logFactory.Info)
	}()
}
