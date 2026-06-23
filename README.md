# unix-adm

Managing secrets manually is painful. That's the reason this service exists — and as a Unix systems administration showcase project, of course.

A lightweight secrets registry for self-managed hosting. Register your projects, configure their `.env` paths, then push or remove secrets via a simple HTTP API. The service writes directly to the `.env` files on the host.

## How it works

- Projects are stored in MongoDB with their `.env` file paths per environment (`staging` / `prod`)
- `POST /projects/:name/secrets` — batch upsert keys into the `.env` file
- `PATCH /projects/:name/secrets` — batch remove keys from the `.env` file
- `GET /projects/:name/keys` — list registered key names (no values exposed)
- All routes are protected by JWT auth

## Setup

Copy `.env.sample` to `.env` and fill in the values, then:

```sh
docker compose up -d
```

Generate a JWT secret:
```sh
openssl rand -hex 32
```
