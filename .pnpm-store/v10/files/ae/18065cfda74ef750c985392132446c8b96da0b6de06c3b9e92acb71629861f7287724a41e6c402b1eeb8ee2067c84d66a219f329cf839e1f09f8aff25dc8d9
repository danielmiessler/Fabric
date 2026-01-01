import React from 'react';
import MarkdownPageView from '../../components/debug/MarkdownPageView.jsx';
import Transformation from './Transformation.jsx';
import ParseResult from '../ParseResult.jsx';

export default class ToMarkdown extends Transformation {

    constructor() {
        super("To Markdown", "String");
    }

    createPageView(page, modificationsOnly) { // eslint-disable-line no-unused-vars
        return <MarkdownPageView key={ page.index } page={ page } />;
    }

    transform(parseResult:ParseResult) {
        parseResult.pages.forEach(page => {
            var text = '';
            page.items.forEach(block => {
                text += block.text + '\n';
            });
            page.items = [text];
        });
        return new ParseResult({
            ...parseResult,
        });
    }

}