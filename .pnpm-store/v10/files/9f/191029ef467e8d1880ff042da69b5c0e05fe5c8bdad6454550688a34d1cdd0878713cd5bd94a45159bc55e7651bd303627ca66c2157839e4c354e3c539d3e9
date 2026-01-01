<div class="container">
  <h1 data-svelte-h="svelte-wr14qn">Youtube Video Embed</h1>
  <p data-svelte-h="svelte-13nlixd">
    The package provides faster page load while embedding youtube videos since
    it loads the youtube thumbnail istead of the entire iframe.
  </p>
  <p data-svelte-h="svelte-1hlxsk3">
    The package will load the iframe and play the video only upong clicking the
    play button.
  </p>
  <h2 data-svelte-h="svelte-136tnhu">Installation</h2>
  <pre
    class="svelte-w3w8f2"
  ><!-- HTML_TAG_START -->npm i -D svelte-youtube-embed<!-- HTML_TAG_END -->
</pre>
  <h2 data-svelte-h="svelte-dnbfhr">Usage</h2>
  <pre
    class="svelte-w3w8f2"
  ><!-- HTML_TAG_START --><span class="token tag"><span class="token tag"><span class="token punctuation">&lt;</span>script</span><span class="token punctuation">></span></span><span class="token script"><span class="token language-javascript">
  <span class="token keyword">import</span> Youtube <span class="token keyword">from</span> <span class="token string">"svelte-youtube-embed"</span><span class="token punctuation">;</span>
</span></span><span class="token tag"><span class="token tag"><span class="token punctuation">&lt;/</span>script</span><span class="token punctuation">></span></span>

<span class="token tag"><span class="token tag"><span class="token punctuation">&lt;</span>Youtube</span> <span class="token attr-name">id</span><span class="token attr-value"><span class="token punctuation">=</span><span class="token punctuation">"</span>EBtsu6naB8g<span class="token punctuation">"</span></span> <span class="token punctuation">/></span></span>

<!-- HTML_TAG_END -->
</pre>
  <h2 data-svelte-h="svelte-zcecrh">Default</h2>
  <pre
    class="svelte-w3w8f2"
  ><!-- HTML_TAG_START --><span class="token tag"><span class="token tag"><span class="token punctuation">&lt;</span>Youtube</span> <span class="token attr-name">id</span><span class="token attr-value"><span class="token punctuation">=</span><span class="token punctuation">"</span>EBtsu6naB8g<span class="token punctuation">"</span></span> <span class="token punctuation">/></span></span><!-- HTML_TAG_END -->
</pre>
  <h2 data-svelte-h="svelte-nf84rh">With custom play button</h2>

  <pre
    class="svelte-w3w8f2"
  ><!-- HTML_TAG_START --><span class="token tag"><span class="token tag"><span class="token punctuation">&lt;</span>Youtube</span> <span class="token attr-name">id</span><span class="token attr-value"><span class="token punctuation">=</span><span class="token punctuation">"</span>EBtsu6naB8g<span class="token punctuation">"</span></span><span class="token punctuation">></span></span>
  <span class="token tag"><span class="token tag"><span class="token punctuation">&lt;</span>button</span><span class="token punctuation">></span></span>play<span class="token tag"><span class="token tag"><span class="token punctuation">&lt;/</span>button</span><span class="token punctuation">></span></span>
<span class="token tag"><span class="token tag"><span class="token punctuation">&lt;/</span>Youtube</span><span class="token punctuation">></span></span><!-- HTML_TAG_END -->
</pre>
  <h2 data-svelte-h="svelte-1r92b2h">Using custom overlay</h2>
  <pre
    class="svelte-w3w8f2"
  ><!-- HTML_TAG_START --><span class="token tag"><span class="token tag"><span class="token punctuation">&lt;</span>Youtube</span>
  <span class="token attr-name">id</span><span class="token attr-value"><span class="token punctuation">=</span><span class="token punctuation">"</span>EBtsu6naB8g<span class="token punctuation">"</span></span>
  <span class="token attr-name">--overlay-bg-color="</span><span class="token language-javascript"><span class="token punctuation">{</span>overlayBGColor<span class="token punctuation">}</span></span><span class="token attr-name">30"</span>
  <span class="token attr-name">--overlay-transition="all</span> <span class="token language-javascript"><span class="token punctuation">{</span>overlayTransitionDuration<span class="token punctuation">}</span></span><span class="token attr-name">ms</span> <span class="token attr-name">linear"</span>
  <span class="token punctuation">/></span></span><!-- HTML_TAG_END -->
