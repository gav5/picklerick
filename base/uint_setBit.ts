import uint_mask from './uint_mask'

/// Returns the transformed value by setting the given value at the given index.
export default function uint_setBit(value: number, index: number, bitVal: boolean = true): number {
  return bitVal ? (value | uint_mask(index)) : value
}
