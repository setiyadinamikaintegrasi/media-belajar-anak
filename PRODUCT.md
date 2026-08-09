# Media Belajar Anak

**Status:** Discovery baseline — target usia dan kurikulum akan dikonfirmasi sebelum implementasi fitur utama.

## Vision

Media belajar digital yang aman, menyenangkan, dan mudah digunakan anak bersama pendampingnya.

## Problem

Materi belajar anak sering tersebar, tidak terstruktur, dan kurang memberikan umpan balik yang sesuai kemampuan anak. Project ini menguji kemampuan template-ai-native untuk mengelola produk edukasi dengan batasan keamanan, konten, dan pengalaman pengguna yang jelas.

## Target pengguna awal

- Anak sebagai pengguna utama.
- Orang tua/guru sebagai pendamping dan pengelola progres.

## Kandidat MVP

- Katalog materi belajar berdasarkan topik dan tingkat kesulitan.
- Materi interaktif berbasis teks, gambar, dan audio.
- Latihan/kuis dengan umpan balik langsung.
- Progress belajar per anak.
- Kontrol pendamping untuk memilih materi.
- Desain mobile-first dan aksesibel.

## Out of scope awal

- Iklan pihak ketiga.
- Chat terbuka antara anak dan pengguna lain.
- Pengumpulan data pribadi anak tanpa persetujuan pendamping.
- Personalisasi berbasis AI sebelum safety dan evaluasinya disetujui.

## Success criteria eksperimen template

1. Project berhasil dibuat dari GitHub Use this template.
2. Monorepo v2 berjalan tanpa symlink.
3. Backend dan frontend memiliki component checks sendiri.
4. `make setup`, `make test-scripts`, `make docs-check`, dan `make ci` berhasil.
5. PR bootstrap menjalankan reusable component workflow dari branch dasar.
