package helper

import (
	"fmt"
	"os"

	"encoding/json"
	"time"

	"net/http"

	"github.com/CloudyKit/jet"
	"github.com/golang/groupcache/lru"
	"github.com/parnurzeal/gorequest"
	"gopkg.in/gin-gonic/gin.v1"
)

var host string = os.Getenv("STATIC_ADDRESS")
var cache *lru.Cache = lru.New(8)

type appAsset struct {
	CSS string `json:"css"`
	JS  string `json:"js"`
}

type assetCache struct {
	Asset  *appAsset
	Expire time.Time
}

func getAssetsFromRemote() (map[string]appAsset, error) {
	var data map[string]appAsset
	request := gorequest.New()
	url := fmt.Sprintf("%s/assets.json", host)
	_, body, errs := request.Get(url).End()
	if len(errs) != 0 {
		return nil, errs[0]
	}

	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func hitCache(name string, now time.Time) (*appAsset, bool) {
	data, exists := cache.Get(name)
	if !exists {
		return nil, false
	}

	cacheAsset := data.(*assetCache)
	if cacheAsset.Expire.Before(now) {
		return nil, false
	}

	return cacheAsset.Asset, true
}

func getAssets(name string) ([]*appAsset, error) {
	assets := make([]*appAsset, 0, 3)
	requests := []string{"manifest", "vendor", name}
	allFound := true
	now := time.Now()
	for _, request := range requests {
		asset, exists := hitCache(request, now)
		if !exists {
			allFound = false
			break
		}

		assets = append(assets, asset)
	}

	if allFound {
		return assets, nil
	}

	remoteAssets, err := getAssetsFromRemote()
	if err != nil {
		return nil, err
	}

	for _, request := range requests {
		if asset, found := remoteAssets[request]; found {
			cacheAssets := &assetCache{
				Asset:  &asset,
				Expire: now.Add(10 * time.Minute),
			}

			cache.Add(request, cacheAssets)
			assets = append(assets, &asset)
		}
	}

	return assets, nil
}

func getAllJS(assets []*appAsset) (js []string) {
	js = make([]string, 0, 3)
	for _, asset := range assets {
		if asset.JS != "" {
			js = append(js, asset.JS)
		}
	}
	return
}

func getAllCSS(assets []*appAsset) (css []string) {
	css = make([]string, 0, 3)
	for _, asset := range assets {
		if asset.CSS != "" {
			css = append(css, asset.CSS)
		}
	}
	return
}

func RenderAppView(c *gin.Context, code int, view, title string) {
	assets, err := getAssets(view)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		panic(err)
	}

	vars := make(jet.VarMap)
	staticPath := fmt.Sprintf("%s", host)

	vars.Set("scripts", getAllJS(assets))
	vars.Set("styles", getAllCSS(assets))

	vars.Set("title", title)
	vars.Set("view", view)
	vars.Set("script", fmt.Sprintf("%s/%s/app.js", staticPath, view))
	vars.Set("style", fmt.Sprintf("%s/%s/app.css", staticPath, view))
	vars.Set("staticPath", staticPath)
	c.HTML(code, "app.jet", vars)
}
