import { groupByBlock, groupByLine } from '../support/groupingUtils';
import Item from '../Item';
import { Converter } from '../convert';
import { TextType, discardTokenTypes, headlineLevel, isHeadline } from '../text-types';
import {
  TokenType,
  attachWithoutWhitespace,
  endSymbol,
  plainTextFormat,
  startSymbol,
  tokenToText,
} from '../token-types';

export default class MarkdownConverter implements Converter {
  convert(items: Item[]) {
    let content = '';

    groupByBlock(items).forEach((blockItems) => {
      const blockTypes: TextType[] = blockItems[0].data['types'] || [];
      let blockContent = '';
      groupByLine(blockItems).forEach((lineItems) => {
        blockContent += lineToText(lineItems, blockTypes);
        blockContent += '\n';
      });
      content += elementToText(blockContent, blockTypes[0]);
    });

    return content;
  }
}

function elementToText(text: string, type: TextType) {
  switch (type) {
    case 'H1':
    case 'H2':
    case 'H3':
    case 'H4':
    case 'H5':
    case 'H6':
      return '#'.repeat(headlineLevel(type)) + ' ' + text + '\n';
    case 'CODE':
      return '```\n' + text.trim() + '\n```\n\n';
    case 'FOOTNOTES':
      return text;
    default:
      return text + '\n';
  }
}

function toWords(text: string): string[] {
  return text.split(' ').filter((string) => string.trim().length > 0);
}

export function lineToText(lineItems: Item[], blockTypes: TextType[]) {
  let text = '';
  let openFormat: TokenType;

  const closeFormat = () => {
    text += endSymbol(openFormat);
    openFormat = null;
  };

  let lastLineItem: Item = null;
  lineItems.forEach((lineItem, indexInLine) => {
    const words = toWords(lineItem.data['str']);
    words.forEach((word, wordIndex) => {
      if (indexInLine === 0 && wordIndex === 0) {
        if (lineItem.listLevel) {
          word = ' '.repeat(lineItem.listLevel * 2) + word;
        }
      }
      const wordFormat = lineItem.tokenTypes[0]; // bold, oblique, footnote etc...
      if (openFormat && (!wordFormat || wordFormat !== openFormat)) {
        closeFormat();
      }
      if (
        (wordIndex > 0 || indexInLine > 0) &&
        !(wordFormat && attachWithoutWhitespace(wordFormat)) &&
        !isPunctationCharacter(word)
      ) {
        let insertWhitespace = true;
        if (indexInLine > 0 && wordIndex == 0) {
          const xDistance = lineItem.data['x'] - lastLineItem.data['x'] - lastLineItem.data['width'];
          if (xDistance < 2 && !lastLineItem.data['str']?.endsWith(' ') && !lineItem.data['str']?.startsWith(' ')) {
            insertWhitespace = false;
          }
        }
        if (insertWhitespace) {
          text += ' ';
        }
      }

      const disableInlineFormats =
        wordFormat && (discardTokenTypes(blockTypes) || (blockTypes.find(isHeadline) && wordFormat === 'BOLD'));
      if (wordFormat && !openFormat && !disableInlineFormats) {
        openFormat = wordFormat;
        text += startSymbol(openFormat);
      }

      if (wordFormat && (!disableInlineFormats || plainTextFormat(wordFormat))) {
        text += tokenToText(word, wordFormat);
      } else {
        text += word;
      }
    });
    if (openFormat && (indexInLine == lineItems.length - 1 || firstFormat(lineItems[indexInLine + 1]) !== openFormat)) {
      closeFormat();
    }
    lastLineItem = lineItem;
  });

  return text;
}

function firstFormat(lineItem: Item) {
  if (lineItem.data['str'].length == 0) {
    return null;
  }
  return lineItem.tokenTypes[0];
}

function isPunctationCharacter(value: string) {
  if (value.length != 1) {
    return false;
  }
  return value[0] === '.' || value[0] === '!' || value[0] === '?';
}
