import { compareTwoStrings } from 'string-similarity';

import ItemTransformer from './ItemTransformer';
import GlobalDefinition from '../GlobalDefinition';
import Item from '../Item';
import ItemResult from '../ItemResult';
import TransformContext from './TransformContext';
import LineItemMerger from '../debug/LineItemMerger';
import { descending, groupByLine, groupByPage, onlyUniques } from '../support/groupingUtils';
import { MOST_USED_HEIGHT, PAGE_MAPPING } from './CacluclateStatistics';
import {
  extractEndingNumber,
  filterOut,
  DASHS_CHAR_CODES,
  TAB_CHAR_CODE,
  WHITESPACE_CHAR_CODE,
  PERIOD_CHAR_CODES,
} from '../support/stringFunctions';
import { numbersAreConsecutive } from '../support/numberFunctions';
import TOC from '../TOC';
import FontType, { declaredFontTypes } from '../FontType';
import { flatten, groupBy } from '../support/functional';
import { getHeight, getText, getFontName, itemWithType } from '../support/items';
import { HeadlineType, toHeadlineType } from '../text-types';

export interface HeadlineRange {
  min: number;
  max: number;
}

const config = {
  // How many characters a line with a ending number needs to have minimally to be a valid link
  linkMinLength: 4,

  // How much bigger (height) then a 'normal' text a headline must be
  // TODO sync with DetectHeadline ??
  minHeadlineDistance: 1.5,
};

export const TOC_GLOBAL = new GlobalDefinition<TOC>('toc');
export const HEADLINE_TYPE_TO_HEIGHT_RANGE = new GlobalDefinition<Record<string, HeadlineRange>>(
  'headlineTypeToHeightRange',
);

export default class DetectToc extends ItemTransformer {
  constructor() {
    super(
      'Detect TOC',
      'Detect table of contents.',
      {
        requireColumns: ['x', 'y', 'str', 'line'],
        producesGlobels: [TOC_GLOBAL.key, HEADLINE_TYPE_TO_HEIGHT_RANGE.key],
        debug: {
          itemMerger: new LineItemMerger(),
        },
      },
      (incomingSchema) => {
        return incomingSchema.reduce((schema, column) => {
          if (column === 'x') {
            return [...schema, 'types', 'x'];
          }
          return [...schema, column];
        }, new Array<string>());
      },
    );
  }

  transform(context: TransformContext, inputItems: Item[]): ItemResult {
    const pageMapping = context.getGlobal(PAGE_MAPPING);
    const mostUsedHeight = context.getGlobal(MOST_USED_HEIGHT);
    const maxPageToEvaluate = Math.min(context.pageCount / 2, 5 + Math.abs(pageMapping.pageFactor));
    const pagesToEvaluate = groupByPage(inputItems.filter((item) => item.page <= maxPageToEvaluate));
    const maxPageToBeLinkedTo = context.pageCount + pageMapping.pageFactor - 1;

    const tocArea = findTocArea(pagesToEvaluate, context.pageCount, maxPageToBeLinkedTo);

    if (!tocArea) {
      return { items: inputItems, messages: ['No Table of Contents found!'] };
    }

    const itemsInTocArea = inputItems.filter((item) => tocArea.pages.includes(item.page));
    const rawTocEntries = selectRawTocEntries(tocArea, itemsInTocArea);
    const headlineLevels = findTocEntryHeadlineLevels(rawTocEntries);
    const tocItemUuids = new Set<string>(
      flatten(flatten(rawTocEntries.map((e) => e.entryLines))).map((item) => item.uuid),
    );

    const tocHeadline = findTocHeadline(context.fontMap, mostUsedHeight, tocArea, itemsInTocArea, tocItemUuids);

    const notFoundHeadlines: RawTocEntry[] = [];
    const foundHeadlines: Headline[] = [];
    const headlineTypeToHeightRange: Record<string, HeadlineRange> = {}; //H1={min:23, max:25}

    rawTocEntries.forEach((rawEntry, index) => {
      const itemType = headlineLevels[index];
      const uuids = findHeadline(
        context.fontMap,
        inputItems,
        mostUsedHeight,
        rawEntry.linkedPage,
        rawEntry.linkedPage - pageMapping.pageFactor,
        rawEntry.entryLines,
      );
      if (uuids) {
        foundHeadlines.push({ level: itemType, uuids });

        // add headline height
        const headlineHeight = inputItems
          .filter((item) => uuids.has(item.uuid))
          .reduce((maxHeight, item) => Math.max(maxHeight, item.data['height']), 0);
        let range = headlineTypeToHeightRange[itemType];
        if (range) {
          range.min = Math.min(range.min, headlineHeight);
          range.max = Math.max(range.max, headlineHeight);
        } else {
          range = {
            min: headlineHeight,
            max: headlineHeight,
          };
          headlineTypeToHeightRange[itemType] = range;
        }
      } else {
        notFoundHeadlines.push(rawEntry);
      }
    });

    const headlineUuidToLevelMap = foundHeadlines.reduce((uidToLevel, headline) => {
      headline.uuids.forEach((uuid) => {
        uidToLevel.set(uuid, headline.level);
      });
      return uidToLevel;
    }, new Map<string, HeadlineType>());

    const headlineTypes = foundHeadlines.reduce((allLevels, headline) => {
      allLevels.add(headline.level);
      return allLevels;
    }, new Set<HeadlineType>());
    const tocHeadlineUuids = new Set(tocHeadline.map((item) => item.uuid));

    return {
      items: inputItems
        .filter((item) => !tocHeadlineUuids.has(item.uuid))
        .filter((item) => !tocArea.pages.includes(item.page) || !tocItemUuids.has(item.uuid))
        .map((item) => {
          const itemType = headlineUuidToLevelMap.get(item.uuid);
          if (itemType) {
            return itemWithType(item, itemType);
          }
          return item;
        }),
      messages: [
        `Detected and removed ${rawTocEntries.length} TOC entries`,
        `Found ${foundHeadlines.length} matching headlines`,
      ],
      globals: [
        TOC_GLOBAL.value(new TOC(tocHeadline, tocArea.pages, headlineTypes)),
        HEADLINE_TYPE_TO_HEIGHT_RANGE.value(headlineTypeToHeightRange),
      ],
    };
  }
}

