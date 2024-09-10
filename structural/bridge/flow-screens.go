package bridge

import (
	"fmt"
	"strings"
)

type IScreen interface {
	RenderProducts(categoryId string) string
}

type BaseScreen struct {
	ecommerce IEcommerce
}

func (s *BaseScreen) RenderProducts(categoryId string) string {
	products := s.ecommerce.ListProductsByCategory(categoryId)
	render := make([]string, 0, len(products))
	for i := range products {
		var b strings.Builder

		b.WriteString("*")
		b.WriteString(products[i].Id)
		b.WriteString("\n")
		b.WriteString("\t* name: ")
		b.WriteString(products[i].Name)
		b.WriteString("\n")
		b.WriteString("\t* slug: ")
		b.WriteString(products[i].Slug)
		b.WriteString("\n")
		b.WriteString("\t* price: R$")
		b.WriteString(fmt.Sprint(products[i].Price))
		b.WriteString("\n")
		render = append(render, b.String())
	}

	return strings.Join(render, "\n")
}

type UserScreen struct {
	BaseScreen
}

type BackofficeScreen struct {
	BaseScreen
}

func (s *BackofficeScreen) RenderProducts(categoryId string) string {
	products := s.BaseScreen.RenderProducts(categoryId)

	return strings.ToUpper(products)
}

func NewScreen(screen string, ecommerce IEcommerce) IScreen {
	switch true {
	case screen == "backoffice":
		return &BackofficeScreen{
			BaseScreen: BaseScreen{
				ecommerce: ecommerce,
			},
		}
	}
	return &UserScreen{
		BaseScreen{
			ecommerce: ecommerce,
		},
	}
}
