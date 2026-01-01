/* global document */

import { adapter as ProxyAdapterDom } from '../proxy-adapter-dom'

import { patchShowModal, getModalData } from './patch-page-show-modal'

patchShowModal()

// Svelte Native support
// =====================
//
// Rerendering Svelte Native page proves challenging...
//
// In NativeScript, pages are the top level component. They are normally
// introduced into NativeScript's runtime by its `navigate` function. This
// is how Svelte Natives handles it: it renders the Page component to a
// dummy fragment, and "navigate" to the page element thus created.
//
// As long as modifications only impact child components of the page, then
// we can keep the existing page and replace its content for HMR.
//
// However, if the page component itself is modified (including its system
// title bar), things get hairy...
//
// Apparently, the sole way of introducing a new page in a NS application is
// to navigate to it (no way to just replace it in its parent "element", for
// example). This is how it is done in NS's own "core" HMR.
//
// NOTE The last paragraph has not really been confirmed with NS6.
//
// Unfortunately the API they're using to do that is not public... Its various
// parts remain exposed though (but documented as private), so this exploratory
// work now relies on it. It might be fragile...
//
// The problem is that there is no public API that can navigate to a page and
// replace (like location.replace) the current history entry. Actually there
// is an active issue at NS asking for that. Incidentally, members of
// NativeScript-Vue have commented on the issue to weight in for it -- they
// probably face some similar challenge.
//
// https://github.com/NativeScript/NativeScript/issues/6283

const getNavTransition = ({ transition }) => {
  if (typeof transition === 'string') {
    transition = { name: transition }
  }
  return transition ? { animated: true, transition } : { animated: false }
}

export const adapter = class ProxyAdapterNative extends ProxyAdapterDom {
  constructor(instance) {
    super(instance)

    this.nativePageElement = null
  }

  dispose() {
    super.dispose()
    this.releaseNativePageElement()
  }

  releaseNativePageElement() {
    if (this.nativePageElement) {
      // native cleaning will happen when navigating back from the page
      this.nativePageElement = null
    }
  }

  afterMount(target, anchor) {
    // nativePageElement needs to be updated each time (only for page
    // components, native component that are not pages follow normal flow)
    //
    // TODO quid of components that are initially a page, but then have the
    // <page> tag removed while running? or the opposite?
    //
    // insertionPoint needs to be updated _only when the target changes_ --
    // i.e. when the component is mount, i.e. (in svelte3) when the component
    // is _created_, and svelte3 doesn't allow it to move afterward -- that
    // is, insertionPoint only needs to be created once when the component is
    // first mounted.
    //
    // TODO is it really true that components' elements cannot move in the
    // DOM? what about keyed list?
    //

    const isNativePage =
      (target.tagName === 'fragment' || target.tagName === 'frame') &&
      target.firstChild &&
      target.firstChild.tagName == 'page'
    if (isNativePage) {
      const nativePageElement = target.firstChild
      this.nativePageElement = nativePageElement
    } else {
      // try to protect against components changing from page to no-page
      // or vice versa -- see DEBUG 1 above. NOT TESTED so prolly not working
      this.nativePageElement = null
      super.afterMount(target, anchor)
    }
  }

  rerender() {
    const { nativePageElement } = this
    if (nativePageElement) {
      this.rerenderNative()
    } else {
      super.rerender()
    }
  }

  rerenderNative() {
    const { nativePageElement: oldPageElement } = this
    const nativeView = oldPageElement.nativeView
    const frame = nativeView.frame
    if (frame) {
      return this.rerenderPage(frame, nativeView)
    }
    const modalParent = nativeView._modalParent // FIXME private API
    if (modalParent) {
      return this.rerenderModal(modalParent, nativeView)
    }
    // wtf? hopefully a race condition with a destroyed component, so
    // we have nothing more to do here
    //
    // for once, it happens when hot reloading dev deps, like this file
    //
  }

  rerenderPage(frame, previousPageView) {
    const isCurrentPage = frame.currentPage === previousPageView
    if (isCurrentPage) {
      const {
        instance: { hotOptions },
      } = this
      const newPageElement = this.createPage()
      if (!newPageElement) {
        throw new Error('Failed to create updated page')
      }
      const isFirstPage = !frame.canGoBack()
      const nativeView = newPageElement.nativeView
      const navigationEntry = Object.assign(
        {},
        {
          create: () => nativeView,
          clearHistory: true,
        },
        getNavTransition(hotOptions)
      )

      if (isFirstPage) {
        // NOTE not so sure of bellow with the new NS6 method for replace
        //
        // The "replacePage" strategy does not work on the first page
        // of the stack.
        //
        // Resulting bug:
        // - launch
        // - change first page => HMR
        // - navigate to other page
        // - back
        //   => actual: back to OS
        //   => expected: back to page 1
        //
        // Fortunately, we can overwrite history in this case.
        //
        frame.navigate(navigationEntry)
      } else {
        frame.replacePage(navigationEntry)
      }
    } else {
      const backEntry = frame.backStack.find(
        ({ resolvedPage: page }) => page === previousPageView
      )
      if (!backEntry) {
        // well... looks like we didn't make it to history after all
        return
      }
      // replace existing nativeView
      const newPageElement = this.createPage()
      if (newPageElement) {
        backEntry.resolvedPage = newPageElement.nativeView
      } else {
        throw new Error('Failed to create updated page')
      }
    }
  }

  // modalParent is the page on which showModal(...) was called
  // oldPageElement is the modal content, that we're actually trying to reload
  rerenderModal(modalParent, modalView) {
    const modalData = getModalData(modalView)

    modalData.closeCallback = () => {
      const nativePageElement = this.createPage()
      if (!nativePageElement) {
        throw new Error('Failed to created updated modal page')
      }
      const { nativeView } = nativePageElement
      const { originalOptions } = modalData
      // Options will get monkey patched again, the only work left for us
      // is to try to reduce visual disturbances.
      //
      // FIXME Even that proves too much unfortunately... Apparently TNS
      // does not respect the `animated` option in this context:
      // https://docs.nativescript.org/api-reference/interfaces/_ui_core_view_base_.showmodaloptions#animated
      //
      const options = Object.assign({}, originalOptions, { animated: false })
      modalParent.showModal(nativeView, options)
    }

    modalView.closeModal()
  }

  createPage() {
    const {
      instance: { refreshComponent },
    } = this
    const { nativePageElement: oldNativePageElement } = this
    const oldNativeView = oldNativePageElement.nativeView
    // rerender
    const target = document.createElement('fragment')

    // not using conservative for now, since there's nothing in place here to
    // leverage it (yet?) -- and it might be easier to miss breakages in native
    // only code paths
    refreshComponent(target, null)

    // this.nativePageElement is updated in afterMount, triggered by proxy / hooks
    const newPageElement = this.nativePageElement

    // svelte-native uses navigateFrom event + e.isBackNavigation to know when to $destroy the component.
    // To keep that behaviour after refresh, we move event handler from old native view to the new one using
    // __navigateFromHandler property that svelte-native provides us with.
    const navigateFromHandler = oldNativeView.__navigateFromHandler
    if (navigateFromHandler) {
      oldNativeView.off('navigatedFrom', navigateFromHandler)
      newPageElement.nativeView.on('navigatedFrom', navigateFromHandler)
      newPageElement.nativeView.__navigateFromHandler = navigateFromHandler
      delete oldNativeView.__navigateFromHandler
    }

    return newPageElement
  }

  renderError(err /* , target, anchor */) {
    // TODO fallback on TNS error handler for now... at least our error
    // is more informative
    throw err
  }
}
