import React from 'react';
import PageView from './PageView.jsx';
import Remarkable from 'remarkable';

export default class MarkdownPageView extends PageView {

    createItemViews(items, showWhitespaces) { // eslint-disable-line no-unused-vars
        const remarkable = new Remarkable({
            breaks: true
        });
        const html = remarkable.render(items[0]);
        return <div>
                 <div dangerouslySetInnerHTML={ { __html: html } } />
               </div>
    }

}