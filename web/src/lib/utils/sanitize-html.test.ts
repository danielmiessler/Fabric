// @vitest-environment jsdom
import { describe, expect, it, vi } from 'vitest';
import { marked } from 'marked';
import { renderMarkdown as md, sanitizeHtml } from './sanitize-html';

describe('sanitizeHtml', () => {
  it('removes a script element', () => {
    const out = sanitizeHtml('<p>hi</p><script>window.x = 1</script>');
    expect(out).toBe('<p>hi</p>');
  });

  it('removes event attributes', () => {
    const out = sanitizeHtml('<img src="x" onerror="window.x = 1"><a href="/" onclick="f()">a</a>');
    expect(out).not.toMatch(/onerror|onclick/);
    expect(out).toContain('<img src="x">');
  });

  it('removes unsafe URL schemes and keeps safe ones', () => {
    const out = md(
      '[a](javascript:alert(1)) [b](https://example.com) [c](mailto:a@example.com) [d](/local)'
    );
    expect(out).not.toContain('javascript:');
    expect(out).toContain('href="https://example.com"');
    expect(out).toContain('href="mailto:a@example.com"');
    expect(out).toContain('href="/local"');
    expect(sanitizeHtml('<iframe src="data:text/html,x"></iframe>')).toBe('');
  });

  it('removes SVG, MathML and style', () => {
    const out = sanitizeHtml(
      '<svg onload="f()"><script>f()</script></svg><math><mi>x</mi></math>' +
        '<style>p{}</style><p style="color:red">t</p>'
    );
    expect(out).toBe('<p>t</p>');
  });

  it('keeps normal Markdown HTML', () => {
    const out = md(
      '# Title\n\nSome **bold** text and a [link](https://example.com).\n\n```go\nx := 1\n```'
    );
    expect(out).toContain('<h1>Title</h1>');
    expect(out).toContain('<strong>bold</strong>');
    expect(out).toContain('<a href="https://example.com">link</a>');
    expect(out).toContain('<pre><code class="language-go">x := 1\n</code></pre>');
  });

  // The chat view renders the reply again after each streamed chunk. Each
  // partial text must be safe, and the full text must be safe.
  it('cleans each part of a streamed reply', () => {
    const chunks = ['Text <img src=x on', 'error="window.x=1"> more <scr', 'ipt>f()</script>'];
    let content = '';
    for (const chunk of chunks) {
      content += chunk;
      const out = md(content);
      expect(out).not.toMatch(/onerror|<script/);
    }
  });

  // The save routes take a text/plain body, and the web UI origin passes the
  // Origin check. Thus a form in a reply must not survive.
  it('removes a form and keeps task list check boxes', () => {
    const out = sanitizeHtml(
      '<form action="/api/patterns/name" method="post" enctype="text/plain">' +
        '<input type="hidden" name="a" value="b"><button>Show</button></form>'
    );
    expect(out).not.toContain('<form');
    expect(out).toContain('<button>Show</button>');
    expect(md('- [x] done')).toContain('<input checked="" disabled="" type="checkbox">');
  });

  // A session restored from localStorage goes through the same path.
  it('cleans a restored message', () => {
    const stored = JSON.parse(
      JSON.stringify({ role: 'assistant', content: '<details open ontoggle="f()">x</details>' })
    );
    expect(md(stored.content)).not.toContain('ontoggle');
  });

  // ChatMessages.svelte renders replies with renderMarkdown.
  it('gives escaped text when marked fails', () => {
    const spy = vi.spyOn(marked, 'parse').mockImplementation(() => {
      throw new Error('bad markdown');
    });
    vi.spyOn(console, 'error').mockImplementation(() => {});
    try {
      expect(md('<img src=x onerror="f()"> & text')).toBe(
        '&#60;img src=x onerror=&#34;f()&#34;&#62; &#38; text'
      );
    } finally {
      vi.restoreAllMocks();
    }
    expect(spy).toHaveBeenCalled();
  });
});
