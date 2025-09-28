import React, { useState } from 'react';
import { BrowserRouter, Routes, Route, Navigate, Link, useNavigate } from 'react-router-dom';

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

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('userId');
    setIsLoggedIn(false);
  };

  return (
    <BrowserRouter>
      <div className="container">
        {isLoggedIn && <Nav onLogout={handleLogout} />}
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
      </div>
    </BrowserRouter>
  );
}

const Nav = ({ onLogout }) => {
  const userId = localStorage.getItem('userId');
  const navigate = useNavigate();

  const handleLogoutClick = () => {
    onLogout();
    navigate('/login');
  };

  return (
    <nav>
      <Link to="/">Timeline</Link>
      | <Link to={`/profile/${userId}`}>Meu Perfil</Link>
      | <button onClick={handleLogoutClick} style={{ marginLeft: '10px', background: 'none', border: 'none', color: 'blue', cursor: 'pointer', textDecoration: 'underline' }}>
        Sair
      </button>
    </nav>
  );
};

export default App;