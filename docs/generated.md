
# Список методов


## Общедоступные


<details><summary>Список форумов</summary>
<p>



**GET** [/v1/forums](../sources/server/internal/endpoints/show_forums.go#L13)


Схема ответа:

```
{
  forum_blocks: [{               # список блоков форумов
    id: uint64                   # id блока форумов
    title: string                # название
    forums: [{                   # форумы
      id: uint64                 # id форума
      title: string              # название
      forum_description: string  # описание
      moderators: [{             # модераторы
        id: uint64               # id пользователя
        login: string            # логин
        name: string             # имя
        gender: int32            # пол
        avatar: string           # аватар
        class: uint64            # класс
        sign: string             # подпись на форуме
      }]
      stats: {                   # статистика
        topic_count: uint64      # количество тем
        message_count: uint64    # количество сообщений
      }
      last_message: {            # последнее сообщение
        id: uint64               # id сообщения
        topic: {                 # тема, в которую входит сообщение
          id: uint64             # id темы
          title: string          # название
        }
        user: {                  # автор
          id: uint64             # id пользователя
          login: string          # логин
          name: string           # имя
          gender: int32          # пол
          avatar: string         # аватар
          class: uint64          # класс
          sign: string           # подпись на форуме
        }
        text: string             # текст
        date: {                  # дата и время создания
          seconds: int64
          nanos: int32
        }
      }
    }]
  }]
}
```
---

</p>
</details>

<details><summary>Список тем форума</summary>
<p>



**GET** [/v1/forums/{id}](../sources/server/internal/endpoints/show_forum_topics.go#L16)

Параметры запроса:


* **id** (path, uint64) - айди форума


* **page** (query, uint64) - номер страницы (по умолчанию - 1)


* **limit** (query, uint64) - кол-во записей на странице (по умолчанию - 20)




Схема ответа:

```
{
  topics: [{                 # список тем
    id: uint64               # id темы
    title: string            # название
    topic_type: int32        # тип
    creation: {              # данные о создании
      user: {                # пользователь
        id: uint64           # id пользователя
        login: string        # логин
        name: string         # имя
        gender: int32        # пол
        avatar: string       # аватар
        class: uint64        # класс
        sign: string         # подпись на форуме
      }
      date: {                # дата создания
        seconds: int64
        nanos: int32
      }
    }
    is_closed: bool          # тема закрыта?
    is_pinned: bool          # тема закреплена?
    stats: {                 # статистика
      message_count: uint64  # количество сообщений
      view_count: uint64     # количество просмотров
    }
    last_message: {          # последнее сообщение
      id: uint64             # id сообщения
      topic: {               # тема, в которую входит сообщение
        id: uint64           # id темы
        title: string        # название
      }
      user: {                # автор
        id: uint64           # id пользователя
        login: string        # логин
        name: string         # имя
        gender: int32        # пол
        avatar: string       # аватар
        class: uint64        # класс
        sign: string         # подпись на форуме
      }
      text: string           # текст
      date: {                # дата и время создания
        seconds: int64
        nanos: int32
      }
    }
  }]
  pages: {                   # страницы
    current: uint64          # текущая
    count: uint64            # количество
  }
}
```
---

</p>
</details>

<details><summary>Сообщения в теме форума</summary>
<p>



**GET** [/v1/topics/{id}](../sources/server/internal/endpoints/show_topic_messages.go#L16)

Параметры запроса:


* **id** (path, uint64) - айди темы


* **page** (query, uint64) - номер страницы (по умолчанию - 1)


* **limit** (query, uint64) - кол-во записей на странице (по умолчанию - 20)


* **order** (query, string) - порядок выдачи (asc - по умолчанию, desc)




Схема ответа:

```
{
  topic: {                     # тема
    id: uint64                 # id темы
    title: string              # название
    topic_type: int32          # тип
    creation: {                # данные о создании
      user: {                  # пользователь
        id: uint64             # id пользователя
        login: string          # логин
        name: string           # имя
        gender: int32          # пол
        avatar: string         # аватар
        class: uint64          # класс
        sign: string           # подпись на форуме
      }
      date: {                  # дата создания
        seconds: int64
        nanos: int32
      }
    }
    is_closed: bool            # тема закрыта?
    is_pinned: bool            # тема закреплена?
    stats: {                   # статистика
      message_count: uint64    # количество сообщений
      view_count: uint64       # количество просмотров
    }
    last_message: {            # последнее сообщение
      id: uint64               # id сообщения
      topic: {                 # тема, в которую входит сообщение
        id: uint64             # id темы
        title: string          # название
      }
      user: {                  # автор
        id: uint64             # id пользователя
        login: string          # логин
        name: string           # имя
        gender: int32          # пол
        avatar: string         # аватар
        class: uint64          # класс
        sign: string           # подпись на форуме
      }
      text: string             # текст
      date: {                  # дата и время создания
        seconds: int64
        nanos: int32
      }
    }
  }
  forum: {                     # форум, в который входит тема
    id: uint64                 # id форума
    title: string              # название
    forum_description: string  # описание
    moderators: [{             # модераторы
      id: uint64               # id пользователя
      login: string            # логин
      name: string             # имя
      gender: int32            # пол
      avatar: string           # аватар
      class: uint64            # класс
      sign: string             # подпись на форуме
    }]
    stats: {                   # статистика
      topic_count: uint64      # количество тем
      message_count: uint64    # количество сообщений
    }
    last_message: {            # последнее сообщение
      id: uint64               # id сообщения
      topic: {                 # тема, в которую входит сообщение
        id: uint64             # id темы
        title: string          # название
      }
      user: {                  # автор
        id: uint64             # id пользователя
        login: string          # логин
        name: string           # имя
        gender: int32          # пол
        avatar: string         # аватар
        class: uint64          # класс
        sign: string           # подпись на форуме
      }
      text: string             # текст
      date: {                  # дата и время создания
        seconds: int64
        nanos: int32
      }
    }
  }
  messages: [{                 # сообщения
    id: uint64                 # id сообщения
    creation: {                # данные о создании
      user: {                  # пользователь
        id: uint64             # id пользователя
        login: string          # логин
        name: string           # имя
        gender: int32          # пол
        avatar: string         # аватар
        class: uint64          # класс
        sign: string           # подпись на форуме
      }
      date: {                  # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string               # текст
    is_censored: bool          # текст изъят модератором?
    stats: {                   # статистика
      rating: int64            # рейтинг
    }
  }]
  pages: {                     # страницы
    current: uint64            # текущая
    count: uint64              # количество
  }
}
```
---

</p>
</details>

<details><summary>Список сообществ</summary>
<p>



**GET** [/v1/communities](../sources/server/internal/endpoints/show_communities.go#L11)


Схема ответа:

```
{
  main: [{                         # основные рубрики
    id: uint64                     # id рубрики
    title: string                  # название
    community_description: string  # описание
    rules: string                  # правила
    avatar: string                 # аватар
    stats: {                       # статистика
      article_count: uint64        # количество статей
      subscriber_count: uint64     # количество подписчиков
    }
    last_article: {                # последняя статья
      id: uint64                   # id статьи
      title: string                # название
      user: {                      # автор
        id: uint64                 # id пользователя
        login: string              # логин
        name: string               # имя
        gender: int32              # пол
        avatar: string             # аватар
        class: uint64              # класс
        sign: string               # подпись на форуме
      }
      date: {                      # дата создания
        seconds: int64
        nanos: int32
      }
    }
  }]
  additional: [{                   # дополнительные рубрики
    id: uint64                     # id рубрики
    title: string                  # название
    community_description: string  # описание
    rules: string                  # правила
    avatar: string                 # аватар
    stats: {                       # статистика
      article_count: uint64        # количество статей
      subscriber_count: uint64     # количество подписчиков
    }
    last_article: {                # последняя статья
      id: uint64                   # id статьи
      title: string                # название
      user: {                      # автор
        id: uint64                 # id пользователя
        login: string              # логин
        name: string               # имя
        gender: int32              # пол
        avatar: string             # аватар
        class: uint64              # класс
        sign: string               # подпись на форуме
      }
      date: {                      # дата создания
        seconds: int64
        nanos: int32
      }
    }
  }]
}
```
---

</p>
</details>

<details><summary>Информация о сообществе</summary>
<p>



**GET** [/v1/communities/{id}](../sources/server/internal/endpoints/show_community.go#L14)

Параметры запроса:


* **id** (path, uint64) - айди сообщества


* **page** (query, uint64) - номер страницы (по умолчанию - 1)


* **limit** (query, uint64) - кол-во записей на странице (по умолчанию - 5)




Схема ответа:

```
{
  community: {                     # рубрика
    id: uint64                     # id рубрики
    title: string                  # название
    community_description: string  # описание
    rules: string                  # правила
    avatar: string                 # аватар
    stats: {                       # статистика
      article_count: uint64        # количество статей
      subscriber_count: uint64     # количество подписчиков
    }
    last_article: {                # последняя статья
      id: uint64                   # id статьи
      title: string                # название
      user: {                      # автор
        id: uint64                 # id пользователя
        login: string              # логин
        name: string               # имя
        gender: int32              # пол
        avatar: string             # аватар
        class: uint64              # класс
        sign: string               # подпись на форуме
      }
      date: {                      # дата создания
        seconds: int64
        nanos: int32
      }
    }
  }
  moderators: [{                   # модераторы
    id: uint64                     # id пользователя
    login: string                  # логин
    name: string                   # имя
    gender: int32                  # пол
    avatar: string                 # аватар
    class: uint64                  # класс
    sign: string                   # подпись на форуме
  }]
  authors: [{                      # авторы
    id: uint64                     # id пользователя
    login: string                  # логин
    name: string                   # имя
    gender: int32                  # пол
    avatar: string                 # аватар
    class: uint64                  # класс
    sign: string                   # подпись на форуме
  }]
  articles: [{                     # статьи
    id: uint64                     # id статьи
    title: string                  # название
    creation: {                    # данные о создании
      user: {                      # пользователь
        id: uint64                 # id пользователя
        login: string              # логин
        name: string               # имя
        gender: int32              # пол
        avatar: string             # аватар
        class: uint64              # класс
        sign: string               # подпись на форуме
      }
      date: {                      # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string                   # текст
    tags: string                   # теги
    stats: {                       # статистика
      like_count: uint64           # количество лайков
      view_count: uint64           # количество просмотров
      comment_count: uint64        # количество комментариев
    }
  }]
  pages: {                         # страницы
    current: uint64                # текущая
    count: uint64                  # количество
  }
}
```
---

</p>
</details>

<details><summary>Список блогов</summary>
<p>



**GET** [/v1/blogs](../sources/server/internal/endpoints/show_blogs.go#L12)

Параметры запроса:


* **page** (query, uint64) - номер страницы (по умолчанию - 1)


* **limit** (query, uint64) - кол-во записей на странице (по умолчанию - 5)


* **sort** (query, string) - сортировать по (кол-ву тем в блоге - article, кол-ву подписчиков - subscriber, дате обновления - update (по умолчанию))




Схема ответа:

```
{
  blogs: [{                     # блоги
    id: uint64                  # id блога
    user: {                     # автор
      id: uint64                # id пользователя
      login: string             # логин
      name: string              # имя
      gender: int32             # пол
      avatar: string            # аватар
      class: uint64             # класс
      sign: string              # подпись на форуме
    }
    is_closed: bool             # блог закрыт?
    stats: {                    # статистика
      article_count: uint64     # количество статей
      subscriber_count: uint64  # количество подписчиков
    }
    last_article: {             # последняя статья
      id: uint64                # id статьи
      title: string             # название
      user: {                   # автор
        id: uint64              # id пользователя
        login: string           # логин
        name: string            # имя
        gender: int32           # пол
        avatar: string          # аватар
        class: uint64           # класс
        sign: string            # подпись на форуме
      }
      date: {                   # дата создания
        seconds: int64
        nanos: int32
      }
    }
  }]
  pages: {                      # страницы
    current: uint64             # текущая
    count: uint64               # количество
  }
}
```
---

</p>
</details>

<details><summary>Список статей в блоге</summary>
<p>



**GET** [/v1/blogs/{id}](../sources/server/internal/endpoints/show_blog.go#L14)

Параметры запроса:


* **id** (path, uint64) - айди блога


* **page** (query, uint64) - номер страницы (по умолчанию - 1)


* **limit** (query, uint64) - кол-во записей на странице (по умолчанию - 20)




Схема ответа:

```
{
  articles: [{               # статьи
    id: uint64               # id статьи
    title: string            # название
    creation: {              # данные о создании
      user: {                # пользователь
        id: uint64           # id пользователя
        login: string        # логин
        name: string         # имя
        gender: int32        # пол
        avatar: string       # аватар
        class: uint64        # класс
        sign: string         # подпись на форуме
      }
      date: {                # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string             # текст
    tags: string             # теги
    stats: {                 # статистика
      like_count: uint64     # количество лайков
      view_count: uint64     # количество просмотров
      comment_count: uint64  # количество комментариев
    }
  }]
  pages: {                   # страницы
    current: uint64          # текущая
    count: uint64            # количество
  }
}
```
---

</p>
</details>

<details><summary>Статья в блоге</summary>
<p>



**GET** [/v1/blog_articles/{id}](../sources/server/internal/endpoints/show_article.go#L13)

Параметры запроса:


* **id** (path, uint64) - айди статьи




Схема ответа:

```
{
  article: {                 # статья
    id: uint64               # id статьи
    title: string            # название
    creation: {              # данные о создании
      user: {                # пользователь
        id: uint64           # id пользователя
        login: string        # логин
        name: string         # имя
        gender: int32        # пол
        avatar: string       # аватар
        class: uint64        # класс
        sign: string         # подпись на форуме
      }
      date: {                # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string             # текст
    tags: string             # теги
    stats: {                 # статистика
      like_count: uint64     # количество лайков
      view_count: uint64     # количество просмотров
      comment_count: uint64  # количество комментариев
    }
  }
}
```
---

</p>
</details>

<details><summary>Список жанров</summary>
<p>



**GET** [/v1/allgenres](../sources/server/internal/endpoints/show_genres.go#L11)


Схема ответа:

```
{
  groups: [{              # группы жанров
    id: uint64            # id группы жанров
    name: string          # название
    genres: [{            # жанры
      id: uint64          # id жанра
      name: string        # название
      info: string        # информация
      subgenres: [...]    # поджанры
      work_count: uint64  # количество произведений (опционально)
      vote_count: uint64  # количество голосов (опционально)
    }]
  }]
}
```
---

</p>
</details>

<details><summary>Классификация произведения</summary>
<p>



**GET** [/v1/work/{id}/classification](../sources/server/internal/endpoints/get_work_classification.go#L11)

Параметры запроса:


* **id** (path, uint64) - айди произведения




Схема ответа:

```
{
  groups: [{                    # группы жанров
    id: uint64                  # id группы жанров
    name: string                # название
    genres: [{                  # жанры
      id: uint64                # id жанра
      name: string              # название
      info: string              # информация
      subgenres: [...]          # поджанры
      work_count: uint64        # количество произведений (опционально)
      vote_count: uint64        # количество голосов (опционально)
    }]
  }]
  classification_count: uint64  # сколько раз пользователи классифицировали произведение
}
```
---

</p>
</details>

<details><summary>Иерархия произведений, входящих в запрашиваемое</summary>
<p>



**GET** [/v1/work/{id}/subworks](../sources/server/internal/endpoints/get_work_subworks.go#L11)

Параметры запроса:


* **id** (path, uint64) - айди произведения


* **depth** (query, uint8) - глубина дерева (1 - 5, по умолчанию - 4)




Схема ответа:

```
{
  work_id: uint64            # айди произведения, для которого был запрос
  subworks: [{               # произведения, входящие в запрашиваемое
    id: uint64               # идентификатор произведения
    orig_name: string        # оригинальное название
    rus_name: string         # название на русском
    year: uint64             # год публикации
    work_type: int32         # тип произведения
    rating: float64          # рейтинг
    marks: uint64            # кол-во оценок
    reviews: uint64          # кол-во отзывов
    plus: bool               # является ли произведение дополнительным
    publish_status: [int32]  # статус публикации (не закончено, в планах, etc.)
    subworks: [...]          # дочерние произведения
  }]
}
```
---

</p>
</details>


## Для анонимов


<details><summary>Логин</summary>
<p>

Создает новую сессию пользователя


**POST** [/v1/login](../sources/server/internal/endpoints/login.go#L14)

Параметры запроса:


* **login** (form, string) - никнейм пользователя


* **password** (form, string) - пароль




Схема ответа:

```
{
  user_id: uint64        # id пользователя, с которым связана созданная сессия
  session_token: string  # токен сессии -> X-Session
}
```
---

</p>
</details>


## Для авторизованных пользователей


<details><summary>Разлогин</summary>
<p>

Удаляет текущую сессию пользователя


**DELETE** [/v1/logout](../sources/server/internal/endpoints/logout.go#L11)


Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Классификация произведения пользователем</summary>
<p>



**GET** [/v1/work/{id}/userclassification](../sources/server/internal/endpoints/get_user_work_genres.go#L11)

Параметры запроса:


* **id** (path, uint64) - айди произведения




Схема ответа:

```
{
  groups: [{              # группы жанров
    id: uint64            # id группы жанров
    name: string          # название
    genres: [{            # жанры
      id: uint64          # id жанра
      name: string        # название
      info: string        # информация
      subgenres: [...]    # поджанры
      work_count: uint64  # количество произведений (опционально)
      vote_count: uint64  # количество голосов (опционально)
    }]
  }]
}
```
---

</p>
</details>


## Для авторизованных незабаненных пользователей


<details><summary>Подписка на тему форума</summary>
<p>



**POST** [/v1/topics/{id}/subscription](../sources/server/internal/endpoints/subscribe_forum_topic.go#L14)

Параметры запроса:


* **id** (path, uint64) - айди темы




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Отписка от темы форума</summary>
<p>



**DELETE** [/v1/topics/{id}/subscription](../sources/server/internal/endpoints/unsubscribe_forum_topic.go#L14)

Параметры запроса:


* **id** (path, uint64) - айди темы




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Вступление в сообщество</summary>
<p>



**POST** [/v1/communities/{id}/subscription](../sources/server/internal/endpoints/subscribe_community.go#L12)

Параметры запроса:


* **id** (path, uint64) - айди сообщества




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Выход из сообщества</summary>
<p>



**DELETE** [/v1/communities/{id}/subscription](../sources/server/internal/endpoints/unsubscribe_community.go#L12)

Параметры запроса:


* **id** (path, uint64) - айди сообщества




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Подписка на блог</summary>
<p>



**POST** [/v1/blogs/{id}/subscription](../sources/server/internal/endpoints/subscribe_blog.go#L12)

Параметры запроса:


* **id** (path, uint64) - айди блога




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Отписка от блога</summary>
<p>



**DELETE** [/v1/blogs/{id}/subscription](../sources/server/internal/endpoints/unsubscribe_blog.go#L12)

Параметры запроса:


* **id** (path, uint64) - айди блога




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Подписка на статью в блоге</summary>
<p>



**POST** [/v1/blog_articles/{id}/subscription](../sources/server/internal/endpoints/subscribe_article.go#L10)

Параметры запроса:


* **id** (path, uint64) - айди статьи




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Отписка от статьи в блоге</summary>
<p>



**DELETE** [/v1/blog_articles/{id}/subscription](../sources/server/internal/endpoints/unsubscribe_article.go#L10)

Параметры запроса:


* **id** (path, uint64) - айди статьи




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Лайк статьи в блоге</summary>
<p>



**POST** [/v1/blog_articles/{id}/like](../sources/server/internal/endpoints/like_article.go#L13)

Параметры запроса:


* **id** (path, uint64) - айди статьи




Схема ответа:

```
{
  like_count: uint64  # количество лайков
}
```
---

</p>
</details>

<details><summary>Дизлайк статьи в блоге</summary>
<p>



**DELETE** [/v1/blog_articles/{id}/like](../sources/server/internal/endpoints/dislike_article.go#L12)

Параметры запроса:


* **id** (path, uint64) - айди статьи




Схема ответа:

```
{
  like_count: uint64  # количество лайков
}
```
---

</p>
</details>


## Для незабаненных философов


<details><summary>Классификация произведения пользователем</summary>
<p>



**PUT** [/v1/work/{id}/userclassification](../sources/server/internal/endpoints/set_work_genres.go#L14)

Параметры запроса:


* **id** (path, uint64) - айди произведения


* **genres** (form, string) - айди жанров, разделённые запятыми




Схема ответа:

```
{}
```
---

</p>
</details>


