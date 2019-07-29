package genresapi

import (
	"errors"
	"fantlab/db"
	"fantlab/pb"
	"strings"
)

type genreIdNode struct {
	id       int32
	children []*genreIdNode
	parent   *genreIdNode
}

func (node *genreIdNode) append(child *genreIdNode) *genreIdNode {
	child.parent = node
	node.children = append(node.children, child)
	return node
}

type genreTree struct {
	root  *genreIdNode
	table map[int32]*genreIdNode
}

func getGenres(dbResponse *db.WorkGenresDBResponse) *pb.Genre_Response {
	genresTable := make(map[uint16]*pb.Genre_Genre, len(dbResponse.Genres))

	for _, dbGenre := range dbResponse.Genres {
		genresTable[dbGenre.Id] = &pb.Genre_Genre{
			Id:        uint32(dbGenre.Id),
			Name:      dbGenre.Name,
			Info:      dbGenre.Info,
			WorkCount: dbGenre.WorkCount,
		}
	}

	groupsTable := make(map[uint16]*pb.Genre_Genre)

	for _, dbGenre := range dbResponse.Genres {
		root := groupsTable[dbGenre.GroupId]

		if root == nil {
			root = &pb.Genre_Genre{}
		}

		genre := genresTable[dbGenre.Id]
		parentGenre := genresTable[dbGenre.ParentId]

		if parentGenre != nil {
			parentGenre.Subgenres = append(parentGenre.Subgenres, genre)
		} else {
			root.Subgenres = append(root.Subgenres, genre)
		}

		groupsTable[dbGenre.GroupId] = root
	}

	genreGroups := make([]*pb.Genre_Group, len(dbResponse.GenreGroups))

	for index, dbGroup := range dbResponse.GenreGroups {
		rootGenre := groupsTable[dbGroup.Id]

		if rootGenre == nil {
			continue
		}

		group := &pb.Genre_Group{
			Id:     uint32(dbGroup.Id),
			Name:   dbGroup.Name,
			Genres: rootGenre.Subgenres,
		}

		genreGroups[index] = group
	}

	return &pb.Genre_Response{
		Groups: genreGroups,
	}
}

func makeGenreTree(dbResponse *db.WorkGenresDBResponse) *genreTree {
	root := &genreIdNode{}

	genresTable := make(map[int32]*genreIdNode, len(dbResponse.Genres)+len(dbResponse.GenreGroups))

	// для удобства, группы жанров - тоже жанры, но с отрицательными идентификаторами

	for _, dbGroup := range dbResponse.GenreGroups {
		groupNode := &genreIdNode{id: -int32(dbGroup.Id)}
		genresTable[groupNode.id] = groupNode
		root.append(groupNode)
	}

	for _, dbGenre := range dbResponse.Genres {
		nodeId := int32(dbGenre.Id)
		genresTable[nodeId] = &genreIdNode{id: nodeId}
	}

	for _, dbGenre := range dbResponse.Genres {
		genre := genresTable[int32(dbGenre.Id)]

		parentGenre := genresTable[int32(dbGenre.ParentId)]

		if parentGenre != nil {
			parentGenre.append(genre)
		} else {
			groupNode := genresTable[-int32(dbGenre.GroupId)]

			if groupNode != nil {
				groupNode.append(genre)
			}
		}
	}

	tree := &genreTree{
		root:  root,
		table: genresTable,
	}

	return tree
}

func checkRequiredGroupsForGenreIds(genreIds []uint64, tree *genreTree) error {
	var requiredGenreGroups = map[int32]string{
		-1: "Жанры/поджанры",
		-3: "Место действия",
		-4: "Время действия",
		-5: "Возраст читателя",
	}

	for _, genreId := range genreIds {
		node := tree.table[int32(genreId)]

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

func selectGenreIdsWithParents(genreIds []uint64, tree *genreTree) []int32 {
	var result []int32

	for _, genreId := range genreIds {
		node := tree.table[int32(genreId)]

		for node != nil {
			if node.id > 0 {
				result = append(result, node.id)
			}

			node = node.parent
		}
	}

	return result
}
