import { compareTwoStrings } from 'string-similarity';

import Item from '../Item';
import ItemResult from '../ItemResult';
import ItemTransformer from './ItemTransformer';
import TransformContext from './TransformContext';
import LineItemMerger from '../debug/LineItemMerger';
import { MIN_Y, MAX_Y, PAGE_MAPPING } from './CacluclateStatistics';
import {
  ascending,
  flatMap,
  groupByLine,
  groupByPage,
  median,
  mostFrequent,
  onlyUniques,
  transformGroupedByPageAndLine,
} from '../support/groupingUtils';
import { filterOutDigits } from '../support/stringFunctions';
import { extractNumbers } from '../support/stringFunctions';

const config = {
  // From the absolute fringe elements (min/max y) how much y can item deviate before beeing disregarded.
  maxDistanceFromFringeElements: 45,

  // Max neighbour taken (in one direction) for detecting neighbour similarity.
  // Choosen number might be more effectful for PDFs with a strong odd/evan page differernce.
  neighbourReach: 2,

  minScore: 0.7,
};

export default class RemoveRepetitiveItems extends ItemTransformer {
  constructor() {
    super('Remove Repetitive Items', 'Remove things like page numbers or license footers.', {
      requireColumns: ['x', 'y', 'str', 'line'],
      debug: {
        itemMerger: new LineItemMerger(),
      },
    });
  }

  transform(context: TransformContext, inputItems: Item[]): ItemResult {
    const minY = context.getGlobal(MIN_Y);
    const maxY = context.getGlobal(MAX_Y);
    const bottomMaxY = minY + config.maxDistanceFromFringeElements;
    const topMinY = maxY - config.maxDistanceFromFringeElements;
    // console.log('bottomMaxY', bottomMaxY, 'topMinY', topMinY);

    const fringeItems = inputItems.filter((item) => {
      const y = item.data['y'];
      return y <= bottomMaxY || y >= topMinY;
    });

    const fringeLines = flatMap(
      groupByPage(fringeItems).map((pageItems) =>
        groupByLine(pageItems)
          .map((lineItems) => {
            const lineY = yFromLineItems(lineItems);
            return new PageLine(pageItems[0].page, lineY, lineItems);
          })
          .sort((a, b) => a.y - b.y),
      ),
      (e) => e,
    );

    const fringeYs = fringeLines
      .map((line) => line.y)
      .filter(onlyUniques)
      .sort(ascending);

    const yScoreMap = calculateScores(context, fringeYs, fringeLines);
    // console.log('uniqueYs', uniqueYs);

    let removalCount = 0;
    const removedY = [...yScoreMap.entries()]
      .filter(([_, value]) => value.value >= config.minScore)
      .map(([key, _]) => key)
      .join('||');
    return {
      items: transformGroupedByPageAndLine(inputItems, (_, __, lineItems) => {
        const itemsY = yFromLineItems(lineItems);
        const score = yScoreMap.get(itemsY);
        if (score) {
          lineItems.forEach((item) => context.trackEvaluation(item, score.description));
          if (score.value >= config.minScore) {
            removalCount++;
            return [];
          }
        }
        return lineItems;
      }),
      messages: [`Filtered out ${removalCount} items with y == ${removedY}`],
    };
  }
}

function calculateScores(context: TransformContext, fringeYs: number[], fringeLines: PageLine[]): Map<number, Score> {
  const pageMapping = context.getGlobal(PAGE_MAPPING);
  const map: Map<number, Score> = new Map();
  fringeYs.forEach((y) => {
    const yLines = fringeLines.filter((line) => line.y == y);
    if (yLines.length < 2) {
      map.set(y, new Score(0, `0 (only on ${yLines.length} page(s))`));
    } else {
      const pageNumberScore: number = pageMapping.detectedOnPage
        ? calculatePageNumerScore(context.pageCount, pageMapping.pageFactor, yLines)
        : 0;
      const textSimilarityScore: number = textSimilarity(yLines);

      // TODO possibly refine with:
      // - exclude headlines (higher height, e.g art of speaking)
      // - structural similarity (x, y, height, etc)
      // - contain chapter headings

      const totalScore = pageNumberScore + textSimilarityScore;
      map.set(
        y,
        new Score(
          totalScore,
          `${totalScore.toFixed(2)}: (${pageNumberScore.toFixed(2)} + ${textSimilarityScore.toFixed(2)})`,
        ),
      );
    }
  });

  return map;
}

function calculatePageNumerScore(pageCount: number, pageFactor: number, lines: PageLine[]): number {
  const maxPageNumbers = pageCount + pageFactor;
  const linesWithPageNumbers = lines.filter((line) => extractNumbers(line.text()).includes(line.page + pageFactor))
    .length;
  return linesWithPageNumbers / Math.min(maxPageNumbers, lines.length);
}

function textSimilarity(lines: PageLine[]): number {
  const similarities = flatMap(lines, (line, idx) =>
    adiacentLines(lines, idx).map((adiacentLine) => calculateSimilarity(line, adiacentLine)),
  );
  return median(similarities);
}

function calculateSimilarity(line1: PageLine, line2: PageLine): number {
  if (line1.textWithoutNumbers().length === 0) {
    return 0;
  }
  return compareTwoStrings(line1.textWithoutNumbers(), line2.textWithoutNumbers());
}

function adiacentLines(lines: PageLine[], index: number): PageLine[] {
  // Prefer to either collect x downstream OR x upstream neighbours (not a mix) in order to better catch odd/even page differences
  let neighbours: PageLine[];
  if (index + config.neighbourReach < lines.length) {
    neighbours = lines.slice(index + 1, index + config.neighbourReach + 1);
  } else if (index - config.neighbourReach >= 0) {
    neighbours = lines.slice(index - config.neighbourReach - 1, index - 1);
  } else {
    neighbours = lines.filter((_, idx) => idx !== index);
  }

  return neighbours;
}

function yFromLineItems(lineItems: Item[]): number {
  return Math.round(mostFrequent(lineItems, 'y') as number);
}

class Score {
  constructor(public value: number, public description: string) {}
}

/**
 * A number of Items on a line (~same y) on a page.
 */
class PageLine {
  private _text: string | undefined;
  private _textWithoutNumbers: string | undefined;

  constructor(public page: number, public y: number, public items: Item[]) {}

  text(): string {
    if (!this._text) {
      this._text = this.items.reduce((all, item) => all + item.data['str'], '');
    }
    return this._text;
  }

  textWithoutNumbers(): string {
    if (!this._textWithoutNumbers) {
      this._textWithoutNumbers = filterOutDigits(this.text());
    }
    return this._textWithoutNumbers;
  }
}
