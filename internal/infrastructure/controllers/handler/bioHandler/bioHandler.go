package bioHandler

import (
	"em-test-task/internal/domain/entity"
	"em-test-task/internal/domain/services"
	"em-test-task/internal/domain/storages/dto"
	httpController "em-test-task/internal/infrastructure/controllers/handler"
	"em-test-task/internal/infrastructure/gateways/ageGateway"
	"em-test-task/internal/infrastructure/gateways/genderGateway"
	"em-test-task/internal/infrastructure/gateways/nationalGateway"
	"encoding/json"
	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"strconv"
	"sync"
)

type BioHandler struct {
	bioService    services.BioInfoService
	filterService services.FilterService

	ageGateway      ageGateway.AgeGateway
	genderGateway   genderGateway.GenderGateway
	nationalGateway nationalGateway.NationalGateway

	log *logrus.Entry
}

func (bH *BioHandler) RegisterRouter(group *router.Group) {
	group.GET("/get", bH.Get)

	group.POST("/add", bH.Add)

	group.POST("/updateById", bH.UpdateById)

	group.DELETE("/deleteById", bH.DeleteById)
}

func NewBioHandler(bioService services.BioInfoService, filterService services.FilterService, ageGateway ageGateway.AgeGateway, genderGateway genderGateway.GenderGateway, nationalGateway nationalGateway.NationalGateway, log *logrus.Entry) httpController.HTTPHandler {
	return &BioHandler{bioService: bioService, filterService: filterService, ageGateway: ageGateway, genderGateway: genderGateway, nationalGateway: nationalGateway, log: log}
}

func (bH *BioHandler) Get(ctx *fasthttp.RequestCtx) {
	args := ctx.QueryArgs()
	if args.Len() == 0 {
		bIs, err := bH.bioService.GetAll(ctx)
		if err != nil {
			ctx.Error("error getting all info", fasthttp.StatusInternalServerError)
			return
		}
		resp, err := json.Marshal(bIs)
		if err != nil {
			ctx.Error("error while preparing response", fasthttp.StatusInternalServerError)
			return
		}

		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Success(httpController.JsonType, resp)
	}

	rawFilters := make(map[string]string)
	args.VisitAll(func(key, value []byte) {
		nameFilter := string(key)
		valueFilter := string(value)
		rawFilters[nameFilter] = valueFilter
	})

	filters, _, err := bH.filterService.ValidateFilters(rawFilters)
	if err != nil {
		ctx.Error("error processing filters", fasthttp.StatusBadRequest)
		return
	}

	bIs, err := bH.bioService.GetByFilters(ctx, filters)
	if err != nil {
		ctx.Error("error getting information", fasthttp.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(bIs)
	if err != nil {
		ctx.Error("error while preparing responce", fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Success(httpController.JsonType, resp)
}

func (bH *BioHandler) UpdateById(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	var bI *entity.BioInfo
	err := json.Unmarshal(body, &bI)
	if err != nil {
		bH.log.Errorln(err)
		ctx.Error("invalid format", fasthttp.StatusBadRequest)
		return
	}

	err = bH.bioService.UpdateById(ctx, bI)
	if err != nil {
		ctx.Error("could not be changed due to a server error", fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SuccessString(httpController.JsonType, "")
}

func (bH *BioHandler) DeleteById(ctx *fasthttp.RequestCtx) {
	idQ := ctx.QueryArgs().Peek("id")
	idC := string(idQ)
	id, err := strconv.Atoi(idC)
	if err != nil {
		bH.log.Errorln(err)
		ctx.Error("invalid id type", fasthttp.StatusBadRequest)
		return
	}

	bI, err := bH.bioService.DeleteById(ctx, int(id))
	if err != nil {
		ctx.Error("the user was not deleted due to a server error", fasthttp.StatusInternalServerError)
		return
	}
	bH.log.Debugln("delete user: ", bI)

	resp, err := json.Marshal(bI)
	if err != nil {
		bH.log.Errorln(err)
		ctx.Error("error while preparing response", fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Success(httpController.JsonType, resp)
}

func (bH *BioHandler) Add(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()

	var bioInfo *dto.BioInfoCreate
	err := json.Unmarshal(body, &bioInfo)
	if err != nil {
		bH.log.Errorln(err)
		ctx.Error("invalid json body format", fasthttp.StatusUnprocessableEntity)
		return
	}
	bH.log.Debugln("request body: ", bioInfo)

	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	wg.Add(1)
	go func() {
		age, err := bH.ageGateway.GetAgeByName(ctx, bioInfo.Name)
		if err != nil {
			ctx.Error("error get age by name", fasthttp.StatusInternalServerError)
			return
		}
		m.Lock()
		bioInfo.Age = age
		m.Unlock()
		bH.log.Debugln("age: ", age)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		gender, err := bH.genderGateway.GetGenderByName(ctx, bioInfo.Name)
		if err != nil {
			ctx.Error("error get gender by name", fasthttp.StatusInternalServerError)
			return
		}
		m.Lock()
		bioInfo.Gender = gender
		m.Unlock()
		bH.log.Debugln("gender: ", gender)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		national, err := bH.nationalGateway.GetNationalByName(ctx, bioInfo.Name)
		if err != nil {
			ctx.Error("error get national by name", fasthttp.StatusInternalServerError)
			return
		}
		m.Lock()
		bioInfo.National = national
		m.Unlock()
		bH.log.Debugln("national: ", national)
		wg.Done()
	}()

	wg.Wait()

	bI, err := bH.bioService.Add(ctx, bioInfo)
	if err != nil {
		ctx.Error("error storing to database", fasthttp.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(bI)
	if err != nil {
		ctx.Error("error while preparing response", fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Success(httpController.JsonType, resp)
}
