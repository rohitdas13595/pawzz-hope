package docs

import (
	"net/http"

	"github.com/go-swagno/swagno"
	"github.com/go-swagno/swagno-http/swagger"
	"github.com/go-swagno/swagno/components/endpoint"
	"github.com/go-swagno/swagno/components/http/response"
	"github.com/go-swagno/swagno/components/mime"
	"github.com/go-swagno/swagno/components/parameter"
	"github.com/go-swagno/swagno/example/models"
	"github.com/rohitdas13595/pawzz-hope/results"
)

func SwaggerHandler() http.Handler {
	sw := swagno.New(swagno.Config{Title: "Testing API", Version: "v1.0.0", Description: "Testing API", Host: "localhost:8080/api/v1"})

	endpoints := []*endpoint.EndPoint{
		endpoint.New(
			endpoint.GET,
			"/product",
			endpoint.WithTags("product"),
			endpoint.WithSuccessfulReturns([]response.Response{response.New(models.Product{}, "200", "OK")}),
			endpoint.WithErrors([]response.Response{response.New(models.UnsuccessfulResponse{}, "400", "Bad Request")}),
			endpoint.WithDescription("Get all products"),
			endpoint.WithProduce([]mime.MIME{mime.JSON, mime.XML}),
			endpoint.WithConsume([]mime.MIME{mime.JSON}),
			endpoint.WithSummary("Get all products"),
		),
		endpoint.New(
			endpoint.GET,
			"/product/{id}",
			endpoint.WithTags("product"),
			endpoint.WithParams(parameter.IntParam("id", parameter.Path, parameter.WithRequired())),
			endpoint.WithSuccessfulReturns([]response.Response{response.New(results.APIResponse[any]{}, "200", "Request Accepted")}),
			endpoint.WithErrors([]response.Response{response.New(models.UnsuccessfulResponse{}, "400", "Bad Request")}),
			endpoint.WithProduce([]mime.MIME{mime.JSON, mime.XML}),
		),
	}

	sw.AddEndpoints(endpoints)
	return swagger.SwaggerHandler(sw.MustToJson())

}
