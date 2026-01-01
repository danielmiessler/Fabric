/// <reference lib="dom" />

declare namespace svelteHTML {

  // Every namespace eligible for use needs to implement the following two functions
  /**
   * @internal do not use
   */
  function mapElementTag<K extends keyof ElementTagNameMap>(
    tag: K
  ): ElementTagNameMap[K];
  function mapElementTag<K extends keyof SVGElementTagNameMap>(
    tag: K
  ): SVGElementTagNameMap[K];
  function mapElementTag(
    tag: any
  ): any; // needs to be any because used in context of <svelte:element>

  /**
   * @internal do not use
   */
  function createElement<Elements extends IntrinsicElements, Key extends keyof Elements>(
    // "undefined | null" because of <svelte:element>
    element: Key | undefined | null, attrs: string extends Key ? import('svelte/elements').HTMLAttributes<any> : Elements[Key]
  ): Key extends keyof ElementTagNameMap ? ElementTagNameMap[Key] : Key extends keyof SVGElementTagNameMap ? SVGElementTagNameMap[Key] : any;
  function createElement<Elements extends IntrinsicElements, Key extends keyof Elements, T>(
    // "undefined | null" because of <svelte:element>
    element: Key | undefined | null, attrsEnhancers: T, attrs: (string extends Key ? import('svelte/elements').HTMLAttributes<any> : Elements[Key]) & T
  ): Key extends keyof ElementTagNameMap ? ElementTagNameMap[Key] : Key extends keyof SVGElementTagNameMap ? SVGElementTagNameMap[Key] : any;

  // For backwards-compatibility and ease-of-use, in case someone enhanced the typings from import('svelte/elements').HTMLAttributes/SVGAttributes
  interface HTMLAttributes<T extends EventTarget = any> {}
  interface SVGAttributes<T extends EventTarget = any> {}

  /**
   * @internal do not use
   */
  type EventsWithColon<T> = {[Property in keyof T as Property extends `on${infer Key}` ? `on:${Key}` : Property]?: T[Property] }
  /**
   * @internal do not use
   */
  type HTMLProps<Property extends string, Override> =
    // This omit chain ensures that properties manually defined in the new transformation take precedence
    // over those manually defined in the old, taking precendence over the defaults, to make sth like this possible
    // https://github.com/sveltejs/language-tools/issues/1352#issuecomment-1248627516
    // The AttributeNames Omit is necessary because the old transformation only has HTMLAttributes on which types for all
    // elements are defined, which would silence type errors in the new transformation.
    Omit<
      Omit<import('svelte/elements').SvelteHTMLElements[Property], keyof EventsWithColon<Omit<svelte.JSX.IntrinsicElements[Property & string], svelte.JSX.AttributeNames>>> & EventsWithColon<Omit<svelte.JSX.IntrinsicElements[Property & string], svelte.JSX.AttributeNames>>,
      keyof Override
    > & Override;
  /**
   * @internal do not use
   */
  type RemoveIndex<T> = {
    [ K in keyof T as string extends K ? never : K ] : T[K]
  };

  // the following type construct makes sure that we can use the new typings while maintaining backwards-compatibility in case someone enhanced the old typings
  interface IntrinsicElements extends Omit<RemoveIndex<svelte.JSX.IntrinsicElements>, keyof RemoveIndex<import('svelte/elements').SvelteHTMLElements>> {
    a: HTMLProps<'a', HTMLAttributes>;
    abbr: HTMLProps<'abbr', HTMLAttributes>;
    address: HTMLProps<'address', HTMLAttributes>;
    area: HTMLProps<'area', HTMLAttributes>;
    article: HTMLProps<'article', HTMLAttributes>;
    aside: HTMLProps<'aside', HTMLAttributes>;
    audio: HTMLProps<'audio', HTMLAttributes>;
    b: HTMLProps<'b', HTMLAttributes>;
    base: HTMLProps<'base', HTMLAttributes>;
    bdi: HTMLProps<'bdi', HTMLAttributes>;
    bdo: HTMLProps<'bdo', HTMLAttributes>;
    big: HTMLProps<'big', HTMLAttributes>;
    blockquote: HTMLProps<'blockquote', HTMLAttributes>;
    body: HTMLProps<'body', HTMLAttributes>;
    br: HTMLProps<'br', HTMLAttributes>;
    button: HTMLProps<'button', HTMLAttributes>;
    canvas: HTMLProps<'canvas', HTMLAttributes>;
    caption: HTMLProps<'caption', HTMLAttributes>;
    cite: HTMLProps<'cite', HTMLAttributes>;
    code: HTMLProps<'code', HTMLAttributes>;
    col: HTMLProps<'col', HTMLAttributes>;
    colgroup: HTMLProps<'colgroup', HTMLAttributes>;
    data: HTMLProps<'data', HTMLAttributes>;
    datalist: HTMLProps<'datalist', HTMLAttributes>;
    dd: HTMLProps<'dd', HTMLAttributes>;
    del: HTMLProps<'del', HTMLAttributes>;
    details: HTMLProps<'details', HTMLAttributes>;
    dfn: HTMLProps<'dfn', HTMLAttributes>;
    dialog: HTMLProps<'dialog', HTMLAttributes>;
    div: HTMLProps<'div', HTMLAttributes>;
    dl: HTMLProps<'dl', HTMLAttributes>;
    dt: HTMLProps<'dt', HTMLAttributes>;
    em: HTMLProps<'em', HTMLAttributes>;
    embed: HTMLProps<'embed', HTMLAttributes>;
    fieldset: HTMLProps<'fieldset', HTMLAttributes>;
    figcaption: HTMLProps<'figcaption', HTMLAttributes>;
    figure: HTMLProps<'figure', HTMLAttributes>;
    footer: HTMLProps<'footer', HTMLAttributes>;
    form: HTMLProps<'form', HTMLAttributes>;
    h1: HTMLProps<'h1', HTMLAttributes>;
    h2: HTMLProps<'h2', HTMLAttributes>;
    h3: HTMLProps<'h3', HTMLAttributes>;
    h4: HTMLProps<'h4', HTMLAttributes>;
    h5: HTMLProps<'h5', HTMLAttributes>;
    h6: HTMLProps<'h6', HTMLAttributes>;
    head: HTMLProps<'head', HTMLAttributes>;
    header: HTMLProps<'header', HTMLAttributes>;
    hgroup: HTMLProps<'hgroup', HTMLAttributes>;
    hr: HTMLProps<'hr', HTMLAttributes>;
    html: HTMLProps<'html', HTMLAttributes>;
    i: HTMLProps<'i', HTMLAttributes>;
    iframe: HTMLProps<'iframe', HTMLAttributes>;
    img: HTMLProps<'img', HTMLAttributes>;
    input: HTMLProps<'input', HTMLAttributes>;
    ins: HTMLProps<'ins', HTMLAttributes>;
    kbd: HTMLProps<'kbd', HTMLAttributes>;
    keygen: HTMLProps<'keygen', HTMLAttributes>;
    label: HTMLProps<'label', HTMLAttributes>;
    legend: HTMLProps<'legend', HTMLAttributes>;
    li: HTMLProps<'li', HTMLAttributes>;
    link: HTMLProps<'link', HTMLAttributes>;
    main: HTMLProps<'main', HTMLAttributes>;
    map: HTMLProps<'map', HTMLAttributes>;
    mark: HTMLProps<'mark', HTMLAttributes>;
    menu: HTMLProps<'menu', HTMLAttributes>;
    menuitem: HTMLProps<'menuitem', HTMLAttributes>;
    meta: HTMLProps<'meta', HTMLAttributes>;
    meter: HTMLProps<'meter', HTMLAttributes>;
    nav: HTMLProps<'nav', HTMLAttributes>;
    noscript: HTMLProps<'noscript', HTMLAttributes>;
    object: HTMLProps<'object', HTMLAttributes>;
    ol: HTMLProps<'ol', HTMLAttributes>;
    optgroup: HTMLProps<'optgroup', HTMLAttributes>;
    option: HTMLProps<'option', HTMLAttributes>;
    output: HTMLProps<'output', HTMLAttributes>;
    p: HTMLProps<'p', HTMLAttributes>;
    param: HTMLProps<'param', HTMLAttributes>;
    picture: HTMLProps<'picture', HTMLAttributes>;
    pre: HTMLProps<'pre', HTMLAttributes>;
    progress: HTMLProps<'progress', HTMLAttributes>;
    q: HTMLProps<'q', HTMLAttributes>;
    rp: HTMLProps<'rp', HTMLAttributes>;
    rt: HTMLProps<'rt', HTMLAttributes>;
    ruby: HTMLProps<'ruby', HTMLAttributes>;
    s: HTMLProps<'s', HTMLAttributes>;
    samp: HTMLProps<'samp', HTMLAttributes>;
    slot: HTMLProps<'slot', HTMLAttributes>;
    script: HTMLProps<'script', HTMLAttributes>;
    section: HTMLProps<'section', HTMLAttributes>;
    select: HTMLProps<'select', HTMLAttributes>;
    small: HTMLProps<'small', HTMLAttributes>;
    source: HTMLProps<'source', HTMLAttributes>;
    span: HTMLProps<'span', HTMLAttributes>;
    strong: HTMLProps<'strong', HTMLAttributes>;
    style: HTMLProps<'style', HTMLAttributes>;
    sub: HTMLProps<'sub', HTMLAttributes>;
    summary: HTMLProps<'summary', HTMLAttributes>;
    sup: HTMLProps<'sup', HTMLAttributes>;
    table: HTMLProps<'table', HTMLAttributes>;
    template: HTMLProps<'template', HTMLAttributes>;
    tbody: HTMLProps<'tbody', HTMLAttributes>;
    td: HTMLProps<'td', HTMLAttributes>;
    textarea: HTMLProps<'textarea', HTMLAttributes>;
    tfoot: HTMLProps<'tfoot', HTMLAttributes>;
    th: HTMLProps<'th', HTMLAttributes>;
    thead: HTMLProps<'thead', HTMLAttributes>;
    time: HTMLProps<'time', HTMLAttributes>;
    title: HTMLProps<'title', HTMLAttributes>;
    tr: HTMLProps<'tr', HTMLAttributes>;
    track: HTMLProps<'track', HTMLAttributes>;
    u: HTMLProps<'u', HTMLAttributes>;
    ul: HTMLProps<'ul', HTMLAttributes>;
    var: HTMLProps<'var', HTMLAttributes>;
    video: HTMLProps<'video', HTMLAttributes>;
    wbr: HTMLProps<'wbr', HTMLAttributes>;
    webview: HTMLProps<'webview', HTMLAttributes>;
    // SVG
    svg: HTMLProps<'svg', SVGAttributes>;

