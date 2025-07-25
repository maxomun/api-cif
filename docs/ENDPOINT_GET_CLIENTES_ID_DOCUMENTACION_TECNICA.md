# Documentación Técnica - Endpoint GET /clientes/{id}

## 1. Descripción General del Endpoint

### Propósito del Endpoint
El endpoint `GET /clientes/{id}` permite consultar la información detallada de un cliente específico basado en su identificador único.

### Significado del Parámetro {id}
El parámetro `{id}` representa el identificador único del cliente en la base de datos. Es un valor numérico entero positivo que debe ser proporcionado en la URL.

### Recurso Devuelto
El endpoint devuelve la información completa del cliente incluyendo:
- Datos básicos del cliente (ID, nombre, documento)
- Información de clasificación (tipo de documento, estado, tipo de entidad)
- Datos de actividad económica
- Información de composición institucional
- Códigos de contraparte
- Indicadores de validez
- Relaciones cliente-institución
- Tipo de sociedad
- Fecha de ingreso

## 2. Especificaciones de la Solicitud

### Método
`GET`

### Ruta
`/clientes/{id}`

### Tipo de Parámetro
- **Tipo**: Path Parameter
- **Nombre**: `id`
- **Tipo de dato**: Integer (int64)
- **Obligatorio**: Sí

### Validaciones sobre id
- Debe ser un número entero positivo
- Debe ser mayor que 0
- No puede ser nulo o vacío
- No puede contener caracteres no numéricos

### Headers Requeridos
- `Content-Type: application/json` (se establece automáticamente)

## 3. Estructura de Respuesta Esperada

### Formato
`application/json`

### Respuesta Exitosa (HTTP 200)
```json
{
  "id": 123,
  "idDocumento": "12345678-9",
  "tipoDocumento": {
    "id": 1,
    "descripcion": "RUT"
  },
  "estado": {
    "id": 1,
    "descripcion": "Activo"
  },
  "tipoEntidad": {
    "id": 1,
    "descripcion": "Persona Natural"
  },
  "actividadEconomica": {
    "id": 1,
    "descripcion": "Comercio al por menor",
    "codigoCIIU": 4711
  },
  "composicionInstitucional": {
    "id": 1,
    "descripcion": "Empresa Privada"
  },
  "codigoContraparte": {
    "id": 1,
    "descripcion": "Cliente Regular"
  },
  "indicadorValidez": {
    "id": "V",
    "descripcion": "Válido"
  },
  "relacionClienteInstitucion": {
    "id": 1,
    "descripcion": "Cliente Directo"
  },
  "tipoSociedad": {
    "id": 1,
    "descripcion": "Sociedad Anónima",
    "idKYC": 1,
    "descripcionKYC": "Sociedad Anónima"
  },
  "nombre": "EMPRESA EJEMPLO SPA",
  "fechaIngreso": "2024-01-15T10:30:00Z"
}
```

### Descripción de Campos

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | int64 | Identificador único del cliente |
| `idDocumento` | string | Número de documento de identidad |
| `tipoDocumento` | object | Información del tipo de documento (id, descripcion) |
| `estado` | object | Estado actual del cliente (id, descripcion) |
| `tipoEntidad` | object | Tipo de entidad del cliente (id, descripcion) |
| `actividadEconomica` | object | Actividad económica con código CIIU |
| `composicionInstitucional` | object | Composición institucional del cliente |
| `codigoContraparte` | object | Código de clasificación como contraparte |
| `indicadorValidez` | object | Indicador de validez del cliente |
| `relacionClienteInstitucion` | object | Relación con la institución |
| `tipoSociedad` | object | Tipo de sociedad con información KYC |
| `nombre` | string | Nombre o razón social del cliente |
| `fechaIngreso` | string | Fecha de ingreso del cliente (formato ISO 8601) |

### Estructura Anidada
La respuesta es un objeto anidado que contiene múltiples sub-objetos para las diferentes clasificaciones del cliente. Todos los campos de tipo objeto pueden ser `null` si no tienen información asociada.

## 4. Códigos de Estado Esperados

