package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"github.com/ulixes-bloom/ya-metrics/internal/pkg/headers"
	"github.com/ulixes-bloom/ya-metrics/internal/pkg/metrics"
)

func (a *api) GetMetric(res http.ResponseWriter, req *http.Request) {
	mtype := chi.URLParam(req, "mtype")
	mname := chi.URLParam(req, "mname")
	if mtype == "" || mname == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	mval, err := a.service.GetMetric(mtype, mname)
	if err != nil {
		log.Error().Msg(err.Error())
		http.Error(res, err.Error(), http.StatusNotFound)
	}

	res.Header().Add(headers.ContentType, "text/plain")
	res.Write([]byte(mval))
}

func (a *api) UpdateMetric(res http.ResponseWriter, req *http.Request) {
	mtype := chi.URLParam(req, "mtype")
	mname := chi.URLParam(req, "mname")
	mval := chi.URLParam(req, "mval")
	if mtype == "" || mname == "" || mval == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err := a.service.UpdateMetric(mtype, mname, mval)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
}

func (a *api) UpdateMetrics(res http.ResponseWriter, req *http.Request) {
	var m []metrics.Metric
	dec := json.NewDecoder(req.Body)
	if err := dec.Decode(&m); err != nil {
		log.Error().Msg(err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	err := a.service.UpdateMetrics(m)
	if err != nil {
		log.Error().Msg(err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Add(headers.ContentType, "application/json")
	res.WriteHeader(http.StatusOK)
}

func (a *api) GetMetricsHTMLTable(res http.ResponseWriter, req *http.Request) {
	table, err := a.service.GetMetricsHTMLTable()
	if err != nil {
		log.Error().Msg(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	res.Header().Add(headers.ContentType, "text/html; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write(table)
}

func (a *api) GetJSONMetric(res http.ResponseWriter, req *http.Request) {
	var m metrics.Metric
	dec := json.NewDecoder(req.Body)
	if err := dec.Decode(&m); err != nil {
		log.Error().Msg(err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	metric, err := a.service.GetJSONMetric(m)
	if err != nil {
		log.Error().Msg(err.Error())
		http.Error(res, err.Error(), http.StatusNotFound)
	}

	res.Header().Add(headers.ContentType, "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(metric)
}

func (a *api) UpdateJSONMetric(res http.ResponseWriter, req *http.Request) {
	var m metrics.Metric
	dec := json.NewDecoder(req.Body)
	if err := dec.Decode(&m); err != nil {
		log.Error().Msg(err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	metric, err := a.service.UpdateJSONMetric(m)
	if err != nil {
		log.Error().Msg(err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Add(headers.ContentType, "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(metric)
}

func (a *api) PingDB(res http.ResponseWriter, req *http.Request) {
	err := a.service.PingDB(a.conf.DatabaseDSN)
	if err != nil {
		log.Error().Msg(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
}
