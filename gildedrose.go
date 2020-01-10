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

func updateSulfuras(sulfuras *Item) {}

func updateNormalItem(item *Item) {
	item.SellIn--
	if item.Quality == 0 {
		return
	}

	item.Quality--
	if item.SellIn < 0 {
		item.Quality--
	}
}

func updateBrie(brie *Item) {
	brie.SellIn--
	if brie.Quality == 50 {
		return
	}

	brie.Quality++
	if brie.SellIn < 0 {
		brie.Quality++
	}
}

func updatePasses(passes *Item) {
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

		if item.Name == "Sulfuras, Hand of Ragnaros" {
			updateSulfuras(item)
		} else if item.Name == "Aged Brie" {
			updateBrie(item)
		} else if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
			updatePasses(item)
		} else {
			updateNormalItem(item)
		}
	}
}
