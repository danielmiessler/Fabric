import React from 'react';

import Table from 'react-bootstrap/lib/Table'

// Displays an array of LineItem as a table
export default class LineItemTable extends React.Component {

    static propTypes = {
        items: React.PropTypes.array.isRequired,
        showWhitespaces: React.PropTypes.bool
    };

    render() {
        const {showWhitespaces, items} = this.props;
        const tableHeader = <thead>
                              <tr>
                                <th>
                                  #
                                </th>
                                <th>
                                  Text
                                </th>
                                <th>
                                  X
                                </th>
                                <th>
                                  Y
                                </th>
                                <th>
                                  Width
                                </th>
                                <th>
                                  Height
                                </th>
                              </tr>
                            </thead>

        const itemRows = items.map((item, i) => <tr key={ i } style={ item.annotation ? {
                          color: item.annotation.color
                      } : null }>
                                                  <td>
                                                    <div style={ { textAlign: 'center' } }>
                                                      { i }
                                                    </div>
                                                    <div style={ { textAlign: 'center' } }>
                                                      { item.annotation ? item.annotation.category : '' }
                                                    </div>
                                                    <div style={ { textAlign: 'center', color: 'brown' } }>
                                                      { item.type ? item.type.name : '' }
                                                    </div>
                                                    <div style={ { textAlign: 'center', color: 'orange' } }>
                                                      { item.parsedElements && item.parsedElements.footnoteLinks.length > 0 ? <div>
                                                                                                                                Footnote-Link
                                                                                                                              </div> : '' }
                                                      { item.parsedElements && item.parsedElements.containLinks ? <div>
                                                                                                                    Link
                                                                                                                  </div> : '' }
                                                      { item.lineFormat ? <div>
                                                                            { item.lineFormat.name }
                                                                          </div> : '' }
                                                      { item.unopenedFormat ? <div>
                                                                                Unopened
                                                                                { ' ' + item.unopenedFormat.name }
                                                                              </div> : '' }
                                                      { item.parsedElements && item.parsedElements.inlineFormats > 0 ? <div>
                                                                                                                         { item.parsedElements.inlineFormats + 'x Bold/Italic' }
                                                                                                                       </div> : '' }
                                                      { item.unclosedFormat ? <div>
                                                                                Unclosed
                                                                                { ' ' + item.unclosedFormat.name }
                                                                              </div> : '' }
                                                    </div>
                                                  </td>
                                                  <td>
                                                    { showWhitespaces ? (
                                                      <pre style={ item.annotation ? {
                                                                       color: item.annotation.color,
                                                                       display: 'inline-block',
                                                                   } : {
                                                                       display: 'inline-block'
                                                                   } }>{ item.text() }</pre>
                                                      ) : (item.text()) }
                                                  </td>
                                                  <td>
                                                    { item.x }
                                                  </td>
                                                  <td>
                                                    { item.y }
                                                  </td>
                                                  <td>
                                                    { item.width }
                                                  </td>
                                                  <td>
                                                    { item.height }
                                                  </td>
                                                </tr>
        )

        return (
            <Table responsive condensed bordered>
              { tableHeader }
              <tbody>
                { itemRows }
              </tbody>
            </Table>
        );
    }
}