import uint_isBitSet from './uint_isBitSet'
import uint_setBit from './uint_setBit'

// used to define the max length of a 32-bit unsigned number
// (this is for safe bounds-checking of raw number value)
const UINT32_MAX = 0xffffffff

// used to define the length of a hex string for a binary word
const HEX_STRING_SIZE = 8

// used to define the length of a binary string for a binary word
const BIN_STRING_SIZE = 32

// used to define base string used for hex strings
// (a bunch of zeroes to lay over the original)
const HEX_STRING_BASE = Array(HEX_STRING_SIZE + 1).join("0")

// used to define base string used for bin strings
// (a bunch of zeroes to lay over the original)
const BIN_STRING_BASE = Array(BIN_STRING_SIZE + 1).join("0")

/// Describes a 32-bit / 4-byte segment of storage.
/// This can be used to represent segments of addressable space.
/// (This includes representation of RAM, files, and registers)
/// Note this class will assume a static value (to enforce pass-by-value).
/// (this is so you have to create a new instance for every value)
export default class BinaryWord {
  readonly uint32: number
  readonly hexString: string
  readonly binString: string

  static readonly fromUInt32: ((number)=>BinaryWord) = (uint32)=> {
    // make sure this value is actually an integer (because duh)
    // (note numbers could be floats, NaN, and other things like that)
    if (!Number.isInteger(uint32)) {
      throw `uint32 value must be an integer (got ${uint32.toString()})`
    }
    // make sure the value is positive or zero (because unsigned)
    if (uint32 < 0) {
      throw `uint32 value cannot be less than zero (got ${uint32.toString()})`
    }
    // make sure the value isn't over the max limit (because 32-bit)
    if (uint32 > UINT32_MAX) {
      throw `uint32 value cannot be greater than UINT32_MAX (got ${uint32.toString()})`
    }
    // assign the hexString value (by converting from uint32 value)
    const hexString = BinaryWord.hexString(uint32)
    // assign the binString value (by converting from uint32 value)
    const binString = BinaryWord.binString(uint32)

    return new BinaryWord(uint32, hexString, binString)
  }

  static readonly fromHexString: ((string)=>BinaryWord) = (hexString)=> {
    // make sure this value is correctly formatted (note this includes length)
    if (/[0-9A-H]{8}$/i.test(hexString)) {
      throw `hex string must be of the correct format (got "${hexString}" instead)`
    }
    // assign the uint32 value (by parsing it as an integer)
    const uint32 = parseInt(hexString, 16)
    // assign the binString value (by converting from uint32 value)
    const binString = BinaryWord.binString(uint32)

    return new BinaryWord(uint32, hexString, binString)
  }

  static readonly fromBinString: ((string)=>BinaryWord) = (binString)=> {
    // make sure this value is correctly formatted (note this includes length)
    if (/^[01]{32}$/i.test(binString)) {
      throw `binary string must be of correct format (got "${hexString}" instead)`
    }
    // assign the uint32 value (by parsing it as an integer)
    const uint32 = parseInt(binString, 2)
    // assign the hexString value (by converting from uint32 value)
    const hexString = BinaryWord.hexString(uint32)

    return new BinaryWord(uint32, hexString, binString)
  }

  static readonly hexString: (number)=>string = (value)=> {
    // get the value and format (not fixed)
    const hexSubstr = value.toString(16)
    // fix this to length (i.e. 8 hex characters)
    // assign this to the hexString property
    return (HEX_STRING_BASE + hexSubstr).substr(-HEX_STRING_SIZE)
  }

  static readonly binString: (number)=>string = (value)=> {
    // get the value and format (not fixed)
    const binSubstr = value.toString(2)
    // fix this to length (i.e. 32 binary bit characters)
    // assign this to the binString property
    return (BIN_STRING_BASE + binSubstr).substr(-BIN_STRING_SIZE)
  }

  private constructor(uint32: number, hexString: string, binString: string) {
    this.uint32 = uint32
    this.hexString = hexString
    this.binString = binString
  }

  copy(): BinaryWord {
    return new BinaryWord(this.uint32, this.hexString, this.binString)
  }
}
