// This module monkey patches Page#showModal in order to be able to
// access from the HMR proxy data passed to `showModal` in svelte-native.
//
// Data are stored in a opaque prop accessible with `getModalData`.
//
// It also switches the `closeCallback` option with a custom brewed one
// in order to give the proxy control over when its own instance will be
// destroyed.
//
// Obviously this method suffer from extreme coupling with the target code
// in svelte-native. So it would be wise to recheck compatibility on SN
// version upgrades.
//
// Relevant code is there (last checked version):
//
// https://github.com/halfnelson/svelte-native/blob/48fdc97d2eb4d3958cfcb4ff6cf5755a220829eb/src/dom/navigation.ts#L132
//

// FIXME should we override ViewBase#showModal instead?
// eslint-disable-next-line import/no-unresolved
import { Page } from '@nativescript/core'

const prop =
  typeof Symbol !== 'undefined'
    ? Symbol('hmr_svelte_native_modal')
    : '___HMR_SVELTE_NATIVE_MODAL___'

const sup = Page.prototype.showModal

let patched = false

export const patchShowModal = () => {
  // guard: already patched
  if (patched) return
  patched = true

  Page.prototype.showModal = function(modalView, options) {
    const modalData = {
      originalOptions: options,
      closeCallback: options.closeCallback,
    }

    modalView[prop] = modalData

    // Proxies to a function that can be swapped on the fly by HMR proxy.
    //
    // The default is still to call the original closeCallback from svelte
    // navtive, which will destroy the modal view & component. This way, if
    // no HMR happens on the modal content, normal behaviour is preserved
    // without the proxy having any work to do.
    //
    const closeCallback = (...args) => {
      return modalData.closeCallback(...args)
    }

    const tamperedOptions = Object.assign({}, options, { closeCallback })

    return sup.call(this, modalView, tamperedOptions)
  }
}

export const getModalData = modalView => modalView[prop]
