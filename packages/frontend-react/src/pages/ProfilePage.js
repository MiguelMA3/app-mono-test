import React, { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router-dom';
import api from '../services/api';

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
    return <div>Carregando perfil...</div>;
  }

  if (error) {
    return <div style={{ color: 'red' }}>{error}</div>;
  }

  return (
    <div>
      {user && (
        <>
          <h1>Perfil de {user.username}</h1>
          <p><strong>Email:</strong> {user.email}</p>
          <p><strong>Bio:</strong> {user.bio}</p>
          <p><strong>Membro desde:</strong> {new Date(user.createdAt).toLocaleDateString()}</p>
        </>
      )}

      <hr />
      <h2>Posts de {user?.username}</h2>
      
      <div>
        {posts.length > 0 ? (
          posts.map(post => (
            <div key={post.id} style={{ border: '1px solid #ccc', margin: '10px', padding: '10px' }}>
              <p><strong>
                <Link to={`/profile/${post.user.id}`}>
                  {post.user.username}
                </Link>
              </strong></p>
              
              <p>{post.content}</p>
              <small>Postado em: {new Date(post.createdAt).toLocaleString()}</small>
              <div>
                <button onClick={() => handleLikePost(post.id)}>
                  Curtir ({post.likes})
                </button>
              </div>
            </div>
          ))
        ) : (
          <p>Este usuário ainda não fez nenhuma postagem.</p>
        )}
      </div>
    </div>
  );
}

export default ProfilePage;