import React from 'react';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';

import LoginPage from './pages/LoginPage';
import TimelinePage from './pages/TimelinePage';
import ProfilePage from './pages/ProfilePage';
import './App.css';

const PrivateRoute = ({ children }) => {
  const token = localStorage.getItem('token');
  return token ? children : <Navigate to="/login" />;
};

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<LoginPage />} />

        <Route
          path="/"
          element={<PrivateRoute><TimelinePage /></PrivateRoute>}
        />

        <Route
          path="/profile"
          element={<PrivateRoute><ProfilePage /></PrivateRoute>}
        />
        
        <Route path="*" element={<Navigate to="/" />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;