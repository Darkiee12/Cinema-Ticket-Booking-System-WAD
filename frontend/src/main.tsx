import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter as Router } from 'react-router-dom';
import App from './App.tsx';
import './index.css';

import createStore from 'react-auth-kit/createStore';
import AuthProvider from 'react-auth-kit';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <AuthProvider
      store={createStore({
        authName: '_auth',
        authType: 'cookie',
        cookieDomain: window.location.hostname,
        cookieSecure: window.location.protocol === 'https:',
      })}
    >
      <Router>
        <App />
      </Router>
    </AuthProvider>
  </React.StrictMode>,
);
