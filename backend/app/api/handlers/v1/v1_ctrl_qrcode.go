package v1

import (
	"bytes"
	"context"
	_ "embed"
	"github.com/go-playground/validator/v10"
	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/hay-kot/homebox/backend/internal/web/adapters"
	"github.com/hay-kot/httpkit/errchain"
	"github.com/signintech/gopdf"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"image/png"
	"io"
	"net/http"
)

//go:embed assets/QRIcon.png
var qrcodeLogo []byte

type query struct {
	// 4,296 characters is the maximum length of a QR code
	Data string `schema:"data" validate:"required,max=4296"`
}

type pageConfigQuery struct {
	Nested    *bool  `schema:"nested" validate:"required"`
	NestLevel int    `schema:"nestLevel" validate:"required,gte=1"`
	PrintType string `schema:"printType" validate:"required,oneof=items locations both"`
	BaseAddr  string `schema:"baseAddr" validate:"required"`
}

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error {
	return nil
}

// HandleGenerateQRCode godoc
//
//	@Summary  Create QR Code
//	@Tags     Items
//	@Produce  json
//	@Param    data      query    string   false "data to be encoded into qrcode"
//	@Success 200 {string} string "image/jpeg"
//	@Router   /v1/qrcode [GET]
//	@Security Bearer
func (ctrl *V1Controller) HandleGenerateQRCode() errchain.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) error {
		q, err := adapters.DecodeQuery[query](r)
		if err != nil {
			return err
		}

		image, err := png.Decode(bytes.NewReader(qrcodeLogo))
		if err != nil {
			panic(err)
		}

		qrc, err := qrcode.New(q.Data)
		if err != nil {
			return err
		}

		toWriteCloser := struct {
			io.Writer
			io.Closer
		}{
			Writer: w,
			Closer: io.NopCloser(nil),
		}

		qrwriter := standard.NewWithWriter(toWriteCloser, standard.WithLogoImage(image))

		// Return the QR code as a jpeg image
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Disposition", "attachment; filename=qrcode.jpg")
		return qrc.Save(qrwriter)
	}
}

// HandleGenerateQRCodeForLocations godoc
//
// @Summary Create PDF of QR codes of a location
// @Tags Items
// @Produce application/pdf
// @Param location_id path string true "UUID of the location"
// @Param page query pageConfigQuery true "query info for how the QR codes are generated"
// @Success 200 {string} string "application/pdf"
// @Router /v1/qrcode/{location_id} [GET]
// @Security Bearer
func (ctrl *V1Controller) HandleGenerateQRCodeForLocations() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {

		data, err := adapters.DecodeQuery[pageConfigQuery](r) //decode the query with the pageConfigQuery struct
		if err != nil {
			return err
		}

		routeUUID, err := ctrl.routeUUID(r, "location_id") //from /qrcode/{location_id}
		if err != nil {
			return err
		}

		err = validator.New().Var(routeUUID.String(), "required,uuid") //validate the UUID, so .Get below won't throw errors
		if err != nil {
			return err
		}

		auth := services.NewContext(r.Context())
		locations, err := ctrl.repo.Locations.Get(auth, routeUUID) //get the location
		if err != nil {
			return err
		}

		var URLs []string
		URLs, err = ctrl.generateQRCodeURLs(locations, URLs, data) //this will append URLS to the URLs slice based on config (data)
		if err != nil {
			return err
		}

		pdf, err := ctrl.generatePDF(URLs) //URLs to PDF
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=qrCodes.pdf")

		_, err = pdf.WriteTo(w)
		if err != nil {
			return err
		}
		return err
	}
}

func (ctrl *V1Controller) generateQRCodeURLs(input repo.LocationOut, output []string, printConfig pageConfigQuery) ([]string, error) {

	// processLocation processes the given location and its items or children based on the printConfig settings
	if printConfig.PrintType == "items" || printConfig.PrintType == "both" {
		for _, item := range input.Items {
			output = append(output, printConfig.BaseAddr+"/item/"+item.ID.String())
		}
	}

	if printConfig.PrintType == "locations" || printConfig.PrintType == "both" {
		for _, child := range input.Children {
			output = append(output, printConfig.BaseAddr+"/location/"+child.ID.String())
		}
	}

	if *printConfig.Nested && printConfig.NestLevel > 1 {
		printConfig.NestLevel-- //decrease nest level for each recursive call
		for _, child := range input.Children {
			childLocation, err := ctrl.repo.Locations.Get(context.Background(), child.ID)
			if err != nil {
				return nil, err
			}
			output, err = ctrl.generateQRCodeURLs(childLocation, output, printConfig)
			if err != nil {
				return nil, err
			}
		}
	}

	return output, nil
}

func (ctrl *V1Controller) generatePDF(URLs []string) (*gopdf.GoPdf, error) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	rowCounter := 0

	const qrPerRow = 4
	const qrPerPage = 24

	for i, thing := range URLs {
		if i%qrPerRow == 0 && i != 0 {
			rowCounter++
			if i%qrPerPage == 0 {
				pdf.AddPage()
			}
		}

		//turn each thing into a image buffer that gopdf accepts
		qrc, err := qrcode.New(thing)
		if err != nil {
			return nil, err
		}

		logo, err := png.Decode(bytes.NewReader(qrcodeLogo))
		if err != nil {
			return nil, err
		}

		img := bytes.NewBuffer(nil)
		wr := standard.NewWithWriter(nopCloser{Writer: img}, standard.WithLogoImage(logo), standard.WithBorderWidth(5), standard.WithQRWidth(7))

		err = qrc.Save(wr)
		if err != nil {
			return nil, err
		}

		//add the image buffer to the pdf
		imgBytes, err := gopdf.ImageHolderByBytes(img.Bytes())
		if err != nil {
			return nil, err
		}

		x := (i % 4) * 150
		y := rowCounter * 150
		err = pdf.ImageByHolder(imgBytes, float64(x), float64(y), nil)
		if err != nil {
			return nil, err
		}
	}
	return &pdf, nil
}
