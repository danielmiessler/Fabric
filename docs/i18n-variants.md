# Language Variants Support in Fabric

## Current Implementation

As of this update, Fabric supports these language variants:

- **Portuguese variants**:
  - `pt-BR` - Brazilian Portuguese
  - `pt-PT` - European Portuguese
  - `pt` - defaults to `pt-BR` for backward compatibility

- **Arabic variants**:
  - `ar` - Modern Standard Arabic (defaults to `ar-BH`)
  - `ar-BH` - Bahrain Arabic (Modern Standard Arabic)
  - `ar-SA` - Saudi Arabic (Modern Standard Arabic)
  - Arabic support added in 2026, with default to Bahrain Arabic variant

## Architecture

The i18n system supports language variants through:

1. **BCP 47 Format**: All locales are normalized to BCP 47 format (language-REGION)
2. **Fallback Chain**: Regional variants fall back to base language, then to configured defaults
3. **Default Variant Mapping**: Languages without base files can specify default regional variants
4. **Flexible Input**: Accepts both underscore (pt_BR) and hyphen (pt-BR) formats

## Recommended Future Variants

Based on user demographics and linguistic differences, these variants would provide the most value:

### High Priority

1. **Chinese Variants**
   - `zh-CN` - Simplified Chinese (Mainland China)
   - `zh-TW` - Traditional Chinese (Taiwan)
   - `zh-HK` - Traditional Chinese (Hong Kong)
   - Default: `zh` → `zh-CN`
   - Rationale: Significant script and vocabulary differences

2. **Spanish Variants**
   - `es-ES` - European Spanish (Spain)
   - `es-MX` - Mexican Spanish
   - `es-AR` - Argentinian Spanish
   - Default: `es` → `es-ES`
   - Rationale: Notable vocabulary and conjugation differences

3. **English Variants**
   - `en-US` - American English
   - `en-GB` - British English
   - `en-AU` - Australian English
   - Default: `en` → `en-US`
   - Rationale: Spelling differences (color/colour, organize/organise)

4. **French Variants**
   - `fr-FR` - France French
   - `fr-CA` - Canadian French
   - Default: `fr` → `fr-FR`
   - Rationale: Some vocabulary and expression differences

5. **Arabic Variants**
   - `ar-SA` - Saudi Arabic (Modern Standard)
   - `ar-EG` - Egyptian Arabic
   - Default: `ar` → `ar-SA`
   - Rationale: Significant dialectal differences

6. **German Variants**
   - `de-DE` - Germany German
   - `de-AT` - Austrian German
   - `de-CH` - Swiss German
   - Default: `de` → `de-DE`
   - Rationale: Minor differences, mostly vocabulary

## Implementation Guidelines

When adding new language variants:

1. **Determine the Base**: Decide which variant should be the default
2. **Create Variant Files**: Copy base file and adjust for regional differences
3. **Update Default Map**: Add to `defaultLanguageVariants` if needed
4. **Focus on Key Differences**:
   - Technical terminology
   - Common UI terms (file/ficheiro, save/guardar)
   - Date/time formats
   - Currency references
   - Formal/informal address conventions

5. **Test Thoroughly**: Ensure fallback chain works correctly

## Adding a New Variant

To add a new language variant:

1. Copy the base language file:

   ```bash
   cp locales/es.json locales/es-MX.json
   ```

2. Adjust translations for regional differences

3. If this is the first variant for a language, update `i18n.go`:

   ```go
   var defaultLanguageVariants = map[string]string{
       "pt": "pt-BR",
       "es": "es-MX",  // Add if Mexican Spanish should be default
   }
   ```

4. Add tests for the new variant

5. Update documentation

## Language Variant Naming Convention

Follow BCP 47 standards:

- Language code: lowercase (pt, es, en)
- Region code: uppercase (BR, PT, US)
- Separator: hyphen (pt-BR, not pt_BR)

Input normalization handles various formats, but files and internal references should use BCP 47.

## Testing Variants

Test each variant with:

```bash
# Direct specification
fabric --help -g=pt-BR
fabric --help -g=pt-PT

# Environment variable
LANG=pt_BR.UTF-8 fabric --help

# Fallback behavior
fabric --help -g=pt  # Should use pt-BR
```

## Maintenance Considerations

When updating translations:

1. Update all variants of a language together
2. Ensure key parity across all variants
3. Test fallback behavior after changes
4. Consider using translation memory tools for consistency
