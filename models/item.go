// структура товара

package models

type Item struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Checked bool   `json:"checked"`
}