function findTocArea(pagesToEvaluate: Item[][], pageCount: number, maxPageToBeLinkedTo: number): TocArea | undefined {
  const linesWithNumber: LineWithNumber[] = [];
  pagesToEvaluate.forEach((pageItems) => {
    const itemsGroupedByLine = groupByLine(pageItems);
    itemsGroupedByLine.forEach((lineItems) => {
      const number = findEndingNumber(lineItems);
      if (
        number &&
        Number.isInteger(number) &&
        number > 0 &&
        number <= maxPageToBeLinkedTo &&
        lineItems.map((item) => item.data['str']).join('').length > config.linkMinLength
      ) {
        const page = lineItems[0].page;
        const startItemUuid = lineItems[0].uuid;
        const y = lineItems[0].data['y'];
        linesWithNumber.push({ page, startItemUuid, y, number });
      }
    });
  });

  if (linesWithNumber.length <= 0) {
    return undefined;
  }

  const lineNumberClusters = linesWithNumber.reduce(
    (arrayOfAscendingNumberArrays: LineWithNumber[][], lineWithNumber) => {
      if (arrayOfAscendingNumberArrays.length == 0) {
        return [[lineWithNumber]];
      }
      const lastArray = arrayOfAscendingNumberArrays[arrayOfAscendingNumberArrays.length - 1];
      const lastNumber = lastArray[lastArray.length - 1];
      if (lineWithNumber.number >= lastNumber.number) {
        lastArray.push(lineWithNumber);
      } else {
        arrayOfAscendingNumberArrays.push([lineWithNumber]);
      }
      return arrayOfAscendingNumberArrays;
    },
    [],
  );

  lineNumberClusters.sort((a, b) => b.length - a.length);
  if (lineNumberClusters[0].length < 3) {
    return undefined;
  }

  const selectedLines = lineNumberClusters[0];
  const pages = selectedLines.map((l) => l.page).filter(onlyUniques);
  if (!numbersAreConsecutive(pages)) {
    return undefined;
  }
  if (pages.length > selectedLines.length / 5) {
    return undefined;
  }

  return {
    pages,
    linesWithNumbers: selectedLines,
  };
}

function findEndingNumber(lineItems: Item[]): number | undefined {
  const text = lineItems
    .reduce((text, item) => {
      return text + item.data['str'];
    }, '')
    .trim();
  return extractEndingNumber(text);
}

