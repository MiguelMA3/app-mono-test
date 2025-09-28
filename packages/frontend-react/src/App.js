import React, { useState } from 'react';
import { BrowserRouter, Routes, Route, Navigate, Link } from 'react-router-dom';

import LoginPage from './pages/LoginPage';
import TimelinePage from './pages/TimelinePage';
import ProfilePage from './pages/ProfilePage';
import './App.css';

const PrivateRoute = ({ children }) => {
  const token = localStorage.getItem('token');
  return token ? children : <Navigate to="/login" />;
};

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(!!localStorage.getItem('token'));

  const handleLoginSuccess = () => {
    setIsLoggedIn(true);
  };
  return (
    <BrowserRouter>
      {isLoggedIn && <Nav />}
      <Routes>
        <Route path="/login" element={<LoginPage onLoginSuccess={handleLoginSuccess} />} />

        <Route
          path="/"
          element={<PrivateRoute><TimelinePage /></PrivateRoute>}
        />
        <Route
          path="/profile/:userId"
          element={<PrivateRoute><ProfilePage /></PrivateRoute>}
        />
        <Route path="*" element={<Navigate to="/" />} />
      </Routes>
    </BrowserRouter>
  );
}

const Nav = () => {
  const userId = localStorage.getItem('userId');
  return (
    <nav>
      <Link to="/">Timeline</Link> 
      | <Link to={`/profile/${userId}`}>Meu Perfil</Link> 
    </nav>
  );
};

export default App;