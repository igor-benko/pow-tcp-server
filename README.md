# TCP Server with Proof of Work

A secure TCP server that delivers random quotes, protected against DDoS attacks using a challenge-response **Proof of Work (PoW)** mechanism. Designed using clean architecture principles and instrumented for observability.

---

## 📁 Project Structure

```
pow-tcp-server/
├── cmd/
│   ├── server/                  # Server entrypoint
│   │   └── main.go
│   └── client/                  # Client entrypoint
│       └── main.go
├── docker/
│   ├── Dockerfile.server        # Dockerfile for the server
│   └── Dockerfile.client        # Dockerfile for the client
├── internal/
│   ├── config/                  # Env and .env config loading
│   ├── delivery/                # delivery layer (tcp, http)
│   |   └── tcp                  # tcp implementation
|   |   └── http                 # http implementation
│   ├── repository/              # repository
│   |   └── client               # external clients
|   |   └── storage              # database clients
│   └── service/                 # business logic (usecase)
│       └── adapters             # service adapters (external calls - db, clients, caches, etc)
│       └── quote                # quote service implementation
|       └── interface.go         # service interface for delivery layer (tcp, http, grpc, etc)
├── pkg/                         # Shared utilities (if any)
│   └── pow/                     # PoW challenge & validation logic
├── go.mod
├── go.sum
├── .env                         # Environment variable config
├── Makefile                     # Build/run shortcuts
└── README.md
```

---

## ⚙️ Configuration

The app reads its config from environment variables and a `.env` file:

| Variable                      | Description                   | Default  |
|-------------------------------|-------------------------------|----------|
| `SERVER_PORT`                 | TCP server port        | ``   |
| `POW_CHALLENGE_DIFFICULTY`    | PoW difficulty (e.g. `3`)     | ``      |

You can override them in `.env` or directly via shell.

---

## 🛠 Build & Run

### 🐳 Docker (Recommended)

#### Server

```bash
docker build -f docker/Dockerfile.server -t pow-tcp-server .
docker run --env-file .env -p 11000:11000 pow-tcp-server
```

#### Client

```bash
docker build -f docker/Dockerfile.client -t pow-tcp-client .
docker run --env-file .env --network host pow-tcp-client
```
