import React from 'react';
import { Link } from 'react-router-dom';
import styles from './Post.module.css';

function Post({ post, onLike }) {
    const formatDate = (dateString) => {
        const options = { year: 'numeric', month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' };

        return new Date(dateString).toLocaleString(undefined, options);
    };

    return (
        <div className={styles.postContainer}>
            <div className={styles.postContent}>
                <div className={styles.postHeader}>
                    <Link to={`/profile/${post.user.id}`} className={styles.authorName}>
                        {post.user ? post.user.username : 'Usuário desconhecido'}
                    </Link>
                    <span className={styles.postTimestamp}>· {formatDate(post.createdAt)}</span>
                </div>

                <div className={styles.postBody}>
                    <p>{post.content}</p>
                </div>

                <div className={styles.postActions}>
                    <button onClick={() => onLike(post.id)} className={styles.actionButton}>
                        ❤️ {post.likes}
                    </button>
                </div>
            </div>
        </div>
    );
}

export default Post;