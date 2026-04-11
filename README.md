# Proxy HTTP con Caché
 
Proyecto final del curso **Sistemas Operativos II** — Universidad Latina de Costa Rica.
 
Implementación de un servidor proxy HTTP con caché desarrollado en Go, que permite reenviar solicitudes a servidores de origen y almacenar las respuestas localmente para optimizar el rendimiento.
 
---
 
## Integrantes
 
| Nombre | Carrera |
|--------|---------|
| José Araya Fernández | Ingeniería en Software |
| Amy Morales Cruz | Ingeniería en Software |
 
## Profesor
 
**Carlos Méndez Rodríguez**  
Universidad Latina de Costa Rica — Primer Período Lectivo 2026
 
---
 
## Descripción del Proyecto
 
Este proyecto implementa un **proxy HTTP con caché** que:
 
- Recibe solicitudes HTTP de clientes mediante sockets TCP
- Reenvía las solicitudes al servidor de origen
- Almacena las respuestas en caché (en memoria o disco)
- Atiende múltiples clientes de forma simultánea usando hilos (threads)
- Evalúa el rendimiento comparando tiempos de respuesta con y sin caché
 
---
