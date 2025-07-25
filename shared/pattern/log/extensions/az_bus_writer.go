// File: az_bus_write.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Escrito que incorpora funcionalidades para enviar log al servicio azure service bus

// Cualquier cambio a la estructura en este archivo, debe ser previa autorizacion del area de arquitectura

package extensions

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

type AzureServiceBusWriter struct {
	objClient *azservicebus.Client
	objSender *azservicebus.Sender
	objCtx    context.Context
}

func NewAzureServiceBusWriter(cnx, nombreCola string) (*AzureServiceBusWriter, error) {

	objClientInterno, objErrorCnx := azservicebus.NewClientFromConnectionString(cnx, nil)

	if objErrorCnx != nil {
		return nil, fmt.Errorf("error creando el cliente de Service Bus: %w", objErrorCnx)
	}

	objSenderInterno, objErrorSender := objClientInterno.NewSender(nombreCola, nil)
	if objErrorSender != nil {
		return nil, fmt.Errorf("error creando el sender de Service Bus: %w", objErrorSender)
	}

	return &AzureServiceBusWriter{
		objClient: objClientInterno,
		objSender: objSenderInterno,
		objCtx:    context.Background(),
	}, nil
}

// Implementa io.Writer y envía los logs a Service Bus
func (objEscritor *AzureServiceBusWriter) Write(mensaje []byte) (numero int, err error) {

	message := string(mensaje)

	objErrorSender := objEscritor.objSender.SendMessage(objEscritor.objCtx, &azservicebus.Message{

		Body: []byte(message),
	}, nil)

	if objErrorSender != nil {
		log.Printf("Error enviando mensaje a Service Bus: %v", objErrorSender)
		return 0, objErrorSender
	}

	return len(mensaje), nil
}

// Close cierra la conexión con Azure Service Bus
func (objEscritor *AzureServiceBusWriter) Close() error {
	return objEscritor.objSender.Close(objEscritor.objCtx)
}
