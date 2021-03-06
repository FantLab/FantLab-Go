
# Константы


<details><summary>Bookcase_BookcaseType</summary>
<p>

| Int | String |
| --- | --- |
| 0 | BOOKCASE_TYPE_UNKNOWN |
| 1 | BOOKCASE_TYPE_READ |
| 2 | BOOKCASE_TYPE_WAIT |
| 3 | BOOKCASE_TYPE_BUY |
| 4 | BOOKCASE_TYPE_SALE |
| 5 | BOOKCASE_TYPE_FREE |
---

</p>
</details>

<details><summary>Common_Gender</summary>
<p>

| Int | String |
| --- | --- |
| 0 | GENDER_UNKNOWN |
| 1 | GENDER_MALE |
| 2 | GENDER_FEMALE |
---

</p>
</details>

<details><summary>Common_UserClass</summary>
<p>

| Int | String |
| --- | --- |
| 0 | USERCLASS_UNKNOWN |
| 1 | USERCLASS_BEGINNER |
| 2 | USERCLASS_ACTIVIST |
| 3 | USERCLASS_AUTHORITY |
| 4 | USERCLASS_PHILOSOPHER |
| 5 | USERCLASS_MASTER |
| 6 | USERCLASS_GRANDMASTER |
| 7 | USERCLASS_PEACEKEEPER |
| 8 | USERCLASS_PEACEMAKER |
---

</p>
</details>

<details><summary>EditionCorrectnessLevel</summary>
<p>

| Int | String |
| --- | --- |
| 0 | EDITION_CORRECTNESS_LEVEL_UNKNOWN |
| 1 | EDITION_CORRECTNESS_LEVEL_GREEN |
| 2 | EDITION_CORRECTNESS_LEVEL_ORANGE |
| 3 | EDITION_CORRECTNESS_LEVEL_RED |
---

</p>
</details>

<details><summary>EditionType</summary>
<p>

| Int | String |
| --- | --- |
| 0 | EDITION_TYPE_UNKNOWN |
| 1 | EDITION_TYPE_AUTHOR_BOOK |
| 2 | EDITION_TYPE_AUTHOR_COMPILATION |
| 3 | EDITION_TYPE_COMPILATION |
| 4 | EDITION_TYPE_ANTHOLOGY |
| 5 | EDITION_TYPE_CHRESTOMATHY |
| 6 | EDITION_TYPE_MAGAZINE |
| 7 | EDITION_TYPE_FANZINE |
| 8 | EDITION_TYPE_ALMANAC |
| 9 | EDITION_TYPE_NEWSPAPER |
| 10 | EDITION_TYPE_AUDIOBOOK |
| 11 | EDITION_TYPE_ILLUSTRATED_ALBUM |
| 12 | EDITION_TYPE_FILM_STRIP |
---

</p>
</details>

<details><summary>FilmType</summary>
<p>

| Int | String |
| --- | --- |
| 0 | FILM_TYPE_UNKNOWN |
| 1 | FILM_TYPE_FILM |
| 2 | FILM_TYPE_SERIES |
| 3 | FILM_TYPE_EPISODE |
| 4 | FILM_TYPE_DOCUMENTARY |
| 5 | FILM_TYPE_ANIMATION |
| 6 | FILM_TYPE_SHORT |
| 7 | FILM_TYPE_SPECTACLE |
---

</p>
</details>

<details><summary>Forum_Topic_Type</summary>
<p>

| Int | String |
| --- | --- |
| 0 | UNKNOWN_TYPE |
| 1 | TOPIC |
| 2 | POLL |
---

</p>
</details>

<details><summary>WorkType</summary>
<p>

| Int | String |
| --- | --- |
| 0 | WORK_TYPE_UNKNOWN |
| 1 | WORK_TYPE_NOVEL |
| 2 | WORK_TYPE_COMPILATION |
| 3 | WORK_TYPE_SERIES |
| 4 | WORK_TYPE_VERSE |
| 5 | WORK_TYPE_OTHER |
| 6 | WORK_TYPE_FAIRY_TALE |
| 7 | WORK_TYPE_ESSAY |
| 8 | WORK_TYPE_ARTICLE |
| 9 | WORK_TYPE_EPIC_NOVEL |
| 10 | WORK_TYPE_ANTHOLOGY |
| 11 | WORK_TYPE_PLAY |
| 12 | WORK_TYPE_SCREENPLAY |
| 13 | WORK_TYPE_DOCUMENTARY |
| 14 | WORK_TYPE_MICROTALE |
| 15 | WORK_TYPE_DISSERTATION |
| 16 | WORK_TYPE_MONOGRAPH |
| 17 | WORK_TYPE_EDUCATIONAL_PUBLICATION |
| 18 | WORK_TYPE_ENCYCLOPEDIA |
| 19 | WORK_TYPE_MAGAZINE |
| 20 | WORK_TYPE_POEM |
| 21 | WORK_TYPE_POETRY |
| 22 | WORK_TYPE_PROSE_VERSE |
| 23 | WORK_TYPE_COMIC_BOOK |
| 24 | WORK_TYPE_MANGA |
| 25 | WORK_TYPE_GRAPHIC_NOVEL |
| 26 | WORK_TYPE_NOVELETTE |
| 27 | WORK_TYPE_STORY |
| 28 | WORK_TYPE_FEATURE_ARTICLE |
| 29 | WORK_TYPE_REPORTAGE |
| 30 | WORK_TYPE_CONDITIONAL_SERIES |
| 31 | WORK_TYPE_EXCERPT |
| 32 | WORK_TYPE_INTERVIEW |
| 33 | WORK_TYPE_REVIEW |
| 34 | WORK_TYPE_POPULAR_SCIENCE_BOOK |
| 35 | WORK_TYPE_ARTBOOK |
| 36 | WORK_TYPE_LIBRETTO |
---

</p>
</details>

<details><summary>Work_PublishStatus</summary>
<p>

