-- Limpiar datos existentes (opcional, útil para evitar duplicados si se ejecuta múltiples veces)
-- TRUNCATE TABLE comments, posts, users RESTART IDENTITY CASCADE;

-- Insertar un usuario semilla
INSERT INTO users (username, password, email)
VALUES ('johndoe', 'hashed_password_placeholder', 'johndoe@example.com')
ON CONFLICT DO NOTHING;

-- Insertar un post semilla para el usuario
INSERT INTO posts (title, content, user_id, tags)
VALUES (
    'Mi Primer Post',
    '¡Hola mundo! Este es mi primer post en la red social.',
    (SELECT id FROM users WHERE username = 'johndoe' LIMIT 1),
    '{"golang", "sql", "backend"}'
)
ON CONFLICT DO NOTHING;

-- Insertar varios comentarios para el post
INSERT INTO comments (post_id, user_id, content)
VALUES
(
    (SELECT id FROM posts WHERE title = 'Mi Primer Post' LIMIT 1),
    (SELECT id FROM users WHERE username = 'johndoe' LIMIT 1),
    '¡Excelente post! Me encantó leerlo.'
),
(
    (SELECT id FROM posts WHERE title = 'Mi Primer Post' LIMIT 1),
    (SELECT id FROM users WHERE username = 'johndoe' LIMIT 1),
    'Muy buena información, ¡gracias por compartir!'
),
(
    (SELECT id FROM posts WHERE title = 'Mi Primer Post' LIMIT 1),
    (SELECT id FROM users WHERE username = 'johndoe' LIMIT 1),
    '¿Podrías hacer otro post sobre cómo conectar Go con PostgreSQL?'
);
