import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom'
import api from '../services/api';

function LoginPage({ onLoginSuccess }) {
    const [username, setUsername] = useState('admin');
    const [password, setPassword] =useState('1234');
    const [error, setError] = useState('');

    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault();

        if (!username || !password) {
            setError('Preencha o usuario e a senha.');
            return;
        }

        try {
            const response = await api.post('/login', { username, password });
            const { token, userId } = response.data;

            localStorage.setItem('token', token);
            localStorage.setItem('userId', userId);

            onLoginSuccess();

            setError('');
            navigate('/');
        } catch (err) {
            setError('Usuario ou senha invalidos. Tente novamente');
            console.error('Falha no login', err);
        }
    };
    return (
        <div>
      <h2>Login</h2>
      <form onSubmit={handleLogin}>
        <div>
          <label>Usuário:</label>
          <input type="text" value={username} onChange={(e) => setUsername(e.target.value)} />
        </div>
        <div>
          <label>Senha:</label>
          <input
            type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
        </div>
        <button type="submit">Entrar</button>
      </form>
      {error && <p style={{ color: 'red' }}>{error}</p>}
    </div>
    );
}

export default LoginPage;