| Int | String |
| --- | --- |
| 0 | PUBLISH_STATUS_UNKNOWN |
| 1 | PUBLISH_STATUS_NOT_FINISHED |
| 2 | PUBLISH_STATUS_NOT_PUBLISHED |
| 3 | PUBLISH_STATUS_NETWORK_PUBLICATION |
| 4 | PUBLISH_STATUS_AVAILABLE_ONLINE |
| 5 | PUBLISH_STATUS_PLANNED_BY_THE_AUTHOR |
---

</p>
</details>


# Эндпойнты


## Общедоступные


<details><summary>Логин</summary>
<p>

Создаёт новый аутентификационный токен для пользователя на основе пары логин/пароль


**POST** [/v1/auth/login](../sources/apiserver/internal/endpoints/login.go#L14)

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

<details><summary>Классификация произведения</summary>
<p>



**GET** [/v1/work/{id}/classification](../sources/apiserver/internal/endpoints/get_work_classification.go#L11)

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



**GET** [/v1/work/{id}/subworks](../sources/apiserver/internal/endpoints/get_work_subworks.go#L11)

Параметры запроса:


* **id** (path, uint64) - айди произведения


* **depth** (query, uint8) - глубина дерева (1 - 5, по умолчанию - 4)




Схема ответа:

```
{
  workId: uint64                                # айди произведения, для которого был запрос
  subworks: [{                                  # произведения, входящие в запрашиваемое
    id: uint64                                  # идентификатор произведения
    origName: string                            # оригинальное название
    rusName: string                             # название на русском
    year: uint64                                # год публикации
    workType: enum (WorkType)                   # тип произведения
    rating: float64                             # рейтинг
    marks: uint64                               # кол-во оценок
    reviews: uint64                             # кол-во отзывов
    plus: bool                                  # является ли произведение дополнительным
    publishStatus: [enum (Work_PublishStatus)]  # статус публикации (не закончено, в планах, etc.)
    subworks: [...]                             # дочерние произведения
  }]
}
```
---

</p>
</details>

<details><summary>Список форумов</summary>
<p>



**GET** [/v1/forums](../sources/apiserver/internal/endpoints/show_forums.go#L11)


Схема ответа:

```
{
  forumBlocks: [{                         # список блоков форумов
    id: uint64                            # id блока форумов
    title: string                         # название
    forums: [{                            # форумы
      id: uint64                          # id форума
      title: string                       # название
      forumDescription: string            # описание
      moderators: [{                      # модераторы
        id: uint64                        # id пользователя
        login: string                     # логин
        name: string                      # имя
        gender: enum (Common_Gender)      # пол
        avatar: string                    # аватар
        class: enum (Common_UserClass)    # класс
        sign: string                      # подпись на форуме
      }]
      stats: {                            # статистика
        topicCount: uint64                # количество тем
        notModeratedTopicCount: uint64    # количество неотмодерированных тем (для модераторов)
        messageCount: uint64              # количество сообщений
        notReadMessageCount: uint64       # количество непрочитанных сообщений (для залогиненных пользователей)
      }
      lastMessage: {                      # последнее сообщение
        id: uint64                        # id сообщения
        topic: {                          # тема, в которую входит сообщение
          id: uint64                      # id темы
          title: string                   # название
        }
        user: {                           # автор
          id: uint64                      # id пользователя
          login: string                   # логин
          name: string                    # имя
          gender: enum (Common_Gender)    # пол
          avatar: string                  # аватар
          class: enum (Common_UserClass)  # класс
          sign: string                    # подпись на форуме
        }
        text: string                      # текст
        date: timestamp                   # дата и время создания
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



**GET** [/v1/forums/{id}](../sources/apiserver/internal/endpoints/show_forum.go#L13)

Параметры запроса:


* **id** (path, uint64) - id форума


* **page** (query, uint64) - номер страницы (по умолчанию - 1)


* **limit** (query, uint64) - кол-во записей на странице (по умолчанию - 20)




Схема ответа:

```
{
  moderators: [{                        # модераторы
    id: uint64                          # id пользователя
    login: string                       # логин
    name: string                        # имя
    gender: enum (Common_Gender)        # пол
    avatar: string                      # аватар
    class: enum (Common_UserClass)      # класс
    sign: string                        # подпись на форуме
  }]
  topics: [{                            # список тем
    id: uint64                          # id темы
    title: string                       # название
    topicType: enum (Forum_Topic_Type)  # тип
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    isClosed: bool                      # тема закрыта?
    isPinned: bool                      # тема закреплена?
    isNotModerated: bool                # тема не отмодерирована? (для модераторов и авторов тем)
    isSubscribed: bool                  # пользователь подписан на тему? (для залогиненных пользователей)
    stats: {                            # статистика
      messageCount: uint64              # количество сообщений
      notReadMessageCount: uint64       # количество непрочитанных сообщений (для залогиненных пользователей)
      viewCount: uint64                 # количество просмотров
    }
    lastMessage: {                      # последнее сообщение
      id: uint64                        # id сообщения
      topic: {                          # тема, в которую входит сообщение
        id: uint64                      # id темы
        title: string                   # название
      }
      user: {                           # автор
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      text: string                      # текст
      date: timestamp                   # дата и время создания
    }
    firstNotReadMessageId: uint64       # id первого непрочитанного сообщения (для залогиненных пользователей)
  }]
  pages: {                              # страницы
    current: uint64                     # текущая
    count: uint64                       # количество
  }
}
```
---

</p>
</details>

<details><summary>Сообщения в теме форума</summary>
<p>



**GET** [/v1/topics/{id}](../sources/apiserver/internal/endpoints/show_forum_topic.go#L15)

Параметры запроса:


* **id** (path, uint64) - id темы


* **page** (query, uint64) - номер страницы (по умолчанию - 1)


* **limit** (query, uint64) - кол-во записей на странице (по умолчанию - 20)


* **sortAsc** (query, uint8) - порядок выдачи (0 - от новых к старым, 1 - наоборот; по умолчанию - 0)




Схема ответа:

```
{
  topic: {                              # тема
    id: uint64                          # id темы
    title: string                       # название
    topicType: enum (Forum_Topic_Type)  # тип
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    isClosed: bool                      # тема закрыта?
    isPinned: bool                      # тема закреплена?
    isNotModerated: bool                # тема не отмодерирована? (для модераторов и авторов тем)
    isSubscribed: bool                  # пользователь подписан на тему? (для залогиненных пользователей)
    stats: {                            # статистика
      messageCount: uint64              # количество сообщений
      notReadMessageCount: uint64       # количество непрочитанных сообщений (для залогиненных пользователей)
      viewCount: uint64                 # количество просмотров
    }
    lastMessage: {                      # последнее сообщение
      id: uint64                        # id сообщения
      topic: {                          # тема, в которую входит сообщение
        id: uint64                      # id темы
        title: string                   # название
      }
      user: {                           # автор
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      text: string                      # текст
      date: timestamp                   # дата и время создания
    }
    firstNotReadMessageId: uint64       # id первого непрочитанного сообщения (для залогиненных пользователей)
  }
  forum: {                              # форум, в который входит тема
    id: uint64                          # id форума
    title: string                       # название
    forumDescription: string            # описание
    moderators: [{                      # модераторы
      id: uint64                        # id пользователя
      login: string                     # логин
      name: string                      # имя
      gender: enum (Common_Gender)      # пол
      avatar: string                    # аватар
      class: enum (Common_UserClass)    # класс
      sign: string                      # подпись на форуме
    }]
    stats: {                            # статистика
      topicCount: uint64                # количество тем
      notModeratedTopicCount: uint64    # количество неотмодерированных тем (для модераторов)
      messageCount: uint64              # количество сообщений
      notReadMessageCount: uint64       # количество непрочитанных сообщений (для залогиненных пользователей)
    }
    lastMessage: {                      # последнее сообщение
      id: uint64                        # id сообщения
      topic: {                          # тема, в которую входит сообщение
        id: uint64                      # id темы
        title: string                   # название
      }
      user: {                           # автор
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      text: string                      # текст
      date: timestamp                   # дата и время создания
    }
  }
  poll: {                               # опрос (если есть)
    answerOptions: [{                   # варианты ответов
      text: string                      # текст
      voterCount: uint64                # количество проголосовавших (если опрос окончен)
    }]
    voters: [{                          # проголосовавшие пользователи (если опрос окончен)
      id: uint64                        # id пользователя
      login: string                     # логин
      name: string                      # имя
      gender: enum (Common_Gender)      # пол
      avatar: string                    # аватар
      class: enum (Common_UserClass)    # класс
      sign: string                      # подпись на форуме
    }]
  }
  pinnedMessage: {                      # закрепленное сообщение, если есть
    id: uint64                          # id сообщения
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    text: string                        # текст
    isCensored: bool                    # текст изъят модератором?
    isUnread: bool                      # сообщение не прочитано? (для залогиненных пользователей)
    attachments: [{                     # аттачи
      url: string                       # ссылка на файл
      size: uint64                      # размер (байт)
    }]
    stats: {                            # статистика
      rating: int64                     # рейтинг
    }
    rights: {                           # права
      canEdit: bool                     # на редактирование
      canDelete: bool                   # на удаление
      canVoteMinus: bool                # на выставление "минуса" сообщению
      canVotePlus: bool                 # на выставление "плюса" сообщению
      canDeleteVotes: bool              # на удаление рейтинга (для модераторов)
    }
  }
  messages: [{                          # сообщения
    id: uint64                          # id сообщения
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    text: string                        # текст
    isCensored: bool                    # текст изъят модератором?
    isUnread: bool                      # сообщение не прочитано? (для залогиненных пользователей)
    attachments: [{                     # аттачи
      url: string                       # ссылка на файл
      size: uint64                      # размер (байт)
    }]
    stats: {                            # статистика
      rating: int64                     # рейтинг
    }
    rights: {                           # права
      canEdit: bool                     # на редактирование
      canDelete: bool                   # на удаление
      canVoteMinus: bool                # на выставление "минуса" сообщению
      canVotePlus: bool                 # на выставление "плюса" сообщению
      canDeleteVotes: bool              # на удаление рейтинга (для модераторов)
    }
  }]
  messageDraft: {                       # черновик, если есть (для залогиненных пользователей)
    topicId: uint64                     # id темы
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    text: string                        # текст
    attachments: [{                     # аттачи
      url: string                       # ссылка на файл
      size: uint64                      # размер (байт)
    }]
  }
  pages: {                              # страницы
    current: uint64                     # текущая
    count: uint64                       # количество
  }
}
```
---

</p>
</details>

<details><summary>Список сообществ</summary>
<p>



**GET** [/v1/communities](../sources/apiserver/internal/endpoints/show_communities.go#L11)


Схема ответа:

```
{
  main: [{                              # основные рубрики
    id: uint64                          # id рубрики
    title: string                       # название
    communityDescription: string        # описание
    rules: string                       # правила
    avatar: string                      # аватар
    stats: {                            # статистика
      articleCount: uint64              # количество статей
      subscriberCount: uint64           # количество подписчиков
    }
    lastArticle: {                      # последняя статья
      id: uint64                        # id статьи
      title: string                     # название
      user: {                           # автор
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
  }]
  additional: [{                        # дополнительные рубрики
    id: uint64                          # id рубрики
    title: string                       # название
    communityDescription: string        # описание
    rules: string                       # правила
    avatar: string                      # аватар
    stats: {                            # статистика
      articleCount: uint64              # количество статей
      subscriberCount: uint64           # количество подписчиков
    }
    lastArticle: {                      # последняя статья
      id: uint64                        # id статьи
      title: string                     # название
      user: {                           # автор
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
  }]
}
```
---

</p>
</details>

<details><summary>Информация о сообществе</summary>
<p>



**GET** [/v1/communities/{id}](../sources/apiserver/internal/endpoints/show_community.go#L14)

Параметры запроса:


* **id** (path, uint64) - айди сообщества


* **page** (query, uint64) - номер страницы (по умолчанию - 1)


* **limit** (query, uint64) - кол-во записей на странице (по умолчанию - 5)




Схема ответа:

```
{
  community: {                          # рубрика
    id: uint64                          # id рубрики
    title: string                       # название
    communityDescription: string        # описание
    rules: string                       # правила
    avatar: string                      # аватар
    stats: {                            # статистика
      articleCount: uint64              # количество статей
      subscriberCount: uint64           # количество подписчиков
    }
    lastArticle: {                      # последняя статья
      id: uint64                        # id статьи
      title: string                     # название
      user: {                           # автор
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
  }
  moderators: [{                        # модераторы
    id: uint64                          # id пользователя
    login: string                       # логин
    name: string                        # имя
    gender: enum (Common_Gender)        # пол
    avatar: string                      # аватар
    class: enum (Common_UserClass)      # класс
    sign: string                        # подпись на форуме
  }]
  authors: [{                           # авторы
    id: uint64                          # id пользователя
    login: string                       # логин
    name: string                        # имя
    gender: enum (Common_Gender)        # пол
    avatar: string                      # аватар
    class: enum (Common_UserClass)      # класс
    sign: string                        # подпись на форуме
  }]
  articles: [{                          # статьи
    id: uint64                          # id статьи
    title: string                       # название
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    text: string                        # текст
    tags: string                        # теги
    attachments: [{                     # аттачи
      url: string                       # ссылка на файл
      size: uint64                      # размер (байт)
    }]
    stats: {                            # статистика
      likeCount: uint64                 # количество лайков
      viewCount: uint64                 # количество просмотров
      commentCount: uint64              # количество комментариев
    }
  }]
  pages: {                              # страницы
    current: uint64                     # текущая
    count: uint64                       # количество
  }
}
```
---

</p>
</details>

<details><summary>Список блогов</summary>
<p>



**GET** [/v1/blogs](../sources/apiserver/internal/endpoints/show_blogs.go#L12)

Параметры запроса:


* **page** (query, uint64) - номер страницы (по умолчанию - 1)


* **limit** (query, uint64) - кол-во записей на странице (по умолчанию - 5)


* **sort** (query, string) - сортировать по (кол-ву тем в блоге - article, кол-ву подписчиков - subscriber, дате обновления от новых к старым - update (по умолчанию))




Схема ответа:

```
{
  blogs: [{                             # блоги
    id: uint64                          # id блога
    user: {                             # автор
      id: uint64                        # id пользователя
      login: string                     # логин
      name: string                      # имя
      gender: enum (Common_Gender)      # пол
      avatar: string                    # аватар
      class: enum (Common_UserClass)    # класс
      sign: string                      # подпись на форуме
    }
    isClosed: bool                      # блог закрыт?
    stats: {                            # статистика
      articleCount: uint64              # количество статей
      subscriberCount: uint64           # количество подписчиков
    }
    lastArticle: {                      # последняя статья
      id: uint64                        # id статьи
      title: string                     # название
      user: {                           # автор
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
  }]
  pages: {                              # страницы
    current: uint64                     # текущая
    count: uint64                       # количество
  }
}
```
---

</p>
</details>

<details><summary>Список статей в блоге</summary>
<p>



**GET** [/v1/blogs/{id}](../sources/apiserver/internal/endpoints/show_blog.go#L14)

Параметры запроса:


* **id** (path, uint64) - айди блога


* **page** (query, uint64) - номер страницы (по умолчанию - 1)


* **limit** (query, uint64) - кол-во записей на странице (по умолчанию - 20)




Схема ответа:

```
{
  articles: [{                          # статьи
    id: uint64                          # id статьи
    title: string                       # название
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    text: string                        # текст
    tags: string                        # теги
    attachments: [{                     # аттачи
      url: string                       # ссылка на файл
      size: uint64                      # размер (байт)
    }]
    stats: {                            # статистика
      likeCount: uint64                 # количество лайков
      viewCount: uint64                 # количество просмотров
      commentCount: uint64              # количество комментариев
    }
  }]
  pages: {                              # страницы
    current: uint64                     # текущая
    count: uint64                       # количество
  }
}
```
---

</p>
</details>

<details><summary>Статья в блоге</summary>
<p>



**GET** [/v1/blog_articles/{id}](../sources/apiserver/internal/endpoints/show_article.go#L14)

Параметры запроса:


* **id** (path, uint64) - айди статьи




Схема ответа:

```
{
  article: {                            # статья
    id: uint64                          # id статьи
    title: string                       # название
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    text: string                        # текст
    tags: string                        # теги
    attachments: [{                     # аттачи
      url: string                       # ссылка на файл
      size: uint64                      # размер (байт)
    }]
    stats: {                            # статистика
      likeCount: uint64                 # количество лайков
      viewCount: uint64                 # количество просмотров
      commentCount: uint64              # количество комментариев
    }
  }
}
```
---

</p>
</details>

<details><summary>Список жанров</summary>
<p>



**GET** [/v1/allgenres](../sources/apiserver/internal/endpoints/show_genres.go#L11)


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

<details><summary>Комментарии к статье в блоге</summary>
<p>



**GET** [/v1/blog_articles/{id}/comments](../sources/apiserver/internal/endpoints/show_blog_article_comments.go#L13)

Параметры запроса:


* **id** (path, uint64) - id статьи


* **after** (query, string) - дата, после которой искать сообщения (в формате RFC3339)


* **count** (query, uint64) - кол-во комментариев верхнего уровня (по умолчанию - 10, [5, 20])


* **sortAsc** (query, uint8) - порядок выдачи (0 - от новых к старым, 1 - наоборот; по умолчанию - 0)




Схема ответа:

```
{
  comments: [{                          # список комментариев
    id: uint64                          # id сообщения
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    text: string                        # текст сообщения
    isCensored: bool                    # текст изъят модератором?
    answers: [...]                      # ответы на комментарий
  }]
  totalCount: uint64                    # общее ко-во комментариев у поста
}
```
---

</p>
</details>

<details><summary>Список книжных полок пользователя</summary>
<p>



**GET** [/v1/users/{id}/bookcases](../sources/apiserver/internal/endpoints/show_bookcases.go#L13)

Параметры запроса:


* **id** (path, uint64) - id пользователя




Схема ответа:

```
{
  bookcaseBlocks: [{                      # список блоков книжных полок
    title: string                         # название блока
    bookcases: [{                         # книжные полки
      id: uint64                          # id книжной полки
      isPrivate: bool                     # приватная?
      type: enum (Bookcase_BookcaseType)  # тип
      title: string                       # название
      comment: string                     # комментарий
      index: uint64                       # порядковый номер
      itemCount: uint64                   # количество элементов
    }]
  }]
}
```
---

</p>
</details>

<details><summary>Содержимое полки с изданиями</summary>
<p>



**GET** [/v1/edition_bookcases/{id}](../sources/apiserver/internal/endpoints/show_edition_bookcase.go#L14)

Параметры запроса:


* **id** (path, uint64) - id книжной полки


* **page** (query, uint64) - номер страницы (>0, по умолчанию - 1)


* **limit** (query, uint64) - кол-во элементов на странице ([5..50], по умолчанию - 50)


* **sort** (query, string) - сортировать по: порядку - order (по умолчанию, если иное не задано в настройках полки), автору - author, названию - title, году - year




Схема ответа:

```
{
  bookcase: {                                         # информация о полке
    id: uint64                                        # id книжной полки
    isPrivate: bool                                   # приватная?
    type: enum (Bookcase_BookcaseType)                # тип
    title: string                                     # название
    comment: string                                   # комментарий
  }
  editions: [{                                        # список изданий на полке
    itemId: uint64                                    # id item-а на полке
    id: uint64                                        # id издания
    type: enum (EditionType)                          # тип (авторская книга/сборник/etc; может отсутствовать, если не задан)
    correctnessLevel: enum (EditionCorrectnessLevel)  # уровень проверенности
    cover: string                                     # URL обложки
    authors: string                                   # авторы
    title: string                                     # название
    year: uint64                                      # год публикации
    publishers: string                                # издательства
    description: string                               # описание
    plannedPublicationDate: string                    # планируемая дата издания (если издание еще не опубликовано)
    offers: {                                         # предложения в магазинах
      ozon: {                                         # предложение на Озоне
        url: string                                   # URL предложения
        price: uint64                                 # цена
      }
      labirint: {                                     # предложение на Лабиринте
        url: string                                   # URL предложения
        price: uint64                                 # цена
      }
    }
    comment: string                                   # комментарий
  }]
  pages: {                                            # страницы
    current: uint64                                   # текущая
    count: uint64                                     # количество
  }
}
```
---

</p>
</details>

<details><summary>Содержимое полки с произведениями</summary>
<p>



**GET** [/v1/work_bookcases/{id}](../sources/apiserver/internal/endpoints/show_work_bookcase.go#L14)

Параметры запроса:


* **id** (path, uint64) - id книжной полки


* **page** (query, uint64) - номер страницы (>0, по умолчанию - 1)


* **limit** (query, uint64) - кол-во элементов на странице ([5..50], по умолчанию - 50)


* **sort** (query, string) - сортировать по: порядку - order (по умолчанию, если иное не задано в настройках полки), автору - author, названию - title, оригинальному названию - orig_title, году - year, количеству оценок - mark_count, средней оценке - avg_mark




Схема ответа:

```
{
  bookcase: {                           # информация о полке
    id: uint64                          # id книжной полки
    isPrivate: bool                     # приватная?
    type: enum (Bookcase_BookcaseType)  # тип
    title: string                       # название
    comment: string                     # комментарий
  }
  works: [{                             # список произведений на полке
    itemId: uint64                      # id item-а на полке
    id: uint64                          # id произведения
    type: enum (WorkType)               # тип (роман/сборник/etc; может отсутствовать)
    authors: [{                         # авторы
      id: uint64                        # id автора
      name: string                      # имя на русском языке
      isOpened: bool                    # страница открыта?
    }]
    title: string                       # название на русском языке
    originalTitle: string               # название в оригинале
    alternativeTitles: string           # альтернативные названия
    note: string                        # примечание
    year: int64                         # год
    description: string                 # описание
    isPublished: bool                   # опубликовано?
    stats: {                            # статистика
      averageMark: float64              # средняя оценка
      markCount: uint64                 # количество оценок
      responseCount: uint64             # количество отзывов
    }
    own: {                              # персональное
      mark: uint64                      # собственная оценка произведению
      isResponsePublished: bool         # опубликован отзыв?
    }
    comment: string                     # комментарий
  }]
  pages: {                              # страницы
    current: uint64                     # текущая
    count: uint64                       # количество
  }
}
```
---

</p>
</details>

<details><summary>Содержимое полки с фильмами</summary>
<p>



**GET** [/v1/film_bookcases/{id}](../sources/apiserver/internal/endpoints/show_film_bookcase.go#L14)

Параметры запроса:


* **id** (path, uint64) - id книжной полки


* **page** (query, uint64) - номер страницы (>0, по умолчанию - 1)


* **limit** (query, uint64) - кол-во элементов на странице ([5..50], по умолчанию - 50)


* **sort** (query, string) - сортировать по: порядку - order (по умолчанию, если иное не задано в настройках полки), названию - title, оригинальному названию - orig_title




Схема ответа:

```
{
  bookcase: {                           # информация о полке
    id: uint64                          # id книжной полки
    isPrivate: bool                     # приватная?
    type: enum (Bookcase_BookcaseType)  # тип
    title: string                       # название
    comment: string                     # комментарий
  }
  films: [{                             # список фильмов на полке
    itemId: uint64                      # id item-а на полке
    id: uint64                          # id фильма
    type: enum (FilmType)               # тип (фильм/сериал/etc; может отсутствовать, если не задан)
    poster: string                      # URL постера
    title: string                       # название на русском языке
    originalTitle: string               # название в оригинале
    year: uint64                        # год выпуска (для всего, кроме сериалов)
    startYear: uint64                   # год старта трансляции (для сериалов)
    endYear: uint64                     # год окончания трансляции (для сериалов)
    countries: string                   # страны производства
    genres: string                      # жанры
    directors: string                   # режиссеры
    screenWriters: string               # сценаристы
    actors: string                      # актеры
    description: string                 # описание
    comment: string                     # комментарий
  }]
  pages: {                              # страницы
    current: uint64                     # текущая
    count: uint64                       # количество
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


**POST** [/v1/auth/refresh](../sources/apiserver/internal/endpoints/refresh_auth.go#L16)

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


## С проверкой на бан


<details><summary>Создание нового сообщения в личной переписке</summary>
<p>



**POST** [/v1/users/{id}/private_message](../sources/apiserver/internal/endpoints/add_private_message.go#L15)

Параметры запроса:


* **id** (path, uint64) - id пользователя, которому отправляется сообщение


* **message** (form, string) - текст сообщения


* **send_copy_via_email** (form, bool) - отправить копию посредством Email?




Схема ответа:

```
{
  message: {                            # сообщение
    id: uint64                          # id сообщения
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    text: string                        # текст
    number: uint64                      # номер сообщения
    isRead: bool                        # прочитано?
  }
}
```
---

</p>
</details>

<details><summary>Классификация произведения пользователем</summary>
<p>



**GET** [/v1/work/{id}/userclassification](../sources/apiserver/internal/endpoints/get_user_work_genres.go#L11)

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

<details><summary>Выставление оценки произведению</summary>
<p>



**POST** [/v1/work/{id}/mark](../sources/apiserver/internal/endpoints/set_mark.go#L12)

Параметры запроса:


* **id** (path, uint64) - id произведения


* **mark** (form, uint8) - оценка (1-10, 0 - удалить)




Схема ответа:

```
{
  averageMark: float64  # средняя оценка произведения (-1, если произведение участвует в Фантлабораторной работе)
  markCount: int64      # количество оценок (-1, если произведение участвует в Фантлабораторной работе)
}
```
---

</p>
</details>

<details><summary>Создание отзыва на произведение</summary>
<p>



**POST** [/v1/work/{id}/response](../sources/apiserver/internal/endpoints/add_response.go#L14)

Параметры запроса:


* **id** (path, uint64) - id произведения


* **response** (form, string) - текст отзыва




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Редактирование отзыва на произведение</summary>
<p>



**PUT** [/v1/response/{id}](../sources/apiserver/internal/endpoints/edit_response.go#L14)

Параметры запроса:


* **id** (path, uint64) - id отзыва


* **response** (form, string) - новый текст отзыва




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Плюс/минус отзыву на произведение</summary>
<p>



**PUT** [/v1/response/{id}/voting](../sources/apiserver/internal/endpoints/vote_response.go#L13)

Параметры запроса:


* **id** (path, uint64) - id отзыва


* **vote_plus** (form, bool) - голос (true - плюс, false - минус)




Схема ответа:

```
{
  rating: int64  # рейтинг отзыва
}
```
---

</p>
</details>

<details><summary>Удаление отзыва на произведение</summary>
<p>



**DELETE** [/v1/response/{id}](../sources/apiserver/internal/endpoints/delete_response.go#L13)

Параметры запроса:


* **id** (path, uint64) - id отзыва




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Создание нового сообщения в форуме</summary>
<p>



**POST** [/v1/topics/{id}/message](../sources/apiserver/internal/endpoints/add_forum_message.go#L14)

Параметры запроса:


* **id** (path, uint64) - id темы


* **message** (form, string) - текст сообщения




Схема ответа:

```
{
  message: {                            # сообщение
    id: uint64                          # id сообщения
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    text: string                        # текст
    isCensored: bool                    # текст изъят модератором?
    isUnread: bool                      # сообщение не прочитано? (для залогиненных пользователей)
    attachments: [{                     # аттачи
      url: string                       # ссылка на файл
      size: uint64                      # размер (байт)
    }]
    stats: {                            # статистика
      rating: int64                     # рейтинг
    }
    rights: {                           # права
      canEdit: bool                     # на редактирование
      canDelete: bool                   # на удаление
      canVoteMinus: bool                # на выставление "минуса" сообщению
      canVotePlus: bool                 # на выставление "плюса" сообщению
      canDeleteVotes: bool              # на удаление рейтинга (для модераторов)
    }
  }
}
```
---

</p>
</details>

<details><summary>Редактирование сообщения в форуме</summary>
<p>



**PUT** [/v1/forum_messages/{id}](../sources/apiserver/internal/endpoints/edit_forum_message.go#L16)

Параметры запроса:


* **id** (path, uint64) - id сообщения


* **message** (form, string) - новый текст сообщения




Схема ответа:

```
{
  message: {                            # сообщение
    id: uint64                          # id сообщения
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    text: string                        # текст
    isCensored: bool                    # текст изъят модератором?
    isUnread: bool                      # сообщение не прочитано? (для залогиненных пользователей)
    attachments: [{                     # аттачи
      url: string                       # ссылка на файл
      size: uint64                      # размер (байт)
    }]
    stats: {                            # статистика
      rating: int64                     # рейтинг
    }
    rights: {                           # права
      canEdit: bool                     # на редактирование
      canDelete: bool                   # на удаление
      canVoteMinus: bool                # на выставление "минуса" сообщению
      canVotePlus: bool                 # на выставление "плюса" сообщению
      canDeleteVotes: bool              # на удаление рейтинга (для модераторов)
    }
  }
}
```
---

</p>
</details>

<details><summary>Удаление сообщения в форуме</summary>
<p>



**DELETE** [/v1/forum_messages/{id}](../sources/apiserver/internal/endpoints/delete_forum_message.go#L14)

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



**GET** [/v1/forum_messages/{id}/file_upload_url](../sources/apiserver/internal/endpoints/get_forum_message_file_upload_url.go#L15)

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



**DELETE** [/v1/forum_messages/{id}/file](../sources/apiserver/internal/endpoints/delete_forum_message_file.go#L15)

Параметры запроса:


* **id** (path, uint64) - id сообщения


* **file_name** (form, string) - полное имя файла (с расширением)




Схема ответа:

```
{
  message: {                            # сообщение
    id: uint64                          # id сообщения
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    text: string                        # текст
    isCensored: bool                    # текст изъят модератором?
    isUnread: bool                      # сообщение не прочитано? (для залогиненных пользователей)
    attachments: [{                     # аттачи
      url: string                       # ссылка на файл
      size: uint64                      # размер (байт)
    }]
    stats: {                            # статистика
      rating: int64                     # рейтинг
    }
    rights: {                           # права
      canEdit: bool                     # на редактирование
      canDelete: bool                   # на удаление
      canVoteMinus: bool                # на выставление "минуса" сообщению
      canVotePlus: bool                 # на выставление "плюса" сообщению
      canDeleteVotes: bool              # на удаление рейтинга (для модераторов)
    }
  }
}
```
---

</p>
</details>

<details><summary>Удаление оценок у сообщения в форуме (для модераторов)</summary>
<p>



**DELETE** [/v1/forum_messages/{id}/voting](../sources/apiserver/internal/endpoints/delete_forum_message_votes.go#L11)

Параметры запроса:


* **id** (path, uint64) - id сообщения




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Сохранение черновика сообщения в форуме</summary>
<p>



**PUT** [/v1/topics/{id}/message_draft](../sources/apiserver/internal/endpoints/save_forum_message_draft.go#L14)

Параметры запроса:


* **id** (path, uint64) - id темы


* **message** (form, string) - текст сообщения




Схема ответа:

```
{
  messageDraft: {                       # черновик сообщения
    topicId: uint64                     # id темы
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    text: string                        # текст
    attachments: [{                     # аттачи
      url: string                       # ссылка на файл
      size: uint64                      # размер (байт)
    }]
  }
}
```
---

</p>
</details>

<details><summary>Подтверждение черновика сообщения в форуме</summary>
<p>



**POST** [/v1/topics/{id}/message_draft](../sources/apiserver/internal/endpoints/confirm_forum_message_draft.go#L17)

Параметры запроса:


* **id** (path, uint64) - id темы




Схема ответа:

```
{
  message: {                            # сообщение
    id: uint64                          # id сообщения
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    text: string                        # текст
    isCensored: bool                    # текст изъят модератором?
    isUnread: bool                      # сообщение не прочитано? (для залогиненных пользователей)
    attachments: [{                     # аттачи
      url: string                       # ссылка на файл
      size: uint64                      # размер (байт)
    }]
    stats: {                            # статистика
      rating: int64                     # рейтинг
    }
    rights: {                           # права
      canEdit: bool                     # на редактирование
      canDelete: bool                   # на удаление
      canVoteMinus: bool                # на выставление "минуса" сообщению
      canVotePlus: bool                 # на выставление "плюса" сообщению
      canDeleteVotes: bool              # на удаление рейтинга (для модераторов)
    }
  }
}
```
---

</p>
</details>

<details><summary>Отмена черновика сообщения в форуме</summary>
<p>



**DELETE** [/v1/topics/{id}/message_draft](../sources/apiserver/internal/endpoints/cancel_forum_message_draft.go#L13)

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



**GET** [/v1/topics/{id}/message_draft/file_upload_url](../sources/apiserver/internal/endpoints/get_forum_message_draft_file_upload_url.go#L15)

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



**DELETE** [/v1/topics/{id}/message_draft/file](../sources/apiserver/internal/endpoints/delete_forum_message_draft_file.go#L15)

Параметры запроса:


* **id** (path, uint64) - id темы


* **file_name** (form, string) - полное имя файла (с расширением)




Схема ответа:

```
{
  messageDraft: {                       # черновик сообщения
    topicId: uint64                     # id темы
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    text: string                        # текст
    attachments: [{                     # аттачи
      url: string                       # ссылка на файл
      size: uint64                      # размер (байт)
    }]
  }
}
```
---

</p>
</details>

<details><summary>Подписка/отписка от темы форума</summary>
<p>



**PUT** [/v1/topics/{id}/subscription](../sources/apiserver/internal/endpoints/toggle_forum_topic_subscription.go#L11)

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



**POST** [/v1/blog_articles/{id}/comment](../sources/apiserver/internal/endpoints/add_blog_article_comment.go#L13)

Параметры запроса:


* **id** (path, uint64) - id статьи


* **comment** (form, string) - текст комментария (непустой)


* **parent_comment_id** (form, uint64) - id родительского комментария (0, если комментарий 1-го уровня вложенности)




Схема ответа:

```
{
  comment: {                            # комментарий
    id: uint64                          # id сообщения
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    text: string                        # текст сообщения
    isCensored: bool                    # текст изъят модератором?
    answers: [...]                      # ответы на комментарий
  }
}
```
---

</p>
</details>

<details><summary>Редактирование комментария к статье в блоге</summary>
<p>



**PUT** [/v1/blog_article_comments/{id}](../sources/apiserver/internal/endpoints/edit_blog_article_comment.go#L14)

Параметры запроса:


* **id** (path, uint64) - id комментария


* **comment** (form, string) - текст комментария (непустой)




Схема ответа:

```
{
  comment: {                            # комментарий
    id: uint64                          # id сообщения
    creation: {                         # данные о создании
      user: {                           # пользователь
        id: uint64                      # id пользователя
        login: string                   # логин
        name: string                    # имя
        gender: enum (Common_Gender)    # пол
        avatar: string                  # аватар
        class: enum (Common_UserClass)  # класс
        sign: string                    # подпись на форуме
      }
      date: timestamp                   # дата создания
    }
    text: string                        # текст сообщения
    isCensored: bool                    # текст изъят модератором?
    answers: [...]                      # ответы на комментарий
  }
}
```
---

</p>
</details>

<details><summary>Удаление комментария к статье в блоге</summary>
<p>



**DELETE** [/v1/blog_article_comments/{id}](../sources/apiserver/internal/endpoints/delete_blog_article_comment.go#L13)

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



**PUT** [/v1/communities/{id}/subscription](../sources/apiserver/internal/endpoints/toggle_community_subscription.go#L12)

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



**PUT** [/v1/blogs/{id}/subscription](../sources/apiserver/internal/endpoints/toogle_blog_subscription.go#L12)

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



**PUT** [/v1/blog_articles/{id}/subscription](../sources/apiserver/internal/endpoints/toogle_article_subscription.go#L12)

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



**PUT** [/v1/blog_articles/{id}/like](../sources/apiserver/internal/endpoints/toggle_article_like.go#L12)

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



**POST** [/v1/bookcases/default](../sources/apiserver/internal/endpoints/create_default_bookcases.go#L11)


Схема ответа:

```
{
  bookcaseBlocks: [{                      # список блоков книжных полок
    title: string                         # название блока
    bookcases: [{                         # книжные полки
      id: uint64                          # id книжной полки
      isPrivate: bool                     # приватная?
      type: enum (Bookcase_BookcaseType)  # тип
      title: string                       # название
      comment: string                     # комментарий
      index: uint64                       # порядковый номер
      itemCount: uint64                   # количество элементов
    }]
  }]
}
```
---

</p>
</details>

<details><summary>Создание книжной полки</summary>
<p>



**POST** [/v1/bookcases](../sources/apiserver/internal/endpoints/add_bookcase.go#L15)

Параметры запроса:


* **title** (form, string) - название


* **group** (form, string) - группа, в которую входит полка (edition - издания, work - произведения, film - фильмы)


* **type** (form, string) - тип полки (sale - на продажу, buy - купить, read - читать, wait - ожидаю, free - прочее)


* **description** (form, string) - описание, до 50 символов (иначе будет обрезано), может быть пустым


* **is_private** (form, bool) - приватная?


* **items** (form, string) - item-ы в формате [{"id1":"comment1"},...,{"idN":"commentN"}], id - это editionId для изданий etc, commentN может быть пустым




Схема ответа:

```
{
  id: uint64  # идентификатор
}
```
---

</p>
</details>

<details><summary>Редактирование книжной полки</summary>
<p>



**PUT** [/v1/bookcases/{id}](../sources/apiserver/internal/endpoints/edit_bookcase.go#L15)

Параметры запроса:


* **id** (path, uint64) - id полки


* **title** (form, string) - название


* **type** (form, string) - тип полки (sale - на продажу, buy - купить, read - читать, wait - ожидаю, free - прочее)


* **description** (form, string) - описание, до 50 символов (иначе будет обрезано), может быть пустым


* **is_private** (form, bool) - приватная?


* **sort** (form, string) - сортировка, издания по: порядку - order (по умолчанию), автору - author, названию - title, году - year
произведения по: порядку - order (по умолчанию), автору - author, названию - title, оригинальному названию - orig_title, году - year, количеству оценок - mark_count, средней оценке - avg_mark
фильмы по: порядку - order (по умолчанию), названию - title, оригинальному названию - orig_title


* **items** (form, string) - item-ы в формате [{"id1":"comment1"},...,{"idN":"commentN"}], id - это editionId для изданий etc, commentN может быть пустым




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Изменение порядка сортировки книжных полок внутри блоков</summary>
<p>



**PUT** [/v1/bookcases/order](../sources/apiserver/internal/endpoints/change_bookcases_order.go#L11)

Параметры запроса:


* **order** (form, string) - новый порядок сортировки в формате {"bookcaseId1":index1,...,"bookcaseIdN":indexN}, indexN > 0




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Добавление item-а на книжную полку</summary>
<p>



**POST** [/v1/bookcases/{id}/items](../sources/apiserver/internal/endpoints/add_bookcase_item.go#L13)

Параметры запроса:


* **id** (path, uint64) - id полки


* **group** (form, string) - группа, в которую входит полка (edition - издания, work - произведения, film - фильмы)


* **item_id** (form, uint64) - id item-а, который необходимо добавить на полку (editionId для изданий etc)




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Редактирование комментария к item-у книжной полки</summary>
<p>



**PUT** [/v1/bookcase_items/{id}/comment](../sources/apiserver/internal/endpoints/edit_bookcase_item_comment.go#L14)

Параметры запроса:


* **id** (path, uint64) - id item-а книжной полки


* **comment** (form, string) - текст комментария




Схема ответа:

```
{
  comment: string  # текст комментария
}
```
---

</p>
</details>

<details><summary>Удаление item-а с книжной полки</summary>
<p>



**DELETE** [/v1/bookcase_items/{id}](../sources/apiserver/internal/endpoints/delete_bookcase_item.go#L12)

Параметры запроса:


* **id** (path, uint64) - id item-а книжной полки




Схема ответа:

```
{}
```
---

</p>
</details>

<details><summary>Удаление книжной полки</summary>
<p>



**DELETE** [/v1/bookcases/{id}](../sources/apiserver/internal/endpoints/delete_bookcase.go#L12)

Параметры запроса:


* **id** (path, uint64) - id книжной полки




Схема ответа:

```
{}
```
---

</p>
</details>


## Для авторитетов


<details><summary>Плюс/минус посту в форуме</summary>
<p>



**PUT** [/v1/forum_messages/{id}/voting](../sources/apiserver/internal/endpoints/vote_forum_message.go#L11)

Параметры запроса:


* **id** (path, uint64) - id сообщения


* **vote_plus** (form, bool) - голос (true - плюс, false - минус)




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



**PUT** [/v1/work/{id}/userclassification](../sources/apiserver/internal/endpoints/set_work_genres.go#L14)

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


