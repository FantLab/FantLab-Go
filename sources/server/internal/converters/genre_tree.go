package converters

import (
	"errors"
	"fantlab/server/internal/db"
	"strings"
)

type genreIdNode struct {
	id       int64
	children []*genreIdNode
	parent   *genreIdNode
}

func (node *genreIdNode) append(child *genreIdNode) *genreIdNode {
	child.parent = node
	node.children = append(node.children, child)
	return node
}

type GenreTree struct {
	data  *db.WorkGenresDBResponse
	root  *genreIdNode
	table map[int64]*genreIdNode
}

func MakeGenreTree(dbResponse *db.WorkGenresDBResponse) *GenreTree {
	root := &genreIdNode{}

	genresTable := make(map[int64]*genreIdNode, len(dbResponse.Genres)+len(dbResponse.GenreGroups))

	// для удобства, группы жанров - тоже жанры, но с отрицательными идентификаторами

	for _, dbGroup := range dbResponse.GenreGroups {
		groupNode := &genreIdNode{id: -int64(dbGroup.Id)}
		genresTable[groupNode.id] = groupNode
		root.append(groupNode)
	}

	for _, dbGenre := range dbResponse.Genres {
		nodeId := int64(dbGenre.Id)
		genresTable[nodeId] = &genreIdNode{id: nodeId}
	}

	for _, dbGenre := range dbResponse.Genres {
		genre := genresTable[int64(dbGenre.Id)]

		parentGenre := genresTable[int64(dbGenre.ParentId)]

		if parentGenre != nil {
			parentGenre.append(genre)
		} else {
			groupNode := genresTable[-int64(dbGenre.GroupId)]

			if groupNode != nil {
				groupNode.append(genre)
			}
		}
	}

	tree := &GenreTree{
		data:  dbResponse,
		root:  root,
		table: genresTable,
	}

	return tree
}

func (tree *GenreTree) Data() *db.WorkGenresDBResponse {
	return tree.data
}

func (tree *GenreTree) SelectGenreIdsWithParents(genreIds []uint64) map[uint64]bool {
	result := make(map[uint64]bool)

	for _, genreId := range genreIds {
		node := tree.table[int64(genreId)]

		for node != nil {
			if node.id > 0 {
				result[uint64(node.id)] = true
			}

			node = node.parent
		}
	}

	return result
}

func (tree *GenreTree) CheckRequiredGroupsForGenreIds(genreIds []uint64) error {
	var requiredGenreGroups = map[int64]string{
		-1: "Жанры/поджанры",
		-3: "Место действия",
		-4: "Время действия",
		-5: "Возраст читателя",
	}

	for _, genreId := range genreIds {
		node := tree.table[int64(genreId)]

		for {
			if node == nil {
				return errors.New("Неизвестная характеристика")
			}

			if node.id < 0 {
				delete(requiredGenreGroups, node.id)
				break
			}

			node = node.parent
		}
	}

	if len(requiredGenreGroups) == 0 {
		return nil
	}

	var groupNames []string

	for _, groupName := range requiredGenreGroups {
		groupNames = append(groupNames, groupName)
	}

	return errors.New("Выберите характеристики из следующих групп: " + strings.Join(groupNames, ", "))
}
