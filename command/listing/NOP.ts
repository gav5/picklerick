import BinaryWord from '../../base/BinaryWord'
import State from '../../cpu/State'
import Base from '../Base'

class NOP extends Base {

  constructor(word: BinaryWord) {
    super(word)
  }

  execute(state: State) {
    // do nothing!
  }

  get assembly(): string {
    return "NOP"
  }
}
