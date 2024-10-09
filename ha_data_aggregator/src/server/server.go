package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/the-kaustubh/ha_data_aggregator/config"
	hacontext "github.com/the-kaustubh/ha_data_aggregator/context"
	"github.com/the-kaustubh/ha_data_aggregator/model"
	"github.com/the-kaustubh/ha_data_aggregator/service"
)

func Init(config config.Configuration) error {
	mux := chi.NewMux()

	mux.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			headers := r.Header
			newCtx := hacontext.Context{
				Context:     ctx,
				Endpoint:    r.URL.Path,
				MachineName: headers.Get("machine"),
				MachineIp:   headers.Get("X-Real-Ip"),
			}
			data, err := io.ReadAll(r.Body)
			if err != nil {
				handleError(newCtx, w, http.StatusInternalServerError, err)
				return
			}
			newCtx.Body = data
			newReq := r.Clone(context.WithValue(ctx, "config", newCtx))
			h.ServeHTTP(w, newReq)
		})
	})

	mux.Post("/temperature", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_haCtx := ctx.Value("config")
		haCtx := _haCtx.(hacontext.Context)
		slog.InfoContext(ctx, "Hi")

		var temperatureModel model.TemperatureModel
		err := json.Unmarshal(haCtx.Body, &temperatureModel)
		if err != nil {
			handleError(haCtx, w, http.StatusBadRequest, err)
			return
		}
		err = service.WriteTemperature(haCtx, temperatureModel.MachineName, temperatureModel.Temperature)
		if err != nil {
			handleError(haCtx, w, http.StatusInternalServerError, err)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"msg":"OK"}`))
	})

	return http.ListenAndServe(":"+config.ServerPort, mux)
}

func handleError(ctx hacontext.Context, w http.ResponseWriter, status int, err error) {
	slog.ErrorContext(ctx, "Error while reading the body", "error", err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
}
