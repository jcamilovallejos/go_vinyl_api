# API Simple de Gestión de Álbumes en Go

API básica para gestionar una colección de álbumes musicales, construida con Gin (framework web de Go).  
Incluye operaciones CRUD básicas y es ideal para aprender los fundamentos de APIs REST.

---

## 📌 Características

- **Métodos HTTP soportados**: GET, POST, PUT, DELETE
- **Endpoints**:
  - Obtener todos los álbumes
  - Obtener un álbum por ID
  - Crear un nuevo álbum
  - Actualizar un álbum existente (reemplazo completo)
  - Eliminar un álbum

---

## ⚙️ Requisitos

- Go 1.20+ instalado
- Dependencia: [Gin Web Framework](https://github.com/gin-gonic/gin)

---

## 🛠️ Instalación

1. Clona el repositorio:
```

git clone https://github.com/tu-usuario/tu-repositorio.git

```
2. Instala Gin:
```

go get -u github.com/gin-gonic/gin

```
3. Ejecuta la API:
```

go run main.go

```

---

## 📡 Endpoints y Ejemplos

### 1. Obtener todos los álbumes
```

GET /albums

```
**Ejemplo**:
```

curl http://localhost:8080/albums

```

### 2. Obtener un álbum por ID
```

GET /albums/:id

```
**Ejemplo**:
```

curl http://localhost:8080/albums/1

```

### 3. Crear un nuevo álbum
```

POST /albums

```
**Ejemplo**:
```

curl -X POST -H "Content-Type: application/json" -d '{
"id": "4",
"title": "Thriller",
"artist": "Michael Jackson",
"price": 19.99
}' http://localhost:8080/albums

```

### 4. Actualizar un álbum (PUT)
```

PUT /albums/:id

```
**Ejemplo**:
```

curl -X PUT -H "Content-Type: application/json" -d '{
"id": "1",
"title": "Blue Train (Edición Especial)",
"artist": "John Doe",
"price": 25.99
}' http://localhost:8080/albums/1

```

### 5. Eliminar un álbum
```

DELETE /albums/:id

```
**Ejemplo**:
```

curl -X DELETE http://localhost:8080/albums/2

```

---

## 🧪 Pruebas

1. Usa **curl** o herramientas como [Postman](https://www.postman.com/).
2. Verifica las respuestas HTTP:
   - `200 OK`: Operación exitosa
   - `201 Created`: Recurso creado
   - `400 Bad Request`: Datos inválidos
   - `404 Not Found`: Recurso no encontrado

---

## 📄 Licencia
MIT License. Proyecto creado con fines educativos.  
**Nota**: Esta es una implementación básica sin persistencia real (los datos se guardan en memoria).

