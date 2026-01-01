/**
 * Injects a `{}*` CSS rule to force Svelte compiler to scope every elements.
 */
export const injectScopeEverythingCssRule = (parse, code) => {
  const { css } = parse(code)
  if (!css) return code
  const {
    content: { end },
  } = css
  return code.slice(0, end) + '*{}' + code.slice(end)
}

export const normalizeJsCode = code => {
  // Svelte now adds locations in dev mode, code locations can change when
  // CSS change, but we're unaffected (not real behaviour changes)
  code = code.replace(/\badd_location\s*\([^)]*\)\s*;?/g, '')
  return code
}
