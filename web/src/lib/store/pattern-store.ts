import { createStorageAPI } from '$lib/api/base';
import type { Pattern, PatternDescription } from '$lib/interfaces/pattern-interface';
import { get, writable, derived } from 'svelte/store';
import { languageStore } from './language-store';

const allPatterns = writable<Pattern[]>([]);

export const patterns = derived(
  [allPatterns, languageStore],
  ([$allPatterns, $language]) => {
    if (!$language) return $allPatterns;
    return $allPatterns.filter(p => {
      if (!$language) return true;

      // A two-letter prefix, for example en_ or fr_, is a language code.
      const match = p.Name.match(/^([a-z]{2})_/);
      if (!match) return true;

      const patternLang = match[1];
      return patternLang === $language;
    });
  }
);

export const systemPrompt = writable<string>('');
export const selectedPatternName = writable<string>('');

export const patternVariables = writable<Record<string, string>>({});

export const setSystemPrompt = (prompt: string) => {
  console.log('Setting system prompt:', prompt);
  systemPrompt.set(prompt);
  console.log('Current system prompt:', get(systemPrompt));
};

export const patternAPI = {
  ...createStorageAPI<Pattern>('patterns'),

  async loadPatterns() {
    try {
      const descriptionsResponse = await fetch('/data/pattern_descriptions.json');
      const descriptionsData = await descriptionsResponse.json();
      const descriptions = descriptionsData.patterns as PatternDescription[];
      console.log("Loaded pattern descriptions:", descriptions.length);

      const response = await fetch(`/api/patterns/names`);
      const data = await response.json();
      console.log("Load Patterns:", data);
      console.log("Loading patterns from API...");

      const patternsPromises = data.map(async (pattern: string) => {
        try {
          console.log(`Loading pattern: ${pattern}`);
          const patternResponse = await fetch(`/api/patterns/${pattern}`);
          const patternData = await patternResponse.json();
          console.log(`Pattern ${pattern} content length:`, patternData.Pattern?.length || 0);

          const desc = descriptions.find(d => d.patternName === pattern);
          if (!desc) {
            console.warn(`No description found for pattern: ${pattern}`);
          }

          return {
            Name: pattern,
            Description: desc?.description || pattern.charAt(0).toUpperCase() + pattern.slice(1),
            Pattern: patternData.Pattern || "",
            tags: desc?.tags || []
          };
        } catch (error) {
          console.error(`Failed to load pattern ${pattern}:`, error);
          // The description is still available when the content fetch fails.
          const desc = descriptions.find(d => d.patternName === pattern);
          return {
            Name: pattern,
            Description: desc?.description || pattern.charAt(0).toUpperCase() + pattern.slice(1),
            Pattern: "",
            tags: desc?.tags || []
          };
        }
      });

      const loadedPatterns = await Promise.all(patternsPromises);
      console.log("Patterns with content:", loadedPatterns);
      allPatterns.set(loadedPatterns);
      return loadedPatterns;
    } catch (error) {
      console.error('Failed to load patterns:', error);
      allPatterns.set([]);
      return [];
    }
  },

  selectPattern(patternName: string) {
    const patterns = get(allPatterns);
    console.log('Selecting pattern:', patternName);
    const selectedPattern = patterns.find(p => p.Name === patternName);
    if (selectedPattern) {
      console.log('Found pattern content (length: ' + selectedPattern.Pattern.length + '):', selectedPattern.Pattern);
      console.log('First 100 chars:', selectedPattern.Pattern.substring(0, 100));
      console.log('Last 100 chars:', selectedPattern.Pattern.substring(selectedPattern.Pattern.length - 100));
      console.log(`Setting system prompt with content length: ${selectedPattern.Pattern.length}`);
      console.log(`Content preview:`, selectedPattern.Pattern.substring(0, 100));
      setSystemPrompt(selectedPattern.Pattern);
      selectedPatternName.set(patternName);
    } else {
      console.log('No pattern found for name:', patternName);
      setSystemPrompt('');
      selectedPatternName.set('');
    }
    console.log('System prompt store value after setting:', get(systemPrompt));
  }
};
