# Proxy HTTP con Caché

Proyecto final del curso **Sistemas Operativos II** — Universidad Latina de Costa Rica.

Implementación de un servidor proxy HTTP con caché desarrollado en Go, que permite reenviar solicitudes a servidores de origen y almacenar las respuestas localmente para optimizar el rendimiento.

---

## Integrantes

| Nombre | Carrera |
|--------|---------|
| José Araya Fernández | Ingeniería del Software |
| Amy Morales Cruz | Ingeniería del Software |

## Profesor

**Carlos Méndez Rodríguez**
Universidad Latina de Costa Rica — Primer Período Lectivo 2026

---

## Descripción del Proyecto

Este proyecto implementa un **proxy HTTP con caché** que:

- Recibe solicitudes HTTP de clientes mediante sockets TCP
- Reenvía las solicitudes al servidor de origen
- Almacena las respuestas en caché en memoria con TTL (tiempo de vida)
- Atiende múltiples clientes de forma simultánea usando goroutines
- Evalúa el rendimiento comparando tiempos de respuesta con y sin caché

---

## Estructura del Proyecto

```
proxy-http-cache/
├── main.go       # Punto de entrada: servidor TCP y gestión de conexiones
├── cache.go      # Almacenamiento en memoria con TTL y limpieza automática
├── handler.go    # Lógica del proxy: análisis HTTP y política de caché
├── go.mod        # Módulo de Go
├── Dockerfile    # Imagen Docker para despliegue reproducible
├── logs/         # Directorio generado en tiempo de ejecución
└── README.md
```

---

## Requisitos

- [Go 1.22+](https://go.dev/dl/) — para ejecución local
- [Docker](https://www.docker.com/) — para despliegue en contenedor

---

## Ejecución Local

### 1. Clonar el repositorio

```bash
git clone https://github.com/AmyMoralesC/Proyecto---SO-II.git
cd Proyecto---SO-II
```

### 2. Compilar y ejecutar

```bash
go build -o proxy .
./proxy
```

El proxy quedará corriendo en el puerto **8080**.

```
Proxy HTTP con Cache corriendo en puerto :8080
Presiona Ctrl+C para detener.
```

---

## Ejecución con Docker

### 1. Construir la imagen

```bash
docker build -t proxy-http-cache .
```

### 2. Ejecutar el contenedor

```bash
docker run -p 8080:8080 proxy-http-cache
```

El proxy estará disponible en `localhost:8080` desde su máquina.

---

## Pruebas

### Opción A — Servidor externo (puede dar 502 por restricciones de red)

```bash
curl -v --proxy http://localhost:8080 http://example.com
```

### Opción B — Servidor local con Python (recomendado para demo)

En una terminal separada, levante un servidor local:

```bash
python3 -m http.server 9090
```

Luego, envie la misma solicitud dos veces para observar el comportamiento del caché:

```bash
# Primera vez: genera CACHE MISS (va al servidor de origen)
curl --proxy http://localhost:8080 http://localhost:9090

# Segunda vez: genera CACHE HIT (responde directo desde caché)
curl --proxy http://localhost:8080 http://localhost:9090
```

### Prueba de tiempos (CACHE MISS vs CACHE HIT)

```bash
time curl --proxy http://localhost:8080 http://localhost:9090 > /dev/null && \
time curl --proxy http://localhost:8080 http://localhost:9090 > /dev/null
```

Debería observar que la segunda solicitud es notablemente más rápida.

---

## Conceptos Aplicados

| Concepto | Implementación |
|----------|----------------|
| Sockets TCP | `net.Listen` / `net.Dial` del paquete estándar de Go |
| Concurrencia | Goroutines (`go handleClient`) por cada conexión |
| Sincronización | `sync.RWMutex` para proteger la caché compartida |
| Caché con TTL | `map[string]CacheEntry` con `time.Time` de expiración |
| Protocolo HTTP | `http.ReadRequest` para análisis de solicitudes |

---

## Licencia

Este proyecto está licenciado bajo la [MIT License](LICENSE).
