import { describe, expect, it } from 'vitest';
import { maxYoutubeUrlLength, youtubeVideoId } from './validators';

describe('youtubeVideoId', () => {
  it('gives the video ID', () => {
    expect(youtubeVideoId('https://www.youtube.com/watch?v=dQw4w9WgXcQ')).toBe('dQw4w9WgXcQ');
    expect(youtubeVideoId('https://youtu.be/dQw4w9WgXcQ?t=5')).toBe('dQw4w9WgXcQ');
    expect(youtubeVideoId('youtube.com/embed/dQw4w9WgXcQ')).toBe('dQw4w9WgXcQ');
    expect(youtubeVideoId('https://m.youtube.com/v/dQw4w9WgXcQ')).toBe('dQw4w9WgXcQ');
    expect(youtubeVideoId('https://WWW.YOUTUBE.COM/watch?v=dQw4w9WgXcQ')).toBe('dQw4w9WgXcQ');
    expect(youtubeVideoId('https://youtube.com:443/watch?v=dQw4w9WgXcQ')).toBe('dQw4w9WgXcQ');
  });

  it('accepts a URL at the length limit', () => {
    const url = 'https://youtu.be/dQw4w9WgXcQ?x=';
    expect(youtubeVideoId(url.padEnd(maxYoutubeUrlLength, 'a'))).toBe('dQw4w9WgXcQ');
    expect(youtubeVideoId(url.padEnd(maxYoutubeUrlLength + 1, 'a'))).toBeNull();
  });

  it('refuses a YouTube path on a different host', () => {
    expect(youtubeVideoId('https://example.com/x/youtube.com/v/dQw4w9WgXcQ')).toBeNull();
    expect(youtubeVideoId('https://notyoutube.com/watch?v=dQw4w9WgXcQ')).toBeNull();
    expect(youtubeVideoId('https://youtube.com.example/watch?v=dQw4w9WgXcQ')).toBeNull();
  });

  it('refuses a value that is not a YouTube URL string', () => {
    expect(youtubeVideoId(42)).toBeNull();
    expect(youtubeVideoId('https://example.com/')).toBeNull();
  });
});