### 200 OK
**Descripción**: Cliente encontrado exitosamente
**Cuerpo**: Objeto JSON con la información completa del cliente
**Headers**: `Content-Type: application/json`

### 404 Not Found
**Descripción**: Cliente no existe en la base de datos
**Cuerpo**: 
```json
{
  "objeto": "",
  "estado": {
    "codigo": 404,
    "mensaje": "Metodo ejecutado sin resultados"
  }
}
```

### 400 Bad Request
**Descripción**: ID inválido (no numérico, nulo, o menor o igual a 0)
**Cuerpo**: Sin cuerpo de respuesta
**Headers**: `Content-Type: application/json`

### 500 Internal Server Error
**Descripción**: Error interno del servidor
**Cuerpo**: Sin cuerpo de respuesta
**Headers**: `Content-Type: application/json`

## 5. Casos de Prueba para QA

### ✅ Casos de Prueba Exitosos

#### 1. Cliente existente con ID válido
- **ID**: 123
- **URL**: `GET http://localhost:8080/clientes/123`
- **Resultado esperado**: HTTP 200 con datos completos del cliente
- **Validaciones**: Verificar que todos los campos estén presentes y con tipos correctos

#### 2. ID válido con respuesta completa
- **ID**: 456
- **URL**: `GET http://localhost:8080/clientes/456`
- **Resultado esperado**: HTTP 200 con objeto JSON válido
- **Validaciones**: 
  - Verificar estructura JSON válida
  - Validar tipos de datos de cada campo
  - Verificar que campos opcionales puedan ser null

### ❌ Casos de Prueba de Error

#### 3. Cliente no existente con ID válido
- **ID**: 999999
- **URL**: `GET http://localhost:8080/clientes/999999`
- **Resultado esperado**: HTTP 404
- **Validaciones**: Verificar mensaje de error apropiado

#### 4. ID no numérico
- **ID**: "abc"
- **URL**: `GET http://localhost:8080/clientes/abc`
- **Resultado esperado**: HTTP 400
- **Validaciones**: Verificar que no se procese el request

#### 5. ID vacío
- **ID**: ""
- **URL**: `GET http://localhost:8080/clientes/`
- **Resultado esperado**: HTTP 400 o 404 (dependiendo del router)
- **Validaciones**: Verificar manejo de parámetro vacío

#### 6. ID cero o negativo
- **ID**: 0
- **URL**: `GET http://localhost:8080/clientes/0`
- **Resultado esperado**: HTTP 400
- **Validaciones**: Verificar validación de ID positivo

- **ID**: -1
- **URL**: `GET http://localhost:8080/clientes/-1`
- **Resultado esperado**: HTTP 400
- **Validaciones**: Verificar validación de ID positivo

### 🔄 Casos de Prueba de Estrés

#### 7. Múltiples requests concurrentes
- **Descripción**: Ejecutar 10 requests simultáneos al mismo cliente
- **URL**: `GET http://localhost:8080/clientes/123`
- **Resultado esperado**: Todos los requests deben responder correctamente
- **Validaciones**: 
  - Verificar tiempos de respuesta consistentes
  - Validar que no haya bloqueos
  - Verificar integridad de datos

#### 8. Requests con diferentes IDs concurrentes
- **Descripción**: Ejecutar requests a diferentes clientes simultáneamente
- **URLs**: 
  - `GET http://localhost:8080/clientes/123`
  - `GET http://localhost:8080/clientes/456`
  - `GET http://localhost:8080/clientes/789`
- **Resultado esperado**: Todos los requests deben responder correctamente
- **Validaciones**: Verificar que cada request devuelva los datos correctos

## 6. Colección de Postman

### Configuración Base
```json
{
  "info": {
    "name": "API CIF - Endpoint GET /clientes/{id}",
    "description": "Colección de pruebas para el endpoint de consulta de clientes",
    "version": "1.0.0"
  },
  "variable": [
    {
      "key": "baseUrl",
      "value": "http://localhost:8080",
      "type": "string"
    }
  ]
}
```

### Casos de Prueba

