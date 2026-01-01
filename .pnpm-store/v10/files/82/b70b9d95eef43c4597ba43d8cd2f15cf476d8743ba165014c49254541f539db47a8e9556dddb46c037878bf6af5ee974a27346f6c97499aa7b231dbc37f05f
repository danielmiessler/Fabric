// eslint-disable-next-line node/no-deprecated-api
const { createRequire, createRequireFromPath } = require('module')

function req (name, rootFile) {
  const create = createRequire || createRequireFromPath
  const require = create(rootFile)
  return require(name)
}

module.exports = req
