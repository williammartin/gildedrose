package gildedrose

type Item struct {
	Name    string
	SellIn  int
	Quality int
}

type Inventory []*Item

type GildedRose struct {
	Inventory Inventory
}

type Sulfuras struct {
	*Item
}

func (s *Sulfuras) Update() {}

type Normal struct {
	Item *Item
}

func (item *Normal) Update() {
	item.Item.SellIn--
	if item.Item.Quality == 0 {
		return
	}

	item.Item.Quality--
	if item.Item.SellIn < 0 {
		item.Item.Quality--
	}
}

type Brie struct {
	*Item
}

func (brie *Brie) Update() {
	brie.SellIn--
	if brie.Quality == 50 {
		return
	}

	brie.Quality++
	if brie.SellIn < 0 {
		brie.Quality++
	}
}

type Passes struct {
	*Item
}

func (passes *Passes) Update() {
	passes.SellIn--
	if passes.SellIn < 0 {
		passes.Quality = 0
		return
	}

	passes.Quality++
	if passes.SellIn < 10 {
		passes.Quality++
	}

	if passes.SellIn < 5 {
		passes.Quality++
	}
}

func (g *GildedRose) UpdateInventory() {
	for i := 0; i < len(g.Inventory); i++ {
		item := g.Inventory[i]
		ItemFactory{}.Wrap(item).Update()
	}
}

type ItemFactory struct{}

func (f ItemFactory) Wrap(item *Item) Updateable {
	if item.Name == "Sulfuras, Hand of Ragnaros" {
		return &Sulfuras{
			Item: item,
		}
	} else if item.Name == "Aged Brie" {
		return &Brie{
			Item: item,
		}
	} else if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		return &Passes{
			Item: item,
		}
	} else {
		return &Normal{
			Item: item,
		}
	}
}

type Updateable interface {
	Update()
}
