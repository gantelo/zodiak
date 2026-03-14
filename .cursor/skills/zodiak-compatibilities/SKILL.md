---
name: zodiak-compatibilities
description: Documents compatibility data structures, map key conventions, categories, and per-pair files in zodiak. Use when adding, editing, or debugging compatibility pairs or compatibility-related features.
---

# Zodiak Compatibilities Data

## Quick start

- Compatibility logic and data live in `internal/compatibilities`.
- The core types are defined in `types.go`:
  - `Friendship`
  - `SummaryPercentage`
  - `Love`
  - `Compatibility`
- All compatibility content is stored as Go data, not external JSON.
- The main map `Compatibilities` is keyed by **concatenated English sign names** in lowercase:
  - `"ariesleo"`, `"leoaries"`, `"taurusgemini"`, etc.

## Core types

Defined in `internal/compatibilities/types.go`:

- `Friendship`:
  - `Match string` — compatibility percentage (for example `"95%"`).
  - `Traits []string` — list of short label constants describing the relationship (for example `GoodRomancePotential`, `DeepMutualUnderstanding`).
  - `Summary string` — long descriptive Spanish text for friendship.

- `SummaryPercentage`:
  - `Text string` — long text describing a dimension (summary, sexual intimacy, trust, etc).
  - `Match string` — percentage for that dimension (for example `"90%"`).

- `Love`:
  - `Match string`
  - `Traits []string`
  - `Summary SummaryPercentage`
  - `SexualIntimacy SummaryPercentage`
  - `Trust SummaryPercentage`
  - `Communication SummaryPercentage`
  - `Emotions SummaryPercentage`
  - `Values SummaryPercentage`
  - `SharedActivities SummaryPercentage`

- `Compatibility`:
  - `Name string` — text label like `"#aries ♈ - #leo ♌"`.
  - `Friendship Friendship`
  - `Love Love`

There are also many reusable text constants such as `CautionRomance`, `FunEnjoyable`, `MutualRespect`, and sign hash‑tag constants like `Aries`, `Leo`, etc.

## Compatibilities map

- `Compatibilities` is a map of `map[string]Compatibility`.
- Keys are concatenated English sign names, always lowercase:
  - `"ariesleo"`
  - `"leoscorpio"`
  - `"virgopisces"`
  - and so on for all pairs.
- Reverse keys map to the same `Compatibility` value by reusing the existing variable:
  - `"ariesleo": Ariesleo`
  - `"leoaries": Ariesleo`

This means:

- **Each unordered pair** should be represented once as a `Compatibility` value.
- The map should have **both directions** (`"ariesleo"` and `"leoaries"`) pointing at that same value.

## Per-pair files

- Each pair has its own file under `internal/compatibilities/`:
  - For example `ariesleo.go` defines `var Ariesleo = Compatibility{ ... }`.
  - The file typically contains:
    - A `Compatibility` value
    - Long string texts for friendship, love summary, sexual intimacy, trust, communication, emotions, values, and shared activities.
- When introducing a new pair, follow these steps:
  1. **Create a file** named `<sign1><sign2>.go` with lowercase English sign names in the filename (for example `ariesleo.go`).
  2. **Export a variable** whose name is the two sign names concatenated with appropriate casing (for example `Ariesleo`):
     - `var Ariesleo = Compatibility{ ... }`
  3. **Populate** the `Compatibility` value:
     - Set `Name` using sign constants like `Aries` and `Leo`.
     - Fill `Friendship` with `Match`, `Traits`, and `Summary`.
     - Fill `Love` and its nested `SummaryPercentage` fields.
  4. **Register in `Compatibilities` map** inside `types.go`:
     - Add an entry for both key orders pointing to the same variable:
       - `"ariesleo": Ariesleo,`
       - `"leoaries": Ariesleo,`

## Categories and explanation text

- In `types.go`, `CompatibilityCategories` lists all available categories:
  - `"Friendship"`, `"Summary"`, `"SexualIntimacy"`, `"Trust"`, `"Communication"`, `"Emotions"`, `"Values"`, `"SharedActivities"`.
- `CompatiblitiesExplanation` is a constant Spanish text explaining how compatibility posts work.

Daily tasks use these values as follows:

- `internal/daily/dailyCompatibilityMapZodiacsToTweet`:
  - Reads the selected `Compatibility` from the `Compatibilities` map.
  - Switches on category to build:
    - A header string describing the context (for example \"en la amistad\", \"y la confianza\").
    - A `Friendship`-shaped struct representing the chosen dimension, containing `Match`, `Summary`, and `Traits`.
  - Passes this data to `x.TweetDailyCompatibilityImg`.

When adding or changing categories:

- Ensure new category labels exist in `CompatibilityCategories`.
- Update the switch in `dailyCompatibilityMapZodiacsToTweet` so it correctly pulls from the appropriate `Love` or `Friendship` sub‑field and sets a meaningful header.

## Adding or editing pairs

When adding a new compatibility pair or adjusting existing data:

1. **Check the map**:
   - Open `internal/compatibilities/types.go`.
   - Confirm which keys already exist for the pair (for both directions).
2. **Create or edit the per-pair file**:
   - If a file does not yet exist, create it and define the `Compatibility` variable.
   - If the file exists, update text, traits, or match percentages there.
3. **Ensure map entries are correct**:
   - For each unordered pair, both directional keys should be present in `Compatibilities`.
   - Both keys should point to the same `Compatibility` variable.
4. **Keep language consistent**:
   - Texts should be in Spanish, in the same tone and style as existing entries.
   - Use existing constants for repeated trait descriptions whenever possible.

## Debugging compatibility issues

- If a daily compatibility tweet fails or looks wrong:
  - Check the keys used in `dailyCompatibilityMapZodiacsToTweet`:
    - It constructs keys as simple concatenations of English sign names from `config.ZodiacSignsArray`.
  - Verify those keys exist in `Compatibilities` and point to the expected variables.
  - Check that the category name returned by `getDailyRandomSignsOfTheDay` is included in `CompatibilityCategories` and handled in the switch.

## References

- Types and map: `internal/compatibilities/types.go`
- Example per-pair content: `internal/compatibilities/ariesleo.go`
- Daily mapping and tweet building: `internal/daily/tasks.go`
