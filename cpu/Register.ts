import BinaryWord from '../base/BinaryWord'

/// Describes a single register state in the processor.
/// This represents both a register in memory as well as the hardware itself.
/// (note this is merely an alias because a register is a 32-bit value)
type Register = BinaryWord

export default Register
