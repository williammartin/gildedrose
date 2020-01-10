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

func (g *GildedRose) UpdateInventory() {
	for i := 0; i < len(g.Inventory); i++ {
		item := g.Inventory[i]

		if item.Name == "Sulfuras, Hand of Ragnaros" {
			updateSulfuras(item)
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
