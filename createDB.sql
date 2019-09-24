DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
	id SERIAL PRIMARY KEY,
	title TEXT,
	summary TEXT,
	body TEXT,
	created DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	deleted DATETIME DEFAULT NULL
) COLLATE='utf8_general_ci';

INSERT INTO posts (title, summary, body) VALUES(
	"Мой первый пост!",
	"Lorem ipsum dolor sit amet, consectetur adipisicing elit. Odio praesentium, quos. Aspernatur assumenda cupiditate deserunt ducimus, eveniet, expedita inventore laboriosam magni modi non odio, officia qui sequi similique unde voluptatem.",
	"Здесь основной текст.
# Markdown!
*Это* **круто**!"
);
INSERT INTO posts (title, summary, body) VALUES(
	"Это уже второй пост!",
	"Lorem ipsum dolor sit amet, consectetur adipisicing elit. Odio praesentium, quos. Aspernatur assumenda cupiditate deserunt ducimus, eveniet, expedita inventore laboriosam magni modi non odio, officia qui sequi similique unde voluptatem.
Lorem ipsum dolor sit amet, consectetur adipisicing elit. Odio praesentium, quos. Aspernatur assumenda cupiditate deserunt ducimus, eveniet, expedita inventore laboriosam magni modi non odio, officia qui sequi similique unde voluptatem.",
	"Разобрался в шаблонах и маркдаунах, как их совместить.

Теперь понять, как переходить на отдельные посты.
# Anybody!
*Hz* **cool**!"
);
INSERT INTO posts (title, summary, body) VALUES(
	"Пример основных вариантов разметки Markdown",
	"Официальное руководство по синтаксису Markdown мне кажется слишком длинным и не слишком наглядным, поэтому я составил краткое руководство, которое поможет выучить или повторить синтаксис Маркдауна за час.",
	"TODO"
);

-- SELECT * FROM posts;
