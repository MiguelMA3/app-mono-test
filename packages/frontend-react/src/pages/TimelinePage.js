import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import api from '../services/api';

function TimelinePage() {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState([]);
  const [newPostContent, setNewPostContent] = useState('');

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const response = await api.get('/posts');
        setPosts(response.data || []);
      } finally {
        setLoading(false);
      }
    };
    fetchPosts();
  }, []);

  const handleCreatePost = async (e) => {
    e.preventDefault();
    if (!newPostContent.trim()) return;

    try {
      const response = await api.post('/posts', {
        content:newPostContent,
      });
      setPosts([response.data, ...posts]);
      setNewPostContent('');
    } catch (error) {
      console.error("Falha ao criar post", error);
    }
  };

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
    return <div>Carregando posts...</div>;
  }

  return (
    <div>
      <nav>
        <Link to="/">Timeline</Link> | <Link to="/profile">Meu Perfil</Link>
      </nav>
      <hr />
      <h1>Timeline</h1>

      <form onSubmit={handleCreatePost}>
        <textarea
          value={newPostContent}
          onChange={(e) => setNewPostContent(e.target.value)}
          placeholder="O que está acontecendo?"
          rows="3"
          style={{ width: '100%', boxSizing: 'border-box' }}
        />
        <button type="submit">Postar</button>
      </form>
      <hr />

      <div>
        {posts.length > 0 ? (
          posts.map(post => (
            <div key={post.id} style={{ border: '1px solid #ccc', margin: '10px', padding: '10px' }}>
              <p>{post.content}</p>
              <small>Postado em: {new Date(post.createdAt).toLocaleString()}</small>
              <div>
                {/* Botão de curtir agora chama a função */}
                <button onClick={() => handleLikePost(post.id)}>
                  Curtir ({post.likes})
                </button>
              </div>
            </div>
          ))
        ) : (
          <p>Nenhum post encontrado. Crie o primeiro!</p>
        )}
      </div>
    </div>
  );
}

export default TimelinePage;