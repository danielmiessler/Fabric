'use strict'
const fs = require('graceful-fs')
const path = require('path')

function klawSync (dir, opts, ls) {
  if (!ls) {
    ls = []
    dir = path.resolve(dir)
    opts = opts || {}
    opts.fs = opts.fs || fs
    if (opts.depthLimit > -1) opts.rootDepth = dir.split(path.sep).length + 1
  }
  const paths = opts.fs.readdirSync(dir).map(p => dir + path.sep + p)
  for (var i = 0; i < paths.length; i += 1) {
    const pi = paths[i]
    const st = opts.fs.statSync(pi)
    const item = {path: pi, stats: st}
    const isUnderDepthLimit = (!opts.rootDepth || pi.split(path.sep).length - opts.rootDepth < opts.depthLimit)
    const filterResult = opts.filter ? opts.filter(item) : true
    const isDir = st.isDirectory()
    const shouldAdd = filterResult && (isDir ? !opts.nodir : !opts.nofile)
    const shouldTraverse = isDir && isUnderDepthLimit && (opts.traverseAll || filterResult)
    if (shouldAdd) ls.push(item)
    if (shouldTraverse) ls = klawSync(pi, opts, ls)
  }
  return ls
}

module.exports = klawSync
