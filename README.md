# go-web-template

Minimal, zero-framework Go (`net/http`) web application template with embedded templates/assets, gzip compression, graceful `SIGTERM` shutdown, and a distroless `Dockerfile` for Cloud Run.

## Using This Template

1. Create a new repository from this template on GitHub (or copy it cleanly).
2. Update the module path in `go.mod`:
   ```sh
   go mod edit -module github.com/<you>/<your-app>
   ```
3. Run locally:
   ```sh
   go run .
   ```
4. Run tests:
   ```sh
   go test -v ./...
   ```
