<script>import { createEventDispatcher } from "svelte";
const dispatch = createEventDispatcher();
import { storeHighlightJs } from "./stores.js";
import { clipboard } from "../../actions/Clipboard/clipboard.js";
export let language = "plaintext";
export let code = "";
export let lineNumbers = false;
export let background = "bg-neutral-900/90";
export let blur = "";
export let text = "text-sm";
export let color = "text-white";
export let rounded = "rounded-container-token";
export let shadow = "shadow";
export let button = "btn btn-sm variant-soft !text-white";
export let buttonLabel = "Copy";
export let buttonCopied = "\u{1F44D}";
const cBase = "overflow-hidden shadow";
const cHeader = "text-xs text-white/50 uppercase flex justify-between items-center p-2 pl-4";
const cPre = "whitespace-pre-wrap break-all p-4 pt-1";
let formatted = false;
let displayCode = code;
let copyState = false;
function languageFormatter(lang) {
  if (lang === "js")
    return "javascript";
  if (lang === "ts")
    return "typescript";
  if (lang === "shell")
    return "terminal";
  return lang;
}
function onCopyClick() {
  copyState = true;
  setTimeout(() => {
    copyState = false;
  }, 2e3);
  dispatch("copy");
}
$:
  if ($storeHighlightJs !== void 0) {
    displayCode = $storeHighlightJs.highlight(code, { language }).value.trim();
    formatted = true;
  }
$:
  if (lineNumbers) {
    displayCode = displayCode.replace(/^/gm, () => {
      return '<span class="line"></span>	';
    });
    formatted = true;
  }
$:
  classesBase = `${cBase} ${background} ${blur} ${text} ${color} ${rounded} ${shadow} ${$$props.class ?? ""}`;
</script>

<!-- prettier-ignore -->
{#if language && code}
<div class="codeblock {classesBase}" data-testid="codeblock">
	<!-- Header -->
	<header class="codeblock-header {cHeader}">
		<!-- Language -->
		<span class="codeblock-language">{languageFormatter(language)}</span>
		<!-- Copy Button -->
		<button type="button" class="codeblock-btn {button}" on:click={onCopyClick} use:clipboard={code}>
			{!copyState ? buttonLabel : buttonCopied}
		</button>
	</header>
	<!-- Pre/Code -->
	<pre class="codeblock-pre {cPre}"><code class="codeblock-code language-{language} lineNumbers">{#if formatted}{@html displayCode}{:else}{code.trim()}{/if}</code></pre>
</div>
{/if}
