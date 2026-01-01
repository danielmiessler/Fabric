import React from 'react';
import PageView from './PageView.jsx';
import Table from 'react-bootstrap/lib/Table'

export default class TextPageView extends PageView {

    createItemViews(items, showWhitespaces) { // eslint-disable-line no-unused-vars
        return <div>
                 <Table responsive>
                   <thead>
                     <tr>
                       <th>
                         #
                       </th>
                       <th>
                         Category
                       </th>
                       <th>
                         Text
                       </th>
                     </tr>
                   </thead>
                   <tbody>
                     { items.map((block, i) => <tr key={ i }>
                                                 <td>
                                                   { i }
                                                 </td>
                                                 <td>
                                                   { block.category }
                                                 </td>
                                                 <td>
                                                   <pre style={ { display: 'inline-block' } }>{ block.text }</pre>
                                                 </td>
                                               </tr>
                       ) }
                   </tbody>
                 </Table>
               </div>
    }

}