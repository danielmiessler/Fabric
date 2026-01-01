import { webcrypto as crypto } from 'node:crypto'
import { urlAlphabet as scopedUrlAlphabet } from './url-alphabet/index.js'
export { urlAlphabet } from './url-alphabet/index.js'
const POOL_SIZE_MULTIPLIER = 128
let pool, poolOffset
function fillPool(bytes) {
  if (!pool || pool.length < bytes) {
    pool = Buffer.allocUnsafe(bytes * POOL_SIZE_MULTIPLIER)
    crypto.getRandomValues(pool)
    poolOffset = 0
  } else if (poolOffset + bytes > pool.length) {
    crypto.getRandomValues(pool)
    poolOffset = 0
  }
  poolOffset += bytes
}
export function random(bytes) {
  fillPool((bytes |= 0))
  return pool.subarray(poolOffset - bytes, poolOffset)
}
export function customRandom(alphabet, defaultSize, getRandom) {
  let mask = (2 << (31 - Math.clz32((alphabet.length - 1) | 1))) - 1
  let step = Math.ceil((1.6 * mask * defaultSize) / alphabet.length)
  return (size = defaultSize) => {
    let id = ''
    while (true) {
      let bytes = getRandom(step)
      let i = step
      while (i--) {
        id += alphabet[bytes[i] & mask] || ''
        if (id.length >= size) return id
      }
    }
  }
}
export function customAlphabet(alphabet, size = 21) {
  return customRandom(alphabet, size, random)
}
export function nanoid(size = 21) {
  fillPool((size |= 0))
  let id = ''
  for (let i = poolOffset - size; i < poolOffset; i++) {
    id += scopedUrlAlphabet[pool[i] & 63]
  }
  return id
}
