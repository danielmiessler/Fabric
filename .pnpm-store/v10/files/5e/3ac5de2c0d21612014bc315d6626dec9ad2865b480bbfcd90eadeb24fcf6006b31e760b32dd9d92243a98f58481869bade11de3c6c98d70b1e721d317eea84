import React from 'react';
import PageView from './PageView.jsx';
import LineItemTable from './LineItemTable.jsx';

// View for a Page which items are of kind LineItemBlock
export default class LineItemBlockPageView extends PageView {

    createItemViews(items, showWhitespaces) {
        const blockTables = items.map((block, i) => {
            const blockType = block.type ? ' - ' + block.type.name : null;
            const blockAnnotation = block.annotation ? <span>{ ' - ' + block.annotation.category }</span>
                : null;
            const borderStyle = block.annotation ? {
                marginBottom: "20px",
                border: "solid thin " + block.annotation.color
            } : null;
            const colorStyle = block.annotation ? {
                color: block.annotation.color
            } : null;
            var footnoteLinks;
            var footnotes;
            if (block.parsedElements) {
                if (block.parsedElements.footnoteLinks.length > 0) {
                    footnoteLinks = <div>
                                      { 'Footnote-Links: ' + block.parsedElements.footnoteLinks }
                                    </div>;
                }
                if (block.parsedElements.footnotes.length > 0) {
                    footnotes = <div>
                                  { 'Footnotes: ' + block.parsedElements.footnotes }
                                </div>;
                }
            }

            return <div key={ i }>
                     <div style={ colorStyle }>
                       <b>Block { i + 1 }</b><i>{ blockType } { blockAnnotation }</i>
                     </div>
                     <div style={ borderStyle }>
                       <LineItemTable items={ block.items } showWhitespaces={ showWhitespaces } />
                       { footnoteLinks }
                       { footnotes }
                     </div>
                   </div>
        });
        return blockTables;
    }

}