    animate: HTMLProps<'animate', SVGAttributes>;
    animateMotion: HTMLProps<'animateMotion', SVGAttributes>;
    animateTransform: HTMLProps<'animateTransform', SVGAttributes>;
    circle: HTMLProps<'circle', SVGAttributes>;
    clipPath: HTMLProps<'clipPath', SVGAttributes>;
    defs: HTMLProps<'defs', SVGAttributes>;
    desc: HTMLProps<'desc', SVGAttributes>;
    ellipse: HTMLProps<'ellipse', SVGAttributes>;
    feBlend: HTMLProps<'feBlend', SVGAttributes>;
    feColorMatrix: HTMLProps<'feColorMatrix', SVGAttributes>;
    feComponentTransfer: HTMLProps<'feComponentTransfer', SVGAttributes>;
    feComposite: HTMLProps<'feComposite', SVGAttributes>;
    feConvolveMatrix: HTMLProps<'feConvolveMatrix', SVGAttributes>;
    feDiffuseLighting: HTMLProps<'feDiffuseLighting', SVGAttributes>;
    feDisplacementMap: HTMLProps<'feDisplacementMap', SVGAttributes>;
    feDistantLight: HTMLProps<'feDistantLight', SVGAttributes>;
    feDropShadow: HTMLProps<'feDropShadow', SVGAttributes>;
    feFlood: HTMLProps<'feFlood', SVGAttributes>;
    feFuncA: HTMLProps<'feFuncA', SVGAttributes>;
    feFuncB: HTMLProps<'feFuncB', SVGAttributes>;
    feFuncG: HTMLProps<'feFuncG', SVGAttributes>;
    feFuncR: HTMLProps<'feFuncR', SVGAttributes>;
    feGaussianBlur: HTMLProps<'feGaussianBlur', SVGAttributes>;
    feImage: HTMLProps<'feImage', SVGAttributes>;
    feMerge: HTMLProps<'feMerge', SVGAttributes>;
    feMergeNode: HTMLProps<'feMergeNode', SVGAttributes>;
    feMorphology: HTMLProps<'feMorphology', SVGAttributes>;
    feOffset: HTMLProps<'feOffset', SVGAttributes>;
    fePointLight: HTMLProps<'fePointLight', SVGAttributes>;
    feSpecularLighting: HTMLProps<'feSpecularLighting', SVGAttributes>;
    feSpotLight: HTMLProps<'feSpotLight', SVGAttributes>;
    feTile: HTMLProps<'feTile', SVGAttributes>;
    feTurbulence: HTMLProps<'feTurbulence', SVGAttributes>;
    filter: HTMLProps<'filter', SVGAttributes>;
    foreignObject: HTMLProps<'foreignObject', SVGAttributes>;
    g: HTMLProps<'g', SVGAttributes>;
    image: HTMLProps<'image', SVGAttributes>;
    line: HTMLProps<'line', SVGAttributes>;
    linearGradient: HTMLProps<'linearGradient', SVGAttributes>;
    marker: HTMLProps<'marker', SVGAttributes>;
    mask: HTMLProps<'mask', SVGAttributes>;
    metadata: HTMLProps<'metadata', SVGAttributes>;
    mpath: HTMLProps<'mpath', SVGAttributes>;
    path: HTMLProps<'path', SVGAttributes>;
    pattern: HTMLProps<'pattern', SVGAttributes>;
    polygon: HTMLProps<'polygon', SVGAttributes>;
    polyline: HTMLProps<'polyline', SVGAttributes>;
    radialGradient: HTMLProps<'radialGradient', SVGAttributes>;
    rect: HTMLProps<'rect', SVGAttributes>;
    stop: HTMLProps<'stop', SVGAttributes>;
    switch: HTMLProps<'switch', SVGAttributes>;
    symbol: HTMLProps<'symbol', SVGAttributes>;
    text: HTMLProps<'text', SVGAttributes>;
    textPath: HTMLProps<'textPath', SVGAttributes>;
    tspan: HTMLProps<'tspan', SVGAttributes>;
    use: HTMLProps<'use', SVGAttributes>;
    view: HTMLProps<'view', SVGAttributes>;

    // Svelte specific
    'svelte:window': HTMLProps<'svelte:window', HTMLAttributes>;
    'svelte:body': HTMLProps<'svelte:body', HTMLAttributes>;
    'svelte:document': HTMLProps<'svelte:document', HTMLAttributes>;
    'svelte:fragment': { slot?: string };
    'svelte:options': { [name: string]: any };
    'svelte:head': { [name: string]: any };

    [name: string]: { [name: string]: any };
  }

}

// Keep svelte.JSX for backwards compatibility, in case someone enhanced it with their own typings,
// which we can transform to the new svelteHTML namespace.
/**
 * @deprecated use the types from `svelte/elements` instead, or the `svelteHTML` namespace.
 * For more info see https://github.com/sveltejs/language-tools/blob/master/docs/preprocessors/typescript.md#im-getting-deprecation-warnings-for-sveltejsx--i-want-to-migrate-to-the-new-typings
 */
declare namespace svelte.JSX {

  /* svelte specific */
  interface ElementClass {
      $$prop_def: any;
  }

  interface ElementAttributesProperty {
      $$prop_def: any; // specify the property name to use
  }

  /* html jsx */

  interface IntrinsicAttributes {
    slot?: string;
  }

  //
  // Event Handler Types
  // ----------------------------------------------------------------------
  type EventHandler<E extends Event = Event, T extends EventTarget = HTMLElement> =
    (event: E & { currentTarget: EventTarget & T}) => any;

  type ClipboardEventHandler<T extends EventTarget> = EventHandler<ClipboardEvent, T>;
  type CompositionEventHandler<T extends EventTarget> = EventHandler<CompositionEvent, T>;
  type DragEventHandler<T extends EventTarget> = EventHandler<DragEvent, T>;
  type FocusEventHandler<T extends EventTarget> = EventHandler<FocusEvent, T>;
  type FormEventHandler<T extends EventTarget> = EventHandler<Event, T>;
  type ChangeEventHandler<T extends EventTarget> = EventHandler<Event, T>;
  type KeyboardEventHandler<T extends EventTarget> = EventHandler<KeyboardEvent, T>;
  type MouseEventHandler<T extends EventTarget> = EventHandler<MouseEvent, T>;
  type TouchEventHandler<T extends EventTarget> = EventHandler<TouchEvent, T>;
  type PointerEventHandler<T extends EventTarget> = EventHandler<PointerEvent, T>;
  type UIEventHandler<T extends EventTarget> = EventHandler<UIEvent, T>;
  type WheelEventHandler<T extends EventTarget> = EventHandler<WheelEvent, T>;
  type AnimationEventHandler<T extends EventTarget> = EventHandler<AnimationEvent, T>;
  type TransitionEventHandler<T extends EventTarget> = EventHandler<TransitionEvent, T>;
  type MessageEventHandler<T extends EventTarget> = EventHandler<MessageEvent, T>;

