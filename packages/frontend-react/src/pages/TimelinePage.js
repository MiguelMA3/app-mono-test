import React, { useState, useEffect } from 'react';
import api from '../services/api';
import Post from '../components/Post/Post';

function TimelinePage() {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [newPostContent, setNewPostContent] = useState('');

  useEffect(() => {
    const fetchPosts = async () => {
      setLoading(true);
      try {
        const response = await api.get('/posts');
        setPosts(response.data || []);
      } catch (error) {
        console.error("Falha ao buscar posts", error);
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
        content: newPostContent,
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
    return <div className="page-content">Carregando posts...</div>;
  }

  return (
    <div className="page-content">
      <h1>Timeline</h1>

      <form onSubmit={handleCreatePost}>
        <textarea
          value={newPostContent}
          onChange={(e) => setNewPostContent(e.target.value)}
          placeholder="O que está acontecendo?"
          rows="3"
        />
        <button type="submit">Postar</button>
      </form>
      <hr style={{margin: '20px 0'}}/>

      <div>
        {posts.length > 0 ? (
          posts.map(post => (
            <Post key={post.id} post={post} onLike={handleLikePost} />
          ))
        ) : (
          <p>Nenhum post encontrado. Crie o primeiro!</p>
        )}
      </div>
    </div>
  );
}

export default TimelinePage;