function selectRawTocEntries(tocArea: TocArea, itemsInTocArea: Item[]): RawTocEntry[] {
  const numbersByStartUuid = tocArea.linesWithNumbers.reduce((map: Map<string, number>, l) => {
    map.set(l.startItemUuid, l.number);
    return map;
  }, new Map());

  const itemsInTocAreaByLine = groupByLine(itemsInTocArea);

  const maxHeightOfNumberedLines = Math.max(
    ...itemsInTocAreaByLine
      .reduce((lineHeights: number[], lineItems) => {
        if (numbersByStartUuid.has(lineItems[0].uuid)) {
          lineHeights.push(Math.max(...lineItems.map((line) => line.data['height'])));
        }
        return lineHeights;
      }, [])
      .filter(onlyUniques),
  );
  const maxLinesBetweenLinesWithNumbers = Math.max(
    ...itemsInTocAreaByLine
      .reduce((lineDistance: number[], lineItems) => {
        if (numbersByStartUuid.has(lineItems[0].uuid)) {
          lineDistance.push(-1);
        }
        if (lineDistance.length > 0) {
          lineDistance[lineDistance.length - 1]++;
        }
        return lineDistance;
      }, [])
      .filter(onlyUniques),
  );
  const linesWithNumbersByPage = groupBy(tocArea.linesWithNumbers, (line) => line.page);
  const maxYBetweenLinesWithNumbers = Math.max(
    ...linesWithNumbersByPage.map((pageLines) => {
      return pageLines.reduce(
        (previous: { y: number; distance: number }, line) => {
          const y = line.y;
          if (previous.y == -1) {
            return { y, distance: -1 };
          }
          return {
            y,
            distance: Math.max(Math.abs(y - previous.y), previous.distance),
          };
        },
        { y: -1, distance: -1 },
      ).distance;
    }),
  );
  const rawTocEntries: RawTocEntry[] = [];
  itemsInTocAreaByLine.reduce((beforeLines: Item[][], lineItems) => {
    const number = numbersByStartUuid.get(lineItems[0].uuid);
    if (!number) {
      beforeLines.push(lineItems);
      return beforeLines;
    }
    const validBeforeLines = beforeLines.filter((beforLine, beforeIndex) => {
      const yDistance = Math.abs(beforLine[0].data['y'] - lineItems[0].data['y']);
      const beforLineHeight = Math.max(...beforLine.map((item) => item.data['height']));
      const beforeLineMuchLarger = beforLineHeight > maxHeightOfNumberedLines;
      return (
        !beforeLineMuchLarger &&
        beforeLines.length - beforeIndex <= maxLinesBetweenLinesWithNumbers &&
        yDistance <= maxYBetweenLinesWithNumbers
      );
    });
    const entryLines = [...validBeforeLines, lineItems];
    rawTocEntries.push({
      linkedPage: number,
      entryLines,
    });
    return [];
  }, []);

  return rawTocEntries;
}

function findTocHeadline(
  fontMap: Map<string, object>,
  mostUsedHeight: number,
  tocArea: TocArea,
  itemsInTocArea: Item[],
  tocItemUuids: Set<string>,
): Item[] {
  const firstPageNonTocItems = itemsInTocArea
    .filter((item) => item.page == tocArea.pages[0])
    .filter((item) => !tocItemUuids.has(item.uuid));
  const itemsGroupedByLine = groupByLine(firstPageNonTocItems).filter((lineItems) =>
    hasHeadlineSymptoms(fontMap, mostUsedHeight, lineItems),
  );
  if (itemsGroupedByLine.length == 0) {
    return [];
  }
  return itemsGroupedByLine[itemsGroupedByLine.length - 1];
}

function findTocEntryHeadlineLevels(rawTocEntries: RawTocEntry[]): HeadlineType[] {
  // We focus on heights since it seems the most consistent metric to determining levels so far.
  //Other options would be looking at X-coordinates (per page), or at leading numbering (e.g. /^(\d)+.(\d)+.(\d)+/).
  const height = (entry: RawTocEntry) => Math.round(getHeight(entry.entryLines[0][0]));
  const allHeights = rawTocEntries.map(height).filter(onlyUniques).sort(descending);

  // we start with H2 (H1 is reserved for the document title)
  if (allHeights.length > 3) {
    return rawTocEntries.map(() => 'H2');
  }

  return rawTocEntries.map((entry) => {
    const index = allHeights.indexOf(height(entry));
    return toHeadlineType(index + 2);
  });
}