  /** @deprecated DO NOT USE, WILL BE REMOVED SOON */
  type AttributeNames =
  |'oncopy'
  |'oncut'
  |'onpaste'
  |'oncompositionend'
  |'oncompositionstart'
  |'oncompositionupdate'
  |'onfocus'
  |'onfocusin'
  |'onfocusout'
  |'onblur'
  |'onchange'
  |'oninput'
  |'onreset'
  |'onsubmit'
  |'oninvalid'
  |'onbeforeinput'
  |'onload'
  |'onerror'
  |'ontoggle'
  |'onkeydown'
  |'onkeypress'
  |'onkeyup'
  |'onabort'
  |'oncanplay'
  |'oncanplaythrough'
  |'oncuechange'
  |'ondurationchange'
  |'onemptied'
  |'onencrypted'
  |'onended'
  |'onloadeddata'
  |'onloadedmetadata'
  |'onloadstart'
  |'onpause'
  |'onplay'
  |'onplaying'
  |'onprogress'
  |'onratechange'
  |'onseeked'
  |'onseeking'
  |'onstalled'
  |'onsuspend'
  |'ontimeupdate'
  |'onvolumechange'
  |'onwaiting'
  |'onauxclick'
  |'onclick'
  |'oncontextmenu'
  |'ondblclick'
  |'ondrag'
  |'ondragend'
  |'ondragenter'
  |'ondragexit'
  |'ondragleave'
  |'ondragover'
  |'ondragstart'
  |'ondrop'
  |'onmousedown'
  |'onmouseenter'
  |'onmouseleave'
  |'onmousemove'
  |'onmouseout'
  |'onmouseover'
  |'onmouseup'
  |'onselect'
  |'onselectionchange'
  |'onselectstart'
  |'ontouchcancel'
  |'ontouchend'
  |'ontouchmove'
  |'ontouchstart'
  |'ongotpointercapture'
  |'onpointercancel'
  |'onpointerdown'
  |'onpointerenter'
  |'onpointerleave'
  |'onpointermove'
  |'onpointerout'
  |'onpointerover'
  |'onpointerup'
  |'onlostpointercapture'
  |'onscroll'
  |'onresize'
  |'onwheel'
  |'onanimationstart'
  |'onanimationend'
  |'onanimationiteration'
  |'ontransitionstart'
  |'ontransitionrun'
  |'ontransitionend'
  |'ontransitioncancel'
  |'onoutrostart'
  |'onoutroend'
  |'onintrostart'
  |'onintroend'
  |'onmessage'
  |'onmessageerror'
  |'oncancel'
  |'onclose'
  |'onfullscreenchange'
  |'onfullscreenerror'
  |'class'
  |'dataset'
  |'accept'
  |'acceptcharset'
  |'accesskey'
  |'action'
  |'allow'
  |'allowfullscreen'
  |'allowtransparency'
  |'allowpaymentrequest'
  |'alt'
  |'as'
  |'async'
  |'autocomplete'
  |'autofocus'
  |'autoplay'
  |'capture'
  |'cellpadding'
  |'cellspacing'
  |'charset'
  |'challenge'
  |'checked'
  |'cite'
  |'classid'
  |'cols'
  |'colspan'
  |'content'
  |'contenteditable'
  |'innerHTML'
  |'textContent'
  |'contextmenu'
  |'controls'
  |'coords'
  |'crossorigin'
  |'currenttime'
  |'decoding'
  |'data'
  |'datetime'
  |'default'
  |'defaultmuted'
  |'defaultplaybackrate'
  |'defer'
  |'dir'
  |'dirname'
  |'disabled'
  |'download'
  |'draggable'
  |'enctype'
  |'enterkeyhint'
  |'for'
  |'form'
  |'formaction'
  |'formenctype'
  |'formmethod'
  |'formnovalidate'
  |'formtarget'
  |'frameborder'
  |'headers'
  |'height'
  |'hidden'
  |'high'
  |'href'
  |'hreflang'
  |'htmlfor'
  |'httpequiv'
  |'id'
  |'inputmode'
  |'integrity'
  |'is'
  |'ismap'
  |'keyparams'
  |'keytype'
  |'kind'
  |'label'
  |'lang'
  |'list'
  |'loading'
  |'loop'
  |'low'
  |'manifest'
  |'marginheight'
  |'marginwidth'
  |'max'
  |'maxlength'
  |'media'
  |'mediagroup'
  |'method'
  |'min'
  |'minlength'
  |'multiple'
  |'muted'
  |'name'
  |'nonce'
  |'novalidate'
  |'open'
  |'optimum'
  |'part'
  |'pattern'
  |'placeholder'
  |'playsinline'
  |'ping'
  |'poster'
  |'preload'
  |'radiogroup'
  |'readonly'
  |'referrerpolicy'
  |'rel'
  |'required'
  |'reversed'
  |'role'
  |'rows'
  |'rowspan'
  |'sandbox'
  |'scope'
  |'scoped'
  |'scrolling'
  |'seamless'
  |'selected'
  |'shape'
  |'size'
  |'sizes'
  |'slot'
  |'span'
  |'spellcheck'
  |'src'
  |'srcdoc'
  |'srclang'
  |'srcset'
  |'start'
  |'step'
  |'style'
  |'summary'
  |'tabindex'
  |'target'
  |'title'
  |'translate'
  |'type'
  |'usemap'
  |'value'
  |'volume'
  |'width'
  |'wmode'
  |'wrap'
  |'about'
  |'datatype'
  |'inlist'
  |'prefix'
  |'property'
  |'resource'
  |'typeof'
  |'vocab'
  |'autocapitalize'
  |'autocorrect'
  |'autosave'
  |'color'
  |'controlslist'
  |'inert'
  |'itemprop'
  |'itemscope'
  |'itemtype'
  |'itemid'
  |'itemref'
  |'results'
  |'security'
  |'unselectable';

