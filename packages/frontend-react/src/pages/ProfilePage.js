import React, { useState, useEffect } from 'react';
import api from '../services/api';

function ProfilePage() {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');

    useEffect(() => {
        const fetchUserData = async () => {
            const userId = localStorage.getItem('userId');

            if (!userId) {
                setError('ID do usuario nao encontrado. Necessario novo login');
                setLoading(false);
                return;
            }

            try {
                const response = await api.get(`/users/${userId}`);
                setUser(response.data);
            } catch (err) {
                setError('Não foi possível carregar os dados do perfil.');
                console.error(err);
            } finally {
                setLoading(false);
            }
        };

        fetchUserData();
    }, []);

    if (loading) {
        return <div>Carregando perfil...</div>;
    }

    if (error) {
    return <div style={{ color: 'red' }}>{error}</div>;
  }

  if (!user) {
    return <div>Nenhum dado de usuário para exibir.</div>;
  }

  return (
    <div>
      <h1>Perfil de {user.username}</h1>
      <p><strong>Email:</strong> {user.email}</p>
      <p><strong>Bio:</strong> {user.bio}</p>
      <p><strong>Membro desde:</strong> {new Date(user.createdAt).toLocaleDateString()}</p>
    </div>
  );
}

export default ProfilePage;