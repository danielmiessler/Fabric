import React from 'react';
import Transformation from './Transformation.jsx';
import ParseResult from '../ParseResult.jsx';
import LineItemBlock from '../LineItemBlock.jsx';
import LineItemBlockPageView from '../../components/debug/LineItemBlockPageView.jsx';
import { REMOVED_ANNOTATION } from '../Annotation.jsx';

// Abstract class for transformations producing LineItemBlock(s) to be shown in the LineItemBlockPageView
export default class ToLineItemBlockTransformation extends Transformation {

    constructor(name) {
        super(name, LineItemBlock.name);
        if (this.constructor === ToLineItemBlockTransformation) {
            throw new TypeError("Can not construct abstract class.");
        }
        this.showWhitespaces = false;
    }

    showModificationCheckbox() {
        return true;
    }

    createPageView(page, modificationsOnly) {
        return <LineItemBlockPageView
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