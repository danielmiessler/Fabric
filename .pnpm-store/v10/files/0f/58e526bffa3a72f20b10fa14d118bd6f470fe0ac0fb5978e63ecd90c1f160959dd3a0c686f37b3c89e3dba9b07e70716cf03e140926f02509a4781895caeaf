import type {
  MarkedExtension as MarkedConfig,
  Tokens,
  TokensList,
} from 'marked'
import type { SvelteComponentTyped } from 'svelte'

type MarkedRendererProps<
  T extends { type: string; raw: string; text?: string }
> = Partial<Omit<T, 'type'>>

type InstantiableSvelteComponentTyped<
  Props extends Record<string, any> = any,
  Events extends Record<string, any> = any,
  Slots extends Record<string, any> = any
> = new (...args: any[]) => SvelteComponentTyped<Props, Events, Slots>

type Renderers = {
  text: InstantiableSvelteComponentTyped<MarkedRendererProps<Tokens.Text>>
  paragraph: InstantiableSvelteComponentTyped<
    MarkedRendererProps<Tokens.Paragraph>
  >
  em: InstantiableSvelteComponentTyped<Omit<Tokens.Em, 'type' | 'text'>>
  strong: InstantiableSvelteComponentTyped<MarkedRendererProps<Tokens.Strong>>
  hr: InstantiableSvelteComponentTyped<MarkedRendererProps<Tokens.Hr>>
  blockquote: InstantiableSvelteComponentTyped<
    MarkedRendererProps<Tokens.Blockquote>
  >
  del: InstantiableSvelteComponentTyped<MarkedRendererProps<Tokens.Del>>
  link: InstantiableSvelteComponentTyped<MarkedRendererProps<Tokens.Link>>
  image: InstantiableSvelteComponentTyped<MarkedRendererProps<Tokens.Image>>
  table: InstantiableSvelteComponentTyped<{}>
  tablehead: InstantiableSvelteComponentTyped<{}>
  tablebody: InstantiableSvelteComponentTyped<{}>
  tablerow: InstantiableSvelteComponentTyped<{}>
  tablecell: InstantiableSvelteComponentTyped<{}>
  list: InstantiableSvelteComponentTyped<MarkedRendererProps<Tokens.List>>

  // Technically, listItem includes {type: string, tokens: []} in the props,
  // but that's probably unintentional
  listitem: InstantiableSvelteComponentTyped<
    MarkedRendererProps<Tokens.ListItem>
  >
  heading: InstantiableSvelteComponentTyped<MarkedRendererProps<Tokens.Heading>>
  codespan: InstantiableSvelteComponentTyped<
    MarkedRendererProps<Tokens.Codespan>
  >
  code: InstantiableSvelteComponentTyped<MarkedRendererProps<Tokens.Code>>

  // Note there's some weird behaviour involved with this ATM.
  /**
   * If the html is an inline element, it *should* be a Tag token (e.g `span`, `a`) If it's a
   * block element, it *should* be an HTML token (e.g `div`, `p`).
   *
   * @see {@link https://developer.mozilla.org/en-US/docs/Web/HTML/Block-level_elements#elements}
   * @see {@link https://developer.mozilla.org/en-US/docs/Web/HTML/Inline_elements#elements}
   */
  html: InstantiableSvelteComponentTyped<
    MarkedRendererProps<Tokens.HTML | Tokens.Tag>
  >
}
type Props = {
  /**
   * The Markdown source to be parsed, or an array of tokens to be rendered directly.
   */
  source: string | TokensList

  /**
   * An object where the keys represent a node type and the value is a Svelte component. This
   * object will be merged with the default renderers.
   */
  renderers?: Partial<Renderers>

  /**
   * Options for [marked](https://marked.js.org/using_advanced#options)
   */
  options?: MarkedConfig

  /**
   * To use [inline markdown](https://marked.js.org/using_advanced#inline), you can assign the prop `isInline` to the component.
   */
  isInline?: boolean
}
export default class SvelteMarkdown extends SvelteComponentTyped<
  Props,
  { parsed: CustomEvent<{ tokens: TokensList }> },
  { default: {} }
> {}
