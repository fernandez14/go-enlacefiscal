package efiscal

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("Create API wrapper", t, func() {

		api := API{"testuser", "testkey", false}

		Convey("Then create a test invoice using API wrapper", func() {

			invoice := api.Invoice("XAXX010101000", "TEST", "1")

			So(invoice.Series, ShouldEqual, "TEST")
			So(invoice.Folio, ShouldEqual, "1")

			Convey("Then add a new invoice item and check totals", func() {

				item := Item{
					Quantity:    5,
					Value:       15.0,
					Unit:        "producto",
					Description: "Producto de Prueba",
				}

				invoice.AddItem(item)

				So(invoice.Items.List, ShouldNotBeEmpty)
				So(invoice.GetSubtotal(), ShouldEqual, 75.0)

				Convey("After item adding transfer IVA taxes and check totals", func() {

					invoice.TransferIVA(16)

					So(invoice.Items.List, ShouldNotBeEmpty)
					So(invoice.Taxes.Transfers, ShouldNotBeEmpty)
					So(invoice.GetSubtotal(), ShouldEqual, 75.0)
					So(invoice.GetTotal(), ShouldEqual, 87.0)
				})
			})
		})
	})
}
