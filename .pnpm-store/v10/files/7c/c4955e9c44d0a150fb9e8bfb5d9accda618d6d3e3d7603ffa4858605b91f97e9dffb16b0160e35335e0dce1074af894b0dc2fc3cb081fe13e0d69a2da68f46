// Metadata of the PDF document
export default class Metadata {

    constructor(originalMetadata) {
        if (originalMetadata.metadata) {
            this.title = originalMetadata.metadata.get('dc:title');
            this.creator = originalMetadata.metadata.get('xap:creatortool')
            this.producer = originalMetadata.metadata.get('pdf:producer')
        } else {
            this.title = originalMetadata.info.Title;
            this.author = originalMetadata.info.Author;
            this.creator = originalMetadata.info.Creator;
            this.producer = originalMetadata.info.Producer;
        }
    }

}