  /**
   * @deprecated use the types from `svelte/elements` instead
   * For more info see https://github.com/sveltejs/language-tools/blob/master/docs/preprocessors/typescript.md#im-getting-deprecation-warnings-for-sveltejsx--i-want-to-migrate-to-the-new-typings
   */
  interface DOMAttributes<T extends EventTarget> {
    oncopy?: ClipboardEventHandler<T> | undefined | null;
    oncut?: ClipboardEventHandler<T> | undefined | null;
    onpaste?: ClipboardEventHandler<T> | undefined | null;
    oncompositionend?: CompositionEventHandler<T> | undefined | null;
    oncompositionstart?: CompositionEventHandler<T> | undefined | null;
    oncompositionupdate?: CompositionEventHandler<T> | undefined | null;
    onfocus?: FocusEventHandler<T> | undefined | null;
    onfocusin?: FocusEventHandler<T> | undefined | null;
    onfocusout?: FocusEventHandler<T> | undefined | null;
    onblur?: FocusEventHandler<T> | undefined | null;
    onchange?: FormEventHandler<T> | undefined | null;
    oninput?: FormEventHandler<T> | undefined | null;
    onreset?: FormEventHandler<T> | undefined | null;
    onsubmit?: EventHandler<SubmitEvent, T> | undefined | null;
    oninvalid?: EventHandler<Event, T> | undefined | null;
    onbeforeinput?: EventHandler<InputEvent, T> | undefined | null;
    onload?: EventHandler | undefined | null;
    onerror?: EventHandler | undefined | null; // also a Media Event
    ontoggle?: EventHandler<Event, T> | undefined | null;
    onkeydown?: KeyboardEventHandler<T> | undefined | null;
    onkeypress?: KeyboardEventHandler<T> | undefined | null;
    onkeyup?: KeyboardEventHandler<T> | undefined | null;
    onabort?: EventHandler<Event, T> | undefined | null;
    oncanplay?: EventHandler<Event, T> | undefined | null;
    oncanplaythrough?: EventHandler<Event, T> | undefined | null;
    oncuechange?: EventHandler<Event, T> | undefined | null;
    ondurationchange?: EventHandler<Event, T> | undefined | null;
    onemptied?: EventHandler<Event, T> | undefined | null;
    onencrypted?: EventHandler<Event, T> | undefined | null;
    onended?: EventHandler<Event, T> | undefined | null;
    onloadeddata?: EventHandler<Event, T> | undefined | null;
    onloadedmetadata?: EventHandler<Event, T> | undefined | null;
    onloadstart?: EventHandler<Event, T> | undefined | null;
    onpause?: EventHandler<Event, T> | undefined | null;
    onplay?: EventHandler<Event, T> | undefined | null;
    onplaying?: EventHandler<Event, T> | undefined | null;
    onprogress?: EventHandler<Event, T> | undefined | null;
    onratechange?: EventHandler<Event, T> | undefined | null;
    onseeked?: EventHandler<Event, T> | undefined | null;
    onseeking?: EventHandler<Event, T> | undefined | null;
    onstalled?: EventHandler<Event, T> | undefined | null;
    onsuspend?: EventHandler<Event, T> | undefined | null;
    ontimeupdate?: EventHandler<Event, T> | undefined | null;
    onvolumechange?: EventHandler<Event, T> | undefined | null;
    onwaiting?: EventHandler<Event, T> | undefined | null;
    onauxclick?: MouseEventHandler<T> | undefined | null;
    onclick?: MouseEventHandler<T> | undefined | null;
    oncontextmenu?: MouseEventHandler<T> | undefined | null;
    ondblclick?: MouseEventHandler<T> | undefined | null;
    ondrag?: DragEventHandler<T> | undefined | null;
    ondragend?: DragEventHandler<T> | undefined | null;
    ondragenter?: DragEventHandler<T> | undefined | null;
    ondragexit?: DragEventHandler<T> | undefined | null;
    ondragleave?: DragEventHandler<T> | undefined | null;
    ondragover?: DragEventHandler<T> | undefined | null;
    ondragstart?: DragEventHandler<T> | undefined | null;
    ondrop?: DragEventHandler<T> | undefined | null;
    onmousedown?: MouseEventHandler<T> | undefined | null;
    onmouseenter?: MouseEventHandler<T> | undefined | null;
    onmouseleave?: MouseEventHandler<T> | undefined | null;
    onmousemove?: MouseEventHandler<T> | undefined | null;
    onmouseout?: MouseEventHandler<T> | undefined | null;
    onmouseover?: MouseEventHandler<T> | undefined | null;
    onmouseup?: MouseEventHandler<T> | undefined | null;
    onselect?: EventHandler<Event, T> | undefined | null;
    onselectionchange?: EventHandler<Event, T> | undefined | null;
    onselectstart?: EventHandler<Event, T> | undefined | null;
    ontouchcancel?: TouchEventHandler<T> | undefined | null;
    ontouchend?: TouchEventHandler<T> | undefined | null;
    ontouchmove?: TouchEventHandler<T> | undefined | null;
    ontouchstart?: TouchEventHandler<T> | undefined | null;
    ongotpointercapture?: PointerEventHandler<T> | undefined | null;
    onpointercancel?: PointerEventHandler<T> | undefined | null;
    onpointerdown?: PointerEventHandler<T> | undefined | null;
    onpointerenter?: PointerEventHandler<T> | undefined | null;
    onpointerleave?: PointerEventHandler<T> | undefined | null;
    onpointermove?: PointerEventHandler<T> | undefined | null;
    onpointerout?: PointerEventHandler<T> | undefined | null;
    onpointerover?: PointerEventHandler<T> | undefined | null;
    onpointerup?: PointerEventHandler<T> | undefined | null;
    onlostpointercapture?: PointerEventHandler<T> | undefined | null;
    onscroll?: UIEventHandler<T> | undefined | null;
    onresize?: UIEventHandler<T> | undefined | null;
    onwheel?: WheelEventHandler<T> | undefined | null;
    onanimationstart?: AnimationEventHandler<T> | undefined | null;
    onanimationend?: AnimationEventHandler<T> | undefined | null;
    onanimationiteration?: AnimationEventHandler<T> | undefined | null;
    ontransitionstart?: TransitionEventHandler<T> | undefined | null;
    ontransitionrun?: TransitionEventHandler<T> | undefined | null;
    ontransitionend?: TransitionEventHandler<T> | undefined | null;
    ontransitioncancel?: TransitionEventHandler<T> | undefined | null;
    onoutrostart?: EventHandler<CustomEvent<null>, T> | undefined | null;
    onoutroend?: EventHandler<CustomEvent<null>, T> | undefined | null;
    onintrostart?: EventHandler<CustomEvent<null>, T> | undefined | null;
    onintroend?: EventHandler<CustomEvent<null>, T> | undefined | null;
    onmessage?: MessageEventHandler<T> | undefined | null;
    onmessageerror?: MessageEventHandler<T> | undefined | null;
    oncancel?: EventHandler<Event, T> | undefined | null;
    onclose?: EventHandler<Event, T> | undefined | null;
    onfullscreenchange?: EventHandler<Event, T> | undefined | null;
    onfullscreenerror?: EventHandler<Event, T> | undefined | null;
  }

  /**
   * @deprecated use the types from `svelte/elements` instead
   * For more info see https://github.com/sveltejs/language-tools/blob/master/docs/preprocessors/typescript.md#im-getting-deprecation-warnings-for-sveltejsx--i-want-to-migrate-to-the-new-typings
   */
  interface AriaAttributes {
      'aria-activedescendant'?: string | undefined | null;
      'aria-atomic'?: boolean | 'false' | 'true' | undefined | null;
      'aria-autocomplete'?: 'none' | 'inline' | 'list' | 'both' | undefined | null;
      'aria-busy'?: boolean | 'false' | 'true' | undefined | null;
      'aria-checked'?: boolean | 'false' | 'mixed' | 'true' | undefined | null;
      'aria-colcount'?: number | undefined | null;
      'aria-colindex'?: number | undefined | null;
      'aria-colspan'?: number | undefined | null;
      'aria-controls'?: string | undefined | null;
      'aria-current'?: boolean | 'false' | 'true' | 'page' | 'step' | 'location' | 'date' | 'time' | undefined | null;
      'aria-describedby'?: string | undefined | null;
      'aria-details'?: string | undefined | null;
      'aria-disabled'?: boolean | 'false' | 'true' | undefined | null;
      'aria-dropeffect'?: 'none' | 'copy' | 'execute' | 'link' | 'move' | 'popup' | undefined | null;
      'aria-errormessage'?: string | undefined | null;
      'aria-expanded'?: boolean | 'false' | 'true' | undefined | null;
      'aria-flowto'?: string | undefined | null;
      'aria-grabbed'?: boolean | 'false' | 'true' | undefined | null;
      'aria-haspopup'?: boolean | 'false' | 'true' | 'menu' | 'listbox' | 'tree' | 'grid' | 'dialog' | undefined | null;
      'aria-hidden'?: boolean | 'false' | 'true' | undefined | null;
      'aria-invalid'?: boolean | 'false' | 'true' | 'grammar' | 'spelling' | undefined | null;
      'aria-keyshortcuts'?: string | undefined | null;
      'aria-label'?: string | undefined | null;
      'aria-labelledby'?: string | undefined | null;
      'aria-level'?: number | undefined | null;
      'aria-live'?: 'off' | 'assertive' | 'polite' | undefined | null;
      'aria-modal'?: boolean | 'false' | 'true' | undefined | null;
      'aria-multiline'?: boolean | 'false' | 'true' | undefined | null;
      'aria-multiselectable'?: boolean | 'false' | 'true' | undefined | null;
      'aria-orientation'?: 'horizontal' | 'vertical' | undefined | null;
      'aria-owns'?: string | undefined | null;
      'aria-placeholder'?: string | undefined | null;
      'aria-posinset'?: number | undefined | null;
      'aria-pressed'?: boolean | 'false' | 'mixed' | 'true' | undefined | null;
      'aria-readonly'?: boolean | 'false' | 'true' | undefined | null;
      'aria-relevant'?: 'additions' | 'additions removals' | 'additions text' | 'all' | 'removals' | 'removals additions' | 'removals text' | 'text' | 'text additions' | 'text removals' | undefined | null;
      'aria-required'?: boolean | 'false' | 'true' | undefined | null;
      'aria-roledescription'?: string | undefined | null;
      'aria-rowcount'?: number | undefined | null;
      'aria-rowindex'?: number | undefined | null;
      'aria-rowspan'?: number | undefined | null;
      'aria-selected'?: boolean | 'false' | 'true' | undefined | null;
      'aria-setsize'?: number | undefined | null;
      'aria-sort'?: 'none' | 'ascending' | 'descending' | 'other' | undefined | null;
      'aria-valuemax'?: number | undefined | null;
      'aria-valuemin'?: number | undefined | null;
      'aria-valuenow'?: number | undefined | null;
      'aria-valuetext'?: string | undefined | null;
  }

