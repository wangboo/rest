// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Create", args).Url
}

func (_ tApp) All(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.All", args).Url
}

func (_ tApp) CreateImage(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.CreateImage", args).Url
}

func (_ tApp) Image(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Image", args).Url
}

func (_ tApp) Download(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Download", args).Url
}

func (_ tApp) Grape(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Grape", args).Url
}

func (_ tApp) Joke(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Joke", args).Url
}


type tJoke struct {}
var Joke tJoke


func (_ tJoke) Index(
		page int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "page", page)
	return revel.MainRouter.Reverse("Joke.Index", args).Url
}

func (_ tJoke) Image(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Joke.Image", args).Url
}

func (_ tJoke) Reply(
		id string,
		msg string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "msg", msg)
	return revel.MainRouter.Reverse("Joke.Reply", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tJobs struct {}
var Jobs tJobs


func (_ tJobs) Status(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Jobs.Status", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


