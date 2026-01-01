// contains synchronous API only so it can be exported as CJS and ESM

/** @type {import('..').isDepIncluded} */
function isDepIncluded(dep, optimizeDepsInclude) {
  return optimizeDepsInclude.some((id) => parseIncludeStr(id) === dep)
}

/** @type {import('..').isDepExcluded} */
function isDepExcluded(dep, optimizeDepsExclude) {
  dep = parseIncludeStr(dep)
  return optimizeDepsExclude.some(
    (id) => id === dep || dep.startsWith(`${id}/`)
  )
}

/** @type {import('..').isDepNoExternaled} */
function isDepNoExternaled(dep, ssrNoExternal) {
  if (ssrNoExternal === true) {
    return true
  } else {
    return isMatch(dep, ssrNoExternal)
  }
}

/** @type {import('..').isDepExternaled} */
function isDepExternaled(dep, ssrExternal) {
  return ssrExternal.includes(dep)
}

/**
 * @param {string} raw could be "foo" or "foo > bar" etc
 */
function parseIncludeStr(raw) {
  const lastArrow = raw.lastIndexOf('>')
  return lastArrow === -1 ? raw : raw.slice(lastArrow + 1).trim()
}

/**
 * @param {string} target
 * @param {string | RegExp | (string | RegExp)[]} pattern
 */
function isMatch(target, pattern) {
  if (Array.isArray(pattern)) {
    return pattern.some((p) => isMatch(target, p))
  } else if (typeof pattern === 'string') {
    return target === pattern
  } else if (pattern instanceof RegExp) {
    return pattern.test(target)
  }
}

module.exports = {
  isDepIncluded,
  isDepExcluded,
  isDepNoExternaled,
  isDepExternaled
}
