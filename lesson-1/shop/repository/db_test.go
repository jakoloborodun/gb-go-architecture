package repository

import (
	"github.com/stretchr/testify/require"
	"shop/models"
	"shop/util"
	"testing"
)

var mDB = mapDB{
	db:    make(map[int32]*models.Item, 5),
	maxID: 0,
}

func clearMockDB() {
	mDB = mapDB{
		db:    make(map[int32]*models.Item, 5),
		maxID: 0,
	}
}

// Creates shop item entity with random values.
func createRandomItem(t *testing.T) *models.Item {
	itemArgs := &models.Item{
		Name:  util.RandomName(),
		Price: util.RandomPrice(),
	}

	item, err := mDB.CreateItem(itemArgs)
	require.NoError(t, err)
	require.NotEmpty(t, item)
	require.NotZero(t, item.ID)

	require.Equal(t, itemArgs.Name, item.Name)
	require.Equal(t, itemArgs.Price, item.Price)

	return item
}

func TestMapDBCreateItem(t *testing.T) {
	mDB := mapDB{
		db:    make(map[int32]*models.Item, 5),
		maxID: 0,
	}

	currentID := int32(1)
	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_1",
		Price: 1000,
	}
	currentID++

	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_2",
		Price: 1500,
	}
	currentID++

	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_3",
		Price: 2000,
	}
	// TEST BEGINS HERE

	mDB.maxID = currentID

	newItem := &models.Item{
		Name:  "TestName_4",
		Price: 2500,
	}

	createdItem, err := mDB.CreateItem(newItem)
	if err != nil {
		t.Fatal(err)
	}
	currentID++

	if createdItem.ID != currentID {
		t.Fatal("expected id == ")
	}
	if createdItem.Name != newItem.Name {
		t.Fatal("expected name == ")
	}
	if createdItem.Price != newItem.Price {
		t.Fatal("expected name == ")
	}

	existingItem := mDB.db[currentID]
	if existingItem == nil {
		t.Fatal("got nil item")
	}

	if existingItem.ID != currentID {
		t.Fatal("expected id == ")
	}
	if existingItem.Name != newItem.Name {
		t.Fatal("expected name == ")
	}
	if existingItem.Price != newItem.Price {
		t.Fatal("expected name == ")
	}
}

func TestMapDB_GetItem(t *testing.T) {
	item1 := createRandomItem(t)
	item2, err := mDB.GetItem(item1.ID)

	//item1.Name = "wrongName" // Uncomment this line to TEST FAIL

	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, item1.Name, item2.Name)
	require.Equal(t, item1.Price, item2.Price)
}

func TestMapDB_UpdateItem(t *testing.T) {
	item1 := createRandomItem(t)

	newParams := &models.Item{
		ID:    item1.ID,
		Name:  util.RandomName(),
		Price: util.RandomPrice(),
	}

	item2, err := mDB.UpdateItem(newParams)

	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, newParams.Name, item2.Name)
	require.Equal(t, newParams.Price, item2.Price)

	//require.Equal(t, item1.Price, item2.Price) // Uncomment this line to TEST FAIL
}

func TestMapDB_DeleteItem(t *testing.T) {
	item1 := createRandomItem(t)
	err := mDB.DeleteItem(item1.ID)
	require.NoError(t, err)

	// Make sure Item was deleted. GetItem should return an error and item2 should be empty.
	item2, err := mDB.GetItem(item1.ID)
	require.Error(t, err)
	require.Empty(t, item2)
}

func TestMapDB_GetAllItems(t *testing.T) {
	clearMockDB()

	for i := 0; i < 5; i++ {
		createRandomItem(t)
	}

	items, err := mDB.GetAllItems()
	require.NoError(t, err)
	require.Len(t, items, 5)

	for _, item := range items {
		require.NotEmpty(t, item)
	}
}
