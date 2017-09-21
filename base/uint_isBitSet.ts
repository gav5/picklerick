import uint_maskedValue from './uint_maskedValue'

/// Return the boolean value of the given bit for the given unsigned integer.
export default function uint_isBitSet(value: number, index: number): boolean {
  return uint_maskedValue(value, index) !== 0
}
