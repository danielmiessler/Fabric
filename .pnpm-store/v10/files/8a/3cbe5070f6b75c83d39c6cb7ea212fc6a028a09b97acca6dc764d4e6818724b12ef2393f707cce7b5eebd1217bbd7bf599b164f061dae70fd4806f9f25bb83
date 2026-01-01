import React from 'react';
import Transformation from './Transformation.jsx';
import ParseResult from '../ParseResult.jsx';
import LineItem from '../LineItem.jsx';
import LineItemPageView from '../../components/debug/LineItemPageView.jsx';
import { REMOVED_ANNOTATION } from '../Annotation.jsx';

// Abstract class for transformations producing LineItem(s) to be shown in the LineItemPageView
export default class ToLineItemTransformation extends Transformation {

    constructor(name) {
        super(name, LineItem.name);
        if (this.constructor === ToLineItemTransformation) {
            throw new TypeError("Can not construct abstract class.");
        }
        this.showWhitespaces = false;
    }

    showModificationCheckbox() {
        return true;
    }

    createPageView(page, modificationsOnly) {
        return <LineItemPageView
                                 key={ page.index }
                                 page={ page }
                                 modificationsOnly={ modificationsOnly }
                                 showWhitespaces={ this.showWhitespaces } />;
    }

    completeTransform(parseResult:ParseResult) {
        // The usual cleanup
        parseResult.messages = [];
        parseResult.pages.forEach(page => {
            page.items = page.items.filter(item => !item.annotation || item.annotation !== REMOVED_ANNOTATION);
            page.items.forEach(item => item.annotation = null);
        });
        return parseResult;
    }


}