  /**
   * @deprecated use the types from `svelte/elements` instead
   * For more info see https://github.com/sveltejs/language-tools/blob/master/docs/preprocessors/typescript.md#im-getting-deprecation-warnings-for-sveltejsx--i-want-to-migrate-to-the-new-typings
   */
  interface HTMLAttributes<T extends EventTarget> extends AriaAttributes, DOMAttributes<T> {
    class?: string | undefined | null;
    dataset?: object | undefined | null;
    accept?: string | undefined | null;
    acceptcharset?: string | undefined | null;
    accesskey?: string | undefined | null;
    action?: string | undefined | null;
    allow?: string | undefined | null;
    allowfullscreen?: boolean | undefined | null;
    allowtransparency?: boolean | undefined | null;
    allowpaymentrequest?: boolean | undefined | null;
    alt?: string | undefined | null;
    as?: string | undefined | null;
    async?: boolean | undefined | null;
    autocomplete?: string | undefined | null;
    autofocus?: boolean | undefined | null;
    autoplay?: boolean | undefined | null;
    capture?: 'environment' | 'user' | boolean | undefined | null;
    cellpadding?: number | string | undefined | null;
    cellspacing?: number | string | undefined | null;
    charset?: string | undefined | null;
    challenge?: string | undefined | null;
    checked?: boolean | undefined | null;
    cite?: string | undefined | null;
    classid?: string | undefined | null;
    cols?: number | undefined | null;
    colspan?: number | undefined | null;
    content?: string | undefined | null;
    contenteditable?: 'true' | 'false' | boolean | undefined | null;
    innerHTML?: string | undefined | null;
    textContent?: string | undefined | null;
    contextmenu?: string | undefined | null;
    controls?: boolean | undefined | null;
    coords?: string | undefined | null;
    crossorigin?: string | undefined | null;
    currenttime?: number | undefined | null;
    decoding?: 'async' | 'sync' | 'auto' | undefined | null;
    data?: string | undefined | null;
    datetime?: string | undefined | null;
    default?: boolean | undefined | null;
    defaultmuted?: boolean | undefined | null;
    defaultplaybackrate?: number | undefined | null;
    defer?: boolean | undefined | null;
    dir?: string | undefined | null;
    dirname?: string | undefined | null;
    disabled?: boolean | undefined | null;
    download?: any | undefined | null;
    draggable?: boolean | 'true' | 'false' | undefined | null;
    enctype?: string | undefined | null;
    enterkeyhint?: 'enter' | 'done' | 'go' | 'next' | 'previous' | 'search' | 'send' | undefined | null;
    for?: string | undefined | null;
    form?: string | undefined | null;
    formaction?: string | undefined | null;
    formenctype?: string | undefined | null;
    formmethod?: string | undefined | null;
    formnovalidate?: boolean | undefined | null;
    formtarget?: string | undefined | null;
    frameborder?: number | string | undefined | null;
    headers?: string | undefined | null;
    height?: number | string | undefined | null;
    hidden?: boolean | undefined | null;
    high?: number | undefined | null;
    href?: string | undefined | null;
    hreflang?: string | undefined | null;
    htmlfor?: string | undefined | null;
    httpequiv?: string | undefined | null;
    id?: string | undefined | null;
    inputmode?: string | undefined | null;
    integrity?: string | undefined | null;
    is?: string | undefined | null;
    ismap?: boolean | undefined | null;
    keyparams?: string | undefined | null;
    keytype?: string | undefined | null;
    kind?: string | undefined | null;
    label?: string | undefined | null;
    lang?: string | undefined | null;
    list?: string | undefined | null;
    loading?: string | undefined | null;
    loop?: boolean | undefined | null;
    low?: number | undefined | null;
    manifest?: string | undefined | null;
    marginheight?: number | undefined | null;
    marginwidth?: number | undefined | null;
    max?: number | string | undefined | null;
    maxlength?: number | undefined | null;
    media?: string | undefined | null;
    mediagroup?: string | undefined | null;
    method?: string | undefined | null;
    min?: number | string | undefined | null;
    minlength?: number | undefined | null;
    multiple?: boolean | undefined | null;
    muted?: boolean | undefined | null;
    name?: string | undefined | null;
    nonce?: string | undefined | null;
    novalidate?: boolean | undefined | null;
    open?: boolean | undefined | null;
    optimum?: number | undefined | null;
    part?: string | undefined | null;
    pattern?: string | undefined | null;
    placeholder?: string | undefined | null;
    playsinline?: boolean | undefined | null;
    ping?: string | undefined | null;
    poster?: string | undefined | null;
    preload?: string | undefined | null;
    radiogroup?: string | undefined | null;
    readonly?: boolean | undefined | null;
    referrerpolicy?: string | undefined | null;
    rel?: string | undefined | null;
    required?: boolean | undefined | null;
    reversed?: boolean | undefined | null;
    role?: string | undefined | null;
    rows?: number | undefined | null;
    rowspan?: number | undefined | null;
    sandbox?: string | undefined | null;
    scope?: string | undefined | null;
    scoped?: boolean | undefined | null;
    scrolling?: string | undefined | null;
    seamless?: boolean | undefined | null;
    selected?: boolean | undefined | null;
    shape?: string | undefined | null;
    size?: number | undefined | null;
    sizes?: string | undefined | null;
    slot?: string | undefined | null;
    span?: number | undefined | null;
    spellcheck?: boolean | 'true' | 'false' | undefined | null;
    src?: string | undefined | null;
    srcdoc?: string | undefined | null;
    srclang?: string | undefined | null;
    srcset?: string | undefined | null;
    start?: number | undefined | null;
    step?: number | string | undefined | null;
    style?: string | undefined | null;
    summary?: string | undefined | null;
    tabindex?: number | undefined | null;
    target?: string | undefined | null;
    title?: string | undefined | null;
    translate?: "yes" | "no" | "" | undefined | null;
    type?: string | undefined | null;
    usemap?: string | undefined | null;
    value?: any | undefined | null;
    volume?: number | undefined | null;
    width?: number | string | undefined | null;
    wmode?: string | undefined | null;
    wrap?: string | undefined | null;
    about?: string | undefined | null;
    datatype?: string | undefined | null;
    inlist?: any | undefined | null;
    prefix?: string | undefined | null;
    property?: string | undefined | null;
    resource?: string | undefined | null;
    typeof?: string | undefined | null;
    vocab?: string | undefined | null;
    autocapitalize?: string | undefined | null;
    autocorrect?: string | undefined | null;
    autosave?: string | undefined | null;
    color?: string | undefined | null;
    controlslist?: 'nodownload' | 'nofullscreen' | 'noplaybackrate' | 'noremoteplayback';
    inert?: boolean | undefined | null;
    itemprop?: string | undefined | null;
    itemscope?: boolean | undefined | null;
    itemtype?: string | undefined | null;
    itemid?: string | undefined | null;
    itemref?: string | undefined | null;
    results?: number | undefined | null;
    security?: string | undefined | null;
    unselectable?: boolean | undefined | null;

    'data-sveltekit-keepfocus'?: true | '' | 'off' | undefined | null;
    'data-sveltekit-noscroll'?: true | '' | 'off' | undefined | null;
    'data-sveltekit-preload-code'?: true | '' | 'eager' | 'viewport' | 'hover' | 'tap' | 'off' | undefined | null;
    'data-sveltekit-preload-data'?: true | '' | 'hover' | 'tap' | 'off' | undefined | null;
    'data-sveltekit-reload'?: true | '' | 'off' | undefined | null;
    'data-sveltekit-replacestate'?: true | '' | 'off' | undefined | null;
  }

