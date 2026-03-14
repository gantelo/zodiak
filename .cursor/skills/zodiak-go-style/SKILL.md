---
name: zodiak-go-style
description: Go code style and project layout for the zodiak service, covering internal packages, config and env usage, naming, and file placement. Use when editing Go code, refactoring, or adding new internal packages.
---

# Zodiak Go Style and Layout

## Quick start

- Prefer small, focused packages under `internal/` (for example `config`, `daily`, `scrap`, `deepl`, `x`, `images`, `compatibilities`, `ctypes`, `stringUtils`, `http`).
- Keep package names short and lowercase; avoid stutter (for example `config.GetEnvVar`, not `config.ConfigGetEnvVar`).
- Put all app entrypoint logic in `main.go` and all business logic in `internal/` packages.
- Read configuration through the `config` package and environment variables, not hardcoded literals in other packages.

## Project layout

- **Entry point**
  - `main.go`:
    - Embeds `templates/*`, `assets/*`, and `data/*`.
    - Sets `images.Assets` and `daily.Data` so internal packages can read embedded files.
    - Registers HTTP handlers (currently only `/`) and starts the HTTP server.
    - Configures and starts the `gocron` scheduler with daily jobs (horoscope, compatibility, moon phase, and \"best at\" tasks).

- **Internal packages**
  - `internal/config`:
    - `constants.go`: shared constants for zodiac lists/maps, schedule hours, font paths, colors, font sizes, and helper like `GetImgPath(sign string)`.
    - `getEnvVar.go`: `GetEnvVar(varName string) string` for required environment variables.
  - `internal/daily`:
    - Orchestrates daily tasks such as `Horoscope`, `Compatibility`, `CompatibilityAndExplanation`, `DailyMoonPhase`, and `SignsBestAt`.
    - Each public function is designed to be callable from the scheduler and should encapsulate a complete unit of work.
  - `internal/scrap`:
    - HTML scraping logic using goquery; for example `UrlToDailyHoroscope(url string) string`.
  - `internal/deepl`:
    - DeepL API client and `DeepLService` wrapper for high‑level translation helpers.
  - `internal/x`:
    - Posting tweets and media to X/Twitter, including helpers to generate text and images for different content types.
  - `internal/images`:
    - Image generation pipeline built on `gg` for drawing text onto template images.
  - `internal/compatibilities`:
    - Data model and per‑pair compatibility content.
  - `internal/ctypes`:
    - Small shared enums and types for cross‑package coordination (for example `ImgGen`).
  - `internal/stringUtils`:
    - String helpers for title‑casing, parsing \"best at\" prompts, and other domain‑specific text manipulation.
  - `internal/http`:
    - Thin interface abstraction around `*http.Client` to allow dependency injection in services such as DeepL.

When adding new behavior, prefer adding a new small package under `internal/` or extending an existing one with cohesive functionality instead of growing `main.go`.

## Config and env usage

- Always access environment variables through `config.GetEnvVar` so missing configuration fails fast with a helpful message.
- Add new env vars to `.env.template` and document them in `README.md` when introducing new configuration.
- Put constant values and shared configuration (such as schedule hours, colors, font sizes, and image paths) in `internal/config/constants.go`:
  - For example:
    - `START_DAILY_TASK_HOUR`
    - `TIME_BETWEEN_POSTS`
    - `FONT_PATH`
    - `IMG_OUTPUT_PATH`
- Use helper functions like `config.GetImgPath(sign string)` instead of duplicating path logic across packages.

## Naming conventions

- Use clear, descriptive function names that indicate behavior:
  - `Horoscope`, `Compatibility`, `DailyMoonPhase`, `SignsBestAt` for scheduled entrypoints.
  - Helpers like `dailyTask`, `getDailyRandomSignsOfTheDay`, `dailyCompatibilityMapZodiacsToTweet` for internal flows.
- For exported symbols:
  - Use PascalCase and group related types and functions in the same file where practical.
  - Keep names concise but specific to their domain (for example `DeepLService`, `GenerateImageFromTemplate`, `TweetDailyHoroscope`).
- For unexported helpers:
  - Use lowerCamelCase and keep scope as small as possible.
- Keep type and file names aligned:
  - A `type Compatibility` lives in `internal/compatibilities/types.go`.
  - A variable `Ariesleo` lives in `internal/compatibilities/ariesleo.go`.

## Error handling and logging

- Use `log.Println` or `log.Printf` to trace major task boundaries and important events:
  - For example, log when a daily task starts and ends, and when external calls succeed or fail.
- Prefer early returns for error conditions rather than deeply nested `if` blocks.
- For configuration or asset problems that prevent the service from functioning, it is acceptable to `log.Fatal` during initialization (for example missing fonts or required env vars).
- For runtime data issues (such as empty scraped content), log the problem and return without crashing the process.

## HTTP and scheduling responsibilities

- `main.go` should:
  - Configure the HTTP server: routes, templates, and bind address/port.
  - Configure the `gocron` scheduler with specific times from `config`.
  - Avoid embedding domain logic beyond wiring calls to `daily` package functions.
- `internal/daily` should:
  - Own the orchestration of individual workflows: scraping, translating, generating images, and posting to X.
  - Provide exported functions that can be safely called from scheduled jobs or potentially from other entrypoints (for example CLI or manual triggers) in the future.

## When adding new code

When implementing new features or refactoring:

- **Place code** in the appropriate `internal/` package or create a new package if the logic is conceptually distinct.
- **Reuse existing patterns**:
  - Use the DeepL client and `DeepLService` for translations instead of creating a new HTTP client.
  - Use the `images` package and `ImgGen` types for image creation.
  - Use the `x` package for all X/Twitter interactions.
- **Wire through `main.go`** only for new entrypoints, HTTP handlers, or scheduled tasks; keep business logic out of `main.go`.

## References

- Entry point and scheduler: `main.go`
- Config and env helpers: `internal/config/constants.go`, `internal/config/getEnvVar.go`
- Daily jobs orchestration: `internal/daily/tasks.go`
- Scraping: `internal/scrap/horoscope.go`, `internal/scrap/moon.go`
- DeepL client: `internal/deepl/`
- X/Twitter integration: `internal/x/x.go`
- Image generation: `internal/images/generator.go`
- Compatibility data model and map: `internal/compatibilities/types.go`
