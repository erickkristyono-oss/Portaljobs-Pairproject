# Skema Keamanan — Portal Job

Ringkasan lapisan keamanan yang dipakai backend ini, apa fungsinya, dan di file
mana kodenya.

## 1. Authentication (siapa kamu) — JWT + bcrypt

- **Password di-hash dengan bcrypt** (`golang.org/x/crypto/bcrypt`), tidak pernah
  disimpan/plaintext. Lihat `internal/usecase/auth_usecase.go`.
  Field `Password` di entity `User` diberi tag `json:"-"` → hash **tidak pernah**
  ikut ter-serialize ke response API.
- **Login mengeluarkan JWT (HS256)** ditandatangani dengan `JWT_SECRET`.
  Token berisi `user_id`, `role`, dan `exp`. Lihat `pkg/jwt/jwt.go`.
- **Pesan error login disamakan** untuk "email salah" maupun "password salah"
  (`ErrInvalidCredential`) supaya penyerang tidak bisa menebak email mana yang
  terdaftar (mencegah user enumeration).
- **Akun suspended diblokir saat login** (`ErrSuspended` → 403).

## 2. Authorization (kamu boleh apa) — role guard

- Setiap request ke endpoint terproteksi melewati middleware `Auth` yang
  memverifikasi tanda tangan + masa berlaku JWT, lalu menaruh `user_id`/`role`
  ke context. Lihat `internal/middleware/auth.go`.
- `RequireRole("jobseeker"|"company"|"admin")` membatasi endpoint per peran.
  Contoh: hanya `company` yang bisa `POST /jobs`, hanya `jobseeker` yang bisa
  melamar, hanya `admin` yang bisa suspend user.
- **Ownership check di usecase** (bukan cuma role): jobseeker hanya bisa hapus
  skill/portfolio miliknya; company hanya bisa melihat/menilai pelamar untuk
  lowongan miliknya sendiri. Lihat `application_usecase.go` (`ownsJob`) dan
  `skill_usecase.go`/`portfolio_usecase.go` (cek `UserID`).

## 3. Header & transport hardening — `internal/middleware/security.go`

- **CORS** dikonfigurasi lewat `CORS_ALLOW_ORIGIN` (default `*` untuk dev;
  **wajib dipersempit** ke domain frontend di produksi).
- **Secure headers**: `X-Content-Type-Options: nosniff`, `X-Frame-Options: DENY`,
  `Referrer-Policy: no-referrer`.
- **Body limit** 1 MB (`http.MaxBytesReader`) — guard payload berlebihan.
- **Rate limit** per-IP (default 120/menit, `RATE_LIMIT_PER_MIN`) — mengurangi
  brute-force & abuse.

## 4. Input & database

- **Validasi input** di DTO via tag `binding` (`required`, `email`, `oneof`,
  `min=6`). Lihat `internal/handler/dto.go`. Request tidak valid → `400`.
- **Tanpa SQL injection**: semua akses DB lewat GORM dengan parameter binding
  (`Where("email = ?", email)`), tidak pernah menyusun query dari string mentah.
- **Status code konsisten**: 201/200/400/401/403/404/409/429 dipetakan terpusat
  di `internal/handler/httperr.go`.

## 5. Secrets

- Semua rahasia (`JWT_SECRET`, kredensial DB, token WA) lewat **environment
  variable** (`.env` di-`.gitignore`, tidak pernah di-commit). Lihat `.env.example`.
- **Ganti `JWT_SECRET`** dengan string acak panjang di produksi, dan rotasi bila
  bocor.

## Yang belum & rencana lanjutan (jujur)

- **Belum ada refresh token / blacklist**: logout tidak bisa mencabut token yang
  masih berlaku. Mitigasi sekarang: masa berlaku token pendek (`JWT_EXPIRE_HOURS`).
  Lanjutan: simpan daftar token dicabut di Redis, atau pakai refresh token.
- **Rate limiter in-memory** menyimpan IP di map yang tumbuh terus (lihat catatan
  kendala di README). Untuk produksi: cleanup berkala atau Redis.
- **HTTPS** di-terminate di layer deploy (Cloud Run sudah otomatis TLS).
