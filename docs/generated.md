
# Список методов


## Общедоступные


<details><summary>Список форумов</summary>
<p>

**GET** [/v1/forums](../app/api/internal/endpoints/show_forums.go#L14)

Нет описания

---

</p>
</details>

<details><summary>Список тем форума</summary>
<p>

**GET** [/v1/forums/{id}](../app/api/internal/endpoints/show_forum_topics.go#L16)

Нет описания

---

</p>
</details>

<details><summary>Сообщения в теме форума</summary>
<p>

**GET** [/v1/topics/{id}](../app/api/internal/endpoints/show_topic_messages.go#L15)

Нет описания

---

</p>
</details>

<details><summary>Список сообществ</summary>
<p>

**GET** [/v1/communities](../app/api/internal/endpoints/show_communities.go#L11)

Нет описания

---

</p>
</details>

<details><summary>Информация о сообществе</summary>
<p>

**GET** [/v1/communities/{id}](../app/api/internal/endpoints/show_community.go#L14)

Нет описания

---

</p>
</details>

<details><summary>Список блогов</summary>
<p>

**GET** [/v1/blogs](../app/api/internal/endpoints/show_blogs.go#L13)

Нет описания

---

</p>
</details>

<details><summary>Список статей в блоге</summary>
<p>

**GET** [/v1/blogs/{id}](../app/api/internal/endpoints/show_blog.go#L14)

Нет описания

---

</p>
</details>

<details><summary>Статья в блоге</summary>
<p>

**GET** [/v1/blog_articles/{id}](../app/api/internal/endpoints/show_article.go#L13)

Нет описания

---

</p>
</details>

<details><summary>Список жанров</summary>
<p>

**GET** [/v1/allgenres](../app/api/internal/endpoints/show_genres.go#L11)

Нет описания

---

</p>
</details>


## Для анонимов


<details><summary>Логин</summary>
<p>

**POST** [/v1/login](../app/api/internal/endpoints/login.go#L16)

Создает новую сессию пользователя.
Параметры (form) - **login** и **password**.
В случае успеха в ответе придёт токен свежесозданной сессии.


---

</p>
</details>


## Для авторизованных пользователей


<details><summary>Разлогин</summary>
<p>

**DELETE** [/v1/logout](../app/api/internal/endpoints/logout.go#L11)

Удаляет текущую сессию пользователя


---

</p>
</details>


## Для авторизованных незабаненных пользователей


<details><summary>Подписка на тему форума</summary>
<p>

**POST** [/v1/topics/{id}/subscription](../app/api/internal/endpoints/subscribe_forum_topic.go#L15)

Нет описания

---

</p>
</details>

<details><summary>Отписка от темы форума</summary>
<p>

**DELETE** [/v1/topics/{id}/subscription](../app/api/internal/endpoints/unsubscribe_forum_topic.go#L15)

Нет описания

---

</p>
</details>

<details><summary>Вступление в сообщество</summary>
<p>

**POST** [/v1/communities/{id}/subscription](../app/api/internal/endpoints/subscribe_community.go#L12)

Нет описания

---

</p>
</details>

<details><summary>Выход из сообщества</summary>
<p>

**DELETE** [/v1/communities/{id}/subscription](../app/api/internal/endpoints/unsubscribe_community.go#L12)

Нет описания

---

</p>
</details>

<details><summary>Подписка на блог</summary>
<p>

**POST** [/v1/blogs/{id}/subscription](../app/api/internal/endpoints/subscribe_blog.go#L12)

Нет описания

---

</p>
</details>

<details><summary>Отписка от блога</summary>
<p>

**DELETE** [/v1/blogs/{id}/subscription](../app/api/internal/endpoints/unsubscribe_blog.go#L12)

Нет описания

---

</p>
</details>

<details><summary>Подписка на статью в блоге</summary>
<p>

**POST** [/v1/blog_articles/{id}/subscription](../app/api/internal/endpoints/subscribe_article.go#L10)

Нет описания

---

</p>
</details>

<details><summary>Отписка от статьи в блоге</summary>
<p>

**DELETE** [/v1/blog_articles/{id}/subscription](../app/api/internal/endpoints/unsubscribe_article.go#L10)

Нет описания

---

</p>
</details>

<details><summary>Лайк статьи в блоге</summary>
<p>

**POST** [/v1/blog_articles/{id}/like](../app/api/internal/endpoints/like_article.go#L13)

Нет описания

---

</p>
</details>

<details><summary>Дизлайк статьи в блоге</summary>
<p>

**DELETE** [/v1/blog_articles/{id}/like](../app/api/internal/endpoints/dislike_article.go#L12)

Нет описания

---

</p>
</details>

<details><summary>Классификация произведения</summary>
<p>

**PUT** [/v1/work/{id}/genres](../app/api/internal/endpoints/set_work_genres.go#L14)

Нет описания

---

</p>
</details>


