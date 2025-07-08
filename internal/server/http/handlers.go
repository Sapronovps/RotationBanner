package http

import (
	"encoding/json"
	"github.com/Sapronovps/RotationBanner/internal/app"
	"github.com/Sapronovps/RotationBanner/internal/model"
	"github.com/Sapronovps/RotationBanner/internal/server/http/request"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"sync"
)

var (
	lock sync.Mutex
)

func home(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Microservice Rotation Banner"))
}

func addSlot(w http.ResponseWriter, r *http.Request, app *app.App) {
	lock.Lock()
	defer lock.Unlock()

	var slot model.Slot
	err := json.NewDecoder(r.Body).Decode(&slot)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = app.AddSlot(&slot)
	if err != nil {
		http.Error(w, "Error create slot: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(slot)
}

func getSlot(w http.ResponseWriter, r *http.Request, app *app.App) {
	lock.Lock()
	defer lock.Unlock()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid item ID "+err.Error(), http.StatusBadRequest)
		return
	}
	slot, err := app.GetSlot(id)
	if err != nil {
		http.Error(w, "Slot not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(slot)
}

func addBanner(w http.ResponseWriter, r *http.Request, app *app.App) {
	lock.Lock()
	defer lock.Unlock()

	var banner model.Banner
	err := json.NewDecoder(r.Body).Decode(&banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = app.AddBanner(&banner)
	if err != nil {
		http.Error(w, "Error adding banner: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(banner)
}

func getBanner(w http.ResponseWriter, r *http.Request, app *app.App) {
	lock.Lock()
	defer lock.Unlock()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid item ID "+err.Error(), http.StatusBadRequest)
		return
	}
	banner, err := app.GetBanner(id)
	if err != nil {
		http.Error(w, "Banner not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(banner)
}

func addGroup(w http.ResponseWriter, r *http.Request, app *app.App) {
	lock.Lock()
	defer lock.Unlock()

	var group model.Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = app.AddGroup(&group)
	if err != nil {
		http.Error(w, "Error adding group: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(group)
}

func getGroup(w http.ResponseWriter, r *http.Request, app *app.App) {
	lock.Lock()
	defer lock.Unlock()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid item ID "+err.Error(), http.StatusBadRequest)
		return
	}
	group, err := app.GetGroup(id)
	if err != nil {
		http.Error(w, "Group not found: "+err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

func addBannerGroupStats(w http.ResponseWriter, r *http.Request, app *app.App) {
	lock.Lock()
	defer lock.Unlock()

	var bannerGroupStats model.BannerGroupStats
	err := json.NewDecoder(r.Body).Decode(&bannerGroupStats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = app.AddBannerGroupStats(&bannerGroupStats)
	if err != nil {
		http.Error(w, "Error adding banner group stats: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bannerGroupStats)
}

func registerClick(w http.ResponseWriter, r *http.Request, app *app.App) {
	lock.Lock()
	defer lock.Unlock()

	var BannerStatsRequest request.BannerStatsRequest
	err := json.NewDecoder(r.Body).Decode(&BannerStatsRequest)
	if err != nil ||
		BannerStatsRequest.SlotID == 0 ||
		BannerStatsRequest.BannerID == 0 ||
		BannerStatsRequest.GroupID == 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = app.RegisterClick(BannerStatsRequest.SlotID, BannerStatsRequest.BannerID, BannerStatsRequest.GroupID)
	if err != nil {
		http.Error(w, "Error registering click: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{success: true}"))
}
