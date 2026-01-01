export type FontType = 'BOLD' | 'OBLIQUE';
export type TokenType = 'LINK' | 'FOOTNOTE' | 'FOOTNOTE_LINK' | FontType;

export const startSymbol = (type: TokenType) => {
  switch (type) {
    case 'BOLD':
      return '**';
    case 'OBLIQUE':
      return '_';
    default:
      return '';
    //   throw new Error(`No start symbol defined for type: ${type}`);
  }
};

export const endSymbol = (type: TokenType) => {
  switch (type) {
    case 'BOLD':
      return '**';
    case 'OBLIQUE':
      return '_';
    default:
      return '';
    //   throw new Error(`No end symbol defined for type: ${type}`);
  }
};

export const attachWithoutWhitespace = (type: TokenType) => {
  switch (type) {
    case 'FOOTNOTE_LINK':
      return true;
    default:
      return false;
  }
};

export const plainTextFormat = (type: TokenType) => {
  switch (type) {
    case 'FOOTNOTE_LINK':
      return true;
    default:
      return false;
  }
};

export function tokenToText(token: string, type: TokenType) {
  switch (type) {
    case 'LINK':
      return `[${token}](${token})`;
    case 'FOOTNOTE_LINK':
      return `[^${token}]`;
    case 'FOOTNOTE':
      return `[^${token}]:`;
    default:
      return token;
  }
}
