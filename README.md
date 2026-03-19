# 🍔 API de Pedidos

API REST para gerenciamento de pedidos de comida, com autenticação JWT e front-end integrado.

---

## 🎥 Demonstração

<p align="center">
  <a href="https://www.youtube.com/watch?v=Nhbe_erhdKU">
    <img src="https://img.youtube.com/vi/Nhbe_erhdKU/maxresdefault.jpg" width="700px"/>
  </a>
</p>

<p align="center">
  ▶️ Clique para assistir a demonstração completa
</p>

---

## 🚀 Funcionalidades

- 📦 Criar pedidos  
- 📋 Listar pedidos  
- 🔄 Atualizar status do pedido  
- 🔐 Autenticação com JWT  
- 💻 Front-end SPA integrado  

---

## 🛠️ Tecnologias

- Go (Golang)
- MySQL / SQLite
- HTML, CSS, JavaScript
- Docker

---

## 🧠 Arquitetura

- API REST  
- Arquitetura em Camadas  
- Cliente-Servidor  
- Autenticação via JWT  

---

## ⚙️ Como executar

### 🔧 Backend

```bash
go run main.go
```

## Endpoints
- `POST /api/auth/login` (body: `email`, `password`) - token JWT
- `GET /api/orders` (header `Authorization: Bearer <token>`)
- `POST /api/orders` (body: `customer_name`, `items`, `total`) + token
- `PATCH /api/orders/{orderID}/status` (body: `status`) + token

### Usuário inicial
- email: `admin@api.com`
- password: `password`

## Executar com Docker (MySQL)
```powershell
docker run --name pedidos-mysql -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=api_pedidos -p 3306:3306 -d mysql:8
```

## Observações
O serviço inicializa tabelas e insere usuário padrão. Use as rotas para validar fluxo de pedidos.