/**
 * @param fontMap
 * @param items
 * @param mostUsedHeight
 * @param targetPage
 * @param targetPageIndex
 * @param entryLines
 * @returns set of uuids
 */
function findHeadline(
  fontMap: Map<string, object>,
  items: Item[],
  mostUsedHeight: number,
  targetPage: number,
  targetPageIndex: number,
  entryLines: Item[][],
): Set<string> | undefined {
  const tocEntryText = normalizeHeadlineChars(entryLines)
    .replace(new RegExp(targetPage + '$', 'g'), '')
    .replace(new RegExp('\\.\\.*$', 'g'), '');
  const pageItems = items.filter((item) => item.page == targetPageIndex);
  const canditate = fineMatchingHeadlineCanditate(tocEntryText, pageItems, fontMap, mostUsedHeight);
  if (canditate.length > 0) {
    return canditate.reduce((itemUuids, lineItems) => {
      lineItems.forEach((item) => itemUuids.add(item.uuid));
      return itemUuids;
    }, new Set<string>());
  }
  return undefined;
}

function fineMatchingHeadlineCanditate(
  tocEntryText: string,
  pageItems: Item[],
  fontMap: Map<string, object>,
  mostUsedHeight: number,
): Item[][] {
  const itemsByLine = groupByLine(pageItems);
  let headlineCanditates: { score: number; lines: Item[][] }[] = [];
  let currentLines: Item[][] = [];
  let currentScore = 0;
  let currentText = '';
  for (let lineIdx = 0; lineIdx < itemsByLine.length; lineIdx++) {
    const lineItems = itemsByLine[lineIdx];
    const lineText = normalizeHeadlineChars([lineItems]);
    const lineInLink = tocEntryText.includes(lineText);
    const headlineSymptoms = hasHeadlineSymptoms(fontMap, mostUsedHeight, lineItems);
    if (lineInLink && headlineSymptoms) {
      const newText = currentText + lineText;
      const newScore = compareTwoStrings(newText, tocEntryText);
      if (newScore > currentScore) {
        currentScore = newScore;
        currentText = newText;
        currentLines.push(lineItems);
        if (newScore == 1) {
          return currentLines;
        }
      } else if (currentScore > 0.95) {
        return currentLines;
      }
    } else {
      if (currentLines.length > 0) {
        headlineCanditates.push({ score: currentScore, lines: currentLines });
        currentLines = [];
        currentScore = 0;
        currentText = '';
      }
    }
  }
  // console.log(
  //   'headlineCanditates',
  //   tocEntryText,
  //   pageItems[0].page,
  //   headlineCanditates
  //     .sort((a, b) => a.score - b.score)
  //     .map((canditate) => canditate.score + ': ' + joinText(flatten(canditate.lines), '')),
  // );
  headlineCanditates = headlineCanditates.filter((candidate) => candidate.score > 0.5);
  if (headlineCanditates.length == 0) {
    return [];
  }

  return headlineCanditates.sort((a, b) => a.score - b.score)[0].lines;
}

function hasHeadlineSymptoms(fontMap: Map<string, object>, mostUsedHeight: number, lineItems: Item[]): boolean {
  return (
    getHeight(lineItems[0]) >= mostUsedHeight + config.minHeadlineDistance ||
    declaredFontTypes(getFontName(fontMap, lineItems[0])).includes(FontType.BOLD)
  );
}

function normalizeHeadlineChars(lines: Item[][]) {
  const text = flatten(lines)
    .map((item) => getText(item))
    .join('');
  return filterOut(text, [
    WHITESPACE_CHAR_CODE,
    TAB_CHAR_CODE,
    ...DASHS_CHAR_CODES,
    ...PERIOD_CHAR_CODES,
  ]).toLowerCase();
}

/**
 * Pointer to pages/items which classified as TOC.
 */
interface TocArea {
  pages: number[];
  linesWithNumbers: LineWithNumber[];
}

/**
 * Contains the page number and all attached lines of items,
 */
interface RawTocEntry {
  linkedPage: number;
  entryLines: Item[][];
}

interface Headline {
  level: HeadlineType;
  uuids: Set<string>;
}

/**
 * A (item[]) line which ends with a number.
 */
interface LineWithNumber {
  page: number;
  startItemUuid: string;
  y: number;
  number: number;
}
