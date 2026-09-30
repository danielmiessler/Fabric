import type { StorageEntity } from './storage-interface';

// One entry of the patterns array in pattern_descriptions.json.
export interface PatternDescription {
  patternName: string;
  description: string;
  tags?: string[];
}

// StorageEntity requires the uppercase Name field.
export interface Pattern extends StorageEntity {
  Name: string;        // patternName in the JSON
  Description: string; // description in the JSON
  Pattern: string;     // pattern content from the API
  tags: string[];
}