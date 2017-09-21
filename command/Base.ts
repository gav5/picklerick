import BinaryWord from '../base/BinaryWord'
import State from '../cpu/State'

/// Describes the bare minimum requirements for being a command to the CPU.
abstract class Base {

  /// All commands are initialized from a BinaryWord.
  constructor(word: BinaryWord) { }

  /// All commands must take an input state and transform it internally.
  /// (you will be working with a copy, so don't worry about shared state)
  abstract execute(State)

  /// All commands must provide what their assembly would look like.
  /// (this is for debugging and testing purposes, of course)
  abstract get assembly(): string
}

export default Base
