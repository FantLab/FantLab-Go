package postprocessing

import (
	"fantlab/core/db"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	authorTagsRegex = regexp.MustCompile(`\[autor=(\d+)](.+?)\[/autor]`)
	workTagsRegex   = regexp.MustCompile(`\[work=(\d+)](.+?)\[/work]`)
)

func SortBookcaseEditionsByAuthor(editions []db.BookcaseEdition, fetchAuthors func([]uint64) ([]db.Autor, error)) error {
	sortIndexes := make(map[uint64]int, len(editions))
	// Для каждого издания составляем список авторов (если у автора нет страницы, пишем 0)
	editionsAuthorIds := make([][]uint64, len(editions))

	for index, edition := range editions {
		sortIndexes[edition.EditionId] = index

		if len(edition.Autors) > 0 {
			authors := strings.Split(edition.Autors, ", ")
			for _, author := range authors {
				if authorTagsRegex.MatchString(author) {
					// В теге autor, по идее, всегда валидный id, поэтому ошибки конвертации быть не может
					authorId, _ := strconv.ParseUint(authorTagsRegex.ReplaceAllString(author, "$1"), 10, 0)
					editionsAuthorIds[index] = append(editionsAuthorIds[index], authorId)
				} else {
					editionsAuthorIds[index] = append(editionsAuthorIds[index], 0)
				}
			}
		}
	}

	// Составляем список уникальных авторов
	uniqueAuthorIdsMap := make(map[uint64]bool)
	for _, editionAuthorIds := range editionsAuthorIds {
		for _, editionAuthorId := range editionAuthorIds {
			uniqueAuthorIdsMap[editionAuthorId] = editionAuthorId > 0
		}
	}

	// Перегоняем в слайс
	var uniqueAuthorIds []uint64
	for uniqueAuthorId := range uniqueAuthorIdsMap {
		uniqueAuthorIds = append(uniqueAuthorIds, uniqueAuthorId)
	}

	// Запрашиваем сортировочные имена по каждому автору, у которого есть страница
	dbAuthors, err := fetchAuthors(uniqueAuthorIds)

	if err != nil {
		return err
	}

	// Загоняем сортировочные имена в мапу для быстрого доступа
	authorsSortNames := make(map[uint64]string, len(dbAuthors))
	for _, dbAuthor := range dbAuthors {
		authorsSortNames[dbAuthor.AutorId] = dbAuthor.ShortRusName
	}

	authorNames := make(map[uint64]string, len(editions))
	for editionIndex, edition := range editions {
		if len(editionsAuthorIds[editionIndex]) == 0 {
			// Костыль для вывода в самом верху списка
			authorNames[edition.EditionId] = "   "
		} else {
			// NOTE В Perl-бэке логика значительно отличается:
			// 1. случаи с авторами в тегах и без обрабатываются по-разному
			// 2. если есть автор в тегах, то учитывается только он, все остальные игнорятся
			// 3. если все авторы без тегов, то бьются по запятой и реверсируются для порядка "фамилия имя"
			//    (но логика кривая, см. https://fantlab.ru/edition3066)
			// 4. есть (условно) отдельная обработка антологий
			authors := strings.Split(edition.Autors, ", ")
			names := make([]string, len(authors))
			for authorIndex, author := range authors {
				if editionsAuthorIds[editionIndex][authorIndex] > 0 {
					authorId := editionsAuthorIds[editionIndex][authorIndex]
					authorSortName := authorsSortNames[authorId]
					if len(authorSortName) > 0 {
						names[authorIndex] = authorSortName
						continue
					}
				}
				// Ориентируемся на порядок "имя фамилия". Попадаем сюда или когда у автора нет страницы,
				// или когда нет сортировочного имени
				dividedName := strings.Split(author, " ")
				revertedName := []string{dividedName[len(dividedName)-1]}
				revertedName = append(revertedName, dividedName[:len(dividedName)-1]...)
				names[authorIndex] = strings.Join(revertedName, " ")
			}
			authorNames[edition.EditionId] = strings.Join(names, ", ")
		}
	}

	// Сортируем: 1. по автору, 2. по порядку сортировки
	sort.Slice(editions, func(i, j int) bool {
		e1 := editions[i]
		e2 := editions[j]
		if authorNames[e1.EditionId] == authorNames[e2.EditionId] {
			return sortIndexes[e1.EditionId] < sortIndexes[e2.EditionId]
		} else {
			return authorNames[e1.EditionId] < authorNames[e2.EditionId]
		}
	})

	return nil
}

func SortBookcaseEditionsByTitle(editions []db.BookcaseEdition) {
	sortIndexes := make(map[uint64]int, len(editions))
	editionSortTitles := make(map[uint64]string, len(editions))

	for index, edition := range editions {
		sortIndexes[edition.EditionId] = index

		var title string
		if workTagsRegex.MatchString(edition.Name) {
			title = workTagsRegex.ReplaceAllString(edition.Name, "$2")
		} else {
			title = edition.Name
		}
		editionSortTitles[edition.EditionId] = title
	}

	// Сортируем: 1. по названию, 2. по порядку сортировки
	sort.Slice(editions, func(i, j int) bool {
		e1 := editions[i]
		e2 := editions[j]
		if editionSortTitles[e1.EditionId] == editionSortTitles[e2.EditionId] {
			return sortIndexes[e1.EditionId] < sortIndexes[e2.EditionId]
		} else {
			return editionSortTitles[e1.EditionId] < editionSortTitles[e2.EditionId]
		}
	})
}

func SortBookcaseEditionsByYear(editions []db.BookcaseEdition) {
	sortIndexes := make(map[uint64]int, len(editions))

	for index, edition := range editions {
		sortIndexes[edition.EditionId] = index
	}

	// Сортируем: 1. по году, 2. по порядку сортировки
	sort.Slice(editions, func(i, j int) bool {
		e1 := editions[i]
		e2 := editions[j]
		if e1.Year == e2.Year {
			return sortIndexes[e1.EditionId] < sortIndexes[e2.EditionId]
		} else {
			return e1.Year < e2.Year
		}
	})
}
