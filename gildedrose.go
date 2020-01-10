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
			continue
		} else if item.Name == "Normal" {
			updateNormalItem(item)
			continue
		} else if item.Name == "Aged Brie" {
			updateBrie(item)
			continue
		} else if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
			updatePasses(item)
			continue
		}

		if item.Name != "Aged Brie" && item.Name != "Backstage passes to a TAFKAL80ETC concert" {
			if item.Quality > 0 {
				if item.Name != "Sulfuras, Hand of Ragnaros" {
					item.Quality = item.Quality - 1
				}
			}
		} else {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
				if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
					if item.SellIn < 11 {
						if item.Quality < 50 {
							item.Quality = item.Quality + 1
						}
					}
					if item.SellIn < 6 {
						if item.Quality < 50 {
							item.Quality = item.Quality + 1
						}
					}
				}
			}
		}

		if item.Name != "Sulfuras, Hand of Ragnaros" {
			item.SellIn = item.SellIn - 1
		}

		if item.SellIn < 0 {
			if item.Name != "Aged Brie" {
				if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
					if item.Quality > 0 {
						if item.Name != "Sulfuras, Hand of Ragnaros" {
							item.Quality = item.Quality - 1
						}
					}
				} else {
					item.Quality = item.Quality - item.Quality
				}
			} else {
				if item.Quality < 50 {
					item.Quality = item.Quality + 1
				}
			}
		}
	}
}
