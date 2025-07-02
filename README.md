# Sistema de autenticación en Go

Este sistema implementa un flujo completo de autenticación y autorización utilizando tokens JWT para acceso seguro y escalable. Está desarrollado en Go e incluye access tokens de corta duración junto con refresh tokens persistentes para renovación segura. 

![Licencia](https://img.shields.io/badge/Licencia-MIT-blue) ![Estado](https://img.shields.io/badge/Estado-en%20desarrollo-yellow)

---

## 🧠 Tabla de Contenidos

- [Tecnologías](#tecnologías)
- [Funcionalidades](#funcionalidades)
- [Características](#características)
- [Instalación](#instalación)

---

## 📦 Funcionalidades

* Inicio de sesión por email y password
	- Si las credenciales son válidas el sistema genera un token temporal y un refresh token.
	- Se devuelve el token al cliente.
	- Refresh token se guarda en BD para poder renovar autorización del usuario. 

* Registro mediante datos básicos del usuario
	- Si las credenciales son válidas, el usuario es creado en el sistema y se genera un token temporal y un refresh token.
	- Se devuelve el token al cliente.
	- Refresh token se guarda en BD para poder renovar autorización del usuario. 

## 💻 Tecnologías utilizadas en el Backend

![Go](https://img.shields.io/badge/go-00ADD8.svg?style=for-the-badge&logo=go&logoColor=white) 

![postgresql](https://img.shields.io/badge/postgresql-4169E1.svg?style=for-the-badge&logo=postgresql&logoColor=white)

![docker](https://img.shields.io/badge/docker-2496ED.svg?style=for-the-badge&logo=docker&logoColor=white)

![gin](https://img.shields.io/badge/gin-008ECF.svg?style=for-the-badge&logo=gin&logoColor=white)

---

## ✨ Características

- Escalabilidad
- Mantenibilidad
- Arquitectura hexagonal
- JWT
- Autenticación
- Autorización
- Refresh tokens

---

## 🛠️ Instalación

### Requisitos

- [Just](https://github.com/casey/just) — para automatizar tareas comunes del proyecto
- [Docker](https://www.docker.com/) — para desplegar la app de manera consistente

### 📦 Clonación e instalación

```bash
git clone https://github.com/jorgerr9011/auth-golang.git
cd auth-golang
just install
just migrate