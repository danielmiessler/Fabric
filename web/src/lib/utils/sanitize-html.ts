import DOMPurify from 'dompurify';
import { marked } from 'marked';

// Model output is not trusted. Call this function on all HTML before you give
// it to {@html}. The function keeps only HTML elements: it removes SVG,
// MathML, <style> elements, style attributes, event attributes and unsafe URL
// schemes such as javascript:. It also removes <form> elements, because one
// click on a form button can send a request to the Fabric API.
export function sanitizeHtml(html: string): string {
  // DOMPurify needs a DOM. Without one (server rendering), it returns its
  // input unchanged, so give an empty string.
  if (!DOMPurify.isSupported) return '';
  return DOMPurify.sanitize(html, {
    USE_PROFILES: { html: true },
    FORBID_TAGS: ['style', 'form'],
    FORBID_ATTR: ['style']
  });
}

// escapeHtml gives text with HTML escapes.
function escapeHtml(text: string): string {
  return text.replace(/[&<>"']/g, (c) => `&#${c.charCodeAt(0)};`);
}

// renderMarkdown gives model output as safe HTML. If marked fails, or if
// there is no DOM (server rendering), it gives the text with HTML escapes, so
// that the reply does not show as empty.
export function renderMarkdown(content: string): string {
  if (!DOMPurify.isSupported) return escapeHtml(content);
  try {
    return sanitizeHtml(marked.parse(content, { async: false }) as string);
  } catch (error) {
    console.error('Error rendering markdown:', error);
    return escapeHtml(content);
  }
}
