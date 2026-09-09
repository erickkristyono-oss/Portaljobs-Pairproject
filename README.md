# Portal Job — Backend Foundation (Phase 0)

REST API foundation for the Portal Job project.
Stack: **Go + Gin + PostgreSQL + GORM + JWT + Swagger/OpenAPI + Docker**.

This phase gives you: register/login for 3 roles (jobseeker/company/admin),
JWT auth + role guard, a health check, Swagger UI, and a **jobs** example domain
that Dev A & Dev B copy as the pattern for every other feature.

## Run it (one command)

```bash
cp .env.example .env      # optional, only needed for local (non-Docker) runs
docker compose up --build
```

Then:
- API base: `http://localhost:8080`
- Health:   `http://localhost:8080/health`
- Swagger:  `http://localhost:8080/swagger`

The database schema is created automatically on startup (GORM AutoMigrate).

## Try the flow

```bash
# 1. register a company
curl -X POST localhost:8080/api/v1/register \
  -H 'Content-Type: application/json' \
  -d '{"nama":"PT Maju","email":"hr@maju.com","password":"secret123","role":"company"}'

# 2. login -> copy the token from the response
curl -X POST localhost:8080/api/v1/login \
  -H 'Content-Type: application/json' \
  -d '{"email":"hr@maju.com","password":"secret123"}'

# 3. create a job (company only) — replace <TOKEN>
curl -X POST localhost:8080/api/v1/jobs \
  -H 'Authorization: Bearer <TOKEN>' \
  -H 'Content-Type: application/json' \
  -d '{"judul":"Backend Engineer","lokasi":"Jakarta","gaji":12000000}'

# 4. list & detail (public)
curl localhost:8080/api/v1/jobs
curl localhost:8080/api/v1/jobs/1
```

A jobseeker token calling `POST /jobs` gets **403** — that is the role guard working.

## Project structure (clean architecture)

Dependencies point inward: handler -> usecase -> repository(interface) -> db.

```
cmd/api/main.go                 wiring / entry point
internal/
  config/                       env config
  database/                     GORM connection + AutoMigrate
  domain/                       entities + repository INTERFACES + errors
  repository/postgres/          GORM implementations of the interfaces
  usecase/                      business logic (+ unit test with mock)
  handler/                      Gin handlers + request DTOs
  middleware/                   JWT auth + RequireRole guard
  router/                       route wiring
pkg/
  jwt/                          token generate/parse
  response/                     standard JSON envelope
docs/                           embedded OpenAPI spec + Swagger UI
```

## The pattern to copy (for every new feature)

To add a feature (e.g. applications, companies, contracts, reports), follow the
same 4 files the `jobs` domain uses:

1. `domain/<name>.go` — entity + `<Name>Repository` interface
2. `repository/postgres/<name>_repository.go` — GORM implementation
3. `usecase/<name>_usecase.go` — business logic
4. `handler/<name>_handler.go` — Gin handlers (+ DTOs in `handler/dto.go`)

Then: register the model in `database/postgres.go` (AutoMigrate), wire the
handler in `router/router.go`, and add the paths to `docs/swagger.json`.

### Suggested split for a 2-person team
- **Dev A (jobseeker side):** profile/skill/portfolio, apply, my-applications — copies the jobs pattern.
- **Dev B (company side + WA):** companies, contracts, and the WhatsApp notifier (behind an interface so it can be mocked in tests).

## Testing (mocking)

```bash
go test ./... -v
```

`internal/usecase/auth_usecase_test.go` shows the mocking approach: a fake
`UserRepository` lets you test business logic with no database. Every usecase
should get tests this way (rubric requirement).

## Notes
- `Job.CompanyID` currently stores the company **user's** id. When you add the
  `companies` table (phase 1), change it to reference `companies.id`.
- `go.sum` is generated on the first `docker compose up` / `go mod tidy`.
- Change `JWT_SECRET` before deploying anywhere real.

---

## v2 — Modul M1–M6 (security + core loop + admin + CI)

Dibangun di atas fondasi. Jalankan seperti biasa:

