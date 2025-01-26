package handlers

import (
	"context"
	"net/http"

	"github.com/plrone/celeritas"
)

func (h *Handlers) render(w http.ResponseWriter, r *http.Request, tmpl string, variables, data interface{}) error {
	return h.App.Render.Page(w, r, tmpl, variables, data)
}

func (h *Handlers) sessionPut(ctx context.Context, key string, data interface{}) {
	h.App.Session.Put(ctx, key, data)
}

func (h *Handlers) sessionHas(ctx context.Context, key string) bool {
	return h.App.Session.Exists(ctx, key)
}

func (h *Handlers) sessionGet(ctx context.Context, key string) interface{} {
	return h.App.Session.Get(ctx, key)
}

func (h *Handlers) sessionRemove(ctx context.Context, key string) {
	h.App.Session.Remove(ctx, key)
}

func (h *Handlers) sessionRenew(ctx context.Context) error {
	return h.App.Session.RenewToken(ctx)
}

func (h *Handlers) sessionDestroy(ctx context.Context) error {
	return h.App.Session.Destroy(ctx)
}

func (h *Handlers) randomString(n int) string {
	return h.App.RandomString(n)
}

func (h *Handlers) encrypt(text string) (string, error) {
	enc := celeritas.Encryption{Key: []byte(h.App.EncryptionKey)}
	return enc.Encrypt(text)
}

func (h *Handlers) decrypt(text string) (string, error) {
	dec := celeritas.Encryption{Key: []byte(h.App.EncryptionKey)}
	return dec.Decrypt(text)

}
