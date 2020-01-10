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

func (g *GildedRose) UpdateInventory() {
	for i := 0; i < len(g.Inventory); i++ {

		if g.Inventory[i].Name != "Aged Brie" && g.Inventory[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if g.Inventory[i].Quality > 0 {
				if g.Inventory[i].Name != "Sulfuras, Hand of Ragnaros" {
					g.Inventory[i].Quality = g.Inventory[i].Quality - 1
				}
			}
		} else {
			if g.Inventory[i].Quality < 50 {
				g.Inventory[i].Quality = g.Inventory[i].Quality + 1
				if g.Inventory[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
					if g.Inventory[i].SellIn < 11 {
						if g.Inventory[i].Quality < 50 {
							g.Inventory[i].Quality = g.Inventory[i].Quality + 1
						}
					}
					if g.Inventory[i].SellIn < 6 {
						if g.Inventory[i].Quality < 50 {
							g.Inventory[i].Quality = g.Inventory[i].Quality + 1
						}
					}
				}
			}
		}

		if g.Inventory[i].Name != "Sulfuras, Hand of Ragnaros" {
			g.Inventory[i].SellIn = g.Inventory[i].SellIn - 1
		}

		if g.Inventory[i].SellIn < 0 {
			if g.Inventory[i].Name != "Aged Brie" {
				if g.Inventory[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
					if g.Inventory[i].Quality > 0 {
						if g.Inventory[i].Name != "Sulfuras, Hand of Ragnaros" {
							g.Inventory[i].Quality = g.Inventory[i].Quality - 1
						}
					}
				} else {
					g.Inventory[i].Quality = g.Inventory[i].Quality - g.Inventory[i].Quality
				}
			} else {
				if g.Inventory[i].Quality < 50 {
					g.Inventory[i].Quality = g.Inventory[i].Quality + 1
				}
			}
		}
	}
}
