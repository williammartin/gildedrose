package gildedrose_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/williammartin/gildedrose"
)

var _ = Describe("The Gilded Rose", func() {

	Describe("Sulfuras, Hand of Ragnaros", func() {
		var sulfuras *Item

		BeforeEach(func() {
			sulfuras = &Item{Name: "Sulfuras, Hand of Ragnaros", SellIn: 20, Quality: 80}
			shop := &GildedRose{Inventory: []*Item{sulfuras}}
			shop.UpdateInventory()
		})

		It("always has a quality of 80", func() {
			Expect(sulfuras.Quality).To(Equal(80))
		})

		It("never needs to be sold", func() {
			Expect(sulfuras.SellIn).To(Equal(20))
		})
	})

	Describe("Any normal item", func() {
		var normalItem *Item

		BeforeEach(func() {
			normalItem = &Item{Name: "Normal", SellIn: 20, Quality: 10}
			shop := &GildedRose{Inventory: []*Item{normalItem}}
			shop.UpdateInventory()
		})

		It("decreases quality by 1 each day", func() {
			Expect(normalItem.Quality).To(Equal(9))
		})

		It("decreases sell in date by 1 each day", func() {
			Expect(normalItem.SellIn).To(Equal(19))
		})

		When("the sell by date has passed", func() {
			BeforeEach(func() {
				normalItem = &Item{Name: "Normal", SellIn: 0, Quality: 10}
				shop := &GildedRose{Inventory: []*Item{normalItem}}
				shop.UpdateInventory()
			})

			It("decreases the quality by 2 each day", func() {
				Expect(normalItem.Quality).To(Equal(8))
			})
		})

		When("the quality of an item has reached 0", func() {
			BeforeEach(func() {
				normalItem = &Item{Name: "Normal", SellIn: 0, Quality: 0}
				shop := &GildedRose{Inventory: []*Item{normalItem}}
				shop.UpdateInventory()
			})

			It("cannot decrease any further", func() {
				Expect(normalItem.Quality).To(Equal(0))
			})
		})
	})

	Describe("Aged Brie", func() {
		var brie *Item

		BeforeEach(func() {
			brie = &Item{Name: "Aged Brie", SellIn: 20, Quality: 10}
			shop := &GildedRose{Inventory: []*Item{brie}}
			shop.UpdateInventory()
		})

		It("increases in quality by 1 each day", func() {
			Expect(brie.Quality).To(Equal(11))
		})

		It("decreases sell in date by 1 each day", func() {
			Expect(brie.SellIn).To(Equal(19))
		})
	})

	Describe("Concert Passes", func() {
		var passes *Item

		BeforeEach(func() {
			passes = &Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 20, Quality: 10}
			shop := &GildedRose{Inventory: []*Item{passes}}
			shop.UpdateInventory()
		})

		It("increases in quality by 1 each day", func() {
			Expect(passes.Quality).To(Equal(11))
		})

		It("decreases sell in date by 1 each day", func() {
			Expect(passes.SellIn).To(Equal(19))
		})
	})
})
