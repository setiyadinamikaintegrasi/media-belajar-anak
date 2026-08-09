# ADR-0008: Bootstrap Media Belajar Anak sebagai Monorepo

**Status:** Accepted for bootstrap experiment

## Context

Project ini menguji template-ai-native dengan dua runtime: backend API dan
frontend web. Struktur harus merepresentasikan monorepo nyata dan tidak boleh
mengandalkan symlink.

## Decision

Gunakan component contract v2:

- `backend=src/backend:go`
- `frontend=src/frontend:node`

Setiap komponen memiliki build dan test command sendiri. Fitur produk dan stack
final akan diputuskan setelah discovery target usia, konten, dan consent.

## Consequences

- Component-aware CI dapat menguji backend dan frontend secara terisolasi.
- Artifact backend dan frontend terpisah.
- Template tidak perlu menebak nested manifest.
- Ada sedikit overhead konfigurasi dibanding single-stack project.
