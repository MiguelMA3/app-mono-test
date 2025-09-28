import React from 'react';
import { Link } from 'react-router-dom';

function TimelinePage() {
  return (
    <div>
      {/* Adicione um menu de navegação simples */}
      <nav>
        <Link to="/">Timeline</Link> | <Link to="/profile">Meu Perfil</Link>
      </nav>
      <hr />
      
      <h1>Timeline</h1>
      <p>Bem-vindo! Os posts aparecerão aqui.</p>
    </div>
  );
}

export default TimelinePage;