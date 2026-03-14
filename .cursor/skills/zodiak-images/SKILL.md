---
name: zodiak-images
description: Describes the image generation pipeline in zodiak, including ImgGen types, templates, fonts, colors, and layout. Use when changing image layouts, adding image types, or editing template assets.
---

# Zodiak Image Generation

## Quick start

- All image generation is handled by the `internal/images` package, primarily via:
  - `GenerateImageFromTemplate(imgPath, text, maxWidthOffset, title, title2, subtitle, subtitleColor, imgType)`
- Image behavior is controlled by:
  - `ctypes.ImgGen` enum (`Horoscope`, `Compatibility`, `Moon`, `BestAt`).
  - Config values in `internal/config/constants.go` (font sizes, colors, output path).
  - Layout calculations in `internal/images/generator.go` (font size selection and text offsets).
- Image templates live in the embedded `assets/` directory:
  - `assets/compatibility.png`
  - `assets/moon.png`
  - `assets/prompt.png`
  - Per‑sign PNGs (for example `assets/aries.png`, `assets/tauro.png`, etc.).

## ImgGen enum

Defined in `internal/ctypes/types.go`:

- `type ImgGen int64`
- Constants:
  - `Horoscope`
  - `Compatibility`
  - `Moon`
  - `BestAt`

This type is used to:

- Select appropriate font sizes (`fontSizeByLength`).
- Position titles, body text, and subtitles (`calculateOffsets`).
- Choose colors for primary text (`getColorByType`).

## GenerateImageFromTemplate

Signature (simplified):

- `GenerateImageFromTemplate(imgPath string, text string, maxWidthOffset float64, title string, title2 string, subtitle string, subtitleColor color.Color, imgType ctypes.ImgGen)`

Behavior:

1. Opens the template image from `Assets` (embedded FS set in `main.go`).
2. Loads the font from `config.FONT_PATH` via the embedded FS.
3. Chooses font sizes based on text length and `imgType` using `fontSizeByLength`.
4. Draws:
   - Optional `title` and `title2` at positions defined by `calculateOffsets`.
   - Optional `subtitle` in a provided color.
   - Main `text` body with alignment depending on `imgType`:
     - Centered for most types.
     - Left-aligned for `BestAt`.
5. Saves the final image to `config.IMG_OUTPUT_PATH` (for example `out.png`).

When adjusting layouts or adding new types, prefer modifying `fontSizeByLength`, `calculateOffsets`, and `getColorByType` rather than scattering layout constants.

## Fonts, colors, and layout

From `internal/config/constants.go`:

- Font path and output:
  - `FONT_PATH` — path to `.ttf` font inside embedded assets (for example `assets/Timeburner.ttf`).
  - `IMG_OUTPUT_PATH` — path where generated images are written (for example `out.png`).
- Horoscope sizes:
  - `HOROSCOPE_MAX_FONT_SIZE`, `HOROSCOPE_MED_FONT_SIZE`, `HOROSCOPE_MIN_FONT_SIZE`
  - `HOROSCOPE_SUBTITLE_SIZE`
- Compatibility sizes:
  - `COMPAT_MAX_FONT_SIZE`, `COMPAT_MED_FONT_SIZE`, `COMPAT_MIN_FONT_SIZE`
  - `COMPAT_SUBTITLE_SIZE`, `COMPAT_TITLE_SIZE`, `COMPAT_TITLE2_SIZE`
- Moon sizes:
  - `MOON_MAX_FONT_SIZE`, `MOON_MED_FONT_SIZE`, `MOON_MIN_FONT_SIZE`
- Best-at sizes:
  - `BESTAT_TITLE_SIZE`, `BESTAT_MAX_FONT_SIZE`, `BESTAT_MED_FONT_SIZE`, `BESTAT_MIN_FONT_SIZE`
- Colors:
  - `HOROSCOPE_TEXT_COLOR`
  - `COMPAT_TEXT_COLOR`

From `internal/images/generator.go`:

- `fontSizeByLength(len int, imgType ctypes.ImgGen) float64`:
  - Chooses between max, medium, and minimum font sizes based on text length (thresholds around 650 and 900 characters).
  - Uses the appropriate size constants for each `ImgGen` value.
- `calculateOffsets(imgHeight int, imgType ctypes.ImgGen) TextOffsets`:
  - Returns `Title`, `Title2`, `Body`, and `Subtitle` Y offsets (pixels) for each image type.
  - Each image type has a custom layout so changes should keep relative structure in mind.
- `getColorByType(imgType ctypes.ImgGen) color.Color`:
  - Returns main text color based on `imgType` (for example, horoscope vs compatibility).
  - Falls back to a default color if no specific color is defined.

## Usage patterns in X posting

In `internal/x/x.go`:

- `TweetDailyHoroscope(sign string, tweet string, maxWidthOffset float64)`:
  - Uses `config.GetImgPath(sign)` to choose the per‑sign image.
  - Uses `ImgGen.Horoscope` and `config.HOROSCOPE_TEXT_COLOR`.
  - Adds the current day as a subtitle.
- `TweetDailyCompatibilityImg(text string, tweet string, maxWidthOffset float64, title1 string, title2 string, compatibility string)`:
  - Uses `assets/compatibility.png`.
  - Uses `ImgGen.Compatibility`.
  - The subtitle shows the compatibility percentage with a color derived from `compatibility`.
- `TweetDailyMoonPhaseImg(text string, tweet string, maxWidthOffset float64)`:
  - Uses `assets/moon.png`.
  - Uses `ImgGen.Moon`.
- `TweetDailyBestAtImg(text string, body []stringutils.BestAt, title string)`:
  - Uses `assets/prompt.png`.
  - Concatenates `BestAt` entries into a single multi‑line `tweet` string.
  - Uses `ImgGen.BestAt` with left-aligned body text.

Use these as blueprints when introducing new image‑based tweet types to ensure consistent layout and styling.

## Adding a new image type

To add a new `ImgGen` type (for example a new kind of daily card):

1. **Extend `ctypes`**:
   - Add a new constant to `internal/ctypes/types.go`.
2. **Add config constants** (if needed):
   - In `internal/config/constants.go`, add font size and color constants for the new type.
3. **Update image generator**:
   - Extend `fontSizeByLength` to handle the new `ImgGen` with appropriate `maxFs`, `medFs`, and `minFs`.
   - Extend `calculateOffsets` to set Y offsets for title(s), body, and subtitle.
   - Extend `getColorByType` to return the correct color for the new type.
4. **Provide a template asset**:
   - Add a base PNG under `assets/` (ensure `//go:embed assets/*` includes it, which it currently does).
5. **Add a high-level entrypoint**:
   - In `internal/x/x.go` (or another appropriate package), create a helper that calls `GenerateImageFromTemplate` with the new type and coordinates the tweet.

## Editing templates and layout

- When changing template images:
  - Keep overall dimensions consistent if possible so offsets and font sizes remain valid.
  - If dimensions change significantly, adjust `calculateOffsets` to keep text visually centered or well‑positioned.
- When adjusting text size:
  - Prefer tuning the constants in `config` rather than hardcoding numbers in the generator.
  - Keep a reasonable balance between max and min sizes to accommodate both short and long texts.

## References

- Image generator: `internal/images/generator.go`
- ImgGen type: `internal/ctypes/types.go`
- Config and colors: `internal/config/constants.go`
- X posting helpers that use images: `internal/x/x.go`
- Assets: `assets/` (embedded via `//go:embed assets/*` in `main.go`)
