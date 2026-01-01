export default class Metadata {
  original: object;

  constructor(original: object) {
    this.original = original;
  }

  title() {
    return this.extract('Title', 'dc:title');
  }

  author() {
    return this.extract('Author', 'dc:creator');
  }

  private extract(infoName: string, metadataKey: string) {
    const metadata = this.original['metadata'];
    if (metadata) {
      return metadata.get(metadataKey);
    }
    return this.original['info'][infoName];
  }
}
