import { v4 as uuidv4 } from 'uuid';
import { TokenType } from './token-types';

export default class Item {
  page: number;
  data: object;
  uuid: string;
  listLevel = 0;
  tokenTypes: TokenType[] = [];

  constructor(page: number, data: object, tokenTypes: TokenType[] = [], uuid: string = uuidv4()) {
    this.page = page;
    this.data = data;
    this.uuid = uuid;
    this.tokenTypes = tokenTypes;
  }

  value(column: string): object {
    return this.data[column];
  }

  withTokenType(tokenType: TokenType): Item {
    const newItem = new Item(this.page, this.data, this.tokenTypes, this.uuid);
    newItem.tokenTypes.push(tokenType);
    return newItem;
  }

  withTokenTypes(tokenTypes: TokenType[]): Item {
    const newItem = new Item(this.page, this.data, this.tokenTypes, this.uuid);
    tokenTypes.forEach((tt) => newItem.tokenTypes.push(tt));
    return newItem;
  }

  withDataAddition(data: object): Item {
    return this.withData({ ...this.data, ...data });
  }

  withData(data: object): Item {
    return new Item(this.page, data, this.tokenTypes, this.uuid);
  }

  /**
   * Returns the item without a uuid.
   */
  withoutUuid(): Item {
    return new Item(this.page, this.data, this.tokenTypes, '');
  }
}
