import Register from './Register'

/// Represents all of the registers in the processor.
/// This is used both as a memory model as well as a hardware implementation.
export default class RegisterList {
  /// Reg 00: Accumulator
  r0: Register

  /// Reg 01: Zero Register
  /// (this value must always be zero)
  r1: Register & 0

  /// Reg 02: General-Purpose Register
  r2: Register

  /// Reg 03: General-Purpose Register
  r3: Register

  /// Reg 04: General-Purpose Register
  r4: Register

  /// Reg 05: General-Purpose Register
  r5: Register

  /// Reg 06: General-Purpose Register
  r6: Register

  /// Reg 07: General-Purpose Register
  r7: Register

  /// Reg 08: General-Purpose Register
  r8: Register

  /// Reg 09: General-Purpose Register
  r9: Register

  /// Reg 10: General-Purpose Register
  r10: Register

  /// Reg 11: General-Purpose Register
  r11: Register

  /// Reg 12: General-Purpose Register
  r12: Register

  /// Reg 13: General-Purpose Register
  r13: Register

  /// Reg 14: General-Purpose Register
  r14: Register

  /// Reg 15: General-Purpose Register
  r15: Register

  /// Reg 16: General-Purpose Register
  r16: Register
}
