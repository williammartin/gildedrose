package gildedrose_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/williammartin/gildedrose"
)

var _ = Describe("The Gilded Rose", func() {

	var (
		inventory []*Item
		shop      *GildedRose
	)

	BeforeEach(func() {
		inventory = []*Item{}
	})

	JustBeforeEach(func() {
		shop = &GildedRose{Inventory: inventory}
		shop.UpdateInventory()
	})

	Describe("An item with any non-special name", func() {
		var normalItem *Item

		BeforeEach(func() {
			normalItem = &Item{Name: "Non-special", SellIn: 20, Quality: 10}
			inventory = append(inventory, normalItem)
		})

		It("decreases quality by 1 each day", func() {
			Expect(normalItem.Quality).To(Equal(9))
		})

		It("decreases sell in date by 1 each day", func() {
			Expect(normalItem.SellIn).To(Equal(19))
		})

		When("the sell by date has passed", func() {
			BeforeEach(func() {
				normalItem = &Item{Name: "Also Non-Special", SellIn: 0, Quality: 10}
				inventory = append(inventory, normalItem)
			})

			It("decreases the quality by 2 each day", func() {
				Expect(normalItem.Quality).To(Equal(8))
			})
		})

		When("the quality of an item has reached 0", func() {
			BeforeEach(func() {
				normalItem = &Item{Name: "Normal", SellIn: 0, Quality: 0}
				inventory = append(inventory, normalItem)
			})

			It("cannot decrease in quality any further", func() {
				Expect(normalItem.Quality).To(Equal(0))
			})
		})
	})

	Describe("Sulfuras, Hand of Ragnaros", func() {
		var sulfuras *Item

		BeforeEach(func() {
			sulfuras = &Item{Name: "Sulfuras, Hand of Ragnaros", SellIn: 20, Quality: 80}
			inventory = append(inventory, sulfuras)
		})

		It("always has a quality of 80", func() {
			Expect(sulfuras.Quality).To(Equal(80))
		})

		It("never needs to be sold", func() {
			Expect(sulfuras.SellIn).To(Equal(20))
		})
	})

	Describe("Aged Brie", func() {
		var brie *Item

		BeforeEach(func() {
			brie = &Item{Name: "Aged Brie", SellIn: 20, Quality: 10}
			inventory = append(inventory, brie)
		})

		It("increases in quality by 1 each day", func() {
			Expect(brie.Quality).To(Equal(11))
		})

		It("decreases sell in date by 1 each day", func() {
			Expect(brie.SellIn).To(Equal(19))
		})

		When("the quality of brie has reached 50", func() {
			BeforeEach(func() {
				brie = &Item{Name: "Aged Brie", SellIn: 20, Quality: 50}
				inventory = append(inventory, brie)
			})

			It("does not increase in quality any further", func() {
				Expect(brie.Quality).To(Equal(50))
			})
		})

		When("the sell by date has passed", func() {
			BeforeEach(func() {
				brie = &Item{Name: "Aged Brie", SellIn: 0, Quality: 10}
				inventory = append(inventory, brie)
			})

			It("increases the quality by 2 each day", func() {
				Expect(brie.Quality).To(Equal(12))
			})
		})
	})

	Describe("Concert Passes", func() {
		var passes *Item

		BeforeEach(func() {
			passes = &Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 20, Quality: 10}
			inventory = append(inventory, passes)
		})

		It("increases in quality by 1 each day", func() {
			Expect(passes.Quality).To(Equal(11))
		})

		It("decreases sell in date by 1 each day", func() {
			Expect(passes.SellIn).To(Equal(19))
		})

		When("the sell by date is within 10 days", func() {
			BeforeEach(func() {
				passes = &Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 10}
				inventory = append(inventory, passes)
			})

			It("increases the quality by 2 each day", func() {
				Expect(passes.Quality).To(Equal(12))
			})
		})

		When("the sell by date is within 5 days", func() {
			BeforeEach(func() {
				passes = &Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 10}
				inventory = append(inventory, passes)
			})

			It("increases the quality by 3 each day", func() {
				Expect(passes.Quality).To(Equal(13))
			})
		})

		When("the sell by date is passed", func() {
			BeforeEach(func() {
				passes = &Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 0, Quality: 10}
				inventory = append(inventory, passes)
			})

			It("the quality decreases to 0", func() {
				Expect(passes.Quality).To(Equal(0))
			})
		})

		When("the quality has reached 0", func() {
			BeforeEach(func() {
				passes = &Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 0, Quality: 10}
				inventory = append(inventory, passes)
			})

			It("cannot decrease in quality any further", func() {
				Expect(passes.Quality).To(Equal(0))
			})
		})
	})
})