</pre>
  <p data-svelte-h="svelte-vvtw0o">
    Here the number `30` represents the opacity of `0.3`. Remove it for 100%
    opaqueness.
  </p>
  <h2 data-svelte-h="svelte-15qxpdj">Using custom title colors</h2>

  <pre
    class="svelte-w3w8f2"
  ><!-- HTML_TAG_START --><span class="token tag"><span class="token tag"><span class="token punctuation">&lt;</span>Youtube</span>
  <span class="token attr-name">id</span><span class="token attr-value"><span class="token punctuation">=</span><span class="token punctuation">"</span>KrSH82gg7BQ<span class="token punctuation">"</span></span>
  <span class="token attr-name">--title-color=</span><span class="token language-javascript"><span class="token punctuation">{</span>titleColor<span class="token punctuation">}</span></span>
  <span class="token attr-name">--title-shadow-color="</span><span class="token language-javascript"><span class="token punctuation">{</span>titleShadowColor<span class="token punctuation">}</span></span><span class="token attr-name">30"</span>
  <span class="token attr-name">--title-font-family=</span><span class="token language-javascript"><span class="token punctuation">{</span>titleFontFamily<span class="token punctuation">}</span></span>
<span class="token punctuation">/></span></span><!-- HTML_TAG_END -->
</pre>
  <h2 data-svelte-h="svelte-1y3xwn7">Alternative thumbnail</h2>
  <p data-svelte-h="svelte-nph3bp">
    Some videos don&#39;t have custom thumbnail, so you can use this option to
    set the thumbnail to the default one.
  </p>
  <pre
    class="svelte-w3w8f2"
  ><!-- HTML_TAG_START --><span class="token tag"><span class="token tag"><span class="token punctuation">&lt;</span>Youtube</span> <span class="token attr-name">id</span><span class="token attr-value"><span class="token punctuation">=</span><span class="token punctuation">"</span>EBtsu6naB8g<span class="token punctuation">"</span></span> <span class="token attr-name">altThumb=</span><span class="token language-javascript"><span class="token punctuation">{</span><span class="token boolean">true</span><span class="token punctuation">}</span></span> <span class="token punctuation">/></span></span><!-- HTML_TAG_END -->
</pre>
  <h2 data-svelte-h="svelte-wdcdcx">Disable animations</h2>
  <p data-svelte-h="svelte-1nd7yj2">
    You can disable animation that is displayed when the play button is clicked
    by passing `animations` attribute.
  </p>
  <pre
    class="svelte-w3w8f2"
  ><!-- HTML_TAG_START --><span class="token tag"><span class="token tag"><span class="token punctuation">&lt;</span>Youtube</span> <span class="token attr-name">id</span><span class="token attr-value"><span class="token punctuation">=</span><span class="token punctuation">"</span>EBtsu6naB8g<span class="token punctuation">"</span></span> <span class="token attr-name">animations=</span><span class="token language-javascript"><span class="token punctuation">{</span><span class="token boolean">false</span><span class="token punctuation">}</span></span> <span class="token punctuation">/></span></span><!-- HTML_TAG_END -->
</pre>
  <h2 data-svelte-h="svelte-xhxpwp">Custom thumbnail</h2>
  <p data-svelte-h="svelte-e07e6k">
    Similar to how we can use custom play button, we can also use custom
    thumbnails.
  </p>
  <p data-svelte-h="svelte-186y148">
    Use <strong>slot=&quot;thumbnail&quot;</strong> to a image/picture html tag
    and place it inside <em>Yoututbe</em> component.
  </p>
  <pre
    class="svelte-w3w8f2"
  ><!-- HTML_TAG_START --><span class="token tag"><span class="token tag"><span class="token punctuation">&lt;</span>Youtube</span> <span class="token attr-name">id</span><span class="token attr-value"><span class="token punctuation">=</span><span class="token punctuation">"</span>EBtsu6naB8g<span class="token punctuation">"</span></span><span class="token punctuation">></span></span>
  <span class="token tag"><span class="token tag"><span class="token punctuation">&lt;</span>img</span>
    <span class="token attr-name">slot</span><span class="token attr-value"><span class="token punctuation">=</span><span class="token punctuation">"</span>thumbnail<span class="token punctuation">"</span></span>
    <span class="token attr-name">alt</span><span class="token attr-value"><span class="token punctuation">=</span><span class="token punctuation">"</span>A video on svelte<span class="token punctuation">"</span></span>
    <span class="token attr-name">src</span><span class="token attr-value"><span class="token punctuation">=</span><span class="token punctuation">"</span>/path/image.jpg<span class="token punctuation">"</span></span>
    <span class="token attr-name">style</span><span class="token attr-value"><span class="token punctuation">=</span><span class="token punctuation">"</span>width: 100%; height: 100%; object-fit: contain; object-position: center; background: red<span class="token punctuation">"</span></span>
  <span class="token punctuation">/></span></span>
  <span class="token comment">&lt;!-- inline styles are optional --></span>
<span class="token tag"><span class="token tag"><span class="token punctuation">&lt;/</span>Youtube</span><span class="token punctuation">></span></span>
<!-- HTML_TAG_END -->
</pre>
</div>