  /**
   * @deprecated use the types from `svelte/elements` instead
   * For more info see https://github.com/sveltejs/language-tools/blob/master/docs/preprocessors/typescript.md#im-getting-deprecation-warnings-for-sveltejsx--i-want-to-migrate-to-the-new-typings
   */
  interface SVGAttributes<T extends EventTarget> extends AriaAttributes, DOMAttributes<T> {
    className?: string | undefined | null;
    class?: string | undefined | null;
    color?: string | undefined | null;
    height?: number | string | undefined | null;
    id?: string | undefined | null;
    lang?: string | undefined | null;
    max?: number | string | undefined | null;
    media?: string | undefined | null;
    method?: string | undefined | null;
    min?: number | string | undefined | null;
    name?: string | undefined | null;
    style?: string | undefined | null;
    target?: string | undefined | null;
    type?: string | undefined | null;
    width?: number | string | undefined | null;
    role?: string | undefined | null;
    tabindex?: number | undefined | null;
    crossorigin?: 'anonymous' | 'use-credentials' | '' | undefined | null;
    'accent-height'?: number | string | undefined | null;
    accumulate?: 'none' | 'sum' | undefined | null;
    additive?: 'replace' | 'sum' | undefined | null;
    'alignment-baseline'?: 'auto' | 'baseline' | 'before-edge' | 'text-before-edge' | 'middle' |
      'central' | 'after-edge' | 'text-after-edge' | 'ideographic' | 'alphabetic' | 'hanging' |
      'mathematical' | 'inherit' | undefined | null;
    allowReorder?: 'no' | 'yes' | undefined | null;
    alphabetic?: number | string | undefined | null;
    amplitude?: number | string | undefined | null;
    'arabic-form'?: 'initial' | 'medial' | 'terminal' | 'isolated' | undefined | null;
    ascent?: number | string | undefined | null;
    attributeName?: string | undefined | null;
    attributeType?: string | undefined | null;
    autoReverse?: number | string | undefined | null;
    azimuth?: number | string | undefined | null;
    baseFrequency?: number | string | undefined | null;
    'baseline-shift'?: number | string | undefined | null;
    baseProfile?: number | string | undefined | null;
    bbox?: number | string | undefined | null;
    begin?: number | string | undefined | null;
    bias?: number | string | undefined | null;
    by?: number | string | undefined | null;
    calcMode?: number | string | undefined | null;
    'cap-height'?: number | string | undefined | null;
    clip?: number | string | undefined | null;
    'clip-path'?: string | undefined | null;
    clipPathUnits?: number | string | undefined | null;
    'clip-rule'?: number | string | undefined | null;
    'color-interpolation'?: number | string | undefined | null;
    'color-interpolation-filters'?: 'auto' | 'sRGB' | 'linearRGB' | 'inherit' | undefined | null;
    'color-profile'?: number | string | undefined | null;
    'color-rendering'?: number | string | undefined | null;
    contentScriptType?: number | string | undefined | null;
    contentStyleType?: number | string | undefined | null;
    cursor?: number | string | undefined | null;
    cx?: number | string | undefined | null;
    cy?: number | string | undefined | null;
    d?: string | undefined | null;
    decelerate?: number | string | undefined | null;
    descent?: number | string | undefined | null;
    diffuseConstant?: number | string | undefined | null;
    direction?: number | string | undefined | null;
    display?: number | string | undefined | null;
    divisor?: number | string | undefined | null;
    'dominant-baseline'?: number | string | undefined | null;
    dur?: number | string | undefined | null;
    dx?: number | string | undefined | null;
    dy?: number | string | undefined | null;
    edgeMode?: number | string | undefined | null;
    elevation?: number | string | undefined | null;
    'enable-background'?: number | string | undefined | null;
    end?: number | string | undefined | null;
    exponent?: number | string | undefined | null;
    externalResourcesRequired?: number | string | undefined | null;
    fill?: string | undefined | null;
    'fill-opacity'?: number | string | undefined | null;
    'fill-rule'?: 'nonzero' | 'evenodd' | 'inherit' | undefined | null;
    filter?: string | undefined | null;
    filterRes?: number | string | undefined | null;
    filterUnits?: number | string | undefined | null;
    'flood-color'?: number | string | undefined | null;
    'flood-opacity'?: number | string | undefined | null;
    focusable?: number | string | undefined | null;
    'font-family'?: string | undefined | null;
    'font-size'?: number | string | undefined | null;
    'font-size-adjust'?: number | string | undefined | null;
    'font-stretch'?: number | string | undefined | null;
    'font-style'?: number | string | undefined | null;
    'font-variant'?: number | string | undefined | null;
    'font-weight'?: number | string | undefined | null;
    format?: number | string | undefined | null;
    from?: number | string | undefined | null;
    fx?: number | string | undefined | null;
    fy?: number | string | undefined | null;
    g1?: number | string | undefined | null;
    g2?: number | string | undefined | null;
    'glyph-name'?: number | string | undefined | null;
    'glyph-orientation-horizontal'?: number | string | undefined | null;
    'glyph-orientation-vertical'?: number | string | undefined | null;
    glyphRef?: number | string | undefined | null;
    gradientTransform?: string | undefined | null;
    gradientUnits?: string | undefined | null;
    hanging?: number | string | undefined | null;
    href?: string | undefined | null;
    'horiz-adv-x'?: number | string | undefined | null;
    'horiz-origin-x'?: number | string | undefined | null;
    ideographic?: number | string | undefined | null;
    'image-rendering'?: number | string | undefined | null;
    in2?: number | string | undefined | null;
    in?: string | undefined | null;
    intercept?: number | string | undefined | null;
    k1?: number | string | undefined | null;
    k2?: number | string | undefined | null;
    k3?: number | string | undefined | null;
    k4?: number | string | undefined | null;
    k?: number | string | undefined | null;
    kernelMatrix?: number | string | undefined | null;
    kernelUnitLength?: number | string | undefined | null;
    kerning?: number | string | undefined | null;
    keyPoints?: number | string | undefined | null;
    keySplines?: number | string | undefined | null;
    keyTimes?: number | string | undefined | null;
    lengthAdjust?: number | string | undefined | null;
    'letter-spacing'?: number | string | undefined | null;
    'lighting-color'?: number | string | undefined | null;
    limitingConeAngle?: number | string | undefined | null;
    local?: number | string | undefined | null;
    'marker-end'?: string | undefined | null;
    markerHeight?: number | string | undefined | null;
    'marker-mid'?: string | undefined | null;
    'marker-start'?: string | undefined | null;
    markerUnits?: number | string | undefined | null;
    markerWidth?: number | string | undefined | null;
    mask?: string | undefined | null;
    maskContentUnits?: number | string | undefined | null;
    maskUnits?: number | string | undefined | null;
    mathematical?: number | string | undefined | null;
    mode?: number | string | undefined | null;
    numOctaves?: number | string | undefined | null;
    offset?: number | string | undefined | null;
    opacity?: number | string | undefined | null;
    operator?: number | string | undefined | null;
    order?: number | string | undefined | null;
    orient?: number | string | undefined | null;
    orientation?: number | string | undefined | null;
    origin?: number | string | undefined | null;
    overflow?: number | string | undefined | null;
    'overline-position'?: number | string | undefined | null;
    'overline-thickness'?: number | string | undefined | null;
    'paint-order'?: number | string | undefined | null;
    'panose-1'?: number | string | undefined | null;
    path?: string | undefined | null;
    pathLength?: number | string | undefined | null;
    patternContentUnits?: string | undefined | null;
    patternTransform?: number | string | undefined | null;
    patternUnits?: string | undefined | null;
    'pointer-events'?: number | string | undefined | null;
    points?: string | undefined | null;
    pointsAtX?: number | string | undefined | null;
    pointsAtY?: number | string | undefined | null;
    pointsAtZ?: number | string | undefined | null;
    preserveAlpha?: number | string | undefined | null;
    preserveAspectRatio?: string | undefined | null;
    primitiveUnits?: number | string | undefined | null;
    r?: number | string | undefined | null;
    radius?: number | string | undefined | null;
    refX?: number | string | undefined | null;
    refY?: number | string | undefined | null;
    'rendering-intent'?: number | string | undefined | null;
    repeatCount?: number | string | undefined | null;
    repeatDur?: number | string | undefined | null;
    requiredExtensions?: number | string | undefined | null;
    requiredFeatures?: number | string | undefined | null;
    restart?: number | string | undefined | null;
    result?: string | undefined | null;
    rotate?: number | string | undefined | null;
    rx?: number | string | undefined | null;
    ry?: number | string | undefined | null;
    scale?: number | string | undefined | null;
    seed?: number | string | undefined | null;
    'shape-rendering'?: number | string | undefined | null;
    slope?: number | string | undefined | null;
    spacing?: number | string | undefined | null;
    specularConstant?: number | string | undefined | null;
    specularExponent?: number | string | undefined | null;
    speed?: number | string | undefined | null;
    spreadMethod?: string | undefined | null;
    startOffset?: number | string | undefined | null;
    stdDeviation?: number | string | undefined | null;
    stemh?: number | string | undefined | null;
    stemv?: number | string | undefined | null;
    stitchTiles?: number | string | undefined | null;
    'stop-color'?: string | undefined | null;
    'stop-opacity'?: number | string | undefined | null;
    'strikethrough-position'?: number | string | undefined | null;
    'strikethrough-thickness'?: number | string | undefined | null;
    string?: number | string | undefined | null;
    stroke?: string | undefined | null;
    'stroke-dasharray'?: string | number | undefined | null;
    'stroke-dashoffset'?: string | number | undefined | null;
    'stroke-linecap'?: 'butt' | 'round' | 'square' | 'inherit' | undefined | null;
    'stroke-linejoin'?: 'miter' | 'round' | 'bevel' | 'inherit' | undefined | null;
    'stroke-miterlimit'?: string | undefined | null;
    'stroke-opacity'?: number | string | undefined | null;
    'stroke-width'?: number | string | undefined | null;
    surfaceScale?: number | string | undefined | null;
    systemLanguage?: number | string | undefined | null;
    tableValues?: number | string | undefined | null;
    targetX?: number | string | undefined | null;
    targetY?: number | string | undefined | null;
    'text-anchor'?: string | undefined | null;
    'text-decoration'?: number | string | undefined | null;
    textLength?: number | string | undefined | null;
    'text-rendering'?: number | string | undefined | null;
    to?: number | string | undefined | null;
    transform?: string | undefined | null;
    u1?: number | string | undefined | null;
    u2?: number | string | undefined | null;
    'underline-position'?: number | string | undefined | null;
    'underline-thickness'?: number | string | undefined | null;
    unicode?: number | string | undefined | null;
    'unicode-bidi'?: number | string | undefined | null;
    'unicode-range'?: number | string | undefined | null;
    'units-per-em'?: number | string | undefined | null;
    'v-alphabetic'?: number | string | undefined | null;
    values?: string | undefined | null;
    'vector-effect'?: number | string | undefined | null;
    version?: string | undefined | null;
    'vert-adv-y'?: number | string | undefined | null;
    'vert-origin-x'?: number | string | undefined | null;
    'vert-origin-y'?: number | string | undefined | null;
    'v-hanging'?: number | string | undefined | null;
    'v-ideographic'?: number | string | undefined | null;
    viewBox?: string | undefined | null;
    viewTarget?: number | string | undefined | null;
    visibility?: number | string | undefined | null;
    'v-mathematical'?: number | string | undefined | null;
    widths?: number | string | undefined | null;
    'word-spacing'?: number | string | undefined | null;
    'writing-mode'?: number | string | undefined | null;
    x1?: number | string | undefined | null;
    x2?: number | string | undefined | null;
    x?: number | string | undefined | null;
    xChannelSelector?: string | undefined | null;
    'x-height'?: number | string | undefined | null;
    xlinkActuate?: string | undefined | null;
    xlinkArcrole?: string | undefined | null;
    xlinkHref?: string | undefined | null;
    xlinkRole?: string | undefined | null;
    xlinkShow?: string | undefined | null;
    xlinkTitle?: string | undefined | null;
    xlinkType?: string | undefined | null;
    xmlBase?: string | undefined | null;
    xmlLang?: string | undefined | null;
    xmlns?: string | undefined | null;
    xmlnsXlink?: string | undefined | null;
    xmlSpace?: string | undefined | null;
    y1?: number | string | undefined | null;
    y2?: number | string | undefined | null;
    y?: number | string | undefined | null;
    yChannelSelector?: string | undefined | null;
    z?: number | string | undefined | null;
    zoomAndPan?: string | undefined | null;
  }

