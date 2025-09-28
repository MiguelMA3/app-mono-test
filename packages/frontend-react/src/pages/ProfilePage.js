import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import api from '../services/api';
import Post from '../components/Post/Post';

function ProfilePage() {
  const [user, setUser] = useState(null);
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  const { userId } = useParams();

  useEffect(() => {
    const fetchProfileData = async () => {
      setLoading(true);
      setError('');
      try {
        const [userResponse, postsResponse] = await Promise.all([
          api.get(`/users/${userId}`),
          api.get(`/users/${userId}/posts`)
        ]);
        
        setUser(userResponse.data);
        setPosts(postsResponse.data || []);
      } catch (err) {
        setError('Não foi possível carregar os dados do perfil.');
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    fetchProfileData();
  }, [userId]);

  const handleLikePost = async (postId) => {
    try {
      const response = await api.post(`/posts/${postId}/like`);
      const updatedPosts = posts.map(post =>
        post.id === postId ? response.data : post
      );
      setPosts(updatedPosts);
    } catch (error) {
      console.error("Falha ao curtir o post", error);
    }
  };

  if (loading) {
    return <div className="page-content">Carregando perfil...</div>;
  }

  if (error) {
    return <div className="page-content" style={{ color: 'red' }}>{error}</div>;
  }

  return (
    <div className="page-content">
      {user && (
        <>
          <h1>Perfil de {user.username}</h1>
          <p><strong>Bio:</strong> {user.bio}</p>
          <p><strong>Membro desde:</strong> {new Date(user.createdAt).toLocaleDateString()}</p>
        </>
      )}

      <hr style={{margin: '20px 0'}}/>
      <h2>Posts de {user?.username}</h2>
      
      <div>
        {posts.length > 0 ? (
          posts.map(post => (
            <Post key={post.id} post={post} onLike={handleLikePost} />
          ))
        ) : (
          <p>Este usuário ainda não fez nenhuma postagem.</p>
        )}
      </div>
    </div>
  );
}

export default ProfilePage;