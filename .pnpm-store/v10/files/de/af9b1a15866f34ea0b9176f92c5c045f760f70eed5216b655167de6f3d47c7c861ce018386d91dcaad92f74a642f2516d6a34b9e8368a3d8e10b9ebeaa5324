const createMakeHotFactory = require('./lib/make-hot.js')
const { resolve } = require('path')
const { name, version } = require('./package.json')

const resolveAbsoluteImport = target => resolve(__dirname, target)

const createMakeHot = createMakeHotFactory({
  pkg: { name, version },
  resolveAbsoluteImport,
})

module.exports = { createMakeHot }
