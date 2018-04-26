package controllers

import (
	"github.com/astaxie/beego"

	"EZChain/g"
	"EZChain/blockchain"
	"EZChain/utils"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "index.html"
}

func (this *MainController) Post() {
	var bodySlice = make([]byte, 1024)
	bodySlice = this.Ctx.Input.RequestBody
	ret := string(bodySlice)

	g.Mutex.Lock()
	newBlock := blockchain.BC.GenerateBlock(blockchain.BC.Blockchain[len(blockchain.BC.Blockchain)-1], ret)
	g.Mutex.Unlock()

	utils.PrintBlock(newBlock)
	this.Data["json"] = map[string]interface{}{"code": 0, "message": "Done"}
	this.ServeJSON()
	return
}