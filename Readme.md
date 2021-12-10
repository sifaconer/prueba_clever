# Prueba Técnica.

### Base de datos

Motor: sqlite

### API
 
**Documentación:** [swagger](http://127.0.0.1:8182/swagger/index.html)

#### Docker

Build: sudo docker build -t prueba:v1 .
Run: sudo docker run -e API_CURRENTLAYER_KEY='{API KEY}' -p 8182:8182 prueba:v1

### Docker Compose

Run: sudo docker-compose up -d