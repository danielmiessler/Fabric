// youtubeVideoId gives the 11-character video ID in a YouTube URL, or null.
// The host must be youtube.com, a subdomain of youtube.com, or youtu.be. The
// match ignores case and accepts a port. The URL comes from a request. Thus
// check the type and the length before the regex match, so that the match
// time stays small.
export const maxYoutubeUrlLength = 2048;

export function youtubeVideoId(url: unknown): string | null {
  if (typeof url !== 'string' || url.length > maxYoutubeUrlLength) return null;
  const match = url.match(
    /^(?:https?:\/\/)?(?:[a-z0-9-]+\.)*(?:youtube\.com(?::\d+)?\/(?:[^/]+\/.+\/|(?:v|e(?:mbed)?)\/|.*[?&]v=)|youtu\.be(?::\d+)?\/)([^"&?/\s]{11})/i
  );
  return match ? match[1] : null;
}
