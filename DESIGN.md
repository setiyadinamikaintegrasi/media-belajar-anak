# DESIGN.md

**Status:** Discovery baseline; fitur produk final menunggu keputusan target usia dan model konten.

## Batasan eksperimen

Project ini dipakai untuk menguji template-ai-native pada monorepo nyata:

- `src/backend/` — API dan domain service.
- `src/frontend/` — web app mobile-first.
- Komponen dideklarasikan eksplisit di `.template/project.yaml`.
- Tidak boleh menggunakan symlink untuk stack detection.

## Desain awal

```text
Pendamping/Anak → Frontend → Backend API → Content Store
                                      └→ Progress Store
```

## Kandidat domain

- `learning_content`: materi, topik, level, media.
- `activities`: latihan, pertanyaan, opsi jawaban, skor.
- `learner_progress`: progres per anak dan materi.
- `guardian_controls`: persetujuan, profil pendamping, batas akses.

## Acceptance criteria bootstrap

1. Backend `go test -short ./...` berhasil.
2. Backend `go build -o bin/ ./...` menghasilkan artifact.
3. Frontend `npm ci` dan `npm run build` berhasil.
4. Frontend memiliki unit test yang dapat dijalankan Vitest.
5. `project.yaml` valid dan resolver menghasilkan dua komponen.
6. CI tidak menjalankan single-stack detection pada monorepo v2.
7. Tidak ada secret, data anak, atau kredensial di repository.

## Risiko yang harus dikaji sebelum fitur anak dibuat

- Privasi dan klasifikasi data anak.
- Consent pendamping.
- Moderasi dan keamanan konten.
- Aksesibilitas.
- Telemetry dan retensi data.
- Validasi materi oleh ahli/pendidik.
