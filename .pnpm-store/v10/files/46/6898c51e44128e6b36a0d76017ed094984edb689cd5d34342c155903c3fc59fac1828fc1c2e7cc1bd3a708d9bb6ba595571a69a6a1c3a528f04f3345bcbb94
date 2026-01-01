import React from 'react';
import ReactDOM from 'react-dom';

import 'bootstrap/dist/css/bootstrap.css';
import '../css/styles.css';

import App from './components/App.jsx';
import AppState from './models/AppState.jsx';

function render(appState) {
    ReactDOM.render(<App appState={ appState } />, document.getElementById('main'));
}

const appState = new AppState({
    renderFunction: render,
});

appState.render()