  /**
   * @deprecated use the types from `svelte/elements` instead
   * For more info see https://github.com/sveltejs/language-tools/blob/master/docs/preprocessors/typescript.md#im-getting-deprecation-warnings-for-sveltejsx--i-want-to-migrate-to-the-new-typings
   */
  interface HTMLProps<T extends EventTarget> extends HTMLAttributes<T> {}
  /**
   * @deprecated use the types from `svelte/elements` instead
   * For more info see https://github.com/sveltejs/language-tools/blob/master/docs/preprocessors/typescript.md#im-getting-deprecation-warnings-for-sveltejsx--i-want-to-migrate-to-the-new-typings
   */
  interface SVGProps<T extends EventTarget> extends SVGAttributes<T> {}

  /**
   * @deprecated use the types from `svelte/elements` instead
   * For more info see https://github.com/sveltejs/language-tools/blob/master/docs/preprocessors/typescript.md#im-getting-deprecation-warnings-for-sveltejsx--i-want-to-migrate-to-the-new-typings
   */
  interface SvelteInputProps extends HTMLProps<HTMLInputElement> {
    group?: any | undefined | null;
    files?: FileList | undefined | null;
    indeterminate?: boolean | undefined | null;
  }

  /**
   * @deprecated use the types from `svelte/elements` instead
   * For more info see https://github.com/sveltejs/language-tools/blob/master/docs/preprocessors/typescript.md#im-getting-deprecation-warnings-for-sveltejsx--i-want-to-migrate-to-the-new-typings
   */
  interface SvelteWindowProps  {
    readonly innerWidth?: Window['innerWidth'] | undefined | null;
    readonly innerHeight?: Window['innerHeight'] | undefined | null;
    readonly outerWidth?: Window['outerWidth'] | undefined | null;
    readonly outerHeight?: Window['outerHeight'] | undefined | null;
    scrollX?: Window['scrollX'] | undefined | null;
    scrollY?: Window['scrollY'] | undefined | null;
    readonly online?: Window['navigator']['onLine'] | undefined | null;

    // Transformed from on:sveltekit:xy
    'onsveltekit:start'?: EventHandler<CustomEvent, Window> | undefined | null;
    'onsveltekit:navigation-start'?: EventHandler<CustomEvent, Window> | undefined | null;
    'onsveltekit:navigation-end'?: EventHandler<CustomEvent, Window> | undefined | null;

    ondevicelight?: EventHandler<Event, Window> | undefined | null;
    onbeforeinstallprompt?: EventHandler<Event, Window> | undefined | null;
    ondeviceproximity?: EventHandler<Event, Window> | undefined | null;
    onpaint?: EventHandler<Event, Window> | undefined | null;
    onuserproximity?: EventHandler<Event, Window> | undefined | null;
    onbeforeprint?: EventHandler<Event, Window> | undefined | null;
    onafterprint?: EventHandler<Event, Window> | undefined | null;
    onlanguagechange?: EventHandler<Event, Window> | undefined | null;
    onorientationchange?: EventHandler<Event, Window> | undefined | null;
    onmessage?: EventHandler<MessageEvent, Window> | undefined | null;
    onmessageerror?: EventHandler<MessageEvent, Window> | undefined | null;
    onoffline?: EventHandler<Event, Window> | undefined | null;
    ononline?: EventHandler<Event, Window> | undefined | null;
    onbeforeunload?: EventHandler<BeforeUnloadEvent, Window> | undefined | null;
    onunload?: EventHandler<Event, Window> | undefined | null;
    onstorage?: EventHandler<StorageEvent, Window> | undefined | null;
    onhashchange?: EventHandler<HashChangeEvent, Window> | undefined | null;
    onpagehide?: EventHandler<PageTransitionEvent, Window> | undefined | null;
    onpageshow?: EventHandler<PageTransitionEvent, Window> | undefined | null;
    onpopstate?: EventHandler<PopStateEvent, Window> | undefined | null;
    ondevicemotion?: EventHandler<DeviceMotionEvent> | undefined | null;
    ondeviceorientation?: EventHandler<DeviceOrientationEvent, Window> | undefined | null;
    ondeviceorientationabsolute?: EventHandler<DeviceOrientationEvent, Window> | undefined | null;
    onunhandledrejection?: EventHandler<PromiseRejectionEvent, Window> | undefined | null;
    onrejectionhandled?: EventHandler<PromiseRejectionEvent, Window> | undefined | null;
  }

  /**
   * @deprecated use the types from `svelte/elements` instead
   * For more info see https://github.com/sveltejs/language-tools/blob/master/docs/preprocessors/typescript.md#im-getting-deprecation-warnings-for-sveltejsx--i-want-to-migrate-to-the-new-typings
   */
  interface SapperAnchorProps {
      // transformed from sapper:noscroll so it should be camel case
      sapperNoscroll?: true | undefined | null;
      sapperPrefetch?: true | undefined | null;
  }

  /**
   * @deprecated use the types from `svelte/elements` instead
   * For more info see https://github.com/sveltejs/language-tools/blob/master/docs/preprocessors/typescript.md#im-getting-deprecation-warnings-for-sveltejsx--i-want-to-migrate-to-the-new-typings
   */
  interface SvelteMediaTimeRange {
      start: number;
      end: number;
  }

  /**
   * @deprecated use the types from `svelte/elements` instead
   * For more info see https://github.com/sveltejs/language-tools/blob/master/docs/preprocessors/typescript.md#im-getting-deprecation-warnings-for-sveltejsx--i-want-to-migrate-to-the-new-typings
   */
  interface SvelteMediaProps {
      readonly duration?: number | undefined | null;
      readonly buffered?: SvelteMediaTimeRange[] | undefined | null;
      readonly played?: SvelteMediaTimeRange[] | undefined | null;
      readonly seekable?: SvelteMediaTimeRange[] | undefined | null;
      readonly seeking?: boolean | undefined | null;
      readonly ended?: boolean | undefined | null;

      /**
       * the current playback time in the video, in seconds
       */
      currentTime?: number | undefined | null;
      /**
       * the current playback time in the video, in seconds
       */
      currenttime?: number | undefined | null;
      // Doesn't work when used as HTML Attribute
      /**
       * how fast or slow to play the video, where 1 is 'normal'
       */
      playbackRate?: number | undefined | null;

      paused?: boolean | undefined | null;
  }

  /**
   * @deprecated use the types from `svelte/elements` instead
   * For more info see https://github.com/sveltejs/language-tools/blob/master/docs/preprocessors/typescript.md#im-getting-deprecation-warnings-for-sveltejsx--i-want-to-migrate-to-the-new-typings
   */
  interface SvelteVideoProps extends SvelteMediaProps {
      // Binding only, don't need lowercase variant
      readonly videoWidth?: number | undefined | null;
      readonly videoHeight?: number | undefined | null;
  }

