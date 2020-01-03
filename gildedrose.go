package gildedrose

type Item struct {
	name    string
	sellIn  int
	quality int
}

func UpdateInventory(items []*Item) {
	for _, item := range items {
		if item.name == "Anything" {
			UpdateAnything(item)
			continue
		} else if item.name == "Aged Brie" {
			UpdateBrie(item)
			continue
		} else if item.name == "Sulfuras, Hand of Ragnaros" {
			UpdateSulfuras(item)
			continue
		} else if item.name == "Backstage passes to a TAFKAL80ETC concert" {
			UpdateBackstage(item)
			continue
		}

		if item.name != "Aged Brie" && item.name != "Backstage passes to a TAFKAL80ETC concert" {
			if item.quality > 0 {
				if item.name != "Sulfuras, Hand of Ragnaros" {
					item.quality = item.quality - 1
				}
			}
		} else {
			if item.quality < 50 {
				item.quality = item.quality + 1
				if item.name == "Backstage passes to a TAFKAL80ETC concert" {
					if item.sellIn < 11 {
						if item.quality < 50 {
							item.quality = item.quality + 1
						}
					}
					if item.sellIn < 6 {
						if item.quality < 50 {
							item.quality = item.quality + 1
						}
					}
				}
			}
		}

		if item.name != "Sulfuras, Hand of Ragnaros" {
			item.sellIn = item.sellIn - 1
		}

		if item.sellIn < 0 {
			if item.name != "Aged Brie" {
				if item.name != "Backstage passes to a TAFKAL80ETC concert" {
					if item.quality > 0 {
						if item.name != "Sulfuras, Hand of Ragnaros" {
							item.quality = item.quality - 1
						}
					}
				} else {
					item.quality = item.quality - item.quality
				}
			} else {
				if item.quality < 50 {
					item.quality = item.quality + 1
				}
			}
		}
	}
}

func UpdateAnything(item *Item) {
	item.sellIn = item.sellIn - 1
	if item.quality > 0 {
		item.quality = item.quality - 1
	}

	if item.quality > 0 && item.sellIn < 0 {
		item.quality = item.quality - 1
	}
}

func UpdateBrie(item *Item) {
	item.sellIn = item.sellIn - 1
	if item.quality == 50 {
		return
	}

	item.quality = item.quality + 1

	if item.sellIn < 0 {
		item.quality = item.quality + 1
	}
}

func UpdateSulfuras(item *Item) {

}

func UpdateBackstage(item *Item) {
	item.sellIn = item.sellIn - 1
	item.quality = item.quality + 1

	if item.sellIn < 10 {
		item.quality = item.quality + 1
	}

	if item.sellIn < 5 {
		item.quality = item.quality + 1
	}

	if item.sellIn < 0 {
		item.quality = 0
	}

	if item.quality > 50 {
		item.quality = 50
	}
}