#### 1. Cliente Válido Existente
```json
{
  "name": "Cliente válido existente",
  "request": {
    "method": "GET",
    "header": [
      {
        "key": "Content-Type",
        "value": "application/json"
      }
    ],
    "url": {
      "raw": "{{baseUrl}}/clientes/123",
      "host": ["{{baseUrl}}"],
      "path": ["clientes", "123"]
    }
  },
  "response": [
    {
      "name": "Respuesta exitosa",
      "originalRequest": {
        "method": "GET",
        "header": [],
        "url": {
          "raw": "{{baseUrl}}/clientes/123",
          "host": ["{{baseUrl}}"],
          "path": ["clientes", "123"]
        }
      },
      "status": "OK",
      "code": 200,
      "_postman_previewlanguage": "json",
      "header": [
        {
          "key": "Content-Type",
          "value": "application/json"
        }
      ],
      "cookie": [],
      "body": "{\n  \"id\": 123,\n  \"idDocumento\": \"12345678-9\",\n  \"tipoDocumento\": {\n    \"id\": 1,\n    \"descripcion\": \"RUT\"\n  },\n  \"estado\": {\n    \"id\": 1,\n    \"descripcion\": \"Activo\"\n  },\n  \"tipoEntidad\": {\n    \"id\": 1,\n    \"descripcion\": \"Persona Natural\"\n  },\n  \"actividadEconomica\": {\n    \"id\": 1,\n    \"descripcion\": \"Comercio al por menor\",\n    \"codigoCIIU\": 4711\n  },\n  \"composicionInstitucional\": {\n    \"id\": 1,\n    \"descripcion\": \"Empresa Privada\"\n  },\n  \"codigoContraparte\": {\n    \"id\": 1,\n    \"descripcion\": \"Cliente Regular\"\n  },\n  \"indicadorValidez\": {\n    \"id\": \"V\",\n    \"descripcion\": \"Válido\"\n  },\n  \"relacionClienteInstitucion\": {\n    \"id\": 1,\n    \"descripcion\": \"Cliente Directo\"\n  },\n  \"tipoSociedad\": {\n    \"id\": 1,\n    \"descripcion\": \"Sociedad Anónima\",\n    \"idKYC\": 1,\n    \"descripcionKYC\": \"Sociedad Anónima\"\n  },\n  \"nombre\": \"EMPRESA EJEMPLO SPA\",\n  \"fechaIngreso\": \"2024-01-15T10:30:00Z\"\n}"
    }
  ]
}
```

#### 2. Cliente No Existente
```json
{
  "name": "Cliente no existente",
  "request": {
    "method": "GET",
    "header": [
      {
        "key": "Content-Type",
        "value": "application/json"
      }
    ],
    "url": {
      "raw": "{{baseUrl}}/clientes/999999",
      "host": ["{{baseUrl}}"],
      "path": ["clientes", "999999"]
    }
  },
  "response": [
    {
      "name": "Cliente no encontrado",
      "originalRequest": {
        "method": "GET",
        "header": [],
        "url": {
          "raw": "{{baseUrl}}/clientes/999999",
          "host": ["{{baseUrl}}"],
          "path": ["clientes", "999999"]
        }
      },
      "status": "Not Found",
      "code": 404,
      "_postman_previewlanguage": "json",
      "header": [
        {
          "key": "Content-Type",
          "value": "application/json"
        }
      ],
      "cookie": [],
      "body": "{\n  \"objeto\": \"\",\n  \"estado\": {\n    \"codigo\": 404,\n    \"mensaje\": \"Metodo ejecutado sin resultados\"\n  }\n}"
    }
  ]
}
```

#### 3. ID Inválido (No Numérico)
```json
{
  "name": "ID inválido - no numérico",
  "request": {
    "method": "GET",
    "header": [
      {
        "key": "Content-Type",
        "value": "application/json"
      }
    ],
    "url": {
      "raw": "{{baseUrl}}/clientes/abc",
      "host": ["{{baseUrl}}"],
      "path": ["clientes", "abc"]
    }
  },
  "response": [
    {
      "name": "Bad Request",
      "originalRequest": {
        "method": "GET",
        "header": [],
        "url": {
          "raw": "{{baseUrl}}/clientes/abc",
          "host": ["{{baseUrl}}"],
          "path": ["clientes", "abc"]
        }
      },
      "status": "Bad Request",
      "code": 400,
      "_postman_previewlanguage": "json",
      "header": [
        {
          "key": "Content-Type",
          "value": "application/json"
        }
      ],
      "cookie": [],
      "body": ""
    }
  ]
}
```

