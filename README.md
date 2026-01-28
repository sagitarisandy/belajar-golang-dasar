# belajar-golang-dasar

Ringkasan singkat / Short summary

This repository is a collection of small Go example programs and notes used to learn and teach Go (Golang) fundamentals. The examples cover basic language features, functions, pointers, error handling, closures, interfaces, and some standard-library snippets.

Repositori ini berisi contoh-contoh kecil program Go untuk mempelajari dasar-dasar bahasa Go.

## Repository structure

- `golang-dasar/` — core learning examples (many single-file examples). Each `.go` file demonstrates a small concept (functions, pointers, structs, interfaces, errors, etc.).
- `belajar-golang-standard-library/` — examples that focus on the Go standard library (has its own `go.mod`).
- `database/` — examples related to database access (may require third-party drivers or configuration).

(There may be additional files or folders; inspect them to find individual example programs.)

## Requirements

- Go 1.18+ installed (newer versions are recommended).
- No global build required — examples are runnable individually.

## How to run examples

Most examples are simple standalone files. There are two common ways to run them:

1) Run a single file (from the folder that contains the file):

```bash
# change into the folder then run one file
cd "golang-dasar"
# run a specific example file
go run helloworld.go
```

2) Run a package (when folder is a valid module or package):

```bash
# run all files in the package (requires a valid package in the folder)
cd "belajar-golang-standard-library"
go run .

# or build everything from repo root
# this will attempt to build all packages under the repository
go build ./...
```

Notes:
- Some examples may have their own `go.mod` (for example `belajar-golang-standard-library`). Change into that folder before running if needed.
- Code that interacts with external services (e.g., `database/mysql.go`) might require additional setup (database server and driver packages).

## Tips for contributors

- Keep examples small and focused.
- Use clear comments and descriptive names — these files are learning artifacts.
- If you add examples that require third-party modules, add a `go.mod` in the example folder and update README with any setup steps.

## License

This repository doesn't include an explicit license. If you want to allow reuse, add a LICENSE file (for example MIT or Apache-2.0).

---

If you want, I can also:
- Add README content in Indonesian only, or bilingual translations for each section.
- Create per-folder READMEs with example lists and expected outputs.
- Run a quick `go build ./...` to report build status for the repository (would run in your environment).

Tell me which of the above you'd like next.