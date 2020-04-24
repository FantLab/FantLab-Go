
# Список методов


## Общедоступные


<details><summary>Логин</summary>
<p>

Создаёт новый аутентификационный токен для пользователя на основе пары логин/пароль


**POST** [/v1/auth/login](../sources/server/internal/endpoints/login.go#L15)

Параметры запроса:


* **login** (form, string) - логин или почта пользователя


* **password** (form, string) - пароль




Схема ответа:

```
{
  userId: uint64        # id пользователя
  token: string         # токен -> X-Session
  refreshToken: string  # токен для продления сессии
}
```
---

</p>
</details>

<details><summary>Список форумов</summary>
<p>



**GET** [/v1/forums](../sources/server/internal/endpoints/show_forums.go#L11)


Схема ответа:

```
{
  forumBlocks: [{               # список блоков форумов
    id: uint64                  # id блока форумов
    title: string               # название
    forums: [{                  # форумы
      id: uint64                # id форума
      title: string             # название
      forumDescription: string  # описание
      moderators: [{            # модераторы
        id: uint64              # id пользователя
        login: string           # логин
        name: string            # имя
        gender: int32           # пол
        avatar: string          # аватар
        class: int32            # класс
        sign: string            # подпись на форуме
      }]
      stats: {                  # статистика
        topicCount: uint64      # количество тем
        messageCount: uint64    # количество сообщений
      }
      lastMessage: {            # последнее сообщение
        id: uint64              # id сообщения
        topic: {                # тема, в которую входит сообщение
          id: uint64            # id темы
          title: string         # название
        }
        user: {                 # автор
          id: uint64            # id пользователя
          login: string         # логин
          name: string          # имя
          gender: int32         # пол
          avatar: string        # аватар
          class: int32          # класс
          sign: string          # подпись на форуме
        }
        text: string            # текст
        date: {                 # дата и время создания
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



**GET** [/v1/forums/{id}](../sources/server/internal/endpoints/show_forum_topics.go#L14)

Параметры запроса:


* **id** (path, uint64) - айди форума


* **page** (query, uint64) - номер страницы (по умолчанию - 1)


* **limit** (query, uint64) - кол-во записей на странице (по умолчанию - 20)




Схема ответа:

```
{
  topics: [{                # список тем
    id: uint64              # id темы
    title: string           # название
    topicType: int32        # тип
    creation: {             # данные о создании
      user: {               # пользователь
        id: uint64          # id пользователя
        login: string       # логин
        name: string        # имя
        gender: int32       # пол
        avatar: string      # аватар
        class: int32        # класс
        sign: string        # подпись на форуме
      }
      date: {               # дата создания
        seconds: int64
        nanos: int32
      }
    }
    isClosed: bool          # тема закрыта?
    isPinned: bool          # тема закреплена?
    stats: {                # статистика
      messageCount: uint64  # количество сообщений
      viewCount: uint64     # количество просмотров
    }
    lastMessage: {          # последнее сообщение
      id: uint64            # id сообщения
      topic: {              # тема, в которую входит сообщение
        id: uint64          # id темы
        title: string       # название
      }
      user: {               # автор
        id: uint64          # id пользователя
        login: string       # логин
        name: string        # имя
        gender: int32       # пол
        avatar: string      # аватар
        class: int32        # класс
        sign: string        # подпись на форуме
      }
      text: string          # текст
      date: {               # дата и время создания
        seconds: int64
        nanos: int32
      }
    }
  }]
  pages: {                  # страницы
    current: uint64         # текущая
    count: uint64           # количество
  }
}
```
---

</p>
</details>

<details><summary>Сообщения в теме форума</summary>
<p>



**GET** [/v1/topics/{id}](../sources/server/internal/endpoints/show_topic_messages.go#L13)

Параметры запроса:


* **id** (path, uint64) - id темы


* **page** (query, uint64) - номер страницы (по умолчанию - 1)


* **limit** (query, uint64) - кол-во записей на странице (по умолчанию - 20)


* **sortAsc** (query, uint8) - порядок выдачи (0 - от новых к старым, 1 - наоборот; по умолчанию - 0)




Схема ответа:

```
{
  topic: {                    # тема
    id: uint64                # id темы
    title: string             # название
    topicType: int32          # тип
    creation: {               # данные о создании
      user: {                 # пользователь
        id: uint64            # id пользователя
        login: string         # логин
        name: string          # имя
        gender: int32         # пол
        avatar: string        # аватар
        class: int32          # класс
        sign: string          # подпись на форуме
      }
      date: {                 # дата создания
        seconds: int64
        nanos: int32
      }
    }
    isClosed: bool            # тема закрыта?
    isPinned: bool            # тема закреплена?
    stats: {                  # статистика
      messageCount: uint64    # количество сообщений
      viewCount: uint64       # количество просмотров
    }
    lastMessage: {            # последнее сообщение
      id: uint64              # id сообщения
      topic: {                # тема, в которую входит сообщение
        id: uint64            # id темы
        title: string         # название
      }
      user: {                 # автор
        id: uint64            # id пользователя
        login: string         # логин
        name: string          # имя
        gender: int32         # пол
        avatar: string        # аватар
        class: int32          # класс
        sign: string          # подпись на форуме
      }
      text: string            # текст
      date: {                 # дата и время создания
        seconds: int64
        nanos: int32
      }
    }
  }
  forum: {                    # форум, в который входит тема
    id: uint64                # id форума
    title: string             # название
    forumDescription: string  # описание
    moderators: [{            # модераторы
      id: uint64              # id пользователя
      login: string           # логин
      name: string            # имя
      gender: int32           # пол
      avatar: string          # аватар
      class: int32            # класс
      sign: string            # подпись на форуме
    }]
    stats: {                  # статистика
      topicCount: uint64      # количество тем
      messageCount: uint64    # количество сообщений
    }
    lastMessage: {            # последнее сообщение
      id: uint64              # id сообщения
      topic: {                # тема, в которую входит сообщение
        id: uint64            # id темы
        title: string         # название
      }
      user: {                 # автор
        id: uint64            # id пользователя
        login: string         # логин
        name: string          # имя
        gender: int32         # пол
        avatar: string        # аватар
        class: int32          # класс
        sign: string          # подпись на форуме
      }
      text: string            # текст
      date: {                 # дата и время создания
        seconds: int64
        nanos: int32
      }
    }
  }
  pinnedMessage: {            # закрепленное сообщение, если есть
    id: uint64                # id сообщения
    creation: {               # данные о создании
      user: {                 # пользователь
        id: uint64            # id пользователя
        login: string         # логин
        name: string          # имя
        gender: int32         # пол
        avatar: string        # аватар
        class: int32          # класс
        sign: string          # подпись на форуме
      }
      date: {                 # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string              # текст
    isCensored: bool          # текст изъят модератором?
    stats: {                  # статистика
      rating: int64           # рейтинг
    }
  }
  messages: [{                # сообщения
    id: uint64                # id сообщения
    creation: {               # данные о создании
      user: {                 # пользователь
        id: uint64            # id пользователя
        login: string         # логин
        name: string          # имя
        gender: int32         # пол
        avatar: string        # аватар
        class: int32          # класс
        sign: string          # подпись на форуме
      }
      date: {                 # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string              # текст
    isCensored: bool          # текст изъят модератором?
    stats: {                  # статистика
      rating: int64           # рейтинг
    }
  }]
  pages: {                    # страницы
    current: uint64           # текущая
    count: uint64             # количество
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
  main: [{                        # основные рубрики
    id: uint64                    # id рубрики
    title: string                 # название
    communityDescription: string  # описание
    rules: string                 # правила
    avatar: string                # аватар
    stats: {                      # статистика
      articleCount: uint64        # количество статей
      subscriberCount: uint64     # количество подписчиков
    }
    lastArticle: {                # последняя статья
      id: uint64                  # id статьи
      title: string               # название
      user: {                     # автор
        id: uint64                # id пользователя
        login: string             # логин
        name: string              # имя
        gender: int32             # пол
        avatar: string            # аватар
        class: int32              # класс
        sign: string              # подпись на форуме
      }
      date: {                     # дата создания
        seconds: int64
        nanos: int32
      }
    }
  }]
  additional: [{                  # дополнительные рубрики
    id: uint64                    # id рубрики
    title: string                 # название
    communityDescription: string  # описание
    rules: string                 # правила
    avatar: string                # аватар
    stats: {                      # статистика
      articleCount: uint64        # количество статей
      subscriberCount: uint64     # количество подписчиков
    }
    lastArticle: {                # последняя статья
      id: uint64                  # id статьи
      title: string               # название
      user: {                     # автор
        id: uint64                # id пользователя
        login: string             # логин
        name: string              # имя
        gender: int32             # пол
        avatar: string            # аватар
        class: int32              # класс
        sign: string              # подпись на форуме
      }
      date: {                     # дата создания
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
  community: {                    # рубрика
    id: uint64                    # id рубрики
    title: string                 # название
    communityDescription: string  # описание
    rules: string                 # правила
    avatar: string                # аватар
    stats: {                      # статистика
      articleCount: uint64        # количество статей
      subscriberCount: uint64     # количество подписчиков
    }
    lastArticle: {                # последняя статья
      id: uint64                  # id статьи
      title: string               # название
      user: {                     # автор
        id: uint64                # id пользователя
        login: string             # логин
        name: string              # имя
        gender: int32             # пол
        avatar: string            # аватар
        class: int32              # класс
        sign: string              # подпись на форуме
      }
      date: {                     # дата создания
        seconds: int64
        nanos: int32
      }
    }
  }
  moderators: [{                  # модераторы
    id: uint64                    # id пользователя
    login: string                 # логин
    name: string                  # имя
    gender: int32                 # пол
    avatar: string                # аватар
    class: int32                  # класс
    sign: string                  # подпись на форуме
  }]
  authors: [{                     # авторы
    id: uint64                    # id пользователя
    login: string                 # логин
    name: string                  # имя
    gender: int32                 # пол
    avatar: string                # аватар
    class: int32                  # класс
    sign: string                  # подпись на форуме
  }]
  articles: [{                    # статьи
    id: uint64                    # id статьи
    title: string                 # название
    creation: {                   # данные о создании
      user: {                     # пользователь
        id: uint64                # id пользователя
        login: string             # логин
        name: string              # имя
        gender: int32             # пол
        avatar: string            # аватар
        class: int32              # класс
        sign: string              # подпись на форуме
      }
      date: {                     # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string                  # текст
    tags: string                  # теги
    stats: {                      # статистика
      likeCount: uint64           # количество лайков
      viewCount: uint64           # количество просмотров
      commentCount: uint64        # количество комментариев
    }
  }]
  pages: {                        # страницы
    current: uint64               # текущая
    count: uint64                 # количество
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


* **sort** (query, string) - сортировать по (кол-ву тем в блоге - article, кол-ву подписчиков - subscriber, дате обновления от новых к старым - update (по умолчанию))




Схема ответа:

```
{
  blogs: [{                    # блоги
    id: uint64                 # id блога
    user: {                    # автор
      id: uint64               # id пользователя
      login: string            # логин
      name: string             # имя
      gender: int32            # пол
      avatar: string           # аватар
      class: int32             # класс
      sign: string             # подпись на форуме
    }
    isClosed: bool             # блог закрыт?
    stats: {                   # статистика
      articleCount: uint64     # количество статей
      subscriberCount: uint64  # количество подписчиков
    }
    lastArticle: {             # последняя статья
      id: uint64               # id статьи
      title: string            # название
      user: {                  # автор
        id: uint64             # id пользователя
        login: string          # логин
        name: string           # имя
        gender: int32          # пол
        avatar: string         # аватар
        class: int32           # класс
        sign: string           # подпись на форуме
      }
      date: {                  # дата создания
        seconds: int64
        nanos: int32
      }
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
  articles: [{              # статьи
    id: uint64              # id статьи
    title: string           # название
    creation: {             # данные о создании
      user: {               # пользователь
        id: uint64          # id пользователя
        login: string       # логин
        name: string        # имя
        gender: int32       # пол
        avatar: string      # аватар
        class: int32        # класс
        sign: string        # подпись на форуме
      }
      date: {               # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string            # текст
    tags: string            # теги
    stats: {                # статистика
      likeCount: uint64     # количество лайков
      viewCount: uint64     # количество просмотров
      commentCount: uint64  # количество комментариев
    }
  }]
  pages: {                  # страницы
    current: uint64         # текущая
    count: uint64           # количество
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
  article: {                # статья
    id: uint64              # id статьи
    title: string           # название
    creation: {             # данные о создании
      user: {               # пользователь
        id: uint64          # id пользователя
        login: string       # логин
        name: string        # имя
        gender: int32       # пол
        avatar: string      # аватар
        class: int32        # класс
        sign: string        # подпись на форуме
      }
      date: {               # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string            # текст
    tags: string            # теги
    stats: {                # статистика
      likeCount: uint64     # количество лайков
      viewCount: uint64     # количество просмотров
      commentCount: uint64  # количество комментариев
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
  groups: [{             # группы жанров
    id: uint64           # id группы жанров
    name: string         # название
    genres: [{           # жанры
      id: uint64         # id жанра
      name: string       # название
      info: string       # информация
      subgenres: [...]   # поджанры
      workCount: uint64  # количество произведений (опционально)
      voteCount: uint64  # количество голосов (опционально)
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
  groups: [{                   # группы жанров
    id: uint64                 # id группы жанров
    name: string               # название
    genres: [{                 # жанры
      id: uint64               # id жанра
      name: string             # название
      info: string             # информация
      subgenres: [...]         # поджанры
      workCount: uint64        # количество произведений (опционально)
      voteCount: uint64        # количество голосов (опционально)
    }]
  }]
  classificationCount: uint64  # сколько раз пользователи классифицировали произведение
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
  workId: uint64            # айди произведения, для которого был запрос
  subworks: [{              # произведения, входящие в запрашиваемое
    id: uint64              # идентификатор произведения
    origName: string        # оригинальное название
    rusName: string         # название на русском
    year: uint64            # год публикации
    workType: int32         # тип произведения
    rating: float64         # рейтинг
    marks: uint64           # кол-во оценок
    reviews: uint64         # кол-во отзывов
    plus: bool              # является ли произведение дополнительным
    publishStatus: [int32]  # статус публикации (не закончено, в планах, etc.)
    subworks: [...]         # дочерние произведения
  }]
}
```
---

</p>
</details>

<details><summary>Комментарии к статье в блоге</summary>
<p>



**GET** [/v1/blog_articles/{id}/comments](../sources/server/internal/endpoints/show_blog_article_comments.go#L13)

Параметры запроса:


* **id** (path, uint64) - id статьи


* **after** (query, string) - дата, после которой искать сообщения (в формате RFC3339)


* **count** (query, uint64) - кол-во комментариев верхнего уровня (по умолчанию - 10, [5, 20])


* **sortAsc** (query, uint8) - порядок выдачи (0 - от новых к старым, 1 - наоборот; по умолчанию - 0)




Схема ответа:

```
{
  comments: [{          # список комментариев
    id: uint64          # id сообщения
    creation: {         # данные о создании
      user: {           # пользователь
        id: uint64      # id пользователя
        login: string   # логин
        name: string    # имя
        gender: int32   # пол
        avatar: string  # аватар
        class: int32    # класс
        sign: string    # подпись на форуме
      }
      date: {           # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string        # текст сообщения
    isCensored: bool    # текст изъят модератором?
    answers: [...]      # ответы на комментарий
  }]
  totalCount: uint64    # общее ко-во комментариев у поста
}
```
---

</p>
</details>

<details><summary>Список книжных полок</summary>
<p>



**GET** [/v1/users/{id}/bookcases](../sources/server/internal/endpoints/show_bookcases.go#L13)

Параметры запроса:


* **id** (path, uint64) - id пользователя




Схема ответа:

```
{
  bookcaseBlocks: [{     # список блоков книжных полок
    title: string        # название блока
    bookcases: [{        # книжные полки
      id: uint64         # id книжной полки
      isPrivate: bool    # приватная?
      type: int32        # тип
      title: string      # название
      comment: string    # комментарий
      index: uint64      # порядковый номер
      itemCount: uint64  # количество элементов
    }]
  }]
}
```
---

</p>
</details>

<details><summary>Содержимое полки с изданиями</summary>
<p>



**GET** [/v1/edition_bookcases/{id}](../sources/server/internal/endpoints/show_edition_bookcase.go#L15)

Параметры запроса:


* **id** (path, uint64) - id книжной полки


* **page** (query, uint64) - номер страницы (>0, по умолчанию - 1)


* **limit** (query, uint64) - кол-во элементов на странице ([5..50], по умолчанию - 50)


* **sort** (query, string) - сортировать по: порядку - order (по умолчанию), автору - author, названию - title, году - year




Схема ответа:

```
{
  bookcase: {                       # информация о полке
    id: uint64                      # id книжной полки
    isPrivate: bool                 # приватная?
    type: int32                     # тип
    title: string                   # название
    comment: string                 # комментарий
  }
  editions: [{                      # список изданий на полке
    id: uint64                      # id издания
    type: int32                     # тип (авторская книга/сборник/etc; может отсутствовать, если не задан)
    correctnessLevel: int32         # уровень проверенности
    cover: string                   # URL обложки
    authors: string                 # авторы
    title: string                   # название
    year: uint64                    # год публикации
    publishers: string              # издательства
    description: string             # описание
    plannedPublicationDate: string  # планируемая дата издания (если издание еще не опубликовано)
    offers: {                       # предложения в магазинах
      ozon: {                       # предложение на Озоне
        url: string                 # URL предложения
        price: uint64               # цена
      }
      labirint: {                   # предложение на Лабиринте
        url: string                 # URL предложения
        price: uint64               # цена
      }
    }
    comment: string                 # комментарий
  }]
  pages: {                          # страницы
    current: uint64                 # текущая
    count: uint64                   # количество
  }
}
```
---

</p>
</details>

<details><summary>Содержимое полки с произведениями</summary>
<p>



**GET** [/v1/work_bookcases/{id}](../sources/server/internal/endpoints/show_work_bookcase.go#L15)

Параметры запроса:


* **id** (path, uint64) - id книжной полки


* **page** (query, uint64) - номер страницы (>0, по умолчанию - 1)


* **limit** (query, uint64) - кол-во элементов на странице ([5..50], по умолчанию - 50)


* **sort** (query, string) - сортировать по: порядку - order (по умолчанию), автору - author, названию - title, оригинальному названию - orig_title, году - year, количеству оценок - mark_count, средней оценке - avg_mark




Схема ответа:

```
{
  bookcase: {                    # информация о полке
    id: uint64                   # id книжной полки
    isPrivate: bool              # приватная?
    type: int32                  # тип
    title: string                # название
    comment: string              # комментарий
  }
  works: [{                      # список произведений на полке
    id: uint64                   # id произведения
    type: int32                  # тип (роман/сборник/etc; может отсутствовать)
    authors: [{                  # авторы
      id: uint64                 # id автора
      name: string               # имя на русском языке
      isOpened: bool             # страница открыта?
    }]
    title: string                # название на русском языке
    originalTitle: string        # название в оригинале
    alternativeTitles: string    # альтернативные названия
    note: string                 # примечание
    year: int64                  # год
    description: string          # описание
    isPublished: bool            # опубликовано?
    stats: {                     # статистика
      averageMark: float64       # средняя оценка
      markCount: uint64          # количество оценок
      responseCount: uint64      # количество отзывов
    }
    own: {                       # персональное
      mark: uint64               # собственная оценка произведению
      isResponsePublished: bool  # опубликован отзыв?
    }
  }]
  pages: {                       # страницы
    current: uint64              # текущая
    count: uint64                # количество
  }
}
```
---

</p>
</details>

<details><summary>Содержимое кинополки</summary>
<p>



**GET** [/v1/film_bookcases/{id}](../sources/server/internal/endpoints/show_film_bookcase.go#L15)

Параметры запроса:


* **id** (path, uint64) - id книжной полки


* **page** (query, uint64) - номер страницы (>0, по умолчанию - 1)


* **limit** (query, uint64) - кол-во элементов на странице ([5..50], по умолчанию - 50)


* **sort** (query, string) - сортировать по: порядку - order (по умолчанию), названию - title, оригинальному названию - orig_title




Схема ответа:

```
{
  bookcase: {              # информация о полке
    id: uint64             # id книжной полки
    isPrivate: bool        # приватная?
    type: int32            # тип
    title: string          # название
    comment: string        # комментарий
  }
  films: [{                # список фильмов на полке
    id: uint64             # id фильма
    type: int32            # тип (фильм/сериал/etc; может отсутствовать, если не задан)
    poster: string         # URL постера
    title: string          # название на русском языке
    originalTitle: string  # название в оригинале
    year: uint64           # год выпуска (для всего, кроме сериалов)
    startYear: uint64      # год старта трансляции (для сериалов)
    endYear: uint64        # год окончания трансляции (для сериалов)
    countries: string      # страны производства
    genres: string         # жанры
    directors: string      # режиссеры
    screenWriters: string  # сценаристы
    actors: string         # актеры
    description: string    # описание
    comment: string        # комментарий
  }]
  pages: {                 # страницы
    current: uint64        # текущая
    count: uint64          # количество
  }
}
```
---

</p>
</details>


## Для зарегистрированных пользователей


<details><summary>Продление сессии</summary>
<p>

Продлевает сессию с помощью рефреш-токена


**POST** [/v1/auth/refresh](../sources/server/internal/endpoints/refresh_auth.go#L17)

Параметры запроса:


* **refresh_token** (form, string) - рефреш-токен, выданный при логине или предыдущем продлении сессии




Схема ответа:

```
{
  userId: uint64        # id пользователя
  token: string         # токен -> X-Session
  refreshToken: string  # токен для продления сессии
}
```
---

</p>
</details>


## Для пользователей с валидной сессией


<details><summary>Классификация произведения пользователем</summary>
<p>



**GET** [/v1/work/{id}/userclassification](../sources/server/internal/endpoints/get_user_work_genres.go#L11)

Параметры запроса:


* **id** (path, uint64) - айди произведения




Схема ответа:

```
{
  groups: [{             # группы жанров
    id: uint64           # id группы жанров
    name: string         # название
    genres: [{           # жанры
      id: uint64         # id жанра
      name: string       # название
      info: string       # информация
      subgenres: [...]   # поджанры
      workCount: uint64  # количество произведений (опционально)
      voteCount: uint64  # количество голосов (опционально)
    }]
  }]
}
```
---

</p>
</details>


## С проверкой на бан


<details><summary>Создание нового сообщения в форуме</summary>
<p>



**POST** [/v1/topics/{id}/message](../sources/server/internal/endpoints/add_forum_message.go#L15)

Параметры запроса:


* **id** (path, uint64) - id темы


* **message** (form, string) - текст сообщения




Схема ответа:

```
{
  message: {            # сообщение
    id: uint64          # id сообщения
    creation: {         # данные о создании
      user: {           # пользователь
        id: uint64      # id пользователя
        login: string   # логин
        name: string    # имя
        gender: int32   # пол
        avatar: string  # аватар
        class: int32    # класс
        sign: string    # подпись на форуме
      }
      date: {           # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string        # текст
    isCensored: bool    # текст изъят модератором?
    stats: {            # статистика
      rating: int64     # рейтинг
    }
  }
}
```
---

</p>
</details>

<details><summary>Редактирование сообщения в форуме</summary>
<p>



**PUT** [/v1/forum_messages/{id}](../sources/server/internal/endpoints/edit_forum_message.go#L15)

Параметры запроса:


* **id** (path, uint64) - id сообщения


* **message** (form, string) - новый текст сообщения




Схема ответа:

```
{
  message: {            # сообщение
    id: uint64          # id сообщения
    creation: {         # данные о создании
      user: {           # пользователь
        id: uint64      # id пользователя
        login: string   # логин
        name: string    # имя
        gender: int32   # пол
        avatar: string  # аватар
        class: int32    # класс
        sign: string    # подпись на форуме
      }
      date: {           # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string        # текст
    isCensored: bool    # текст изъят модератором?
    stats: {            # статистика
      rating: int64     # рейтинг
    }
  }
}
```
---

</p>
</details>

<details><summary>Удаление сообщения в форуме</summary>
<p>



**DELETE** [/v1/forum_messages/{id}](../sources/server/internal/endpoints/delete_forum_message.go#L13)

Параметры запроса:


* **id** (path, uint64) - id сообщения




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Получение URL для загрузки аттача к сообщению в форуме</summary>
<p>



**GET** [/v1/forum_messages/{id}/file_upload_url](../sources/server/internal/endpoints/get_forum_message_file_upload_url.go#L15)

Параметры запроса:


* **id** (path, uint64) - id сообщения


* **file_name** (query, string) - полное имя файла (с расширением)




Схема ответа:

```
{
  url: string  # URL на загрузку файла
}
```
---

</p>
</details>

<details><summary>Удаление аттача сообщения в форуме</summary>
<p>



**DELETE** [/v1/forum_messages/{id}/file](../sources/server/internal/endpoints/delete_forum_message_file.go#L15)

Параметры запроса:


* **id** (path, uint64) - id сообщения


* **file_name** (form, string) - полное имя файла (с расширением)




Схема ответа:

```
{
  message: {            # сообщение
    id: uint64          # id сообщения
    creation: {         # данные о создании
      user: {           # пользователь
        id: uint64      # id пользователя
        login: string   # логин
        name: string    # имя
        gender: int32   # пол
        avatar: string  # аватар
        class: int32    # класс
        sign: string    # подпись на форуме
      }
      date: {           # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string        # текст
    isCensored: bool    # текст изъят модератором?
    stats: {            # статистика
      rating: int64     # рейтинг
    }
  }
}
```
---

</p>
</details>

<details><summary>Сохранение черновика сообщения в форуме</summary>
<p>



**PUT** [/v1/topics/{id}/message_draft](../sources/server/internal/endpoints/save_forum_message_draft.go#L14)

Параметры запроса:


* **id** (path, uint64) - id темы


* **message** (form, string) - текст сообщения




Схема ответа:

```
{
  messageDraft: {       # черновик сообщения
    topicId: uint64     # id темы
    creation: {         # данные о создании
      user: {           # пользователь
        id: uint64      # id пользователя
        login: string   # логин
        name: string    # имя
        gender: int32   # пол
        avatar: string  # аватар
        class: int32    # класс
        sign: string    # подпись на форуме
      }
      date: {           # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string        # текст
  }
}
```
---

</p>
</details>

<details><summary>Подтверждение черновика сообщения в форуме</summary>
<p>



**POST** [/v1/topics/{id}/message_draft](../sources/server/internal/endpoints/confirm_forum_message_draft.go#L15)

Параметры запроса:


* **id** (path, uint64) - id темы




Схема ответа:

```
{
  message: {            # сообщение
    id: uint64          # id сообщения
    creation: {         # данные о создании
      user: {           # пользователь
        id: uint64      # id пользователя
        login: string   # логин
        name: string    # имя
        gender: int32   # пол
        avatar: string  # аватар
        class: int32    # класс
        sign: string    # подпись на форуме
      }
      date: {           # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string        # текст
    isCensored: bool    # текст изъят модератором?
    stats: {            # статистика
      rating: int64     # рейтинг
    }
  }
}
```
---

</p>
</details>

<details><summary>Отмена черновика сообщения в форуме</summary>
<p>



**DELETE** [/v1/topics/{id}/message_draft](../sources/server/internal/endpoints/cancel_forum_message_draft.go#L13)

Параметры запроса:


* **id** (path, uint64) - id темы




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Получение URL для загрузки аттача к черновику сообщения в форуме</summary>
<p>



**GET** [/v1/topics/{id}/message_draft/file_upload_url](../sources/server/internal/endpoints/get_forum_message_draft_file_upload_url.go#L14)

Параметры запроса:


* **id** (path, uint64) - id темы


* **file_name** (query, string) - полное имя файла (с расширением)




Схема ответа:

```
{
  url: string  # URL на загрузку файла
}
```
---

</p>
</details>

<details><summary>Удаление аттача черновика сообщения в форуме</summary>
<p>



**DELETE** [/v1/topics/{id}/message_draft/file](../sources/server/internal/endpoints/delete_forum_message_draft_file.go#L14)

Параметры запроса:


* **id** (path, uint64) - id темы


* **file_name** (form, string) - полное имя файла (с расширением)




Схема ответа:

```
{
  messageDraft: {       # черновик сообщения
    topicId: uint64     # id темы
    creation: {         # данные о создании
      user: {           # пользователь
        id: uint64      # id пользователя
        login: string   # логин
        name: string    # имя
        gender: int32   # пол
        avatar: string  # аватар
        class: int32    # класс
        sign: string    # подпись на форуме
      }
      date: {           # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string        # текст
  }
}
```
---

</p>
</details>

<details><summary>Подписка/отписка от темы форума</summary>
<p>



**PUT** [/v1/topics/{id}/subscription](../sources/server/internal/endpoints/toggle_forum_topic_subscription.go#L12)

Параметры запроса:


* **id** (path, uint64) - айди темы


* **subscribe** (form, bool) - подписаться - true, отписаться - false




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Создание нового комментария к статье в блоге</summary>
<p>



**POST** [/v1/blog_articles/{id}/comment](../sources/server/internal/endpoints/add_blog_article_comment.go#L13)

Параметры запроса:


* **id** (path, uint64) - айди статьи


* **comment** (form, string) - текст комментария (непустой)


* **parent_comment_id** (form, uint64) - id родительского комментария (0, если комментарий 1-го уровня вложенности)




Схема ответа:

```
{
  comment: {            # комментарий
    id: uint64          # id сообщения
    creation: {         # данные о создании
      user: {           # пользователь
        id: uint64      # id пользователя
        login: string   # логин
        name: string    # имя
        gender: int32   # пол
        avatar: string  # аватар
        class: int32    # класс
        sign: string    # подпись на форуме
      }
      date: {           # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string        # текст сообщения
    isCensored: bool    # текст изъят модератором?
    answers: [...]      # ответы на комментарий
  }
}
```
---

</p>
</details>

<details><summary>Редактирование комментария к статье в блоге</summary>
<p>



**PUT** [/v1/blog_article_comments/{id}](../sources/server/internal/endpoints/edit_blog_article_comment.go#L13)

Параметры запроса:


* **id** (path, uint64) - id комментария


* **comment** (form, string) - текст комментария (непустой)




Схема ответа:

```
{
  comment: {            # комментарий
    id: uint64          # id сообщения
    creation: {         # данные о создании
      user: {           # пользователь
        id: uint64      # id пользователя
        login: string   # логин
        name: string    # имя
        gender: int32   # пол
        avatar: string  # аватар
        class: int32    # класс
        sign: string    # подпись на форуме
      }
      date: {           # дата создания
        seconds: int64
        nanos: int32
      }
    }
    text: string        # текст сообщения
    isCensored: bool    # текст изъят модератором?
    answers: [...]      # ответы на комментарий
  }
}
```
---

</p>
</details>

<details><summary>Удаление комментария к статье в блоге</summary>
<p>



**DELETE** [/v1/blog_article_comments/{id}](../sources/server/internal/endpoints/delete_blog_article_comment.go#L12)

Параметры запроса:


* **id** (path, uint64) - id комментария




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Вступление/выход из сообщества</summary>
<p>



**PUT** [/v1/communities/{id}/subscription](../sources/server/internal/endpoints/toggle_community_subscription.go#L12)

Параметры запроса:


* **id** (path, uint64) - айди сообщества


* **subscribe** (form, bool) - подписаться - true, отписаться - false




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Подписка/отписка от блога</summary>
<p>



**PUT** [/v1/blogs/{id}/subscription](../sources/server/internal/endpoints/toogle_blog_subscription.go#L12)

Параметры запроса:


* **id** (path, uint64) - айди блога


* **subscribe** (form, bool) - подписаться - true, отписаться - false




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Подписка/отписка от статьи в блоге</summary>
<p>



**PUT** [/v1/blog_articles/{id}/subscription](../sources/server/internal/endpoints/toogle_article_subscription.go#L12)

Параметры запроса:


* **id** (path, uint64) - айди статьи


* **subscribe** (form, bool) - подписаться - true, отписаться - false




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Лайк/дизлайк статьи в блоге</summary>
<p>



**PUT** [/v1/blog_articles/{id}/like](../sources/server/internal/endpoints/toggle_article_like.go#L11)

Параметры запроса:


* **id** (path, uint64) - айди статьи


* **like** (form, bool) - лайк - true, dislike - false




Схема ответа:

```
{
  likeCount: uint64  # количество лайков
}
```
---

</p>
</details>

<details><summary>Создание первичных книжных полок</summary>
<p>



**POST** [/v1/bookcases/create](../sources/server/internal/endpoints/create_default_bookcases.go#L10)


Схема ответа:

```
{
  bookcaseBlocks: [{     # список блоков книжных полок
    title: string        # название блока
    bookcases: [{        # книжные полки
      id: uint64         # id книжной полки
      isPrivate: bool    # приватная?
      type: int32        # тип
      title: string      # название
      comment: string    # комментарий
      index: uint64      # порядковый номер
      itemCount: uint64  # количество элементов
    }]
  }]
}
```
---

</p>
</details>

<details><summary>Удаление книжной полки</summary>
<p>



**DELETE** [/v1/bookcases/{id}](../sources/server/internal/endpoints/delete_bookcase.go#L11)

Параметры запроса:


* **id** (path, uint64) - id книжной полки




Схема ответа:

```
{}
```
---

</p>
</details>


## Для философов


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

<details><summary>Плюс/минус посту в форуме</summary>
<p>



**PUT** [/v1/forum_messages/{id}/voting](../sources/server/internal/endpoints/toggle_forum_message_voting.go#L12)

Параметры запроса:


* **id** (path, uint64) - id сообщения


* **vote** (form, string) - плюс посту - plus, минус посту - minus, удалить голос - none (для модераторов)




Схема ответа:

```
{}
```
---

</p>
</details>