  /**
   * @deprecated use the types from `svelte/elements` instead
   * For more info see https://github.com/sveltejs/language-tools/blob/master/docs/preprocessors/typescript.md#im-getting-deprecation-warnings-for-sveltejsx--i-want-to-migrate-to-the-new-typings
   */
  interface IntrinsicElements {
    // HTML
    a: HTMLProps<HTMLAnchorElement> & SapperAnchorProps;
    abbr: HTMLProps<HTMLElement>;
    address: HTMLProps<HTMLElement>;
    area: HTMLProps<HTMLAreaElement>;
    article: HTMLProps<HTMLElement>;
    aside: HTMLProps<HTMLElement>;
    audio: HTMLProps<HTMLAudioElement> & SvelteMediaProps;
    b: HTMLProps<HTMLElement>;
    base: HTMLProps<HTMLBaseElement>;
    bdi: HTMLProps<HTMLElement>;
    bdo: HTMLProps<HTMLElement>;
    big: HTMLProps<HTMLElement>;
    blockquote: HTMLProps<HTMLElement>;
    body: HTMLProps<HTMLBodyElement>;
    br: HTMLProps<HTMLBRElement>;
    button: HTMLProps<HTMLButtonElement>;
    canvas: HTMLProps<HTMLCanvasElement>;
    caption: HTMLProps<HTMLElement>;
    cite: HTMLProps<HTMLElement>;
    code: HTMLProps<HTMLElement>;
    col: HTMLProps<HTMLTableColElement>;
    colgroup: HTMLProps<HTMLTableColElement>;
    data: HTMLProps<HTMLElement>;
    datalist: HTMLProps<HTMLDataListElement>;
    dd: HTMLProps<HTMLElement>;
    del: HTMLProps<HTMLElement>;
    details: HTMLProps<HTMLElement>;
    dfn: HTMLProps<HTMLElement>;
    dialog: HTMLProps<HTMLElement>;
    div: HTMLProps<HTMLDivElement>;
    dl: HTMLProps<HTMLDListElement>;
    dt: HTMLProps<HTMLElement>;
    em: HTMLProps<HTMLElement>;
    embed: HTMLProps<HTMLEmbedElement>;
    fieldset: HTMLProps<HTMLFieldSetElement>;
    figcaption: HTMLProps<HTMLElement>;
    figure: HTMLProps<HTMLElement>;
    footer: HTMLProps<HTMLElement>;
    form: HTMLProps<HTMLFormElement>;
    h1: HTMLProps<HTMLHeadingElement>;
    h2: HTMLProps<HTMLHeadingElement>;
    h3: HTMLProps<HTMLHeadingElement>;
    h4: HTMLProps<HTMLHeadingElement>;
    h5: HTMLProps<HTMLHeadingElement>;
    h6: HTMLProps<HTMLHeadingElement>;
    head: HTMLProps<HTMLHeadElement>;
    header: HTMLProps<HTMLElement>;
    hgroup: HTMLProps<HTMLElement>;
    hr: HTMLProps<HTMLHRElement>;
    html: HTMLProps<HTMLHtmlElement>;
    i: HTMLProps<HTMLElement>;
    iframe: HTMLProps<HTMLIFrameElement>;
    img: HTMLProps<HTMLImageElement>;
    input: SvelteInputProps;
    ins: HTMLProps<HTMLModElement>;
    kbd: HTMLProps<HTMLElement>;
    keygen: HTMLProps<HTMLElement>;
    label: HTMLProps<HTMLLabelElement>;
    legend: HTMLProps<HTMLLegendElement>;
    li: HTMLProps<HTMLLIElement>;
    link: HTMLProps<HTMLLinkElement>;
    main: HTMLProps<HTMLElement>;
    map: HTMLProps<HTMLMapElement>;
    mark: HTMLProps<HTMLElement>;
    menu: HTMLProps<HTMLElement>;
    menuitem: HTMLProps<HTMLElement>;
    meta: HTMLProps<HTMLMetaElement>;
    meter: HTMLProps<HTMLElement>;
    nav: HTMLProps<HTMLElement>;
    noindex: HTMLProps<HTMLElement>;
    noscript: HTMLProps<HTMLElement>;
    object: HTMLProps<HTMLObjectElement>;
    ol: HTMLProps<HTMLOListElement>;
    optgroup: HTMLProps<HTMLOptGroupElement>;
    option: HTMLProps<HTMLOptionElement>;
    output: HTMLProps<HTMLElement>;
    p: HTMLProps<HTMLParagraphElement>;
    param: HTMLProps<HTMLParamElement>;
    picture: HTMLProps<HTMLElement>;
    pre: HTMLProps<HTMLPreElement>;
    progress: HTMLProps<HTMLProgressElement>;
    q: HTMLProps<HTMLQuoteElement>;
    rp: HTMLProps<HTMLElement>;
    rt: HTMLProps<HTMLElement>;
    ruby: HTMLProps<HTMLElement>;
    s: HTMLProps<HTMLElement>;
    samp: HTMLProps<HTMLElement>;
    script: HTMLProps<HTMLElement>;
    section: HTMLProps<HTMLElement>;
    select: HTMLProps<HTMLSelectElement>;
    small: HTMLProps<HTMLElement>;
    source: HTMLProps<HTMLSourceElement>;
    span: HTMLProps<HTMLSpanElement>;
    strong: HTMLProps<HTMLElement>;
    style: HTMLProps<HTMLStyleElement>;
    sub: HTMLProps<HTMLElement>;
    summary: HTMLProps<HTMLElement>;
    sup: HTMLProps<HTMLElement>;
    table: HTMLProps<HTMLTableElement>;
    tbody: HTMLProps<HTMLTableSectionElement>;
    td: HTMLProps<HTMLTableDataCellElement>;
    textarea: HTMLProps<HTMLTextAreaElement>;
    tfoot: HTMLProps<HTMLTableSectionElement>;
    th: HTMLProps<HTMLTableHeaderCellElement>;
    thead: HTMLProps<HTMLTableSectionElement>;
    time: HTMLProps<HTMLElement>;
    title: HTMLProps<HTMLTitleElement>;
    tr: HTMLProps<HTMLTableRowElement>;
    track: HTMLProps<HTMLTrackElement>;
    u: HTMLProps<HTMLElement>;
    ul: HTMLProps<HTMLUListElement>;
    var: HTMLProps<HTMLElement>;
    video: HTMLProps<HTMLVideoElement> & SvelteVideoProps;
    wbr: HTMLProps<HTMLElement>;

    svg: SVGProps<SVGSVGElement>;

    animate: SVGProps<SVGElement>; // @TODO: It is SVGAnimateElement but not in dom.d.ts for now.
    circle: SVGProps<SVGCircleElement>;
    clipPath: SVGProps<SVGClipPathElement>;
    defs: SVGProps<SVGDefsElement>;
    desc: SVGProps<SVGDescElement>;
    ellipse: SVGProps<SVGEllipseElement>;
    feBlend: SVGProps<SVGFEBlendElement>;
    feColorMatrix: SVGProps<SVGFEColorMatrixElement>;
    feComponentTransfer: SVGProps<SVGFEComponentTransferElement>;
    feComposite: SVGProps<SVGFECompositeElement>;
    feConvolveMatrix: SVGProps<SVGFEConvolveMatrixElement>;
    feDiffuseLighting: SVGProps<SVGFEDiffuseLightingElement>;
    feDisplacementMap: SVGProps<SVGFEDisplacementMapElement>;
    feDistantLight: SVGProps<SVGFEDistantLightElement>;
    feFlood: SVGProps<SVGFEFloodElement>;
    feFuncA: SVGProps<SVGFEFuncAElement>;
    feFuncB: SVGProps<SVGFEFuncBElement>;
    feFuncG: SVGProps<SVGFEFuncGElement>;
    feFuncR: SVGProps<SVGFEFuncRElement>;
    feGaussianBlur: SVGProps<SVGFEGaussianBlurElement>;
    feImage: SVGProps<SVGFEImageElement>;
    feMerge: SVGProps<SVGFEMergeElement>;
    feMergeNode: SVGProps<SVGFEMergeNodeElement>;
    feMorphology: SVGProps<SVGFEMorphologyElement>;
    feOffset: SVGProps<SVGFEOffsetElement>;
    fePointLight: SVGProps<SVGFEPointLightElement>;
    feSpecularLighting: SVGProps<SVGFESpecularLightingElement>;
    feSpotLight: SVGProps<SVGFESpotLightElement>;
    feTile: SVGProps<SVGFETileElement>;
    feTurbulence: SVGProps<SVGFETurbulenceElement>;
    filter: SVGProps<SVGFilterElement>;
    foreignObject: SVGProps<SVGForeignObjectElement>;
    g: SVGProps<SVGGElement>;
    image: SVGProps<SVGImageElement>;
    line: SVGProps<SVGLineElement>;
    linearGradient: SVGProps<SVGLinearGradientElement>;
    marker: SVGProps<SVGMarkerElement>;
    mask: SVGProps<SVGMaskElement>;
    metadata: SVGProps<SVGMetadataElement>;
    path: SVGProps<SVGPathElement>;
    pattern: SVGProps<SVGPatternElement>;
    polygon: SVGProps<SVGPolygonElement>;
    polyline: SVGProps<SVGPolylineElement>;
    radialGradient: SVGProps<SVGRadialGradientElement>;
    rect: SVGProps<SVGRectElement>;
    stop: SVGProps<SVGStopElement>;
    switch: SVGProps<SVGSwitchElement>;
    symbol: SVGProps<SVGSymbolElement>;
    text: SVGProps<SVGTextElement>;
    textPath: SVGProps<SVGTextPathElement>;
    tspan: SVGProps<SVGTSpanElement>;
    use: SVGProps<SVGUseElement>;
    view: SVGProps<SVGViewElement>;

    // Svelte specific
    sveltewindow: HTMLProps<Window> & SvelteWindowProps;
    sveltebody: HTMLProps<HTMLElement>;
    sveltefragment: { slot?: string; };
    svelteoptions: { [name: string]: any };
    sveltehead: { [name: string]: any };
    svelteelement: { 'this': string | undefined | null; } & HTMLProps<any> & SVGProps<any> & SapperAnchorProps;
    // Needed due to backwards compatibility type which hits these
    'svelte:window': HTMLProps<Window> & SvelteWindowProps;
    'svelte:body': HTMLProps<HTMLElement>;
    'svelte:fragment': { slot?: string; };
    'svelte:options': { [name: string]: any };
    'svelte:head': { [name: string]: any };
  }
}
