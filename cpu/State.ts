import RegisterList from './RegisterList'

/// Record of enviornment that is saved on interrupt
export default class State {

  /// Describes the state of the registers for the given process.
  /// Used to "freeze" the processors registers during a context switch.
  /// When "unfrozen" the processor's registers will be set back to this state.
  registers: RegisterList

  // TODO: add description of what this does
  // TODO: figure out what type this should be
  permissions: any

  // TODO: add description of what this does
  // TODO: figure out what type this should be
  buffers: any

  // TODO: add description of what this does
  // TODO: figure out what type this should be
  caches: any

  // TODO: add description of what this does
  // TODO: figure out what type this should be
  // NOTE: could also be `activeBlocks` (seems like a design decision)
  activePages: any
}