#### 4. ID Cero
```json
{
  "name": "ID inválido - cero",
  "request": {
    "method": "GET",
    "header": [
      {
        "key": "Content-Type",
        "value": "application/json"
      }
    ],
    "url": {
      "raw": "{{baseUrl}}/clientes/0",
      "host": ["{{baseUrl}}"],
      "path": ["clientes", "0"]
    }
  },
  "response": [
    {
      "name": "Bad Request - ID cero",
      "originalRequest": {
        "method": "GET",
        "header": [],
        "url": {
          "raw": "{{baseUrl}}/clientes/0",
          "host": ["{{baseUrl}}"],
          "path": ["clientes", "0"]
        }
      },
      "status": "Bad Request",
      "code": 400,
      "_postman_previewlanguage": "json",
      "header": [
        {
          "key": "Content-Type",
          "value": "application/json"
        }
      ],
      "cookie": [],
      "body": ""
    }
  ]
}
```

### Scripts de Validación

#### Validación de Respuesta Exitosa
```javascript
pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});

pm.test("Response has correct content type", function () {
    pm.expect(pm.response.headers.get("Content-Type")).to.include("application/json");
});

pm.test("Response has required fields", function () {
    const jsonData = pm.response.json();
    pm.expect(jsonData).to.have.property('id');
    pm.expect(jsonData).to.have.property('idDocumento');
    pm.expect(jsonData).to.have.property('nombre');
    pm.expect(jsonData).to.have.property('fechaIngreso');
});

pm.test("ID is a positive integer", function () {
    const jsonData = pm.response.json();
    pm.expect(jsonData.id).to.be.a('number');
    pm.expect(jsonData.id).to.be.above(0);
});

pm.test("Date format is ISO 8601", function () {
    const jsonData = pm.response.json();
    const dateRegex = /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$/;
    pm.expect(jsonData.fechaIngreso).to.match(dateRegex);
});
```

#### Validación de Respuesta de Error 404
```javascript
pm.test("Status code is 404", function () {
    pm.response.to.have.status(404);
});

pm.test("Error response structure", function () {
    const jsonData = pm.response.json();
    pm.expect(jsonData).to.have.property('objeto');
    pm.expect(jsonData).to.have.property('estado');
    pm.expect(jsonData.estado).to.have.property('codigo');
    pm.expect(jsonData.estado).to.have.property('mensaje');
    pm.expect(jsonData.estado.codigo).to.eql(404);
});
```

#### Validación de Respuesta de Error 400
```javascript
pm.test("Status code is 400", function () {
    pm.response.to.have.status(400);
});

pm.test("Response has correct content type", function () {
    pm.expect(pm.response.headers.get("Content-Type")).to.include("application/json");
});
```

## 7. Información Adicional

### Configuración del Servidor
- **Puerto**: 8080 (configurado en `configuracion.json`)
- **Base URL**: `http://localhost:8080`
- **Framework**: Gin (Go)
- **Base de Datos**: SQL Server

### Logs y Monitoreo
El endpoint registra logs de información y error utilizando ZeroLog:
- **Log de información**: Cuando se ejecuta correctamente
- **Log de error**: Cuando ocurren errores internos o de validación

### Consideraciones de Rendimiento
- El endpoint utiliza conexiones de base de datos con pool configurado
- Máximo 20 conexiones abiertas
- Máximo 10 conexiones inactivas
- Tiempo de vida máximo de 30 minutos por conexión

### Seguridad
- No requiere autenticación en el endpoint actual
- Validación de entrada para prevenir inyección de parámetros
- Headers de Content-Type establecidos automáticamente

---

**Documento generado para el equipo de QA**
**Versión**: 1.0
**Fecha**: 2025-01-09
**Autor**: Sistema de Documentación Automática 