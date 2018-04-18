package main

import (
	. "github.com/cool2645/kotori-ng/kotoriplugin"
	. "github.com/cool2645/kotori-core/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/BurntSushi/toml"
	"github.com/cool2645/kotori-core/model"
	"github.com/cool2645/kotori-core/handler/v2"
)

type CorePlugin struct{}

var pi = PluginInfo{
	Name:    "kotori-core",
	Author:  "2645 Studio",
	Version: "2.0-pre-alpha.1",
	License: "Unlicense",
	URL:     "https://github.com/cool2645/kotori-core",
}

func (p *CorePlugin) GetPluginInfo() PluginInfo {
	return pi
}

func (p *CorePlugin) LoadConfig() error {
	_, err := toml.DecodeFile("conf.d/kotori-core.toml", &GlobCfg)
	return err
}

func (p *CorePlugin) RegRouter(r *mux.Router) error {
	v2api := r.PathPrefix("/v2").Subrouter()
	v2api.Methods("GET").Path("").HandlerFunc(v2.Pong)
	return nil
}

func (p *CorePlugin) InitDB(db *gorm.DB) error {
	db.AutoMigrate(&model.User{}, &model.Category{}, &model.Tag{}, &model.Post{}, &model.Edition{}, &model.Comment{})
	db.Model(&model.Post{}).Related(&model.Tag{}, "Tags")
	return nil
}

var PluginInstance Plugin = &CorePlugin{}
