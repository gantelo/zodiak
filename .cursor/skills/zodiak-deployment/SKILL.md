---
name: zodiak-deployment
description: Documents deployment and configuration for the zodiak service, including Fly.io, environment variables, CI workflows, and README notes. Use when changing deployment, env setup, secrets, or project documentation.
---

# Zodiak Deployment and Configuration

## Quick start

- **Runtime**: Go 1.21 service using embedded assets/templates/data and gocron.
- **Hosting**: Fly.io, configured via `fly.toml`.
- **Configuration**: Environment variables, with `.env` supported locally via `godotenv`.
- **CI**: GitHub Actions workflows for deployment and scheduled prompt generation.

Use this skill when:

- Updating Fly configuration or deployment strategy.
- Adding or changing environment variables.
- Adjusting CI workflows related to deployment or data generation.
- Editing `README.md` or `.env.template` documentation for setup.

## Fly.io configuration

Key settings in `fly.toml`:

- `app = "zodiak-fly"` — Fly application name.
- `primary_region = "scl"` — Primary region for deployment.
- `[build]`:
  - `builder = "paketobuildpacks/builder:base"`
  - `buildpacks = ["gcr.io/paketo-buildpacks/go"]` — Go buildpack for the service.
- `[env]`:
  - `PORT = "8080"` — Port used by the HTTP server.
- `[[statics]]`:
  - `guest_path = "/app/dist"`
  - `url_prefix = "/assets"` — static assets path (currently not heavily used by this service).
- `[http_service]`:
  - `internal_port = 8080`
  - `force_https = true`
  - `auto_stop_machines = true`
  - `auto_start_machines = true`
  - `min_machines_running = 1`

The Go app listens on `0.0.0.0:8080` by default (configurable via `PORT`). Fly maps external requests to this port.

## Environment variables

Configuration is driven by env vars referenced via `config.GetEnvVar` or `os.Getenv`.

From `.env.template`:

- DeepL:
  - `DEEPL_API_KEY`
- Scraping:
  - `SCRAP_WEB` — base URL for daily horoscopes.
  - `SCRAP_CLASS` — CSS selector used by the scraper to find content.
- X/Twitter API:
  - `X_API_KEY`
  - `X_API_KEY_SECRET`
  - `X_ACCESS_TOKEN`
  - `X_ACCESS_TOKEN_SECRET`
  - `X_BEARER_TOKEN`
- Optional:
  - `API_PASS` — documented but currently unused in `main.go`.

Additional runtime env vars:

- `MOON_SCRAP_WEB` — source URL for moon phase scraping (read by `DailyMoonPhase`).
- `FLY_REGION` — used in the root HTTP handler to show region in the template.
- `PORT` — optional override for the HTTP port (falls back to `8080` if empty).

When adding new env vars:

1. Add them to `.env.template` with clear placeholder values.
2. Document their purpose and expected format in `README.md`.
3. Use `config.GetEnvVar` where the value is required to avoid silent misconfiguration.

## Local development and secrets

- On local runs, `init()` in `main.go`:
  - Checks for `X_API_KEY` in the environment.
  - If missing, attempts to load `.env` using `github.com/joho/godotenv`.
  - Exits with `log.Fatalf` if neither env nor `.env` is present.

Workflow for local development:

1. Copy `.env.template` to `.env`.
2. Fill in all required API keys (DeepL, scraping site, X/Twitter credentials).
3. Run the Go app locally; it will load configuration from `.env`.

On Fly, secrets should be set via `fly secrets` rather than committing them to the repo.

## CI and automation

- GitHub Actions workflows live under `.github/workflows/`:
  - `fly.yml` — handles deployment to Fly.io:
    - Typically builds and deploys the Go app when changes are pushed to relevant branches (inspect the workflow for exact triggers).
  - `prompts-schedule.yml` — scheduled job that updates `data/prompts.json` and related JSON files used for the \"best at\" feature.

When modifying deployment:

- Ensure the Fly GitHub Action uses the correct app name and region.
- Keep secrets in GitHub repository settings or environments, not in the repo.

When modifying data generation:

- Ensure that any changes to the structure or naming of `data/*.json` match how `internal/daily` and `stringUtils` read and parse these files.

## HTTP behavior

- `main.go`:
  - Registers a single HTTP route `/` that renders `templates/index.html.tmpl`.
  - Passes `Region` from `FLY_REGION` to the template for simple runtime diagnostics.
- There is no public API or authentication enforced at the HTTP level in the current code.
  - If you introduce new endpoints or authentication (for example using `API_PASS`), document the behavior and configuration in `README.md` and this skill.

## Documentation

- `README.md`:
  - Explains what the bot does, its schedule, and the basic setup steps.
  - Documents needed environment variables (DeepL, scraping site, X/Twitter).
  - Mentions how `SCRAP_WEB` and `SCRAP_CLASS` are used by the scraper.
  - Notes that DeepL free tier uses `https://api-free.deepl.com/v2`.
- Keep the README aligned with changes to:
  - Schedule or gocron configuration.
  - Required environment variables.
  - Deployment or architectural changes that impact how contributors run the app.

## References

- Deployment config: `fly.toml`
- Env template: `.env.template`
- README: `README.md`
- Scheduler and entrypoint: `main.go`
- Daily tasks and data usage: `internal/daily/tasks.go`