```bash
cp .env.example .env      # sesuaikan bila perlu
docker compose up --build # Postgres + API
# atau lokal:
go mod tidy && go run ./cmd/api
```

Test:

```bash
go test ./... -race -cover
```

Swagger: `http://localhost:8080/swagger` (20 endpoint).

### Endpoint baru
- Jobseeker: `GET/PUT /me/profile`, `GET/POST /me/skills`, `DELETE /me/skills/:id`,
  `GET/POST /me/portfolios`, `DELETE /me/portfolios/:id`,
  `POST /jobs/:id/apply`, `GET /me/applications`
- Company: `GET/PUT /company/profile`, `POST /jobs`, `GET /company/jobs`,
  `GET /jobs/:id/applications`, `PATCH /applications/:id`
- Umum (login): `POST /reports`
- Admin: `GET /admin/users`, `PATCH /admin/users/:id/suspend`,
  `GET /admin/reports`, `PATCH /admin/reports/:id`

Keamanan: lihat `SECURITY.md`.

### Catatan kendala yang diketahui (baca sebelum menilai kode)
1. **Dependency belum ter-verify compile di sini** (registry Go diblokir di
   lingkungan pembuatan). Semua file sudah lolos `gofmt`. Di mesinmu jalankan
   `go mod tidy` lalu `go test ./...`; kalau ada versi modul yang gagal, jalankan
   `go get <paket>@latest`.
2. **Rate limiter in-memory** (`internal/middleware/security.go`): map per-IP
   tumbuh terus (potensi memory leak). Fix: goroutine cleanup berkala atau Redis.
3. **Notifikasi WA dipanggil best-effort & sinkron** di usecase (error hanya
   di-log). Untuk produksi bungkus dalam goroutine/queue agar tidak menahan
   response. Interface `domain.Notifier` sudah membuatnya mudah diganti & di-mock.
4. **JWT tanpa refresh/blacklist** — logout belum bisa mencabut token aktif.
5. **CORS default `*`** — persempit ke domain frontend saat produksi.

admin :

{
  "nama": "Admin",
  "email": "admin@portaljob.com",
  "password": "admin123",
  "role": "admin"
}

token :

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJyb2xlIjoiY29tcGFueSIsImV4cCI6MTc4ODU5OTc0MSwiaWF0IjoxNzg4NTEzMzQxfQ.84mXyEXkROwZQSSKbXYxRLRoyHqNE_C50oaIt5C1_E0

create profile company :

{
  "name": "PT Maju Jaya",
  "field_of": "Technology",
  "address": "Jakarta",
  "description": "Perusahaan software"
}

token:


Register jobseeker:

imung
imung@mail.com
imung123

token :
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJyb2xlIjoiY29tcGFueSIsImV4cCI6MTc4ODU5OTc0MSwiaWF0IjoxNzg4NTEzMzQxfQ.84mXyEXkROwZQSSKbXYxRLRoyHqNE_C50oaIt5C1_E0

token login:
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJyb2xlIjoiam9ic2Vla2VyIiwiZXhwIjoxNzg4OTYyNTU3LCJpYXQiOjE3ODg4NzYxNTd9.TwzT9BDTVJevBCXDlyMeAhA8WwIki3uGlC3Sr0HS6j0

POST JOB company

{ "judul": "Fullstack", "lokasi": "Yogyakarta", "required_skills": ["JavaScript Modern", "Framework Front-End", "go","postgresql"] }

{
  "judul": "Fullstack",
  "about_role": "Membangun dan memelihara REST API untuk portal kerja",
  "responsibilities": "Desain endpoint, menulis unit test, code review, deploy ke GCP",
  "deskripsi": "fullstack",
  "lokasi": "Yogyakarta",
  "gaji": 25000000,
  "required_skills": ["JavaScript Modern", "Framework Front-End", "go","postgresql"]
}


PUT company profile
{
  "name": "PT Maju Jaya",
  "field_of": "Technology",
  "address": "Jakarta Selatan",
  "description": "Perusahaan software"
}