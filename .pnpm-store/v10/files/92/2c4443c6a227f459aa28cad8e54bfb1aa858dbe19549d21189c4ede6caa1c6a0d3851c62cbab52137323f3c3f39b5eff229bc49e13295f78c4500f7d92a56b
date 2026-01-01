import { urlAlphabet as scopedUrlAlphabet } from './url-alphabet/index.js'
export { urlAlphabet } from './url-alphabet/index.js'
export let random = bytes => crypto.getRandomValues(new Uint8Array(bytes))
export let customRandom = (alphabet, defaultSize, getRandom) => {
  let mask = (2 << Math.log2(alphabet.length - 1)) - 1
  let step = -~((1.6 * mask * defaultSize) / alphabet.length)
  return (size = defaultSize) => {
    let id = ''
    while (true) {
      let bytes = getRandom(step)
      let j = step | 0
      while (j--) {
        id += alphabet[bytes[j] & mask] || ''
        if (id.length >= size) return id
      }
    }
  }
}
export let customAlphabet = (alphabet, size = 21) =>
  customRandom(alphabet, size | 0, random)
export let nanoid = (size = 21) => {
  let id = ''
  let bytes = crypto.getRandomValues(new Uint8Array((size |= 0)))
  while (size--) {
    id += scopedUrlAlphabet[bytes[size] & 63]
  }
  return id
}
