import Page from './models/Page.js';

// Holds the state of the application
export default {

    state: {
        uploaded: false,
        pagesToUpload: 0,
        uploadedPages: 0,
        rawPages: [],
        pages: [],
    },

    preparePageUpload: function(numPages) {
        this.state.pagesToUpload = numPages;
        for (var i = 0; i <= numPages; i++) {
            this.state.rawPages.push(new Page({
                index: i
            }));
        }
    },

    uploadPage: function(pageIndex, textItems) {
        this.state.rawPages[pageIndex].textItems = textItems;
        this.state.uploadedPages++;
        if (this.state.uploadedPages == this.state.pagesToUpload) {
            this.state.rawPages.map(rawPage => {
                var text = '';
                var line;
                var lineY;
                // console.debug("Page " + rawPage.index + "-------");
                rawPage.textItems.forEach(textItem => {
                    // console.debug(textItem);
                    const yRounded = Math.round(textItem.y);
                    if (!line) {
                        // console.debug("First line: " + textItem.text);
                        lineY = yRounded;
                        line = textItem.text;
                    } else {
                        if (yRounded === lineY) {
                            // console.debug("Add to line: " + line + " / " + textItem.text);
                            line += textItem.text;
                        } else {
                            // console.debug("Start line: " + line + " / " + textItem.text);
                            text += line + '\n';
                            line = textItem.text;
                            lineY = yRounded;
                        }
                    }
                });
                text += line;
                this.state.pages.push(text);
            });
            // this.state.pages = pages;
            this.state.uploaded = true;
        }
    },

    upload: function(pages) {
        // console.debug("Store: upload");
        this.state.uploaded = true;
        this.state.pages = pages;
    },

}