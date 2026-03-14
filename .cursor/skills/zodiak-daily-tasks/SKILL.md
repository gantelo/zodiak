---
name: zodiak-daily-tasks
description: Guides daily cron tasks in the zodiak service, including scraping, DeepL translation, image generation, and X posting. Use when adding, modifying, or debugging daily tasks or scheduler configuration.
---

# Zodiak Daily Tasks and Scheduler

## Quick start

- All recurring jobs are registered in `main.go` using `github.com/go-co-op/gocron` with **UTC** times from `internal/config/constants.go`.
- Each scheduled job calls an exported function in `internal/daily`:
  - `Horoscope`
  - `Compatibility`
  - `CompatibilityAndExplanation`
  - `DailyMoonPhase`
  - `SignsBestAt`
- A typical workflow is:
  1. Get or scrape source data (for example web page, embedded JSON).
  2. Optionally translate content with DeepL (`deepl.NewDeepLService().TranslateToSpanish`).
  3. Optionally generate an image using the `images` package.
  4. Post text or media via the `x` package.

## Scheduler wiring in main.go

- `main.go` sets up the scheduler:
  - Creates a new scheduler: `s := gocron.NewScheduler(time.UTC)`.
  - Registers tasks with times from `internal/config/constants.go`:
    - `START_DAILY_TASK_HOUR`
    - `START_DAILY_COMPATIBILITY_TASK_HOUR`
    - `START_DAILY_COMPATIBILITY_TASK_HOUR_2`
    - `START_DAILY_MOON_PHASE_TASK_HOUR`
    - `START_DAILY_BESTAT_TASK_HOUR`
  - Calls `s.StartBlocking()` so the scheduler runs for the lifetime of the process.

When introducing a new recurring job:

1. Add an exported function to `internal/daily` that encapsulates the job.
2. Add a corresponding schedule entry in `main.go` using one of the existing time constants or a new one.
3. If you introduce a new schedule time, add a constant in `internal/config/constants.go`.

## Daily horoscope workflow

Implemented in `internal/daily`:

- `Horoscope()`:
  - Logs start and end of the task.
  - Iterates over `config.ZodiacSigns` and calls `dailyTask(sign)` for each.
  - Sleeps for `config.TIME_BETWEEN_POSTS` between signs to avoid rate limits and bursty behavior.

- `dailyTask(sign string)`:
  1. Builds the source URL: `SCRAP_WEB` + sign + `WEB_SUFFIX` using `config.GetEnvVar` and `config.WEB_SUFFIX`.
  2. Scrapes text with `scrap.UrlToDailyHoroscope(url)`:
     - Finds a CSS selector from `SCRAP_CLASS`.
     - Extracts the first and last `<p>` elements as text.
  3. Translates the combined text to Spanish using `DeepLService.TranslateToSpanish`.
  4. Maps the English sign to a Spanish label using `config.ZodiacSigns`.
  5. Formats the tweet body, inserting newlines between sentences.
  6. Calls `x.TweetDailyHoroscope` to generate an image and post it.

To add a similar per‑sign workflow, mirror this structure:

- Create a helper like `dailyTask` that:
  - Accepts a sign.
  - Fetches or computes content.
  - Translates if needed.
  - Calls the appropriate `x` image/tweet helper.
- Have the exported function loop over `config.ZodiacSigns` and sleep between posts.

## Compatibility workflows

- `Compatibility()`:
  - Selects two random signs and a random category by calling `getDailyRandomSignsOfTheDay`.
  - Calls `dailyCompatibilityMapZodiacsToTweet` to build and post a compatibility tweet.

- `CompatibilityAndExplanation()`:
  - Performs the same steps as `Compatibility`.
  - After a short delay, posts a static explanation text `compatibilities.CompatiblitiesExplanation` via `x.Tweet`.

- `getDailyRandomSignsOfTheDay()`:
  - Randomly picks two indices from `config.ZodiacSignsArray`.
  - Randomly selects a category from `compatibilities.CompatibilityCategories`.
  - Returns the chosen signs and category.

- `dailyCompatibilityMapZodiacsToTweet(zodiac1, zodiac2, category string)`:
  - Looks up the `Compatibility` struct in `compatibilities.Compatibilities` using keys like `"ariesleo"` or `"leoaries"`.
  - Selects a specific slice of data based on `category` (for example `Friendship`, `Summary`, `SexualIntimacy`, etc.).
  - Constructs a multi‑line header string describing the relationship and overall compatibility.
  - Builds a Spanish name string from `config.ZodiacSigns` and passes data through to `x.TweetDailyCompatibilityImg`.

When extending compatibility behavior:

- Use existing categories from `CompatibilityCategories` or add new ones with clear semantics.
- Ensure new categories are handled both in the `compatibilities` data and in `dailyCompatibilityMapZodiacsToTweet`.

## Moon phase workflow

- `DailyMoonPhase()`:
  - Reads `MOON_SCRAP_WEB` from the environment via `config.GetEnvVar`.
  - Uses `scrap.UrlTomorrowMoon` to get the header and body describing the next day’s moon phase.
  - Translates the body to Spanish via `DeepLService.TranslateToSpanish`.
  - Formats the tweet body with line breaks.
  - Calls `x.TweetDailyMoonPhaseImg` to generate an image and post the tweet.

Use this flow as a template when adding other non‑horoscope daily content that has a single piece of text instead of per‑sign data.

## \"Best at\" workflow

- `SignsBestAt()`:
  - Reads `data/prompts.json` and a weekday‑specific JSON file from the embedded `data` filesystem.
  - Picks a prompt string via `stringUtils.ParseBestAtPrompt`.
  - Calls `signsBestAt(prompt, responses)` to generate content and tweet it.

- `signsBestAt(prompt string, responses []byte)`:
  - Translates the prompt to Spanish with DeepL.
  - Constructs a full prompt string (`"¿Qué tan buenos son los signos en ...?"`).
  - Parses the response JSON with `stringUtils.ParseBestAtArray` into a slice of `BestAt` items.
  - Calls `x.TweetDailyBestAtImg` to generate an image listing all signs with descriptions.

This shows how to combine embedded data (`data/*.json`), translation, and image generation in a daily task.

## Adding a new daily task

When introducing a new scheduled behavior:

1. **Design the pipeline**:
   - Decide what the source data is (scraped HTML, third‑party API, embedded JSON, or computed data).
   - Decide whether it needs translation and/or images.
   - Decide which `x` helpers to reuse or extend (text‑only vs image‑based tweets).
2. **Implement the workflow in `internal/daily`**:
   - Add an exported function that logs start/end and orchestrates the full flow.
   - Add private helpers as needed to keep the exported function readable.
3. **Wire scheduler in `main.go`**:
   - Use `config` constants for schedule times; add new ones in `internal/config/constants.go` if required.
4. **Configure env and data**:
   - Add any new env vars to `.env.template` and `README.md`.
   - Add any embedded data files to the `data/` directory and ensure they are covered by the `//go:embed data/*` directive in `main.go`.

## References

- Scheduler and wiring: `main.go`
- Daily task orchestration: `internal/daily/tasks.go`
- Scraping helpers: `internal/scrap/`
- Translation helpers: `internal/deepl/`
- Image generation: `internal/images/generator.go`
- X/Twitter posting: `internal/x/x.go`
