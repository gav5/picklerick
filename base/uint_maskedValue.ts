import uint_mask from './uint_mask'

/// Get the bit-masked value of the given value.
export default function uint_maskedValue(value: number, index: number): number {
  return value & uint_mask(index)
}
