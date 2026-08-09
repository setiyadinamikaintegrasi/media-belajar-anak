# ARCHITECTURE.md

**Status:** Bootstrap architecture.

## Component topology

```text
src/frontend (Node)
       │ HTTP/JSON
       ▼
src/backend (Go)
       │
       ├── Content service
       ├── Activity service
       └── Progress service
```

## Component contract

Monorepo component contract berada di `.template/project.yaml`. Workflow CI
menjalankan quality, test, dan build dari working directory setiap komponen.

## Initial deployment boundary

- Frontend dapat dideploy sebagai static artifact.
- Backend dideploy sebagai service terpisah.
- Database dan object storage belum dipilih.
- Tidak ada data anak nyata dalam fase bootstrap.

## Keputusan yang ditunda

- Framework backend dan frontend final.
- Database production.
- Authentication/parental consent.
- Content authoring workflow.
- Analytics dan retention policy.
